package formflow

import (
	"docker-db-management/databases"
	"docker-db-management/form"
	"docker-db-management/types"
	"fmt"

	"github.com/fatih/color"
)

var (
	latestVersionEntity = types.LatestVersionEntity{
		Form: types.FormValues[bool]{
			Question:    "Pull the latest image?",
			Description: "If you don't have an image, it will still be pulled.",
			Choice:      false,
		},
	}

	passwordEntity = types.StringEntity{
		Form: types.FormValues[string]{
			Question:    "Set a root password for the database?",
			Description: "If you don't set a password, it will be set to 12345678.",
			Choice:      "",
		},
	}

	databaseNameEntity = types.StringEntity{
		Form: types.FormValues[string]{
			Question:    "Create a database?",
			Description: "You can leave this blank if you don't want to create a database.",
			Choice:      "",
		},
	}
)

func Create(dbHandler databases.DBHandler) {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()

	latestVersionEntity.Form.Choice, _ = form.NewSwitch(form.SwitchPrompt{
		Question:     latestVersionEntity.Form.Question,
		Description:  latestVersionEntity.Form.Description,
		Options:      [2]string{"Yes", "No"},
		DefaultValue: false,
	})

	symbol := red("✗ ")
	if latestVersionEntity.Form.Choice {
		symbol = green("✓ ")
	}
	fmt.Println(symbol, latestVersionEntity.Form.Question)

	passwordEntity.Form.Choice, _ = form.NewInput(form.InputPrompt{
		Question:    passwordEntity.Form.Question,
		Description: passwordEntity.Form.Description,
		Placeholder: "12345678",
	})

	if passwordEntity.Form.Choice == "" {
		passwordEntity.Form.Choice = "12345678"
	}
	fmt.Println(green("✓ "), passwordEntity.Form.Question, blue(passwordEntity.Form.Choice))

	databaseNameEntity.Form.Choice, _ = form.NewInput(form.InputPrompt{
		Question:    databaseNameEntity.Form.Question,
		Description: databaseNameEntity.Form.Description,
	})

	fmt.Println(green("✓ "), databaseNameEntity.Form.Question, blue(databaseNameEntity.Form.Choice))

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
