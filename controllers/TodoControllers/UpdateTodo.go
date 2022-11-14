package TodoControllers

import (
	"github.com/gin-gonic/gin"
	"mvcgo/models"
	"mvcgo/utils"
	"net/http"
)

func HandleUpdateTodo(ctx *gin.Context) {

	updatedTodo, err := models.UpdateTodoTitleById(ctx.Param("id"), ctx.PostForm("title"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HttpResponse[*models.Todo]{
			Success:    false,
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
			Body:       nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.HttpResponse[*models.Todo]{
		Success:    true,
		Message:    "Todo updated successfully",
		StatusCode: http.StatusOK,
		Body:       updatedTodo,
	})

}
