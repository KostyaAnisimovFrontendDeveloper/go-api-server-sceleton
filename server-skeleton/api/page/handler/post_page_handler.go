package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server-skeleton/api/page/dictionary"
	"server-skeleton/api/page/dto/request"
	_ "server-skeleton/api/page/dto/result"
	"server-skeleton/api/page/repository"
	"server-skeleton/utils"
)

//	@title			Create page
//	@version		1.0
//	@description	Create page
//	@contact.name	API Support
//	@contact.email	kostiaGm@gmail.com
//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// CreatePage godoc
// @Summary      Create page
// @Description  Create page
// @Tags         pages
// @Accept       json
// @Produce      json
// @Param        request body request.RequestPageDTO true "Sent date"
// @Success      200 {object} result.PageItemResultDto
// @Failure      400 {object}  map[string]interface{}
// @Failure      500 {object}  map[string]interface{}
// @Router       /page [post]
func CreatePage(c *gin.Context) {
	var requestPagePostDTO request.RequestPageDTO

	errorMessage := fmt.Sprintf(dictionary.CreateError, requestPagePostDTO.Name)

	if err := c.ShouldBindJSON(&requestPagePostDTO); err != nil {
		utils.LogFormat(errorMessage+" Error parsing JSON", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	responsePageDto, err := repository.CreatePageItem(requestPagePostDTO)
	if err != nil {
		utils.LogFormat(errorMessage+" Error saving in db", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusCreated, responsePageDto)
}
