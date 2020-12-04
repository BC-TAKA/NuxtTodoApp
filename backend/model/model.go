package model

type Listup struct {
	ID   int
	Name string
	Todo string
}

type TODO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Todo string `json:"todo"`
}
