package controllers

import (
	"net/http"

	"wb2/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To Website Builder API")

}