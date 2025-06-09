package dto

import "server-skeleton/api/page/dto/result"

type ResultListDTO struct {
	List []result.PageItemResultDto `json:"list"`
}
