package handlers

import (
	"encoding/json"
	"net/http"
	dto "preschool/dto/result"
	teacherdto "preschool/dto/teacher"
	"preschool/models"
	"preschool/repositories"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerTeacher struct {
	TeacherRepository repositories.TeacherRepository
}

func HandlerTeacher(TeacherRepository repositories.TeacherRepository) *handlerTeacher {
	return &handlerTeacher{TeacherRepository}
}

func (h *handlerTeacher) FindTeachers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	teachers, err := h.TeacherRepository.FindTeachers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: teachers}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTeacher) GetTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	teacher, err := h.TeacherRepository.GetTeacher(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}


	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: teacher}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTeacher) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	request := new(teacherdto.CreateTeacher)
	teacher := models.Teacher{
		FullName: request.FullName,
		Email: request.Email,
		Phone: request.Phone,
		Subject: request.Subject,
		Register: request.Register,
	}
	data, err := h.TeacherRepository.CreateTeacher(teacher)
	if err != nil {
		w.Header().Set("Content-type", "aplication/json")
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTeacher) UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	request := new(teacherdto.UpdateTeacher)

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	teacher, _ := h.TeacherRepository.GetTeacher(id)

	if request.FullName!=""{
		teacher.FullName = request.FullName
	}
	if request.Email!=""{
		teacher.Email = request.Email
	}
	if request.Phone!=""{
		teacher.Phone = request.Phone
	}
	if request.Subject!=""{
		teacher.Subject = request.Subject
	}

	data, err := h.TeacherRepository.UpdateTeacher(teacher)
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

func (h *handlerTeacher) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	teacher, err := h.TeacherRepository.GetTeacher(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TeacherRepository.DeleteTeacher(teacher)
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
