package repository

import (
	"fmt"
	"gorm.io/gorm"
	"server-skeleton/api/page/dto"
	"strings"
)

func GetOneById(db *gorm.DB, requestDto dto.RequestPageDTO) (dto.ResponsePageDTO, error) {
	sql := "SELECT * FROM pages WHERE id = $1"
	var page dto.ResponsePageDTO
	err := db.Raw(sql, requestDto.ID).Scan(&page).Error
	return page, err
}

func Get(db *gorm.DB, filterDto *dto.ResponseFilterPageDto) ([]dto.ResponsePageDTO, error) {
	sql := "SELECT * FROM pages"

	where := ""

	limit := " LIMIT %d"

	if filterDto.Limit <= 0 {
		filterDto.Limit = 10
	}

	limit = fmt.Sprintf(limit, filterDto.Limit)

	if len(filterDto.Name) > 0 {
		where += fmt.Sprintf(" AND name IN ('%s')", strings.Join(filterDto.Name, "','"))
	}

	if len(where) > 0 {
		sql += where
	}

	sql += limit

	var pages []dto.ResponsePageDTO

	err := db.Raw(sql).Scan(&pages).Error
	return pages, err
}
