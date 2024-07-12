package entities

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"-"`
}
