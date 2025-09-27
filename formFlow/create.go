package formflow

import (
	"docker-db-management/databases"
	"log"

	"github.com/charmbracelet/huh"
)

var (
	latestVersionEntity = ActionSelection{
		FormValues: FormValues{
			Title:  "What would you like to do?",
			Choice: false,
		},
	}

	databasesEntity = DatabaseSelection{
		Databases: []NameValue{
			{Key: "MySQL", Value: "mysql"},
			{Key: "MariaDB", Value: "mariadb"},
		},
		FormValues: FormValues{
			Title:  "Which database do you choose?",
			Choice: "",
		},
	}
)

func Create(*databases.DBHandler) {
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

}

// latest version yes/no
// root password (if blank than defaul 12345678)
// database name (if blank no database will be created)

// later
// database user (if blank than defaul root)
// database password (if blank than defaul 12345678)
// database port (if blank than defaul 3306)
// database host (if blank than defaul localhost)

// create database
// create user
// create password
// create database
