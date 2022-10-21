package views

import (
	"net/http"
	"text/template"

	"github.com/nabalmes/time_managementv2/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/delete.html"))
	data := map[string]interface{}{}
	database_connection := "root:Allen is Great 200%@tcp(localhost:3306)/time_managementv2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(database_connection), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to the Database")
	}
	username := r.FormValue("username")

	users := []models.Users{}
	if r.Method == "POST" && username != "" {
		db.Where("username = ?", username).Delete(&users)
		data["Users"] = users

		http.Redirect(w, r, "/dashboard/", http.StatusSeeOther)
	}
	data["Title"] = "DELETE | Time Managementv2"
	tmpl.Execute(w, data)
}
