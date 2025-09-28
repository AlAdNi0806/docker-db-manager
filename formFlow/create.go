package formflow

import (
	"docker-db-management/databases"
	"docker-db-management/types"
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
)

var (
	latestVersionEntity = types.LatestVersion{
		Form: types.FormValues[bool]{
			Title:       "Pull the latest image?",
			Description: "If you don't have an image, it will still be pulled.",
			Choice:      false,
		},
	}

	passwordEntity = types.StringEntity{
		Form: types.FormValues[string]{
			Title:       "Set a root password for the database?",
			Description: "If you don't set a password, it will be set to 12345678.",
			Choice:      "",
		},
	}

	databaseNameEntity = types.StringEntity{
		Form: types.FormValues[string]{
			Title:       "Create a database?",
			Description: "You can leave this blank if you don't want to create a database.",
			Choice:      "",
		},
	}

	// databasePortEntity = types.IntEntity{
	// 	Form: types.FormValues[int]{
	// 		Title:       "Set a database port?",
	// 		Description: "If you don't set a database port, it will be set to 3306.",
	// 		Choice:      3306,
	// 	},
	// }

)

func Create(dbHandler databases.DBHandler) {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(latestVersionEntity.Form.Title).
				Description(latestVersionEntity.Form.Description).
				Affirmative("Yes").
				Negative("No.").
				Value(&latestVersionEntity.Form.Choice),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	symbol := red("✗ ")
	if latestVersionEntity.Form.Choice {
		symbol = green("✓ ")
	}
	fmt.Println(symbol, latestVersionEntity.Form.Title)

	form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title(passwordEntity.Form.Title).
				Description(passwordEntity.Form.Description).
				Value(&passwordEntity.Form.Choice),
		),
	)

	err = form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if passwordEntity.Form.Choice == "" {
		passwordEntity.Form.Choice = "12345678"
	}
	fmt.Println(green("✓ "), passwordEntity.Form.Title, blue(passwordEntity.Form.Choice))

	form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title(databaseNameEntity.Form.Title).
				Description(databaseNameEntity.Form.Description).
				Value(&databaseNameEntity.Form.Choice),
		),
	)

	err = form.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(green("✓ "), databaseNameEntity.Form.Title, blue(databaseNameEntity.Form.Choice))

	dbHandler.SetConfig(types.Config{
		LatestImage:  latestVersionEntity.Form.Choice,
		Password:     passwordEntity.Form.Choice,
		DatabaseName: databaseNameEntity.Form.Choice,
	})
	dbHandler.Create()
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
