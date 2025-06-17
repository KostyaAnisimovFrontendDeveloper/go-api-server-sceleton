package page

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"server-skeleton/dictionary"
	_ "server-skeleton/docs"
	"server-skeleton/utils"
	"strings"
	"time"
)

// ================================== Get page by ID ===================================================================

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
// @Success 200 {array} RequestPageDTO
// @Router /page/{id} [get]
func GetPageById(c *gin.Context) {

	requestDto, _ := parseDtoId(c)

	if requestDto.ID == "" {
		return
	}

	resultDto, err := GetOneById(requestDto)

	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	if resultDto.ID == uuid.Nil {
		message := fmt.Sprintf(dictionary.PageByIdNotFound, requestDto.ID)
		c.JSON(http.StatusNotFound, &ErrorResponseDto{
			Message: message,
		})
		return
	}

	c.JSON(http.StatusOK, resultDto)
}

// ================================== Get pages by filter ==============================================================

//	@title			Getting pages
//	@version		1.0
//	@contact.name	API Support
//	@contact.email	kostiaGm@gmail.com
//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// GetPagesListByFilter @Summary Getting pages
// @Description Getting pages
// @Tags pages
// @Accept json
// @Produce json
// @Param names query []string false "Name"
// @Param limit query int false "Limit"
// @Param cursor query string false "cursor (las id uuid)"
// @Param lastTimestamp query string false "lastTimestamp"
// @Param orders[created_at] query []string false "Filter created_at Like min-max (example: 2025-06-11T08:28:51.400404Z)"
// @Success 200 {array} RequestPageDTO
// @Router /page [get]
func GetPagesListByFilter(c *gin.Context) {
	names := c.QueryArray("names")
	lastTimestamp := c.Query("lastTimestamp")
	ordersCreatedAt := c.Query("orders[created_at]")
	var orders map[string]string

	if ordersCreatedAt != "" {
		orders = map[string]string{
			"created_at": strings.ToUpper(ordersCreatedAt),
		}
	} else {
		orders = map[string]string{
			"created_at": "DESC",
		}
	}

	if lastTimestamp != "" {
		_, err := time.Parse(time.RFC3339Nano, lastTimestamp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(dictionary.ErrorParsingFilter, "lastTimestamp", lastTimestamp)})
			return
		}
	}

	if len(names) > 0 {
		names = strings.Split(names[0], ",")
	}

	requestFilterPageDto := &RequestFilterPageDto{
		Name:          names,
		LastTimestamp: lastTimestamp,
		Orders:        orders,
	}

	if err := c.ShouldBindQuery(&requestFilterPageDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.Dump(requestFilterPageDto)
	resultListDTO, err := GetItems(requestFilterPageDto)
	if err != nil {
		utils.LogError(dictionary.SomethingWrong, nil)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	listLength := len(resultListDTO.List)

	if listLength == 0 {
		utils.LogError(dictionary.SomethingWrong, nil)
		c.JSON(http.StatusNotFound, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})

		return
	}
	c.JSON(http.StatusOK, resultListDTO)
}

// ================================== Delete page by ID ================================================================

