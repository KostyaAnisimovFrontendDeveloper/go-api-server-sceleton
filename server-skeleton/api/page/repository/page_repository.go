package repository

import (
	"fmt"
	"server-skeleton/api/page/dto/request"
	pageItemResultDto "server-skeleton/api/page/dto/result"
	"server-skeleton/api_init"
	resultListDto "server-skeleton/pkg/dto"
	"strings"
)

func GetOneById(requestPageIdDTO request.RequestPageIdDTO) (pageItemResultDto.PageItemResultDto, error) {
	sql := "SELECT * FROM pages WHERE id = $1"
	var page pageItemResultDto.PageItemResultDto
	err := api_init.GetDbh().Raw(sql, requestPageIdDTO.ID).Scan(&page).Error
	return page, err
}

func GetItems(filterDto *request.RequestFilterPageDto) (*resultListDto.ResultListDTO, error) {
	sql := "SELECT * FROM pages"

	where := ""
	limit := " LIMIT %d"

	if filterDto.Limit <= 0 {
		filterDto.Limit = 10
	}

	limit = fmt.Sprintf(limit, filterDto.Limit)

	if len(filterDto.Name) > 0 {
		where += fmt.Sprintf("name IN ('%s')", strings.Join(filterDto.Name, "','"))
	}

	if len(where) > 0 {
		sql += " WHERE " + where
	}

	sql += limit

	var pages []pageItemResultDto.PageItemResultDto

	err := api_init.GetDbh().Raw(sql).Scan(&pages).Error

	result := &resultListDto.ResultListDTO{
		List: pages,
	}

	return result, err
}

func CreatePageItem(requestDto request.RequestPageDTO) (bool, error) {

	err := api_init.GetDbh().Create(requestDto).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func PatchPageItem(requestPageIdDTO request.RequestPageIdDTO, requestPageDTO request.RequestPageDTO) (bool, error) {

	err := api_init.GetDbh().Model(requestPageDTO.Name).Where("id=?", requestPageIdDTO.ID).Updates(requestPageDTO).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func PutPageItem(requestPageIdDTO request.RequestPageIdDTO, requestPageDTO request.RequestPageDTO) (bool, error) {

	err := api_init.GetDbh().Save(requestPageDTO.Name).Where("id=?", requestPageIdDTO.ID).Updates(requestPageDTO).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
