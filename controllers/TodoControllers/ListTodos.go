package TodoControllers

import (
	"github.com/gin-gonic/gin"
	"mvcgo/models"
	"mvcgo/utils"
	"net/http"
)

func HandleListTodos(context *gin.Context) {
	context.JSON(http.StatusOK, utils.HttpResponse[[]models.Todo]{
		Success:    true,
		Message:    "Todos fetched successfully",
		StatusCode: http.StatusOK,
		Body:       models.GetTodos(),
	})
}
