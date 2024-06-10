package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ashishjuyal/banking-lib/logger"
	"github.com/ashishjuyal/banking/domain"
	"github.com/ashishjuyal/banking/service"
	"github.com/jmoiron/sqlx"

	"github.com/gorilla/mux"
)

/*
func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}*/

func Start() {

	//sanityCheck()

	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}
	//router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	//router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	//router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	//router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting server
	//log.Fatal(http.ListenAndServe("localhost:8000", router))

	// define routes
	router.
		HandleFunc("/customers", ch.getAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).
		Methods(http.MethodPost).
		Name("NewAccount")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())
	// starting server
	//address := os.Getenv("SERVER_ADDRESS")
	address := "localhost"
	//port := os.Getenv("SERVER_PORT")
	port := "8080"
	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}

func getDbClient() *sqlx.DB {
	/*dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")*/

	/*dbUser := "root"
	//dbPasswd := ""
	dbAddr := "127.0.0.1"
	dbPort := "3306"
	dbName := "banking"
	*/
	//dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	//dataSource := fmt.Sprintf("%s@tcp(%s:%s)/%s", dbUser, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", "root:@tcp(localhost:3306)/banking")
	//client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
