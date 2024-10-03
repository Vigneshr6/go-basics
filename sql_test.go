package main

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const dsn string = "file:/Users/vigneshr/Developer/docs/Car_Database.db"

type Model struct {
	Id      string `db:"model_id"`
	BrandId string `db:"brand_id"`
	Name    string `db:"model_name"`
	Price   int64  `db:"model_base_price"`
}

func TestSelect(t *testing.T) {
	db, err := sqlx.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = db.Close()
		if err != nil {
			panic(err)
		}
	}()
	models := []Model{}
	err = db.Select(&models, "Select * from Models")
	if err != nil {
		panic(err)
	}
	fmt.Println(models)
}
