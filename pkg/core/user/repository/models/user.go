package models

type User struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Alias            string `json:"alias"`
}

// TableName sets the insert table name for this struct type
func (User) TableName() string {
	return "user"
}
