package main

import (
	"fmt"
	"log"

	"docker-db-management/databases"
	formflow "docker-db-management/formFlow"
	"docker-db-management/types"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
)

var (
	actionsEntity = types.ActionSelection{
		Actions: []types.NameValue{
			{Key: "Create database", Value: "create"},
			{Key: "Remove database", Value: "remove"},
		},
		Form: types.FormValues[string]{
			Title:  "What would you like to do?",
			Choice: "",
		},
	}

	databasesEntity = types.DatabaseSelection{
		Databases: []types.NameValue{
			{Key: "MySQL", Value: "mysql"},
			{Key: "MariaDB", Value: "mariadb"},
		},
		Form: types.FormValues[string]{
			Title:  "Which database do you choose?",
			Choice: "",
		},
	}
)

func main_() {
	// mysql := databases.DBHandler(&databases.MySQL{})
	// mysql.Create()

	// os.Exit(1)

	fmt.Println("")
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

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
				Title(actionsEntity.Form.Title).
				Options(actionOptions...).
				Value(&actionsEntity.Form.Choice),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(green("✓ "), actionsEntity.Form.Title, blue(actionsEntity.Form.Choice))

	form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Which database do you choose?").
				Options(dbOptions...).
				Value(&databasesEntity.Form.Choice),
		),
	)

	err = form.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(green("✓ "), databasesEntity.Form.Title, blue(databasesEntity.Form.Choice))

	var dbHandler databases.DBHandler
	switch databasesEntity.Form.Choice {
	case "mysql":
		dbHandler = &databases.MySQL{}
	case "mariadb":
		dbHandler = &databases.MariaDB{}
	default:
		log.Fatal("Unsupported database type")
	}

	switch actionsEntity.Form.Choice {
	case "create":
		formflow.Create(dbHandler)
	case "remove":
		formflow.Remove(dbHandler)
	default:
		log.Fatal("Unsupported action")
	}

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Successfully executed %s on %s database named %s\n", action, database, dbName)
}
