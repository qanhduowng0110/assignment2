// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/earthquake"
	"entdemo/ent/featuretype"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// FeatureType is the model entity for the FeatureType schema.
type FeatureType struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// EarthquakeID holds the value of the "earthquake_id" field.
	EarthquakeID int32 `json:"earthquake_id,omitempty"`
	// FeatureProductType holds the value of the "feature_product_type" field.
	FeatureProductType string `json:"feature_product_type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FeatureTypeQuery when eager-loading is set.
	Edges        FeatureTypeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FeatureTypeEdges holds the relations/edges for other nodes in the graph.
type FeatureTypeEdges struct {
	// Earthquake holds the value of the earthquake edge.
	Earthquake *Earthquake `json:"earthquake,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EarthquakeOrErr returns the Earthquake value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FeatureTypeEdges) EarthquakeOrErr() (*Earthquake, error) {
	if e.Earthquake != nil {
		return e.Earthquake, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: earthquake.Label}
	}
	return nil, &NotLoadedError{edge: "earthquake"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FeatureType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case featuretype.FieldID, featuretype.FieldEarthquakeID:
			values[i] = new(sql.NullInt64)
		case featuretype.FieldFeatureProductType:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FeatureType fields.
func (ft *FeatureType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case featuretype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ft.ID = int32(value.Int64)
		case featuretype.FieldEarthquakeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field earthquake_id", values[i])
			} else if value.Valid {
				ft.EarthquakeID = int32(value.Int64)
			}
		case featuretype.FieldFeatureProductType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field feature_product_type", values[i])
			} else if value.Valid {
				ft.FeatureProductType = value.String
			}
		default:
			ft.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FeatureType.
// This includes values selected through modifiers, order, etc.
func (ft *FeatureType) Value(name string) (ent.Value, error) {
	return ft.selectValues.Get(name)
}

// QueryEarthquake queries the "earthquake" edge of the FeatureType entity.
func (ft *FeatureType) QueryEarthquake() *EarthquakeQuery {
	return NewFeatureTypeClient(ft.config).QueryEarthquake(ft)
}

// Update returns a builder for updating this FeatureType.
// Note that you need to call FeatureType.Unwrap() before calling this method if this FeatureType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ft *FeatureType) Update() *FeatureTypeUpdateOne {
	return NewFeatureTypeClient(ft.config).UpdateOne(ft)
}

// Unwrap unwraps the FeatureType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ft *FeatureType) Unwrap() *FeatureType {
	_tx, ok := ft.config.driver.(*txDriver)
	if !ok {
		panic("ent: FeatureType is not a transactional entity")
	}
	ft.config.driver = _tx.drv
	return ft
}

// String implements the fmt.Stringer.
func (ft *FeatureType) String() string {
	var builder strings.Builder
	builder.WriteString("FeatureType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ft.ID))
	builder.WriteString("earthquake_id=")
	builder.WriteString(fmt.Sprintf("%v", ft.EarthquakeID))
	builder.WriteString(", ")
	builder.WriteString("feature_product_type=")
	builder.WriteString(ft.FeatureProductType)
	builder.WriteByte(')')
	return builder.String()
}

// FeatureTypes is a parsable slice of FeatureType.
type FeatureTypes []*FeatureType