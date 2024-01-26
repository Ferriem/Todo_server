package models

type Auth struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
