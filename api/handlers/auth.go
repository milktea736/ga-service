package auth

import (
	"gaservice/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type jwtRequest struct {
	UserToken string `json:"user_token" example:"example_user_token"`
	SAToken   string `json:"sa_token" example:"example_sa_token"`
}

type refreshRequest struct {
	Token string `json:"token" example:"example_jwt_token"`
}

type jwtResponse struct {
	Token     string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	ExpiresIn int    `json:"expires_in" example:"3600"`
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// IssueJWT handles the issuance of JWT tokens
//
//	@Summary		Issue a JWT token
//	@Description	Issue a JWT token based on provided user_token and sa_token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			jwtRequest	body		jwtRequest			false	"Issued a GA token"
//	@Success		200			{object}	jwtResponse			"Return an GA token"
//	@Failure		400			{object}	HTTPError	"Invalid input or sa_token is required"
//	@Failure		500			{object}	HTTPError	"Internal server error"
//	@Router			/auth/jwt [post]
func IssueJWT(context *gin.Context) {
	var req jwtRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.SAToken == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "sa_token is required"})
		return
	}
	response, err := services.IssueJWT(req.UserToken, req.SAToken)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, response)
}

// Refresh JWT token
//
//	@Summary		Refresh a JWT token
//	@Description	Refresh a JWT token based on provided token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			refreshRequest	body		refreshRequest		true	"Refresh an issued token"
//	@Success		200				{object}	jwtResponse			"Return an GA token"
//	@Failure		400				{object}	HTTPError	"Invalid input"
//	@Failure		500				{object}	HTTPError
//	@Router			/auth/jwt/refresh [post]
func RefreshJWT(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "refresh"})
}

func GetPublicKey(context *gin.Context) {
	// return the public key
}
