package route

import (
	"github.com/gin-gonic/gin"
	"server-skeleton/api/page/handler"
)

const UriPageGetById = "/page/:id"
const UriPageGetByIdS = "/page/%s"

func InitPageRoutes(route *gin.Engine) {
	route.GET(UriPageGetById, handler.GetPageById)
}
