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
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func main() {

	// Connect to database
	db, err := sql.Open("postgres", "dbname=airplanes user=postgres password=postgres")
	dieIf(err)
	defer db.Close()

	err = db.Ping()
	dieIf(err)

	// Insert new pilot
	p := &models.Pilot{Name: randomdata.LastName()}
	err = p.Insert(context.Background(), db, boil.Infer())
	dieIf(err)
	fmt.Printf("Pilot ID: %v Name: %v\n", p.ID, p.Name)

	// Get one pilot
	onePilot, err := models.Pilots().One(context.Background(), db)
	dieIf(err)
	fmt.Printf("One pilot: %+v\n", onePilot)

	// Found pilot on ID
	f, err := models.FindPilot(context.Background(), db, onePilot.ID)
	dieIf(err)
	fmt.Printf("Found pilot: %+v\n", f)

	// Update pilot
	f.Name = randomdata.LastName()
	rows, err := f.Update(context.Background(), db, boil.Whitelist(models.PilotColumns.Name))
	dieIf(err)
	fmt.Printf("Rows: %+v\n", rows)

	// SELECT pilot
	fa, err := models.Pilots(qm.Where("name = ?", f.Name)).One(context.Background(), db)
	dieIf(err)
	fmt.Printf("Found where: %+v\n", fa)

	// Check if pilot exists
	exists, err := models.PilotExists(context.Background(), db, fa.ID)
	dieIf(err)
	fmt.Printf("Pilot exists: %+v\n", exists)

	// Count number of records in table
	count, err := models.Pilots().Count(context.Background(), db)
	dieIf(err)
	fmt.Printf("Number of pilots: %+v\n", count)

	// Get first pilot in table, and show languages and jets for this pilot
	op, err := models.Pilots(qm.Where("id = ?", 1)).One(context.Background(), db)
	dieIf(err)

	languages, err := op.Languages().All(context.Background(), db)
	dieIf(err)
	fmt.Printf("Pilot %v %v knows following lanugages:\n", op.Firstname, op.Name)
	for _, oneLang := range languages {
		fmt.Println(oneLang.Language)
	}

	jets, err := op.Jets().All(context.Background(), db)
	fmt.Printf("Jets for pilot: %v\n", op.Name)
	for _, oneJet := range jets {
		fmt.Println(oneJet.Name)
	}

}

func dieIf(err error) {
	if err != nil {
		panic(err)
	}

}
