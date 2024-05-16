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

type handlerPage struct {
	PageRepository repositories.PageRepository
}

func HandlerPage(PageRepository repositories.PageRepository) *handlerPage {
	return &handlerPage{PageRepository}
}

func (h *handlerPage) FindPages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	teachers, err := h.PageRepository.FindPages()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: teachers}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerPage) GetPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	teacher, err := h.PageRepository.GetPage(id)
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

func (h *handlerPage) CreatePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	request := new(models.CreatePage)
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

	teacher := models.Page{
		Type   : request.Type ,
		Title  : request.Title,
		Desc   : request.Desc ,
		Email  : request.Email,
		Phone  : request.Phone,
		AdminID: 1            ,
	}
	data, err := h.PageRepository.CreatePage(teacher)
	if err != nil {
		w.Header().Set("Content-type", "aplication/json")
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerPage) UpdatePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	request := new(models.UpdatePage)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	teacher, _ := h.PageRepository.GetPage(id)

	if request.Type != "" {
		teacher.Type = request.Type
	}
	if request.Title != "" {
		teacher.Title = request.Title
	}
	if request.Desc != "" {
		teacher.Desc = request.Desc
	}
	if request.Email != "" {
		teacher.Email = request.Email
	}
	if request.Phone != "" {
		teacher.Phone = request.Phone
	}
	teacher.UpdatedAt = time.Now()

	data, err := h.PageRepository.UpdatePage(teacher)
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

func (h *handlerPage) DeletePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	teacher, err := h.PageRepository.GetPage(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.PageRepository.DeletePage(teacher)
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
