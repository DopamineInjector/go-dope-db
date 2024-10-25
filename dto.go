package godopedb;

// Request body specifying requested value
type SelectValueRequest struct {
	Key string `json:"key"` // Value's key in the KV store
	Namespace string `json:"namespace"` // Namespace that the value is stored under
}

// Response returned from the db upon successfull fetching of the value
type ValueResponse struct {
	Value string `json:"value"` // Fetched value
	Checksum string `json:"checksum"` // Checkusum of the entire DB state
}

// Response including the checksum of the entire DB state
type ChecksumResponse struct {
	Checksum string `json:"checksum"` // Checksum of the entire DB state
}

// Request body specifying data to be inserted
type InsertValueRequest struct {
	Key string `json:"key"` // Key to be inserted under
	Value string `json:"value"` // Value to insert
	Namespace string `json:"namespace"` // Namespace to insert the KV pair into
}

// Request body specifying a namespace to be operated on
type SelectNamespaceRequest struct {
	Namespace string `json:"namespace"` // Selected namespace
}
