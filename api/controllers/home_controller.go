package controllers

import (
	"github.com/mgw2168/go-jwt-example/api/responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}