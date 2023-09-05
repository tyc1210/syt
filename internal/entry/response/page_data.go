package response

type PageData struct {
	TotalCount int         `json:"total"`
	Data       interface{} `json:"content"`
}

func NewPageData(count int, data interface{}) PageData {
	return PageData{TotalCount: count, Data: data}
}
