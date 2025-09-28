package types

type NameValue struct {
	Key   string
	Value string
}

type ActionSelection struct {
	Actions []NameValue
	Form    FormValues[string] // or StringForm
}

type DatabaseSelection struct {
	Databases []NameValue
	Form      FormValues[string]
}

type StringEntity struct {
	Entity []NameValue
	Form   FormValues[string]
}

type LatestVersion struct {
	Labels []NameValue
	Form   FormValues[bool]
}

type FormValues[T any] struct {
	Title       string
	Description string
	Choice      T
}

type StringForm = FormValues[string]
type BoolForm = FormValues[bool]
type IntForm = FormValues[int]
type StringsForm = FormValues[[]string]

type Config struct {
	LatestImage  bool
	Password     string
	DatabaseName string
}
