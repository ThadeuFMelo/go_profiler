package database

import "reflect"

type User struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Address         string `json:"address"`
	PictureLocation string `json:"picture_location"`
}

func (User) GetTableName() string {
	return "users"
}

func (u *User) Encode() IScyllaTable {
	user := &ScyllaUser{
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Address:         u.Address,
		PictureLocation: u.PictureLocation,
	}
	return user
}

type ScyllaUser struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Address         string `json:"address"`
	PictureLocation string `json:"picture_location"`
}

func (ScyllaUser) GetTableName() string {
	return "users"
}

func (u ScyllaUser) GetColumns() []string {
	val := reflect.ValueOf(u)
	typeOf := val.Type()
	var columns []string
	for i := 0; i < typeOf.NumField(); i++ {
		columns = append(columns, typeOf.Field(i).Tag.Get("json"))
	}
	return columns
}

func (u ScyllaUser) GetValues() map[string]string {
	val := reflect.ValueOf(u)
	typeOf := val.Type()
	values := map[string]string{}
	for i := 0; i < typeOf.NumField(); i++ {
		values[typeOf.Field(i).Tag.Get("json")] = val.Field(i).String()
	}
	return values
}

func (u *ScyllaUser) BuildSelectQuery(fields []string) string {
	return buildSelectQuery(u, fields)
}

func (u *ScyllaUser) BuildInsertQuery() string {
	return buildInsertQuery(u)
}

func (u *ScyllaUser) Decode() ITable {
	user := &User{
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Address:         u.Address,
		PictureLocation: u.PictureLocation,
	}
	return user
}
