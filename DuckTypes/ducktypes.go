package ducktypes

type ID string

type Message struct {
	ID
	Contents    string
	Respondants []User
}

type User struct {
	ID
}
