package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/NikhilChoudhary001/ibmassignment/domain"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

/* type CustomerController struct {
	store domain.CustomerStore
} */

type Handler struct {
	Repository domain.CustomerStore
	Logger     *zap.Logger // Uber's Zap Logger
}

//HTTP Post - /api/customer
func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	//Flushing any buffered log entries
	defer h.Logger.Sync()
	var customer domain.Customer
	// Decode the incoming customer json
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		h.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if reflect.TypeOf(customer.Parameters).Kind() == reflect.Slice {
		fmt.Println("customer parameter is of array type %v", customer.Parameters)
	} else if reflect.TypeOf(customer.Parameters).Kind() == reflect.Map {
		fmt.Println("customer parameter is of map type %v", customer.Parameters)
	}

	// Create customer
	if err := h.Repository.Create(customer); err != nil {
		h.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.Logger.Info("created customer",
		zap.String("url", r.URL.String()),
	)
	w.WriteHeader(http.StatusCreated)
}

//HTTP Get - /api/customers
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	// Get all
	if customers, err := h.Repository.GetAll(); err != nil {
		h.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	} else {
		j, err := json.Marshal(customers)
		if err != nil {
			h.Logger.Error(err.Error(),
				zap.String("url", r.URL.String()),
			)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

//HTTP Get - /api/customer/{id}
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// Get by id
	if cust, err := h.Repository.GetById(id); err != nil {
		h.Logger.Error(err.Error(),
			zap.String("customer id", id),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(cust)
		if err != nil {
			h.Logger.Error(err.Error(),
				zap.String("url", r.URL.String()),
			)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

//HTTP Put - /api/customer/{id}
func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	defer h.Logger.Sync()
	vars := mux.Vars(r)
	id := vars["id"]
	var customer domain.Customer
	// Decode the incoming customer json
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		h.Logger.Error(err.Error(),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Update
	if err := h.Repository.Update(id, customer); err != nil {
		h.Logger.Error(err.Error(),
			zap.String("customer id", id),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.Logger.Info("updated customer",
		zap.String("customer id", id),
		zap.String("url", r.URL.String()),
	)
	w.WriteHeader(http.StatusNoContent)
}

//HTTP Delete - /api/customer/{id}
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	defer h.Logger.Sync()
	vars := mux.Vars(r)
	id := vars["id"]
	// delete
	if err := h.Repository.Delete(id); err != nil {
		h.Logger.Error(err.Error(),
			zap.String("customer id", id),
			zap.String("url", r.URL.String()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.Logger.Info("deleted customer",
		zap.String("customer id", id),
		zap.String("url", r.URL.String()),
	)
	w.WriteHeader(http.StatusNoContent)
}
