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

	Encode() IScyllaTable
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

	values := table.GetValues()

	columnsString := ""
	valuesString := ""
	for column, value := range values {
		columnsString += column + ","
		valuesString += value + ","
	}

	columnsString = columnsString[:len(columnsString)-1]
	valuesString = valuesString[:len(valuesString)-1]

	return "INSERT INTO " + tableName + " (" + columnsString + ") VALUES (" + valuesString + ")"

}
