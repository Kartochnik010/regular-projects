package main

import (
	inter "crud/internal"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

// createTable := "create table Products (id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY, model VARCHAR(255) NOT NULL, company VARCHAR(255) NOT NULL, price integer NOT NULL);"
// insertToTable := "insert into Products (model, company, price) values ('iphone','apple',240);"

type product struct {
	Id      int    `json:"id"`
	Model   string `json:"model"`
	Company string `json:"company"`
	Price   int    `json:"price"`
}

func select_products(db *sql.DB) []product {
	var products []product

	rows, err := db.Query("SELECT * FROM Products;")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	return products
}

func main() {
	db := inter.DB_conn()
	// result, err := db.Exec("insert into Products (model, company, price) values ('iPhone X', 'Apple', 72000);")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// RowsAffected, _ := result.RowsAffected()
	// fmt.Println("Rows affected: ", RowsAffected)
	artists := select_products(db)
	for _, v := range artists {
		fmt.Println(v.Id, v.Company, v.Model, v.Price)
	}

	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}