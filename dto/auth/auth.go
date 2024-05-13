package authdto

type RequestRegister struct {
	FullName      string `json:"fullname" validate:"required"`
	AdminUserName string `json:"admin_user_name" validate:"required"`
	Email         string `json:"email" validate:"required"`
	Phone         string `json:"phone" validate:"required"`
	Password      string `json:"password" validate:"required"`
	Role      	  string `json:"role" validate:"required"`
}

type RequestLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ResponseLogin struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"fullname"`
	Token    string `json:"token"`
}

type CheckAuthResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
