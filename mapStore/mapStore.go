package mapStore

import (
	"github.com/NikhilChoudhary001/ibmassignment/domain"
)

type MapStore struct {
	store map[string]domain.Customer
}

// This is for caller packages, not for mapstore

func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

func (m *MapStore) Create(c domain.Customer) error {

	m.store[c.ID] = c

	/* if err != nil {
		fmt.Errorf("Error: %s", err)
		return err
	} */
	return nil

}

func (m *MapStore) Update(s string, c domain.Customer) error {
	if _, ok := m.store[s]; ok {
		m.store[s] = c
	}
	return nil
}

func (m *MapStore) Delete(s string) error {
	if _, ok := m.store[s]; ok {
		delete(m.store, s)
	}
	return nil
}

func (m *MapStore) GetById(s string) (domain.Customer, error) {
	var c domain.Customer
	if value, ok := m.store[s]; ok {
		c = value
	}

	return c, nil
}

func (m *MapStore) GetAll() ([]domain.Customer, error) {

	var allCustomers []domain.Customer

	for _, v := range m.store {
		allCustomers = append(allCustomers, v)
	}

	return allCustomers, nil
}
