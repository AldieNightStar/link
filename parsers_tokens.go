package link

type TokenVariable struct {
	Value string
}

type tokEnd struct {
	Value string
}

type TokenCommand struct {
	Name string
	Args []interface{}
}

type TokenCodeBlock struct {
	Commands []*TokenCommand
}
