package handlers

import (
	"encoding/json"
	"net/http"
	dto "preschool/dto/result"
	"preschool/models"
	"preschool/repositories"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerEnrollment struct {
	EnrollmentRepository repositories.EnrollmentRepository
}

func HandlerEnrollment(EnrollmentRepository repositories.EnrollmentRepository) *handlerEnrollment {
	return &handlerEnrollment{EnrollmentRepository}
}

func (h *handlerEnrollment) FindEnrollments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	teachers, err := h.EnrollmentRepository.FindEnrollments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: teachers}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerEnrollment) FindEnrollmentWait(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	status := "Waiting"
	enrollment, err := h.EnrollmentRepository.FindEnrollmentStatus(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: enrollment}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerEnrollment) FindEnrollmentAccepted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	status := "Accept"
	enrollment, err := h.EnrollmentRepository.FindEnrollmentStatus(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: enrollment}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerEnrollment) FindEnrollmentCancel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	status := "Cancel"
	enrollment, err := h.EnrollmentRepository.FindEnrollmentStatus(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: enrollment}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerEnrollment) GetEnrollment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	enrollment, err := h.EnrollmentRepository.GetEnrollment(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: enrollment}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerEnrollment) CreateEnrollment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	request := new(models.CreateEnrollment)
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

	enrollment := models.Enrollment{
		FatherName : request.FatherName,
		MotherName : request.MotherName,
		Email      : request.Email     ,
		Phone      : request.Phone     ,
		ChildName  : request.ChildName ,
		ChildAge   : request.ChildAge  ,
		ClassID    : request.ClassID   ,
		Status     : "Waiting"         ,
		AdminID    : 1                 ,
	}
	data, err := h.EnrollmentRepository.CreateEnrollment(enrollment)
	if err != nil {
		w.Header().Set("Content-type", "aplication/json")
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerEnrollment) UpdateEnrollment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	request := new(models.UpdateEnrollment)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	enrollment, _ := h.EnrollmentRepository.GetEnrollmentUpdate(id)

	if request.Status !=""{
		enrollment.Status = request.Status
	}
	if request.ClassID !=0{
		enrollment.ClassID = request.ClassID
	}
	enrollment.UpdatedAt = time.Now()

	data, err := h.EnrollmentRepository.UpdateEnrollment(enrollment)
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

func (h *handlerEnrollment) DeleteEnrollment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	enrollment, err := h.EnrollmentRepository.GetEnrollment(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.EnrollmentRepository.DeleteEnrollment(enrollment)
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
