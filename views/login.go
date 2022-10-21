package views

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/nabalmes/time_managementv2/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LogInHander(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	fmt.Println("-----LogInHander-----")
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Database connection
		dsn := "root:Allen is Great 200%@tcp(127.0.0.1:3306)/time_managementv2?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Println("Faied to Connect to the Database ", err)
		}

		var modelusers []models.Users
		// SELECT * FROM users WHERE username = ? and password = ?;
		rows := db.Where("username = ? and password = ?", username, password).Find(&modelusers)

		if rows.RowsAffected > 0 {
			db.Exec("USE time_managementv2;")
			db.Debug().Exec("UPDATE users SET time_in = ?, time_in_string = ? WHERE username = ?", time.Now(), time.Now().Format("Mon, 02 Jan 2006 03:04:05 PM"), username)

			expires := time.Now().AddDate(1, 0, 0)
			ck := http.Cookie{
				Name:    "username",
				Path:    "/",
				Expires: expires,
			}
			ck.Value = username

			http.SetCookie(w, &ck)
			http.Redirect(w, r, "/dashboard/", http.StatusSeeOther)
		}
	}

	data["Title"] = "Log in | Time Managementv2"
	tmpl := template.Must(template.ParseFiles("./templates/login.html"))
	tmpl.Execute(w, data)
}
