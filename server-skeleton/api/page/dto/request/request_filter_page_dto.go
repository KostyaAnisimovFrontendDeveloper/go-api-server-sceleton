package request

type RequestFilterPageDto struct {
	Name  []string `form:"names[]"`
	Limit int      `form:"limit"`
}
