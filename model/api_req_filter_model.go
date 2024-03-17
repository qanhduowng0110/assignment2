package model

type ApiReqFilterModel struct {
	ReqParam struct{} `json:"req_param,omitempty"`
	// ReqBody holds the value of the "req_body" field.
	ReqBody EarthquakeFilterModel `json:"req_body,omitempty"`
	// ReqHeaders holds the value of the "req_headers" field.
	ReqHeaders struct{} `json:"req_headers,omitempty"`
	// ReqMetadata holds the value of the "req_metadata" field.
	ReqMetadata struct{} `json:"req_metadata,omitempty"`
}
