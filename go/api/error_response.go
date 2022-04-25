package api

import "encoding/json"

// Response is sent for all API success responses without additional
// data and for all errors
type Response struct {
	OK      bool   `json:"ok"`
	ReplyTo int    `json:"reply_to"`
	Error   string `json:"error,omitempty"`
}

// NewOKResponse creates a JSON encoded ok response
func NewOKResponse(requestID int) []byte {
	okResponse := Response{OK: true, ReplyTo: requestID}
	data, err := json.Marshal(okResponse)
	if err != nil {
		panic("response unable to marshal to json")
	}
	return data
}

// NewErrorResponse creates a JSON encoded error response wrapping the given message
func NewErrorResponse(requestID int, error string) []byte {
	errResponse := Response{OK: false, ReplyTo: requestID, Error: error}
	data, err := json.Marshal(errResponse)
	if err != nil {
		panic("response unable to marshal to json")
	}
	return data
}
