package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"docker-db-management/databases"
	"docker-db-management/form"
	formflow "docker-db-management/formFlow"
	"docker-db-management/types"

	"github.com/fatih/color"
)

var (
	actionsEntity = types.ActionEntity{
		Actions: []form.SelectOption{
			{Label: "Remove database", Value: "remove"},
			{Label: "Create database", Value: "create"},
			{Label: "Manage db containers/image/volumes", Value: "list"},
		},
		Form: types.FormValues[form.SelectOption]{
			Question:    "What would you like to do?",
			Description: "wassup",
		},
	}

	databasesEntity = types.DatabaseEntity{
		Databases: []form.SelectOption{
			{Label: "MySQL", Value: "mysql"},
			{Label: "MariaDB", Value: "mariadb"},
		},
		Form: types.FormValues[form.SelectOption]{
			Question: "Which database do you choose?",
		},
	}
)

func main_() {
	// mysql := databases.DBHandler(&databases.MySQL{})
	// mysql.Create()

	// os.Exit(1)
	//
	form.PrintFullWidthBox("rami")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		form.Cleanup()
		os.Exit(1)
	}()

	// Hide cursor once at start
	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	// Switch TO alternate screen
	fmt.Print("\033[?1049h")
	defer fmt.Print("\033[?1049l")

	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	actionsEntity.Form.Choice, _ = form.NewSelect(form.SelectPrompt{
		Question:    actionsEntity.Form.Question,
		Description: actionsEntity.Form.Description,
		Options:     actionsEntity.Actions,
	})

	fmt.Println("")
	fmt.Println(green("✓ "), actionsEntity.Form.Question, blue(actionsEntity.Form.Choice.Label))

	databasesEntity.Form.Choice, _ = form.NewSelect(form.SelectPrompt{
		Question:    databasesEntity.Form.Question,
		Description: databasesEntity.Form.Description,
		Options:     databasesEntity.Databases,
	})

	fmt.Println(green("✓ "), databasesEntity.Form.Question, blue(databasesEntity.Form.Choice.Label))

	var dbHandler databases.DBHandler
	switch databasesEntity.Form.Choice.Value {
	case "mysql":
		dbHandler = &databases.MySQL{}
	case "mariadb":
		dbHandler = &databases.MariaDB{}
	default:
		log.Fatal("Unsupported database type")
	}

	switch actionsEntity.Form.Choice.Value {
	case "create":
		formflow.Create(dbHandler)
	case "remove":
		formflow.Remove(dbHandler)
	default:
		log.Fatal("Unsupported action")
	}

	// fmt.Printf("Successfully executed %s on %s database named %s\n", action, database, dbName)
}