//	@title			Deleting page by id
//	@version		1.0
//	@contact.name	API Support
//	@contact.email	kostiaGm@gmail.com
//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// DeletePageById @Summary Getting page by id
// @Description Deleting page by id
// @Tags page
// @Accept json
// @Produce json
// @Param id path string true "Page id (UUID)"
// @Success 200 {array} RequestPageDTO
// @Router /page/{id} [delete]
func DeletePageById(c *gin.Context) {
	_, id := parseDtoId(c)

	if id == uuid.Nil {
		utils.LogError(dictionary.SomethingWrong, nil)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	isDeleted, err := DeletePageItemById(id)
	if err != nil || !isDeleted {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	c.JSON(http.StatusOK, &SuccessResponseDto{
		Message: dictionary.PageDeletedSuccessful,
	})
}

// ================================== Patch page by ID =================================================================
//	@title			Patch page
//	@version		1.0
//	@contact.name	API Support
//	@contact.email	kostiaGm@gmail.com
//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// PatchPageById     godoc
// @Summary      Patch page
// @Description  Update only sent fields
// @Tags         page
// @Accept       json
// @Produce      json
// @Param id path string true "Page id (UUID)"
// @Param        request body RequestPageDTO true "Updated data"
// @Success      200 {object}  PageItemResultDto
// @Failure      400 {object}  map[string]interface{}
// @Failure      500 {object}  map[string]interface{}
// @Router       /page/{id} [patch]
func PatchPageById(c *gin.Context) {
	_, id := parseDtoId(c)
	page, _ := parseRequestBody(c)
	page.ID = id

	isUpdated, err := PatchPageItem(page)

	if err != nil || !isUpdated {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	c.JSON(http.StatusOK, &SuccessResponseDto{
		Message: dictionary.SaveSuccessfulMessage,
	})
}

// ================================== Put page by ID ===================================================================
//	@title			Put page
//	@version		1.0
//	@contact.name	API Support
//	@contact.email	kostiaGm@gmail.com
//	@license.name	MIT
//	@license.url	https://opensource.org/license/mit

// PutPageItemById   godoc
// @Summary      Put page
// @Description  Update all sent fields
// @Tags         page
// @Accept       json
// @Produce      json
// @Param id path string true "Page id (UUID)"
// @Param        request body  RequestPageDTO true "Updated data"
// @Success      200 {object}  SuccessResponseDto
// @Failure      400 {object}  map[string]interface{}
// @Failure      500 {object}  map[string]interface{}
// @Router       /page/{id} [put]
func PutPageItemById(c *gin.Context) {
	requestIdDto, _ := parseDtoId(c)
	_, requestPagePostDTO := parseRequestBody(c)
	pageMap := convertRequestPageDTOToMap(c, requestPagePostDTO)
	pageMap["updated_at"] = time.Now()
	isUpdated, err := PutPageItem(requestIdDto, pageMap)

	if err != nil || !isUpdated {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	c.JSON(http.StatusOK, &SuccessResponseDto{
		Message: dictionary.SaveSuccessfulMessage,
	})
}

// ================================== Create (post) page ===============================================================
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
// @Param        request body RequestPageDTO true "Sent date"
// @Success      200 {object}  SuccessResponseDto
// @Failure      400 {object}  map[string]interface{}
// @Failure      500 {object}  map[string]interface{}
// @Router       /page [post]
func CreatePage(c *gin.Context) {

	page, _ := parseRequestBody(c)
	isCreated, err := CreatePageItem(page)

	if err != nil || !isCreated {
		utils.Dump(err)
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	c.JSON(http.StatusCreated, &SuccessResponseDto{
		Message: dictionary.SaveSuccessfulMessage,
	})
}

// === Sys

func parseDtoId(c *gin.Context) (RequestPageIdDTO, uuid.UUID) {
	var requestPageIdDTO RequestPageIdDTO
	var id uuid.UUID
	if err := c.ShouldBindUri(&requestPageIdDTO); err != nil {
		utils.LogError(dictionary.PageByIdNotFound, err)
		c.JSON(http.StatusUnprocessableEntity, &ErrorResponseDto{
			Message: err.Error(),
		})
	} else {
		id, err = uuid.Parse(requestPageIdDTO.ID)
		if err != nil {
			utils.LogError(dictionary.PageByIdNotFound, err)
			c.JSON(http.StatusNotFound, &ErrorResponseDto{
				Message: dictionary.PageByIdNotFound,
			})
		}
	}

	return requestPageIdDTO, id
}

func parseRequestBody(c *gin.Context) (Page, RequestPageDTO) {
	var requestPagePostDTO RequestPageDTO
	if err := c.ShouldBindJSON(&requestPagePostDTO); err != nil {
		utils.LogError(dictionary.ErrorParsingRequestBody, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
	}

	return Page{
			Name: requestPagePostDTO.Name,
		},
		requestPagePostDTO
}
