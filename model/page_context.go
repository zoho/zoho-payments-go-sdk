package model

type PageContext struct {
	Page        int  `json:"page"`
	PerPage     int  `json:"per_page"`
	Total       int  `json:"total"`
	TotalPages  int  `json:"total_pages"`
	HasMorePage bool `json:"has_more_page"`
}
