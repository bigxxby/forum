package models

type DefaultMessage struct {
	Message string `json:"message"`
}
type AuthenticationMessage struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
}
type PostCreationMessage struct {
	Message string `json:"message"`
	PostId  int    `json:"postId"`
}
type BadRequest struct {
	Message string `json:"message"`
}
