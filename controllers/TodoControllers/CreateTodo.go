package TodoControllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mvcgo/models"
	"mvcgo/utils"
	"net/http"
)

func HandleCreateTodo(context *gin.Context) {

	title := context.PostForm("title")
	done := false

	if context.PostForm("done") == "true" {
		done = true
	}

	var newTodo = models.Todo{
		ID:    uuid.New().String(),
		Title: title,
		Done:  done,
	}

	todo := models.CreateNewTodo(newTodo)

	context.JSON(http.StatusOK, utils.HttpResponse[models.Todo]{
		Success: true,
		Message: "Todo created successfully",
		Body:    todo,
	})
}
