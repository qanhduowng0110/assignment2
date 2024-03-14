// Code generated by ent, DO NOT EDIT.

package earthquake_db

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the earthquake_db type in the database.
	Label = "earthquake_db"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the earthquake_db in the database.
	Table = "earthquake_dbs"
)

// Columns holds all SQL columns for earthquake_db fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Earthquake_db queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}