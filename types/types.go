package types

import "docker-db-management/form"

type LabelValue struct {
	Label string
	Value string
}

type ActionEntity struct {
	Actions []form.SelectOption
	Form    FormValues[form.SelectOption]
}

type DatabaseEntity struct {
	Databases []form.SelectOption
	Form      FormValues[form.SelectOption]
}

type StringEntity struct {
	Entity []LabelValue
	Form   FormValues[string]
}

type LatestVersionEntity struct {
	Labels []LabelValue
	Form   FormValues[bool]
}

type FormValues[T any] struct {
	Question    string
	Description string
	Choice      T
}

type SelectOptionForm = FormValues[form.SelectOption]
type BoolForm = FormValues[bool]
type StringsForm = FormValues[[]string]
type StringForm = FormValues[string]
type IntForm = FormValues[int]

type Config struct {
	LatestImage  bool
	Password     string
	DatabaseName string
}
