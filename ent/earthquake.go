// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/earthquake"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Earthquake is the model entity for the Earthquake schema.
type Earthquake struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// FeatureID holds the value of the "feature_id" field.
	FeatureID string `json:"feature_id,omitempty"`
	// Magnitude holds the value of the "magnitude" field.
	Magnitude float64 `json:"magnitude,omitempty"`
	// OccurTime holds the value of the "occur_time" field.
	OccurTime time.Time `json:"occur_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// DetailURL holds the value of the "detail_url" field.
	DetailURL string `json:"detail_url,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Tsunami holds the value of the "tsunami" field.
	Tsunami int32 `json:"tsunami,omitempty"`
	// Sig holds the value of the "sig" field.
	Sig int32 `json:"sig,omitempty"`
	// Net holds the value of the "net" field.
	Net string `json:"net,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Nst holds the value of the "nst" field.
	Nst int32 `json:"nst,omitempty"`
	// Dmin holds the value of the "dmin" field.
	Dmin float64 `json:"dmin,omitempty"`
	// Rms holds the value of the "rms" field.
	Rms float64 `json:"rms,omitempty"`
	// Gap holds the value of the "gap" field.
	Gap float64 `json:"gap,omitempty"`
	// MagType holds the value of the "mag_type" field.
	MagType string `json:"mag_type,omitempty"`
	// EarthquakeType holds the value of the "earthquake_type" field.
	EarthquakeType string `json:"earthquake_type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EarthquakeQuery when eager-loading is set.
	Edges        EarthquakeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// EarthquakeEdges holds the relations/edges for other nodes in the graph.
