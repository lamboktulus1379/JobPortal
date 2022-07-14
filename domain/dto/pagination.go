package dto

type Pagination struct {
	PageNumber  int `json:"page_number,omitempty"`
	TotalRecord int `json:"total_record,omitempty"`
}
