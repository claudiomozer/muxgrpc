package controllers

import (
	"encoding/json"
	"net/http"

	controllers_protocols "github.com/claudiomozer/muxgrpc/internal/presentation/controllers/http/protocols"
	"github.com/gorilla/mux"
)

type FindUserByDocumentController struct {
	find_user_by_document controllers_protocols.FindUserByDocumentUseCase
}

func NewFindUserByDocumentController(
	find_user_by_document controllers_protocols.FindUserByDocumentUseCase,
) *FindUserByDocumentController {
	return &FindUserByDocumentController{
		find_user_by_document: find_user_by_document,
	}
}

func (controller *FindUserByDocumentController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	document := vars["document"]

	if document == "" {
		http.Error(w, "document it is a required field", http.StatusBadRequest)
		return
	}

	userDTO, err := controller.find_user_by_document.FindByDocument(document)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userJSON, err := json.Marshal(&userDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}
