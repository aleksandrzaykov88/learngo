package models

type Image struct {
	ID        int         `json:"id" db:"id"`
	Link      string      `json:"link" db:"link"`
	Alt       string      `json:"alt" db:"alt"`
	CreatedAt string      `json:"created_at" db:"created_at"`
	UpdatedAt string      `json:"updated_at" db:"updated_at"`
	Tags      interface{} `json:"tags" db:"tags"`
}
