package application

import (
	"encoding/json"
	"net/http"
	"time"
)

type UserJSON struct {
	Id         string    `json:"id"`
	Password   string    `json:"password"`
	Username   string    `json:"username"`
	Role       string    `json:"role"`
	Created_at time.Time `json:"date created"`
}

type TmpJSON struct {
	Hello string `json:"hello"`
}

func DisplayHomePage(writer http.ResponseWriter, request *http.Request) {
	// users := []UserJSON{{Id: "6", Username: "minhi1", Role: "Admin", Created_at: time.Now()},
	// 	{Id: "25", Username: "TR", Role: "Admin", Created_at: time.Now()}}

	// jsonData, err := json.Marshal(users)
	// if err != nil {
	// 	http.Error(writer, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	writer.Header().Set("Content-Type", "application/json")

	// _, err = writer.Write(jsonData)
	// if err != nil {
	// 	http.Error(writer, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	tmp := TmpJSON{Hello: "by minhi1"}
	jsonData, err := json.Marshal(tmp)
	if err != nil {

	}
	_, err2 := writer.Write(jsonData)
	if err2 != nil {

	}
}

func DisplayUsers(writer http.ResponseWriter, request *http.Request) {
	users := []UserJSON{{Id: "6", Username: "minhi1", Role: "Admin", Created_at: time.Now()},
		{Id: "25", Username: "TR", Role: "Admin", Created_at: time.Now()}}

	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	_, err = writer.Write(jsonData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
