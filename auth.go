package Nap

type Authentication interface {
	AuthenticationHeader() string // basic <base64-encoded string>
}

type AuthToken struct {
	Token string `json:"token"`
}

type AuthBasic struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
