package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"wb2/api/models"
	"wb2/api/responses"
	"wb2/api/utils/formaterror"
)

func (server *Server) CreateRole(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	role := models.Role{}
	err = json.Unmarshal(body, &role)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	role.Prepare()
	err = role.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	roleCreated, err := role.SaveRole(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, roleCreated.ID))
	responses.JSON(w, http.StatusCreated, roleCreated)
}

func (server *Server) GetRoles(w http.ResponseWriter, r *http.Request) {

	role := models.Role{}

	roles, err := role.FindAllRoles(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, roles)
}

func (server *Server) GetRole(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	rid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	role := models.Role{}
	roleData, err := role.FindRoleByID(server.DB, uint32(rid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, roleData)
}

func (server *Server) UpdateRole(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	rid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	role := models.Role{}
	err = json.Unmarshal(body, &role)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	
	role.Prepare()
	err = role.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedRole, err := role.UpdateARole(server.DB, uint32(rid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedRole)
}

func (server *Server) DeleteRole(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	role := models.Role{}

	rid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = role.DeleteARole(server.DB, uint32(rid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", rid))
	responses.JSON(w, http.StatusNoContent, "")
}