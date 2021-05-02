package controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"assignment/controller"
	"assignment/domain"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

var _ = Describe("CustomerController", func() {
	var r *mux.Router
	var w *httptest.ResponseRecorder
	var store *FakeCustomerStore
	var handler controller.Handler

	BeforeEach(func() {
		r = mux.NewRouter()
		store = newFakeCustomerStore()
		logger, _ := zap.NewProduction()
		handler = controller.Handler{
			// Injecting a test stub
			// In production code, this would be a persistent store
			Repository: store,
			Logger:     logger,
		}
	})

	Describe("Get list of Customers", func() {
		Context("Get all Customers from data store", func() {
			It("Should get list of Customers", func() {
				r.HandleFunc("/api/customers", handler.GetAll).Methods("GET")
				req, err := http.NewRequest("GET", "/api/customers", nil)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(200))
				var customers []domain.Customer
				json.Unmarshal(w.Body.Bytes(), &customers)
				// Verifying mocked data of 2 customers
				Expect(len(customers)).To(Equal(2))
				Expect(customers[0].Name).To(Equal("Nikhil Choudhary"))
			})
		})
	})

	Describe("Get a specific Customer", func() {
		Context("Get a Customer from data store", func() {
			It("Should get a Customer based on customer id provided", func() {
				r.HandleFunc("/api/customer/{id}", handler.Get).Methods("GET")
				customerId := "cust102"

				req, err := http.NewRequest("GET", "/api/customer/"+customerId, nil)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(200))
				var customer domain.Customer
				json.Unmarshal(w.Body.Bytes(), &customer)
				fmt.Printf("Customer based on Id Provided - %v\n", w.Body)
				Expect(customer.Name).To(Equal("Vicky Singh"))
			})
		})
	})

	Describe("Post a new Customer", func() {
		Context("Provide a valid Customer data", func() {
			It("Should create a new Customer and get HTTP Status: 201", func() {
				r.HandleFunc("/api/customer", handler.Post).Methods("POST")
				userJson := `{"id": "Alex", "name": "John", "email": "alex@xyz.com"}`

				req, err := http.NewRequest(
					"POST",
					"/api/customer",
					strings.NewReader(userJson),
				)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(201))
			})
		})
	})

	Describe("Delete a Customer", func() {
		Context("Provide a valid Customer Id", func() {
			It("Should delete a Customer and get HTTP Status: 204", func() {
				r.HandleFunc("/api/customer/{id}", handler.Delete).Methods("DELETE")
				customerId := "cust101"

				req, err := http.NewRequest(
					"DELETE",
					"/api/customer/"+customerId,
					nil,
				)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(204))
			})
		})
	})

	Describe("Update a Customer", func() {
		Context("Provide a valid Customer Id and customer data", func() {
			It("Should update a Customer and get HTTP Status: 204", func() {
				r.HandleFunc("/api/customer/{id}", handler.Put).Methods("PUT")
				customerId := "cust102"
				userJson := `{"id": "cust102", "name": "Sagar", "email": "sagar@xyz.com"}`

				req, err := http.NewRequest(
					"PUT",
					"/api/customer/"+customerId,
					strings.NewReader(userJson),
				)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(204))

			})
		})
	})

})

type FakeCustomerStore struct {
	customerStore map[string]domain.Customer
}

func (store *FakeCustomerStore) Create(customer domain.Customer) error {
	store.customerStore[customer.ID] = customer

	return nil
}

func (store *FakeCustomerStore) Update(id string, customer domain.Customer) error {
	//fmt.Printf("Customer data Inside Update stub - %v\n", store)
	if _, ok := store.customerStore[id]; ok {
		store.customerStore[id] = customer
	}
	//fmt.Printf("Customer data Inside Update stub - %v\n", store)
	return nil
}

func (store *FakeCustomerStore) Delete(id string) error {
	//fmt.Printf("Customer data Inside Delete stub - %v\n", store)
	if _, ok := store.customerStore[id]; ok {
		delete(store.customerStore, id)
	}
	//fmt.Printf("Customer data Inside Delete stub- %v\n", store)
	return nil
}

func (store *FakeCustomerStore) GetById(id string) (domain.Customer, error) {
	var customer domain.Customer

	if value, ok := store.customerStore[id]; ok {
		customer = value
	}

	return customer, nil
}

func (store *FakeCustomerStore) GetAll() ([]domain.Customer, error) {
	var allCustomers []domain.Customer

	for _, v := range store.customerStore {
		allCustomers = append(allCustomers, v)
	}

	return allCustomers, nil
}

func newFakeCustomerStore() *FakeCustomerStore {
	//store := &FakeCustomerStore{}
	store := &FakeCustomerStore{customerStore: make(map[string]domain.Customer)}
	store.Create(domain.Customer{
		ID:    "cust101",
		Name:  "Nikhil Choudhary",
		Email: "nikhil@xyz.com",
	})

	store.Create(domain.Customer{
		ID:    "cust102",
		Name:  "Vicky Singh",
		Email: "vicky@xyz.com",
	})
	return store
}
