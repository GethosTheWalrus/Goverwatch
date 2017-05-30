package models

// Model representing an HTTP response
// Fields: 
// Status: HTTP status (200, 404, 500, etc)
// RequestType: What type of data was requested (PlayerDetail, HeroDetail, etc)
// Data: The response body
type Response struct {
	Status string
	RequestType string
	Data interface{}
}