package page

import (
	"fmt"
	"github.com/google/uuid"
	"server-skeleton/api_init"
	"strings"
)

func GetItems(filterDto *RequestFilterPageDto) (*ResultListDTO, error) {

	query := api_init.GetDbh().Model(&Page{})

	where := "(created_at, id) %s (?, ?)"
	orderCreatedAt, orderCreatedExists := filterDto.Orders["created_at"]

	if orderCreatedExists && strings.ToUpper(orderCreatedAt) != "DESC" {
		query.Order("created_at ASC")
		query.Order("id ASC")
		where = fmt.Sprintf(where, ">")
	} else {
		query.Order("created_at DESC")
		query.Order("id DESC")
		where = fmt.Sprintf(where, "<")
	}

	if filterDto.Cursor != "" && filterDto.LastTimestamp != "" {
		query.Where(where, filterDto.LastTimestamp, filterDto.Cursor)
	}

	if len(filterDto.Name) > 0 {

		var likeConditions []string
		var likeArgs []interface{}
		for _, name := range filterDto.Name {
			likeConditions = append(likeConditions, "name LIKE ?")
			likeArgs = append(likeArgs, name+"%")
		}
		query = query.Where("("+strings.Join(likeConditions, " OR ")+")", likeArgs...)

	}

	if filterDto.Limit <= 0 {
		filterDto.Limit = 10
	}

	var total int64
	var result []PageItemResultDto
	err := query.Count(&total).Error

	if err != nil {
		return nil, err
	}

	query.Limit(filterDto.Limit)
	query.Find(&result)

	lastItem := result[len(result)-1]
	resultDto := ResultListDTO{
		List:          result,
		Cursor:        lastItem.ID,
		LastTimestamp: lastItem.CreatedAt,
		Total:         total,
	}

	return &resultDto, nil
}

func GetOneById(requestPageIdDTO RequestPageIdDTO) (*PageItemResultDto, error) {

	if requestPageIdDTO.ID == "" {
		return nil, nil
	}

	result := PageItemResultDto{}
	err := api_init.GetDbh().Raw("SELECT * FROM pages WHERE id = $1 LIMIT 1", requestPageIdDTO.ID).Scan(&result).Error
	return &result, err
}

func CreatePageItem(page Page) (bool, error) {
	err := api_init.GetDbh().Create(&page).Error
	return result(err)
}

func PutPageItem(requestPageIdDTO RequestPageIdDTO, page map[string]interface{}) (bool, error) {
	err := api_init.GetDbh().Model(&Page{}).Where("id = ?", requestPageIdDTO.ID).Updates(page).Error
	return result(err)
}

func PatchPageItem(page Page) (bool, error) {
	err := api_init.GetDbh().Updates(&page).Error
	return result(err)
}

func DeletePageItemById(id uuid.UUID) (bool, error) {
	err := api_init.GetDbh().Delete(&Page{}, "id = ?", id.String()).Error
	return result(err)
}

func result(err error) (bool, error) {
	if err != nil {
		return false, err
	}
	return true, nil
}
