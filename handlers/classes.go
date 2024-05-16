package handlers

import (
	"encoding/json"
	"net/http"
	dto "preschool/dto/result"
	"preschool/models"
	"preschool/repositories"
	"strconv"
	"time"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerClass struct {
	ClassRepository repositories.ClassRepository
}

func HandlerClass(ClassRepository repositories.ClassRepository) *handlerClass {
	return &handlerClass{ClassRepository}
}

func (h *handlerClass) FindClasses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")

	classes, err := h.ClassRepository.FindClasses()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: classes}
	json.NewEncoder(w).Encode(response)
}


func (h *handlerClass) GetClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	class, err := h.ClassRepository.GetClass(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: class}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerClass) CreateClass(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")

    request := new(models.CreateClass)
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
        json.NewEncoder(w).Encode(response)
        log.Println("Error decoding request:", err) // Log pesan kesalahan
        return
    }

    validation := validator.New()
    err := validation.Struct(request)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
        json.NewEncoder(w).Encode(response)
        log.Println("Validation error:", err) // Log pesan kesalahan
        return
    }

    class := models.Class{
        Name    : request.Name    ,
        Capacity: request.Capacity,
        AdminID : 1               ,
    }
	
    data, err := h.ClassRepository.CreateClass(class)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
        json.NewEncoder(w).Encode(response)
        log.Println("Error creating class:", err) // Log pesan kesalahan
        return
    }

    w.WriteHeader(http.StatusOK)
    response := dto.SuccessResult{Code: http.StatusOK, Data: data}
    json.NewEncoder(w).Encode(response)
}

func (h *handlerClass) UpdateClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	request := new(models.UpdateClass)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	class, _ := h.ClassRepository.GetClassUpdate(id)

	if request.Name != "" {
		class.Name = request.Name
	}
	if request.Capacity != 0 {
		class.Capacity = request.Capacity
	}

	class.UpdatedAt = time.Now()

	data, err := h.ClassRepository.UpdateClass(class)
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

func (h *handlerClass) DeleteClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	class, err := h.ClassRepository.GetClass(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.ClassRepository.DeleteClass(class)
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
