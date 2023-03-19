package models

type OauthRequest struct {
	Code  string `json:"code"`
	State string `json:"state"`
}