type EarthquakeEdges struct {
	// MainEvents holds the value of the main_events edge.
	MainEvents []*AssociatedEvent `json:"main_events,omitempty"`
	// AssociatedEvents holds the value of the associated_events edge.
	AssociatedEvents []*AssociatedEvent `json:"associated_events,omitempty"`
	// EventTypes holds the value of the event_types edge.
	EventTypes []*EventType `json:"event_types,omitempty"`
	// FeatureTypes holds the value of the feature_types edge.
	FeatureTypes []*FeatureType `json:"feature_types,omitempty"`
	// FeltReports holds the value of the felt_reports edge.
	FeltReports []*FeltReport `json:"felt_reports,omitempty"`
	// Geometries holds the value of the geometries edge.
	Geometries []*Geometry `json:"geometries,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// MainEventsOrErr returns the MainEvents value or an error if the edge
// was not loaded in eager-loading.
func (e EarthquakeEdges) MainEventsOrErr() ([]*AssociatedEvent, error) {
	if e.loadedTypes[0] {
		return e.MainEvents, nil
	}
	return nil, &NotLoadedError{edge: "main_events"}
}

// AssociatedEventsOrErr returns the AssociatedEvents value or an error if the edge
// was not loaded in eager-loading.
func (e EarthquakeEdges) AssociatedEventsOrErr() ([]*AssociatedEvent, error) {
	if e.loadedTypes[1] {
		return e.AssociatedEvents, nil
	}
	return nil, &NotLoadedError{edge: "associated_events"}
}

// EventTypesOrErr returns the EventTypes value or an error if the edge
// was not loaded in eager-loading.
func (e EarthquakeEdges) EventTypesOrErr() ([]*EventType, error) {
	if e.loadedTypes[2] {
		return e.EventTypes, nil
	}
	return nil, &NotLoadedError{edge: "event_types"}
}

// FeatureTypesOrErr returns the FeatureTypes value or an error if the edge
// was not loaded in eager-loading.
func (e EarthquakeEdges) FeatureTypesOrErr() ([]*FeatureType, error) {
	if e.loadedTypes[3] {
		return e.FeatureTypes, nil
	}
	return nil, &NotLoadedError{edge: "feature_types"}
}

// FeltReportsOrErr returns the FeltReports value or an error if the edge
// was not loaded in eager-loading.
func (e EarthquakeEdges) FeltReportsOrErr() ([]*FeltReport, error) {
	if e.loadedTypes[4] {
		return e.FeltReports, nil
	}
	return nil, &NotLoadedError{edge: "felt_reports"}
}

// GeometriesOrErr returns the Geometries value or an error if the edge
// was not loaded in eager-loading.
func (e EarthquakeEdges) GeometriesOrErr() ([]*Geometry, error) {
	if e.loadedTypes[5] {
		return e.Geometries, nil
	}
	return nil, &NotLoadedError{edge: "geometries"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Earthquake) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case earthquake.FieldMagnitude, earthquake.FieldDmin, earthquake.FieldRms, earthquake.FieldGap:
			values[i] = new(sql.NullFloat64)
		case earthquake.FieldID, earthquake.FieldTsunami, earthquake.FieldSig, earthquake.FieldNst:
			values[i] = new(sql.NullInt64)
		case earthquake.FieldFeatureID, earthquake.FieldURL, earthquake.FieldDetailURL, earthquake.FieldStatus, earthquake.FieldNet, earthquake.FieldCode, earthquake.FieldMagType, earthquake.FieldEarthquakeType:
			values[i] = new(sql.NullString)
		case earthquake.FieldOccurTime, earthquake.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Earthquake fields.
func (e *Earthquake) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case earthquake.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int32(value.Int64)
		case earthquake.FieldFeatureID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field feature_id", values[i])
			} else if value.Valid {
				e.FeatureID = value.String
			}
		case earthquake.FieldMagnitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field magnitude", values[i])
			} else if value.Valid {
				e.Magnitude = value.Float64
			}
		case earthquake.FieldOccurTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field occur_time", values[i])
			} else if value.Valid {
				e.OccurTime = value.Time
			}
		case earthquake.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				e.UpdateTime = value.Time
			}
		case earthquake.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				e.URL = value.String
			}
		case earthquake.FieldDetailURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detail_url", values[i])
			} else if value.Valid {
				e.DetailURL = value.String
			}
		case earthquake.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				e.Status = value.String
			}
		case earthquake.FieldTsunami:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tsunami", values[i])
			} else if value.Valid {
				e.Tsunami = int32(value.Int64)
			}
		case earthquake.FieldSig:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sig", values[i])
			} else if value.Valid {
				e.Sig = int32(value.Int64)
			}
		case earthquake.FieldNet:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field net", values[i])
			} else if value.Valid {
				e.Net = value.String
			}
		case earthquake.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				e.Code = value.String
			}
		case earthquake.FieldNst:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nst", values[i])
			} else if value.Valid {
				e.Nst = int32(value.Int64)
			}
		case earthquake.FieldDmin:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field dmin", values[i])
			} else if value.Valid {
				e.Dmin = value.Float64
			}
		case earthquake.FieldRms:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field rms", values[i])
			} else if value.Valid {
				e.Rms = value.Float64
			}
		case earthquake.FieldGap:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field gap", values[i])
			} else if value.Valid {
				e.Gap = value.Float64
			}
		case earthquake.FieldMagType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mag_type", values[i])
			} else if value.Valid {
				e.MagType = value.String
			}
		case earthquake.FieldEarthquakeType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field earthquake_type", values[i])
			} else if value.Valid {
				e.EarthquakeType = value.String
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Earthquake.
// This includes values selected through modifiers, order, etc.
func (e *Earthquake) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QueryMainEvents queries the "main_events" edge of the Earthquake entity.
func (e *Earthquake) QueryMainEvents() *AssociatedEventQuery {
	return NewEarthquakeClient(e.config).QueryMainEvents(e)
}

// QueryAssociatedEvents queries the "associated_events" edge of the Earthquake entity.
func (e *Earthquake) QueryAssociatedEvents() *AssociatedEventQuery {
	return NewEarthquakeClient(e.config).QueryAssociatedEvents(e)
}

// QueryEventTypes queries the "event_types" edge of the Earthquake entity.
func (e *Earthquake) QueryEventTypes() *EventTypeQuery {
	return NewEarthquakeClient(e.config).QueryEventTypes(e)
}

// QueryFeatureTypes queries the "feature_types" edge of the Earthquake entity.
func (e *Earthquake) QueryFeatureTypes() *FeatureTypeQuery {
	return NewEarthquakeClient(e.config).QueryFeatureTypes(e)
}

// QueryFeltReports queries the "felt_reports" edge of the Earthquake entity.
func (e *Earthquake) QueryFeltReports() *FeltReportQuery {
	return NewEarthquakeClient(e.config).QueryFeltReports(e)
}

// QueryGeometries queries the "geometries" edge of the Earthquake entity.
func (e *Earthquake) QueryGeometries() *GeometryQuery {
	return NewEarthquakeClient(e.config).QueryGeometries(e)
}

// Update returns a builder for updating this Earthquake.
// Note that you need to call Earthquake.Unwrap() before calling this method if this Earthquake
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Earthquake) Update() *EarthquakeUpdateOne {
	return NewEarthquakeClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Earthquake entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Earthquake) Unwrap() *Earthquake {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Earthquake is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Earthquake) String() string {
	var builder strings.Builder
	builder.WriteString("Earthquake(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("feature_id=")
	builder.WriteString(e.FeatureID)
	builder.WriteString(", ")
	builder.WriteString("magnitude=")
	builder.WriteString(fmt.Sprintf("%v", e.Magnitude))
	builder.WriteString(", ")
	builder.WriteString("occur_time=")
	builder.WriteString(e.OccurTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(e.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(e.URL)
	builder.WriteString(", ")
	builder.WriteString("detail_url=")
	builder.WriteString(e.DetailURL)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(e.Status)
	builder.WriteString(", ")
	builder.WriteString("tsunami=")
	builder.WriteString(fmt.Sprintf("%v", e.Tsunami))
	builder.WriteString(", ")
	builder.WriteString("sig=")
	builder.WriteString(fmt.Sprintf("%v", e.Sig))
	builder.WriteString(", ")
	builder.WriteString("net=")
	builder.WriteString(e.Net)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(e.Code)
	builder.WriteString(", ")
	builder.WriteString("nst=")
	builder.WriteString(fmt.Sprintf("%v", e.Nst))
	builder.WriteString(", ")
	builder.WriteString("dmin=")
	builder.WriteString(fmt.Sprintf("%v", e.Dmin))
	builder.WriteString(", ")
	builder.WriteString("rms=")
	builder.WriteString(fmt.Sprintf("%v", e.Rms))
	builder.WriteString(", ")
	builder.WriteString("gap=")
	builder.WriteString(fmt.Sprintf("%v", e.Gap))
	builder.WriteString(", ")
	builder.WriteString("mag_type=")
	builder.WriteString(e.MagType)
	builder.WriteString(", ")
	builder.WriteString("earthquake_type=")
	builder.WriteString(e.EarthquakeType)
	builder.WriteByte(')')
	return builder.String()
}

// Earthquakes is a parsable slice of Earthquake.
type Earthquakes []*Earthquake