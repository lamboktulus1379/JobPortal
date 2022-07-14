package dto

type ResPositions struct {
	Res
	Data       *[]Position `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type ResPosition struct {
	Res
	Data *Position `json:"data,omitempty"`
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
