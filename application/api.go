package application

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserJSON struct {
	Id         int       `json:"id"`
	Password   string    `json:"password"`
	Username   string    `json:"username"`
	Role       string    `json:"role"`
	Created_at time.Time `json:"date created"`
	Ketqua     string    `json:"ketqua"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginReq LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	connStr := "user=youruser dbname=yourdb password=yourpassword host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user UserJSON
	err = db.QueryRow(`SELECT id, username, password, role, created_at FROM users WHERE username = $1`, loginReq.Username).
		Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.Created_at)

	if err == sql.ErrNoRows {
		json.NewEncoder(w).Encode(map[string]string{
			"ketqua": "tai khoan khong ton tai",
		})
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if user.Password != loginReq.Password {
		json.NewEncoder(w).Encode(map[string]string{
			"ketqua": "sai mat khau",
		})
		return
	}

	// Nếu đúng password → trả về thông tin user + ketqua
	user.Ketqua = "thanh cong"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
