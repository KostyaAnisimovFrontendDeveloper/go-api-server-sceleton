package page

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"server-skeleton/dictionary"
	"server-skeleton/utils"
	"time"
)

func convertRequestPageDTOToPage(c *gin.Context, requestPageIdDTO RequestPageIdDTO, requestPageDTO RequestPageDTO, page *Page) {
	id, err := uuid.Parse(requestPageIdDTO.ID)

	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}

	page.ID = id
	createdAt, err := time.Parse(time.RFC3339, requestPageDTO.CreatedAt)

	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}
	page.CreatedAt = createdAt

	updatedAt, err := time.Parse(time.RFC3339, requestPageDTO.UpdatedAt)

	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}
	page.UpdatedAt = updatedAt

	deletedAt, err := time.Parse(time.RFC3339, requestPageDTO.DeletedAt)

	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return
	}
	page.DeletedAt = deletedAt
}

func convertRequestPageDTOToMap(c *gin.Context, requestPageDTO RequestPageDTO) map[string]interface{} {

	result := map[string]interface{}{}

	createdAt, err := time.Parse(time.RFC3339, requestPageDTO.CreatedAt)

	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return nil
	}

	if createdAt.IsZero() {
		createdAt = time.Time{}
	}

	result["created_at"] = createdAt
	updatedAt, err := time.Parse(time.RFC3339, requestPageDTO.UpdatedAt)
	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return nil
	}

	if updatedAt.IsZero() {
		updatedAt = time.Time{}
	}

	result["updated_at"] = updatedAt

	deletedAt, err := time.Parse(time.RFC3339, requestPageDTO.DeletedAt)
	if err != nil {
		utils.LogError(dictionary.SomethingWrong, err)
		c.JSON(http.StatusInternalServerError, &ErrorResponseDto{
			Message: dictionary.SomethingWrong,
		})
		return nil
	}

	if deletedAt.IsZero() {
		deletedAt = time.Time{}
	}

	result["deleted_at"] = deletedAt

	return result
}
