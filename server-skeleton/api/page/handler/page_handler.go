package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-skeleton/api/page/dto"
	"server-skeleton/api/page/repository"
	"server-skeleton/api_init"
)

func GetPageById(c *gin.Context) {

	var requestDto dto.RequestPageDTO
	if err := c.ShouldBindUri(&requestDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	}

	resultDto, err := repository.GetOneById(api_init.InitGlobal.Dbh, requestDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	}

	c.JSON(http.StatusOK, resultDto)
}

func GetPage(c *gin.Context) {
	names := c.QueryArray("names[]")

}
