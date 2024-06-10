package app

import (
	"log"
	"net/http"

	"github.com/ashishjuyal/banking/domain"
	"github.com/ashishjuyal/banking/service"

	"github.com/gorilla/mux"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	//ch := CustomerHandlers{service.getAllCustomers(domain.NewCustomerRepositoryStub())}

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	//router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	//router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
