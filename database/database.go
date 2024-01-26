package database

import (
	"log"
	"reflect"
)

type IScyllaTable interface {
	// GetTableName returns the table name
	GetTableName() string
	// GetColumns returns the table columns
	GetColumns() []string
	// GetValues returns a map with the values for each column
	GetValues() map[string]string

	// BuildSelectQuery returns a string with the select query
	BuildSelectQuery(fields []string) string

	// BuildInsertQuery returns a string with the insert query
	BuildInsertQuery() string

	// Decode returns a struct of type ITable
	Decode() ITable
}

type ITable interface {
	// GetTableName returns the table name
	GetTableName() string
	// GetColumns returns the table columns
	GetColumns() []string
	// GetValues returns a map with the values for each column
	GetValues() map[string]string
}

func getColumns(s ITable) []string {
	val := reflect.ValueOf(s)
	typeOf := val.Type()
	var columns []string
	for i := 0; i < typeOf.NumField(); i++ {
		columns = append(columns, typeOf.Field(i).Tag.Get("json"))
	}
	return columns
}

func getValues(s ITable) map[string]string {
	val := reflect.ValueOf(s)
	typeOf := val.Type()
	values := map[string]string{}
	for i := 0; i < typeOf.NumField(); i++ {
		values[typeOf.Field(i).Tag.Get("json")] = val.Field(i).String()
	}
	return values
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func buildSelectQuery(table IScyllaTable, fields []string) string {
	tableName := table.GetTableName()

	if len(fields) == 0 {
		return "SELECT * FROM " + tableName
	}
	columns := table.GetColumns()
	// check if all fields are valid
	for _, field := range fields {
		if !contains(columns, field) {
			log.Default().Println("Invalid field: " + field)
			return ""
		}
	}

	fieldsString := ""
	for i := 0; i < len(fields); i++ {
		fieldsString += fields[i]
		if i < len(fields)-1 {
			fieldsString += ","
		}
	}

	return "SELECT " + fieldsString + " FROM " + tableName
}

func buildInsertQuery(table IScyllaTable) string {
	tableName := table.GetTableName()
	columns := table.GetColumns()
	query := "INSERT INTO " + tableName + " ("
	for i := 0; i < len(columns); i++ {
		query += columns[i]
		if i < len(columns)-1 {
			query += ","
		}
	}
	//use string wildcards to build the values part of the query later
	query += ") VALUES ("
	for i := 0; i < len(columns); i++ {
		query += "?"
		if i < len(columns)-1 {
			query += ","
		}
	}
	query += ")"
	return query
}
