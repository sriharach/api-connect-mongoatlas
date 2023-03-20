package models

type SignInInput struct {
	E_mail   string `json:"e_mail"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type AuthToken struct {
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}

type ModuleProfile struct {
	// ID         uuid.UUID `json:"id"`
	User_name  string `json:"user_name"`
	E_mail     string `json:"e_mail"`
	Password   string `json:"-"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Activate   uint8  `json:"activate"`
	Picture    string `json:"picture"`
	Is_oauth   bool   `json:"is_oauth"`
}

type ModuleProfileOauth struct {
	Issuer          string `json:"iss"`
	Subject         string `json:"sub"`
	Audience        string `json:"aud"`
	Expiry          int    `json:"exp"`
	IssuedAt        int    `json:"iat"`
	AtHash          string `json:"at_hash"`
	Hd              string `json:"hd"`
	AuthorizedParty string `json:"azp"`
	Picture         string `json:"picture"`
	Locale          string `json:"locale"`
	Email           string `json:"email"`
	EmailVerified   bool   `json:"email_verified"`
	Name            string `json:"name"`
	FamilyName      string `json:"family_name"`
	GivenName       string `json:"given_name"`
}
