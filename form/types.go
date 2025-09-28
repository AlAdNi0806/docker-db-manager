package form

// We could use an interface, but for later

type SelectPrompt struct {
	Question    string
	Description string
	Options     []SelectOption
}

type SelectOption struct {
	Label string
	Value string
}

type SwitchPrompt struct {
	Question     string
	Description  string
	Options      [2]string
	DefaultValue bool
}

type InputPrompt struct {
	Question    string
	Description string
	Placeholder string
}
