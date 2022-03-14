package model

// Pagination is the data transfer object for pagination.
type Pagination struct {
	Page  int `form:"page,default=1" binding:"omitempty,min=1"`
	Limit int `form:"limit,default=10" binding:"omitempty,min=1,max=50"`
}
