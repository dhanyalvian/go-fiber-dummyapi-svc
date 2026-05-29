//- apps/entities/auth_entity.go

package entities

type AuthToken struct {
	BaseID

	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type RespAuthLogin struct {
	BaseID

	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`

	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
