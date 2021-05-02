package domain

type Customer struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Email      string      `json:"email"`
	Parameters interface{} `json:"parameters,omitempty"`
	//Parameters1 Parameter   `json:"parameters,omitempty"`
}

type Parameter struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type CustomerStore interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	GetById(string) (Customer, error)
	GetAll() ([]Customer, error)
}
