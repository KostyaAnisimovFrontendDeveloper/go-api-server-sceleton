package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-skeleton/api/page/dictionary"
	"server-skeleton/api/page/dto/request"
	"server-skeleton/api/page/repository"
	"server-skeleton/utils"
	"strings"
)

//	@title			Getting pages
//	@version		1.0
//	@contact.name	API Support
//	@contact.email	kostiaGm@gmail.com
//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// GetPagesList @Summary Getting pages
// @Description Getting pages
// @Tags pages
// @Accept json
// @Produce json
// @Param names query []string false "Name"
// @Success 200 {array} request.RequestPageDTO
// @Router /page [get]
func GetPagesList(c *gin.Context) {
	names := c.QueryArray("names")

	if len(names) > 0 {
		names = strings.Split(names[0], ",")
	}

	requestFilterPageDto := &request.RequestFilterPageDto{
		Name: names,
	}

	if err := c.ShouldBindQuery(&requestFilterPageDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultListDTO, err := repository.GetItems(requestFilterPageDto)
	if err != nil {
		utils.LogFormat(dictionary.PageListErrors, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": dictionary.PageListErrors})
		return
	}

	if err == nil && len(resultListDTO.List) == 0 {
		utils.LogFormat(dictionary.PageListNotFound, "")
		c.JSON(http.StatusNotFound, gin.H{"msg": dictionary.PageListNotFound})
		return
	}
	c.JSON(http.StatusOK, resultListDTO)
}
