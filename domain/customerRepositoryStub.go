package domain

type CustomerRepositoryStub struct {
	customers []Customer
	//ById(string) (*Customer, *errs.AppError)
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Alexa", City: "Manizales", Zipcode: "110011", DateofBirth: "1982-12-07", Status: "1"},
		{Id: "1002", Name: "Dani", City: "Medellin", Zipcode: "110011", DateofBirth: "1982-12-07", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
