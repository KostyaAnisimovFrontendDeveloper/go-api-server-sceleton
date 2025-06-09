package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server-skeleton/api/page/dictionary"
	"server-skeleton/api/page/dto/request"
	"server-skeleton/api/page/repository"
	"server-skeleton/utils"
)

//	@title			Put page
//	@version		1.0
//	@contact.name	API Support
//	@contact.email	kostiaGm@gmail.com
//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// PutPage   godoc
// @Summary      Patch page
// @Description  Update all sent fields
// @Tags         page
// @Accept       json
// @Produce      json
// @Param        request body request.RequestPageDTO true "Updated data"
// @Success      200 {object}  result.PageItemResultDto
// @Failure      400 {object}  map[string]interface{}
// @Failure      500 {object}  map[string]interface{}
// @Router       /page/{id} [put]
func PutPage(c *gin.Context) {
	var requestPageIdDTO request.RequestPageIdDTO
	var requestPagePostDTO request.RequestPageDTO

	errorMessage := fmt.Sprintf(dictionary.CreateError, requestPagePostDTO.Name)

	if err := c.ShouldBindQuery(&requestPageIdDTO); err != nil {
		utils.LogFormat(errorMessage+" Error getting id from url param", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	if err := c.ShouldBindJSON(&requestPagePostDTO); err != nil {
		utils.LogFormat(errorMessage+" Error parsing JSON", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	responsePageDto, err := repository.PatchPageItem(requestPageIdDTO, requestPagePostDTO)
	if err != nil {
		utils.LogFormat(errorMessage+" Error saving in db", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, responsePageDto)
}
