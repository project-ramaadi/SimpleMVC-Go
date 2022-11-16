package TodoControllers

import (
	"github.com/gin-gonic/gin"
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

	id, err := utils.GenerateSnowflake()

	if err != nil {
		context.JSON(http.StatusInternalServerError, utils.HttpResponse[*models.Todo]{
			Success:    false,
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
			Body:       nil,
		})
		return
	}

	var newTodo = models.Todo{
		ID:    id,
		Title: title,
		Done:  done,
	}

	todo := models.CreateNewTodo(newTodo)

	context.JSON(http.StatusOK, utils.HttpResponse[models.Todo]{
		Success:    true,
		Message:    "Todo created successfully",
		StatusCode: http.StatusOK,
		Body:       todo,
	})
}
