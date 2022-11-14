package TodoControllers

import (
	"github.com/gin-gonic/gin"
	"mvcgo/models"
	"mvcgo/utils"
	"net/http"
)

func HandleGetTodoById(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := models.GetTodoById(id)

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
		Message:    "Todo found",
		StatusCode: http.StatusOK,
		Body:       todo,
	})

}
