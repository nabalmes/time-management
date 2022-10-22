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

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/dashboard.html"))
	data := map[string]interface{}{}
	database_connection := "root:Allen is Great 200%@tcp(localhost:3306)/time_managementv2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(database_connection), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to the Database")
	}

	users := []models.Users{}
	db.Find(&users)

	// details := models.Users{}
	// details.Username = r.FormValue("username")

	// cookie, _ := r.Cookie("username")

	// cookie, _ := r.Cookie("username")
	// fmt.Print("COOKIE", cookie.Value)

	username := r.FormValue("username")
	if r.Method == "POST" {
		// fmt.Println("username", username)
		// _ = db.Exec("USE time_management;")
		// _ = db.Exec("UPDATE users SET time_out = ? WHERE username = ?", time.Now(), cookie.Value)
		// // rows := db.Exec("SELECT * FROM users WHERE username LIKE % = ?", username)
		// // fmt.Printf("rows %v", rows)

		// ck := http.Cookie{
		// 	Name:    "username",
		// 	Path:    "/",
		// 	Expires: time.Now(),
		// }
		// http.SetCookie(w, &ck)
		// http.Redirect(w, r, "/", http.StatusSeeOther)

		if username != "" {
			fmt.Printf("username %T: %v", username, username)

			username = "%" + username + "%"
			db.Where("username LIKE ?", username).Find(&users)
			// search := []models.Users{}
			// searchusers := db.Debug().Where("username LIKE ?", username).Find(&users)
			// data["Users"] = searchusers
			// fmt.Println("search ", searchusers)
			// tmpl.Execute(w, data)
			// http.Redirect(w, r, "/dashboard/", http.StatusSeeOther)
		}

		// else {
		// 	users := []models.Users{}
		// 	db.Find(&users)
		// 	data["Users"] = users
		// 	// fmt.Println("gege ", users)

		// }
	}
	// if username != "" {
	// 	fmt.Printf("username %T: %v", username, username)

	// 	username = "%" + username + "%"
	// 	if r.Method == "POST" {
	// 		search := []models.Users{}
	// 		searchusers := db.Debug().Where("username LIKE ?", username).Find(&search)
	// 		data["Users"] = searchusers
	// 		fmt.Println("search ", searchusers)
	// 		// tmpl.Execute(w, data)
	// 		// http.Redirect(w, r, "/dashboard/", http.StatusSeeOther)
	// 	} else {
	// 		users := []models.Users{}
	// 		db.Find(&users)
	// 		data["Users"] = users
	// 		// fmt.Println("gege ", users)

	// 	}
	// }
	// POST for cookie
	// if r.Method == "POST" {
	// ck := http.Cookie{
	// Name:    "Username",
	// Path:    "/",
	// Expires: time.Now(),
	// }
	// http.SetCookie(w, &ck)
	// http.Redirect(w, r, "/", http.StatusSeeOther)
	// }

	data["Users"] = users
	tmpl.Execute(w, data)
}

// func SearchUsername(user string) ([]Username, error) {
// 	rows, err := db.Query("SELECT * FROM users WHERE username = ?", user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var usernames []Username

// 	// Loop through rows, using Scan to assign column data to struct fields.
// 	for rows.Next() {
// 		var users Username
// 		if err := rows.Scan(&users.username); err != nil {
// 			return usernames, err
// 		}
// 		usernames = append(usernames, users)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return users, err
// 	}
// 	return users, nil
// }
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-----LogoutHandler-----")
	database_connection := "root:Allen is Great 200%@tcp(localhost:3306)/time_managementv2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(database_connection), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to the Database")
	}

	fmt.Println("-----LogoutHandler POST-----")
	cookie, _ := r.Cookie("username")
	_ = db.Exec("USE time_managementv2;")
	_ = db.Debug().Exec("UPDATE users SET time_out = ?, time_out_string = ? WHERE username = ?", time.Now(), time.Now().Format("Mon, 02 Jan 2006 03:04:05 PM"), cookie.Value)
	// rows := db.Exec("SELECT * FROM users WHERE username LIKE % = ?", username)
	// fmt.Printf("rows %v", rows)

	ck := http.Cookie{
		Name:    "username",
		Path:    "/",
		Expires: time.Now(),
	}

	http.SetCookie(w, &ck)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
