package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/internal/database/mysql"
	"server/internal/models"
)

var dbClient *sql.DB = mysql.DbClient

func HomePage(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Welcome to your Free Speech Social Media Network!")
}

func GetUsers(responseWriter http.ResponseWriter, request *http.Request) {
	result, err := dbClient.Query("select * from users;")

	if err != nil {
		panic(err.Error())
	}

	var users []models.UserDetails

	for result.Next() {
		var id int
		var user models.UserDetails

		err = result.Scan(&id, &user.UserName)

		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	jsonData, error := json.Marshal(users)

	if error != nil {
		panic(error.Error())
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(jsonData)
}

func GetUser(responseWriter http.ResponseWriter, request *http.Request) {
	bytes, err := io.ReadAll(request.Body)

	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	var userDetails models.UserDetailsQuery

	error := json.Unmarshal(bytes, &userDetails)

	if error != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := dbClient.Query("select * from users where username = ?;", userDetails.UserId)

	if err != nil {
		panic(err.Error())
	}

	var user models.UserDetails

	for result.Next() {
		var id int

		err = result.Scan(&id, &user.UserName)

		if err != nil {
			panic(err)
		}
	}

	jsonData, errr := json.Marshal(user)

	if errr != nil {
		panic(errr.Error())
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(jsonData)

}

func CreateUser(responseWriter http.ResponseWriter, request *http.Request) {
	user, err := io.ReadAll(request.Body)

	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	var userDetails models.CreateUser

	errr := json.Unmarshal(user, &userDetails)

	if errr != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	response, errorr := dbClient.Query("select * from users where username = ?;", userDetails.UserName)

	if errorr != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	for response.Next() {
		var id int
		var usernameCheck string

		error := response.Scan(&id, &usernameCheck)

		if error != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(usernameCheck) > 0 {
			responseWriter.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(responseWriter, "Username already exists!")
			return
		}
	}

	query := " Insert into users (`id`, `username`) values (default, ?); "

	_, error := dbClient.ExecContext(context.Background(), query, userDetails.UserName)

	if error != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(responseWriter, "User Created Successfully!")
}

func Posts(responseWriter http.ResponseWriter, request *http.Request) {

	const text = "monisha"

	query := "Insert into users (`id`, `username`) values (default, ?);"

	_, err := dbClient.ExecContext(context.Background(), query, text)

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(responseWriter, "Post Created Successfully!")
}

func Likes(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Likes")
}

func Comments(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Comments")
}
