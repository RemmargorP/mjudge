package web

import (
	"github.com/RemmargorP/memjudge/models"
	"log"
	"net/http"
)

func (wi *WebInstance) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["title"] = TITLE
	data["webinstanceid"] = wi.Id
	data["webinfo"] = map[string]interface{}{"isSignUp": true}
	data["additionalJS"] = []string{"signup.js", "login.js"}

	user := models.GetUserFromCookie(r, wi.DB)

	if user.IsLoggedIn() {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	models.GatherUserInfo(user, &data)

	wi.ParseCookieReason(w, r, &data)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := wi.Templates.ExecuteTemplate(w, "signup.html", data)
	if err != nil {
		log.Println(err)
	}
}
