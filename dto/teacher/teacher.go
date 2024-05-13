package teacherdto

type CreateTeacher struct {
	FullName string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	ClassID  int    `json:"class_id" validate:"required"`
	AdminID  int    `json:"admin_id" validate:"required"`
}

type UpdateTeacher struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	ClassID  int    `json:"class_id"`
	AdminID  int    `json:"admin_id"`
}
