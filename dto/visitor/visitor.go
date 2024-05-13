package vistordto

type CreateVisitor struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Phone     string `json:"guardian_phone" validate:"required"`
	ChildName string `json:"child_name" validate:"required"`
	ChildAge  int    `json:"child_age" validate:"required"`
	Message   string `json:"message" validate:"required"`
	Status    string `json:"status"`
	AdminID   int    `json:"admin_id"`
}

type UpdateVisitor struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"guardian_phone"`
	ChildName string `json:"child_name"`
	ChildAge  int    `json:"child_age"`
	Message   string `json:"message"`
	Status    string `json:"status"`
	AdminID   int    `json:"admin_id"`
}
