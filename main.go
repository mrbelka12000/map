package main

import (
	"fmt"
	"github.com/graincomg/graincom_logistik/db"
	"log"
)

func main() {
	conn, err := db.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println(conn)
}
