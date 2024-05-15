package models

type DefaultMessage struct {
	Message string `json:"message"`
}
type AuthenticationMessage struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
}
type CreationMessage struct {
	Message string `json:"message"`
	Id      int    `json:"id"`
}
type BadRequest struct {
	Message string `json:"message"`
}
