package users

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"../../entity"
	"../common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

type successCreate struct {
	Status string    `json:"status"`
	ID     uuid.UUID `json:"id"`
}

type userResponse struct {
	ID    uuid.UUID `json:"ID"`
	Name  string    `json:"Name"`
	Login string    `json:"Login"`
}

//GetUsersHandler is a handler for getting all Users from DB
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := entity.GetUsers()
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get users", err)
		return
	}

	var uResponse userResponse
	var uResponses []userResponse
	for _, u := range users {
		uResponse.ID = u.ID
		uResponse.Name = u.Name
		uResponse.Login = u.Login
		uResponses = append(uResponses, uResponse)
	}

	common.RenderJSON(w, r, uResponses)
}

//GetUserHandler is a handler for getting User from DB by ID
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idUser, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Converting ID from URL", err)
		return
	}
	user, err := entity.GetUserByID(idUser)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find user with such ID", err)
		return
	}

	var uResponse userResponse
	uResponse.ID = user.ID
	uResponse.Name = user.Name
	uResponse.Login = user.Login

	common.RenderJSON(w, r, uResponse)
}

//AddUserHandler is a handler for creating User
func AddUserHandler(w http.ResponseWriter, r *http.Request) {

	var newUser entity.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't parse POST Body", err)
		return
	}

	valid, errMessage := IsValid(newUser)
	if !valid {
		log.Print(errMessage)
	}

	id, err := entity.AddUser(newUser)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add this user", err)
		return
	}
	common.RenderJSON(w, r, successCreate{Status: "201 Created", ID: id})
}

//IsValid checks if password is valid
func IsValid(user entity.User) (bool, string) {
	errMessage := ""
	var checkPass = regexp.MustCompile(`^[[:graph:]]*$`)
	var checkName = regexp.MustCompile(`^[A-Z]{1}[a-z]+$`)
	var checkLogin = regexp.MustCompile(`^[[:graph:]]*$`)
	var validPass, validName, validLogin bool
	if len(user.Password) >= 8 && checkPass.MatchString(user.Password) {
		validPass = true
	} else {
		errMessage += "Invalid Password"
	}
	if checkName.MatchString(user.Name) && len(user.Name) < 15 {
		validName = true
	} else {
		errMessage += " Invalid Name"
	}
	if checkLogin.MatchString(user.Login) && len(user.Login) < 15 {
		validLogin = true
	} else {
		errMessage += " Invalid Login"
	}
	return validName && validLogin && validPass, errMessage
}

//DeleteUserHandler is a handler for deleting User from DB
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	idUser, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong user ID (can't convert string to int)", err)
		return
	}
	err = entity.DeleteUser(idUser)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't delete this user", err)
		return
	}
	common.RenderJSON(w, r, nil)
}
