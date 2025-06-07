package dto

type ResponseFilterPageDto struct {
	Name  []string `form:"name[]"`
	Limit int      `form:"limit"`
}
