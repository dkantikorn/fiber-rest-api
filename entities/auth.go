package entities

type (
	Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	Token struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)
