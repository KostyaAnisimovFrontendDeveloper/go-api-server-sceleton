package page

import (
	"github.com/gin-gonic/gin"
)

const UriPage = "/page"
const UriPageGetById = UriPage + "/:id"
const UriPageGetByIdS = UriPage + "/%s"

func InitPageRoutes(route *gin.Engine) {
	route.GET(UriPage, GetPagesListByFilter)
	route.GET(UriPageGetById, GetPageById)
	route.POST(UriPage, CreatePage)
	route.PUT(UriPageGetById, PutPageItemById)
	route.PATCH(UriPageGetById, PatchPageById)
	route.DELETE(UriPageGetById, DeletePageById)
}
