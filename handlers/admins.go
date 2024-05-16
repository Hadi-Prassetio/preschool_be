package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	dto "preschool/dto/result"
	"preschool/models"
	"preschool/pkg/bcrypt"
	jwtToken "preschool/pkg/jwt"
	"preschool/repositories"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAdmin struct {
	AdminRepository repositories.AdminRepository
}

func HandlerAdmin(AdminRepository repositories.AdminRepository) *handlerAdmin {
	return &handlerAdmin{AdminRepository}
}

func (h *handlerAdmin) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(models.RequestRegister)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	dataEmail, _ := h.AdminRepository.Login(request.Email)
	if dataEmail.Email != "" {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Email has been used"}
		json.NewEncoder(w).Encode(response)
		return
	}

	admin := models.Admin{
		FullName     : request.FullName     ,
		AdminUserName: request.AdminUserName,
		Email        : request.Email        ,
		Phone        : request.Phone        ,
		Password     : password             ,
		Role         : "Sub"                ,
	}

	data, err := h.AdminRepository.Register(admin)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAdmin) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(models.RequestLogin)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	admin := models.Admin{
		Email:    request.Email,
		Password: request.Password,
	}

	// Check email
	admin, err := h.AdminRepository.Login(admin.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check password
	isValid := bcrypt.CheckPasswordHash(request.Password, admin.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = admin.ID
	claims["fullname"] = admin.FullName
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
		return
	}

	loginResponse := models.ResponseLogin{
		ID      : admin.ID      ,
		FullName: admin.FullName,
		Email   : admin.Email   ,
		Token   : token         ,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerAdmin) CheckAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	adminInfo := r.Context().Value("adminInfo").(jwt.MapClaims)
	adminId := int(adminInfo["id"].(float64))

	// Check Admin by Id
	admin, err := h.AdminRepository.Getadmin(adminId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	CheckAdminResponse := models.CheckAdminResponse{
		ID      : admin.ID      ,
		FullName: admin.FullName,
		Email   : admin.Email   ,
		Role    : admin.Role    ,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: CheckAdminResponse}
	json.NewEncoder(w).Encode(response)
}
