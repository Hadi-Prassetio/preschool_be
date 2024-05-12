package teacherdto

type CreateTeacher struct {
	FullName string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone"`
	Subject  string `json:"subject" validate:"required"`
	Register string `json:"register" validate:"required"`
}

type UpdateTeacher struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Subject  string `json:"subject"`
	Register string `json:"register"`
}
