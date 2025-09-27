package types

type NameValue struct {
	Key   string
	Value string
}

type ActionSelection struct {
	Actions    []NameValue
	FormValues FormValues
}

type DatabaseSelection struct {
	Databases  []NameValue
	FormValues FormValues
}

type LatestVersion struct {
	LatestVersion bool
}

type FormValues struct {
	Title  string
	Choice string
}

type FormValuesBool struct {
	Title  string
	Choice bool
}

type FormValuesInt struct {
	Title  string
	Choice int
}

type FormValuesStrings struct {
	Title   string
	Choices []string
}
