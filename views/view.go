package views

import (
	"encoding/json"
	"fmt"
	"html/template"
	db "login-signup/database"
	"login-signup/model"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var dbConn, _ = db.GetPostGreSql()

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IndexPage(w http.ResponseWriter, r *http.Request) {

	tpl, err := template.ParseFiles("web/index.html")
	if err != nil {
		fmt.Println(err)
	}
	tpl.Execute(w, nil)
}
func LoginPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tpl, err := template.ParseFiles("web/index.html")
		if err != nil {
			fmt.Println(err)
		}
		login := model.ActiveToggle{Login: "checked"}
		tpl.Execute(w, login)
	}

	if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		user := model.User{
			Username: username,
			Password: password}
		var k model.User
		fmt.Println(dbConn.Model(&user).First(&k))
		fmt.Println(k)
		passMatch := CheckPasswordHash(password, k.Password)

		var val interface{}

		if !passMatch {
			val = model.ErrorValue{Error: "username or password is incorrect "}
		} else {
			val = model.Success{Status: "Successfully Logged in"}
		}

		js, err := json.Marshal(val)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	}

}
func RegisterPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tpl, err := template.ParseFiles("web/index.html")
		if err != nil {
			fmt.Println(err)
		}
		signUp := model.ActiveToggle{Login: "", Signup: "checked"}
		tpl.Execute(w, signUp)
	}

	if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		emailAddress := r.Form.Get("email")
		if username == "" || password == "" || emailAddress == "" {
			http.Error(w, "bad request", http.StatusUnprocessableEntity)
			return
		}

		var k model.User
		var val interface{}
		dbConn.Where("username = ? or email = ?", username, emailAddress).First(&k)
		if k.Username == username {
			val = model.ErrorValue{Error: "Username already exist"}
		} else if k.Email == emailAddress {
			val = model.ErrorValue{Error: "User already registered with this email"}

		} else {
			Hashed, err := HashPassword(password)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			user := model.User{
				Username: username,
				Password: Hashed,
				Email:    emailAddress}
			dbConn.Create(&user)
			val = model.Success{Status: "Successfully registered"}
		}
		js, err := json.Marshal(val)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	}
}
