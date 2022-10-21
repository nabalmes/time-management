package views

import (
	"fmt"
	"net/http"
	"text/template"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{}

	// Database connection
	dsn := "root:Allen is Great 200%@tcp(127.0.0.1:3306)/time_managementv2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to Connect to the Database ", err)
	}

	// var modelusers []models.Users
	// // SELECT * FROM users WHERE username = ? and password = ?;
	// rows := db.Where("username = ? and password = ?", username, password).Find(&modelusers)

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		firstname := r.FormValue("first_name")
		lastname := r.FormValue("last_name")

		_ = db.Exec("USE time_managementv2;")
		_ = db.Exec("INSERT INTO users( username, password, first_name, last_name) VALUES(?,?,?,?)", username, password, firstname, lastname)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	data["Title"] = "Sign up | Time Managementv2"
	tmpl := template.Must(template.ParseFiles("./templates/signup.html"))
	tmpl.Execute(w, data)
}
