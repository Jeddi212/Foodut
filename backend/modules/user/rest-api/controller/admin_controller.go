package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	srvc "github.com/Foodut/backend/modules/user/domain/service"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	rspn "github.com/Foodut/backend/responses"
)

func PostAdmin(writer http.ResponseWriter, req *http.Request) {

	// Decode JSON
	var postUserDto dto.PostUser
	err := json.NewDecoder(req.Body).Decode(&postUserDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Send DTO to service
	result := srvc.MapToAdmin(postUserDto)

	// Set response
	var response rspn.Response
	if result.Error == nil {
		response.Response_201()
	} else {
		fmt.Println(result.Error)
		response.Response_400("")
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
