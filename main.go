package main

//go:generate sqlboiler --wipe psql

import (
	"airplanes/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/Pallinder/go-randomdata"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {

	db, err := sql.Open("postgres", "dbname=airplanes user=postgres password=postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	p := &models.Pilot{Name: randomdata.LastName()}
	err = p.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Pilot ID: %v Name: %v\n", p.ID, p.Name)

}
