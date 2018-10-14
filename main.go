package main

import (
	"go-db-example/util"
	"log"
)

func main() {
	specificQuery()
	customQuery()
}

func customQuery() {
	log.Println("Start to do custom query with stock id = 2330...")
	rows, err := myDB.Query("SELECT stock_name, stock_type FROM tbl_stock_list WHERE stock_no = ?", "2330")

	var mName string
	var mType string
	for rows.Next() {
		err := rows.Scan(&mName, &mType)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Stock Name = ", mName, ", Stock Type = ", mType)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func specificQuery() {
	log.Println("Start to do specific query with stock id = 2330...")
	rows, err := myDB.GetStockName("2330")

	var mName string
	var mType string
	for rows.Next() {
		err := rows.Scan(&mName, &mType)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Stock Name = ", mName, ", Stock Type = ", mType)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
