package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/transactional_outbox_pattern/main_service/database"
	"github.com/transactional_outbox_pattern/main_service/model"
)

func createUserHandler(context *gin.Context) {
	var user model.User

	if err := context.BindJSON(&user); err != nil {
		return
	}

	db := database.NewDBConnection()

	db.InsertData(&user)

	context.IndentedJSON(http.StatusCreated, user)
}