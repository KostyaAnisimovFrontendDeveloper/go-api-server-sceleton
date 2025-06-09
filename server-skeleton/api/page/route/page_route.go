package route

import (
	"github.com/gin-gonic/gin"
	"server-skeleton/api/page/handler"
)

const UriPage = "/page"
const UriPageGetById = UriPage + "/:id"
const UriPageGetByIdS = UriPage + "/%s"

func InitPageRoutes(route *gin.Engine) {
	route.GET(UriPageGetById, handler.GetPageById)
	route.GET(UriPage, handler.GetPagesList)
	route.POST(UriPage, handler.CreatePage)
	route.PATCH(UriPage, handler.PatchPage)
	route.PUT(UriPage, handler.PutPage)
}
