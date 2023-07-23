package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)


var Conn *pgx.Conn

func DatabaseConnect() {
	// postgres://user:password@host:port/dbname
	databaseURL := "postgres://postgres:aisyiLL21@localhost:5432/db_personal_web"
	
	var err error
	Conn, err = pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database : %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Succesfully connected to database.")
}