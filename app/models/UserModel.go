package models

type Users struct {
	Model
	User         string `json:"user"`
	Password     string `json:"password"`
	Nickname     string `json:"nickname"`
	Sex          string `json:"sex"`
	Age          int    `json:"age"`
	Region       string `json:"region"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Introduction string `json:"introduction"`
	Status       string `json:"status"`
}
