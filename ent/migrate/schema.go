// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// APIRequestLogsColumns holds the columns for the "API_request_logs" table.
	APIRequestLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "request_datetime", Type: field.TypeTime},
		{Name: "request_parameters", Type: field.TypeJSON},
		{Name: "request_body", Type: field.TypeJSON},
		{Name: "request_headers", Type: field.TypeJSON},
		{Name: "request_metadata", Type: field.TypeJSON},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// APIRequestLogsTable holds the schema information for the "API_request_logs" table.
	APIRequestLogsTable = &schema.Table{
		Name:       "API_request_logs",
		Columns:    APIRequestLogsColumns,
		PrimaryKey: []*schema.Column{APIRequestLogsColumns[0]},
	}
	// AssociatedEventsColumns holds the columns for the "Associated_events" table.
	AssociatedEventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "earthquake_id", Type: field.TypeInt32, Nullable: true},
		{Name: "associate_id", Type: field.TypeInt32, Nullable: true},
	}
	// AssociatedEventsTable holds the schema information for the "Associated_events" table.
	AssociatedEventsTable = &schema.Table{
		Name:       "Associated_events",
		Columns:    AssociatedEventsColumns,
		PrimaryKey: []*schema.Column{AssociatedEventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Associated_events_Earthquakes_main_events",
				Columns:    []*schema.Column{AssociatedEventsColumns[1]},
				RefColumns: []*schema.Column{EarthquakesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "Associated_events_Earthquakes_associated_events",
				Columns:    []*schema.Column{AssociatedEventsColumns[2]},
				RefColumns: []*schema.Column{EarthquakesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EarthquakesColumns holds the columns for the "Earthquakes" table.
	EarthquakesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "feature_id", Type: field.TypeString},
		{Name: "magnitude", Type: field.TypeFloat64},
		{Name: "occur_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "url", Type: field.TypeString},
		{Name: "detail_url", Type: field.TypeString},
		{Name: "status", Type: field.TypeString},
		{Name: "tsunami", Type: field.TypeInt32},
		{Name: "sig", Type: field.TypeInt32},
		{Name: "net", Type: field.TypeString},
		{Name: "code", Type: field.TypeString},
		{Name: "nst", Type: field.TypeInt32},
		{Name: "dmin", Type: field.TypeFloat64},
		{Name: "rms", Type: field.TypeFloat64},
		{Name: "gap", Type: field.TypeFloat64},
		{Name: "mag_type", Type: field.TypeString},
		{Name: "earthquake_type", Type: field.TypeString},
	}
	// EarthquakesTable holds the schema information for the "Earthquakes" table.
	EarthquakesTable = &schema.Table{
		Name:       "Earthquakes",
		Columns:    EarthquakesColumns,
		PrimaryKey: []*schema.Column{EarthquakesColumns[0]},
	}
	// EarthquakeDbsColumns holds the columns for the "earthquake_dbs" table.
	EarthquakeDbsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// EarthquakeDbsTable holds the schema information for the "earthquake_dbs" table.
	EarthquakeDbsTable = &schema.Table{
		Name:       "earthquake_dbs",
		Columns:    EarthquakeDbsColumns,
		PrimaryKey: []*schema.Column{EarthquakeDbsColumns[0]},
	}
	// EventTypesColumns holds the columns for the "Event_types" table.
	EventTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "types", Type: field.TypeString, Nullable: true},
		{Name: "earthquake_id", Type: field.TypeInt32, Nullable: true},
	}
	// EventTypesTable holds the schema information for the "Event_types" table.
	EventTypesTable = &schema.Table{
		Name:       "Event_types",
		Columns:    EventTypesColumns,
		PrimaryKey: []*schema.Column{EventTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Event_types_Earthquakes_event_types",
				Columns:    []*schema.Column{EventTypesColumns[2]},
				RefColumns: []*schema.Column{EarthquakesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FeatureTypesColumns holds the columns for the "Feature_types" table.
	FeatureTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "feature_product_type", Type: field.TypeString},
		{Name: "earthquake_id", Type: field.TypeInt32, Nullable: true},
	}
	// FeatureTypesTable holds the schema information for the "Feature_types" table.
	FeatureTypesTable = &schema.Table{
		Name:       "Feature_types",
		Columns:    FeatureTypesColumns,
		PrimaryKey: []*schema.Column{FeatureTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Feature_types_Earthquakes_feature_types",
				Columns:    []*schema.Column{FeatureTypesColumns[2]},
				RefColumns: []*schema.Column{EarthquakesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FeltReportsColumns holds the columns for the "Felt_reports" table.
	FeltReportsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "felt", Type: field.TypeInt32, Nullable: true},
		{Name: "cdi", Type: field.TypeFloat64, Nullable: true},
		{Name: "mmi", Type: field.TypeFloat64, Nullable: true},
		{Name: "alert", Type: field.TypeString, Nullable: true},
		{Name: "earthquake_id", Type: field.TypeInt32, Nullable: true},
	}
	// FeltReportsTable holds the schema information for the "Felt_reports" table.
	FeltReportsTable = &schema.Table{
		Name:       "Felt_reports",
		Columns:    FeltReportsColumns,
		PrimaryKey: []*schema.Column{FeltReportsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Felt_reports_Earthquakes_felt_reports",
				Columns:    []*schema.Column{FeltReportsColumns[5]},
				RefColumns: []*schema.Column{EarthquakesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GeometryColumns holds the columns for the "Geometry" table.
	GeometryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt32, Increment: true},
		{Name: "longitude", Type: field.TypeFloat64},
		{Name: "latitude", Type: field.TypeFloat64},
		{Name: "depth", Type: field.TypeFloat64},
		{Name: "place", Type: field.TypeString},
		{Name: "earthquake_id", Type: field.TypeInt32, Nullable: true},
	}
	// GeometryTable holds the schema information for the "Geometry" table.
	GeometryTable = &schema.Table{
		Name:       "Geometry",
		Columns:    GeometryColumns,
		PrimaryKey: []*schema.Column{GeometryColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "Geometry_Earthquakes_geometries",
				Columns:    []*schema.Column{GeometryColumns[5]},
				RefColumns: []*schema.Column{EarthquakesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SchemaMigrationsColumns holds the columns for the "schema_migrations" table.
	SchemaMigrationsColumns = []*schema.Column{
		{Name: "version", Type: field.TypeInt, Increment: true},
		{Name: "dirty", Type: field.TypeBool},
	}
	// SchemaMigrationsTable holds the schema information for the "schema_migrations" table.
	SchemaMigrationsTable = &schema.Table{
		Name:       "schema_migrations",
		Columns:    SchemaMigrationsColumns,
		PrimaryKey: []*schema.Column{SchemaMigrationsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		APIRequestLogsTable,
		AssociatedEventsTable,
		EarthquakesTable,
		EarthquakeDbsTable,
		EventTypesTable,
		FeatureTypesTable,
		FeltReportsTable,
		GeometryTable,
		SchemaMigrationsTable,
	}
)

func init() {
	APIRequestLogsTable.Annotation = &entsql.Annotation{
		Table: "API_request_logs",
	}
	AssociatedEventsTable.ForeignKeys[0].RefTable = EarthquakesTable
	AssociatedEventsTable.ForeignKeys[1].RefTable = EarthquakesTable
	AssociatedEventsTable.Annotation = &entsql.Annotation{
		Table: "Associated_events",
	}
	EarthquakesTable.Annotation = &entsql.Annotation{
		Table: "Earthquakes",
	}
	EventTypesTable.ForeignKeys[0].RefTable = EarthquakesTable
	EventTypesTable.Annotation = &entsql.Annotation{
		Table: "Event_types",
	}
	FeatureTypesTable.ForeignKeys[0].RefTable = EarthquakesTable
	FeatureTypesTable.Annotation = &entsql.Annotation{
		Table: "Feature_types",
	}
	FeltReportsTable.ForeignKeys[0].RefTable = EarthquakesTable
	FeltReportsTable.Annotation = &entsql.Annotation{
		Table: "Felt_reports",
	}
	GeometryTable.ForeignKeys[0].RefTable = EarthquakesTable
	GeometryTable.Annotation = &entsql.Annotation{
		Table: "Geometry",
	}
}