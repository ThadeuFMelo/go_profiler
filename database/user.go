package database

type User struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Address         string `json:"address"`
	PictureLocation string `json:"picture_location"`
}
