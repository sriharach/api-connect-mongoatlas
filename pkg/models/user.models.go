package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignInInput struct {
	E_mail   string `json:"e_mail" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthToken struct {
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}

type ModuleProfile struct {
	// ID         uuid.UUID `json:"id"`
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	User_name  string             `json:"user_name"`
	E_mail     string             `json:"e_mail"`
	Password   *string            `json:"-"`
	First_name *string            `json:"first_name"`
	Last_name  *string            `json:"last_name"`
	Activate   uint8              `json:"activate"`
	Picture    *string            `json:"picture"`
	Is_oauth   bool               `json:"-"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
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
