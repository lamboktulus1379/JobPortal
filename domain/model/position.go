package model

type ReqQueryParam struct {
	Description string `json:"description" form:"description"`
	Location    string `json:"location" form:"location"`
	FullTime    string `json:"full_time" form:"full_time"`
	Page        int    `json:"page" form:"page"`
}

type ReqURIParam struct {
	ID string `json:"string" uri:"id" binding:"required,uuid4"`
}
