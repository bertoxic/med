package handler

import (
	"log"

	"github.com/bertoxic/med/services/authentication/internal/models"
	"github.com/bertoxic/med/services/authentication/internal/transport/httputil"
	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	user := &models.Patient{}
	ctx.ShouldBind(user)
	var userD models.UserDetails
	userD.Email = user.Email
	userD.FirstName = user.FirstName
	userD.LastName = user.LastName
	userD.UserType = "patient"

	// opt, err := GenerateOTP(9)
    // if err != nil {
    //     httputil.WriteJson(ctx,false,200, &models.JsonResponse{})	
    //     return
    // }

	var tokens models.Tokens
	tokens, err := GenerateToken(userD)
	if err != nil {
		log.Println(err)
		return
	}
	data := gin.H{
		"token":         tokens.Token,
		"refresh_token": tokens.RefreshToken,
	}
	httputil.WriteJson(ctx, true, 200, &models.JsonResponse{
		Data: data,
		Success: true,
		Message: "suceess obtaining json response",
	})
	log.Println(tokens)
}


