package domain

import (
	"database/sql"
	"log"

	//	"github.com/ashishjuyal/banking-lib/errs"
	//	"github.com/ashishjuyal/banking-lib/logger"

	"github.com/ashishjuyal/banking/errs"
	"github.com/ashishjuyal/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	//	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d *CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	//var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
		//rows, err = d.client.Query(findAllSql)

	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		logger.Info(status)
		err = d.client.Select(&customers, findAllSql, status)
		//rows, err = d.client.Query(findAllSql, status)

	}
	if err != nil {
		log.Println("Error while quering table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected databse error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

/*
	func NewCustomerRepositoryDb() CustomerRepositoryDb {
		dbUser := os.Getenv("DB_USER")
		dbPasswd := os.Getenv("DB_PASSWD")
		dbAddr := os.Getenv("DB_ADDR")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)

		client, err := sqlx.Open("mysql", dataSource)
		//	datasource := fmt.Sprintf()
		//client, err := sqlx.Open("mysql", "root:@tcp(localhost:3306)/banking")
		if err != nil {
			panic(err)
		}
		// See "Important settings" section.
		client.SetConnMaxLifetime(time.Minute * 3)
		client.SetMaxOpenConns(10)
		client.SetMaxIdleConns(10)
		return CustomerRepositoryDb{client: client}
	}
*/
func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}

/*var err error
customers := make([]Customer, 0)

if status == "" {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	err = d.client.Select(&customers, findAllSql)
} else {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
	err = d.client.Select(&customers, findAllSql, status)
}

if err != nil {
	logger.Error("Error while querying customers table " + err.Error())
	return nil, errs.NewUnexpectedError("Unexpected database error")
}

return customers, nil*/

/*
import (
	"database/sql"
	"github.com/ashishjuyal/banking-lib/errs"
	"github.com/ashishjuyal/banking-lib/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}

*/
