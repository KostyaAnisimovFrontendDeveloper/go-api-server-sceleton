package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"server-skeleton/api/page/dictionary"
	"server-skeleton/api/page/dto/request"
	"server-skeleton/api/page/repository"
	"server-skeleton/utils"
)

//	@title			Getting page by id
//	@version		1.0
//	@contact.name	API Support
//	@contact.email	kostiaGm@gmail.com
//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// GetPageById @Summary Getting page by id
// @Description Getting page by id
// @Tags page
// @Accept json
// @Produce json
// @Param id path string true "Page id (UUID)"
// @Success 200 {array} request.RequestPageDTO
// @Router /page/{id} [get]
func GetPageById(c *gin.Context) {

	var requestPageIdDTO request.RequestPageIdDTO
	if err := c.ShouldBindUri(&requestPageIdDTO); err != nil {
		utils.LogFormat(dictionary.ErrorFindById, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": dictionary.ErrorFindById})
		return
	}

	resultDto, err := repository.GetOneById(requestPageIdDTO)

	if err != nil {
		utils.LogFormat(dictionary.ErrorFindById, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": dictionary.ErrorFindById})
		return
	}

	if err == nil && resultDto.ID == uuid.Nil {
		utils.LogFormat(dictionary.NotFound, " - page id: "+requestPageIdDTO.ID)
		c.JSON(http.StatusNotFound, gin.H{"msg": dictionary.NotFound})
		return
	}

	c.JSON(http.StatusOK, resultDto)
}
