package models

type DefaultMessage struct {
	Message string `json:"message"`
}
type AuthenticationMessage struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
}
