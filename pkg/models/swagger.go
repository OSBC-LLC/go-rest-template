package models

// These models are to NOT be used in actual code. This is for the Swagger Docs ONLY!!
// Create a new file in the models package to store any structs you need to use in the code.

type Heartbeat struct {
	RequestID      string `json:"requestId"`
	DatabaseOnline bool   `json:"databaseOnline"`
	AppName        string `json:"appName"`
	ReleaseDate    string `json:"releaseCreatedAt"`
	ReleaseVersion string `json:"releaseVersion"`
	Slug           string `json:"slugCommit"`
	Message        string `json:"message"`
}

type Todo struct {
	RequestID string `json:"requestId"`
	Message   string `json:"message"`
}
