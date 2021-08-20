package models

type Joke struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Joke string `json:"joke"`
}
