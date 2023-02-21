package controllers

import (
	"net/http"

	"github.com/QuocBao92/go-sample/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Its home page")

}
