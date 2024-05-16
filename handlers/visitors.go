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

type handlerVisitor struct {
	VisitorRepository repositories.VisitorRepository
}

func HandlerVisitor(VisitorRepository repositories.VisitorRepository) *handlerVisitor {
	return &handlerVisitor{VisitorRepository}
}

func (h *handlerVisitor) FindVisitors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	visitors, err := h.VisitorRepository.FindVisitors()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: visitors}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerVisitor) FindVisitorAccepted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	status := "Accept"
	visitor, err := h.VisitorRepository.FindVisitorStatus(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: visitor}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerVisitor) FindVisitorWait(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	status := "Waiting"
	visitor, err := h.VisitorRepository.FindVisitorStatus(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: visitor}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerVisitor) FindVisitorCancel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	status := "Cancel"
	visitor, err := h.VisitorRepository.FindVisitorStatus(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: visitor}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerVisitor) GetVisitor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	visitor, err := h.VisitorRepository.GetVisitor(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: visitor}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerVisitor) CreateVisitor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	request := new(models.CreateVisitor)
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

	visitor := models.Visitor{
		Name     : request.Name     ,
		Email    : request.Email    ,
		Phone    : request.Phone    ,
		ChildName: request.ChildName,
		ChildAge : request.ChildAge ,
		Message  : request.Message  ,
		Status   : "Waiting"        ,
		AdminID  : 1                ,
	}
	data, err := h.VisitorRepository.CreateVisitor(visitor)
	if err != nil {
		w.Header().Set("Content-type", "aplication/json")
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerVisitor) UpdateVisitor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	request := new(models.UpdateVisitor)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	visitor, _ := h.VisitorRepository.GetVisitor(id)

	if request.AdminID != 0 {
		visitor.AdminID = request.AdminID
	}
	if request.Status != "" {
		visitor.Status = request.Status
	}

	visitor.UpdatedAt = time.Now()

	data, err := h.VisitorRepository.UpdateVisitor(visitor)
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

func (h *handlerVisitor) DeleteVisitor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	visitor, err := h.VisitorRepository.GetVisitor(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.VisitorRepository.DeleteVisitor(visitor)
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
