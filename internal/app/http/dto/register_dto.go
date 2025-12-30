package dto

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type RegisterDTO struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6"`
}

func (d *RegisterDTO) Validate() map[string]string {
	errs := map[string]string{}

	if err := validate.Struct(d); err != nil {
		if ves, ok := err.(validator.ValidationErrors); ok {
			for _, fe := range ves {
				switch fe.Field() {

				case "Email":
					if fe.Tag() == "required" {
						errs["email"] = "Email is required"
					} else if fe.Tag() == "email" {
						errs["email"] = "Invalid email address"
					}

				case "Password":
					if fe.Tag() == "required" {
						errs["password"] = "Password is required"
					} else if fe.Tag() == "min" {
						errs["password"] = "Password must be at least 6 characters"
					}
				}
			}
		}
	}

	return errs
}
