package main

import (
	"net/http"
	"os"

	"github.com/hanhtrang256/spoonie-testing/application"
)

/**

**/

func main() {
	http.HandleFunc("/", application.DisplayHomePage)
	http.HandleFunc("/users", application.DisplayUsers)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe("0.0.0.0:"+port, nil)
	// http.ListenAndServe("localhost:8080", nil)
	// conn, err := app.ConnectPostgres("postgres", "trang256", "localhost", "5432", "spooniedatabase")
	// if err != nil {
	// 	log.Println("Cannot connect to database!")
	// 	log.Fatal()
	// }
	// authen_bool, authen_err := app.UserAuthentication(conn, "minhi1' OR '1=1", "uia25")
	// if authen_err != nil {
	// 	log.Println("Something wrong with authentication!")
	// 	log.Fatal()
	// } else {
	// 	if !authen_bool {
	// 		fmt.Println("Wrong password idiot")
	// 	} else {
	// 		fmt.Println("You're in!!!")
	// 	}
	// }
}
