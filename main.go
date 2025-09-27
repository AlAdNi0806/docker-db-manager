package main

import (
	"fmt"
	"log"

	"docker-db-management/databases"
	formflow "docker-db-management/formFlow"
	"docker-db-management/types"

	"github.com/charmbracelet/huh"
)

var (
	actionsEntity = types.ActionSelection{
		Actions: []types.NameValue{
			{Key: "Create database", Value: "create"},
			{Key: "Remove database", Value: "remove"},
		},
		FormValues: types.FormValues{
			Title:  "What would you like to do?",
			Choice: "",
		},
	}

	databasesEntity = types.DatabaseSelection{
		Databases: []types.NameValue{
			{Key: "MySQL", Value: "mysql"},
			{Key: "MariaDB", Value: "mariadb"},
		},
		FormValues: types.FormValues{
			Title:  "Which database do you choose?",
			Choice: "",
		},
	}
)

func main() {
	var actionOptions []huh.Option[string]
	for _, a := range actionsEntity.Actions {
		actionOptions = append(actionOptions, huh.NewOption(a.Key, a.Value))
	}

	var dbOptions []huh.Option[string]
	for _, db := range databasesEntity.Databases {
		dbOptions = append(dbOptions, huh.NewOption(db.Key, db.Value))
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title(actionsEntity.FormValues.Title).
				Options(actionOptions...).
				Value(&actionsEntity.FormValues.Choice),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(actionsEntity.FormValues.Title, actionsEntity.FormValues.Choice)

	form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Which database do you choose?").
				Options(dbOptions...).
				Value(&databasesEntity.FormValues.Choice),
		),
	)

	err = form.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(databasesEntity.FormValues.Title, databasesEntity.FormValues.Choice)

	var dbHandler databases.DBHandler
	switch databasesEntity.FormValues.Choice {
	case "mysql":
		dbHandler = databases.MySQL{}
	case "mariadb":
		dbHandler = databases.MariaDB{}
	default:
		log.Fatal("Unsupported database type")
	}

	switch actionsEntity.FormValues.Choice {
	case "create":
		formflow.Create(&dbHandler)
	case "remove":
		formflow.Remove(&dbHandler)
	default:
		log.Fatal("Unsupported action")
	}

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Successfully executed %s on %s database named %s\n", action, database, dbName)
}
