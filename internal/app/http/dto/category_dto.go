package requests

type CreateCategoryDTO struct {
	Name string `form:"name" json:"name" validate:"required,min=2"`
}
