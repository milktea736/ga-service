package services

type ResponseJWTToken struct {
	Token string `json:"token"`
}

func IssueJWT(userToken string, saToken string) (*ResponseJWTToken, error) {
	if userToken == "" {
		// issue a robot token
		return &ResponseJWTToken{Token: "robot_token"}, nil
	} else {
		// issue a user token
		return &ResponseJWTToken{Token: "user_token"}, nil
	}
}

func RefreshJWT() {
	// swagger:route POST /auth/jwt/refresh auth refreshJWT
	// Responses:
	// 200: jwtResponse
}

func GetPublicKey() {
	// swagger:route GET /auth/jwt/public auth getPublicKey
	// Responses:
	// 200: publicKeyResponse
}
