package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server-skeleton/api/page/DTO"
	"server-skeleton/pkg/format"
)

type RequestPage struct {
	name string
}

func Get(w http.ResponseWriter, r *http.Request) {
	var page RequestPage

	if err := json.NewDecoder(r.Body).Decode(&page); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	validationErrors := format.ValidateFormat(page)

	errCode := http.StatusOK

	if len(validationErrors) > 0 {
		errCode = http.StatusBadRequest
	}

	result := format.ResponseFormat(errCode, DTO.ResponsePageDTO{
		Name: "Some Name",
	}, validationErrors)

	w.WriteHeader(errCode)
	fmt.Print(w, result)
}
