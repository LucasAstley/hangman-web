package back

import (
	"net/http"
	"time"
)

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	/*
		This function is used to set the player name and difficulty cookies
		and redirect to the game page
	*/
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		difficulty := r.FormValue("difficulty")
		http.SetCookie(w, &http.Cookie{
			Name:     "playerName",
			Value:    name,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		})
		http.SetCookie(w, &http.Cookie{
			Name:     "difficulty",
			Value:    difficulty,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		})
		http.Redirect(w, r, "/game", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
