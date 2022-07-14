package models

type ReqQueryParam struct {
	Description string `json:"description,omitempty" url:"description,omitempty"`
	Location    string `json:"location,omitempty" url:"location,omitempty"`
	FullTime    string `json:"full_time,omitempty" url:"full_time,omitempty"`
	Page        int    `json:"page,omitempty" url:"page,omitempty"`
}

type Position struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	URL         string `json:"url,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	Company     string `json:"company,omitempty"`
	CompanyURL  string `json:"company_url,omitempty"`
	Location    string `json:"location,omitempty"`
	Title       string `json:"title,omitempty"`
	HowToApply  string `json:"how_to_apply,omitempty"`
	CompanyLogo string `json:"company_logo,omitempty"`
}
