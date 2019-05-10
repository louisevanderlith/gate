package domains

import (
	"net/http"
	"strings"
)

const cookieName = "avosession"

func handleSession(r *http.Request, w http.ResponseWriter) {
	token := ""
	path := r.URL.RequestURI()

	if !strings.Contains(path, "?access_token") {
		return
	}

	_, token = removeToken(path)

	if len(token) > 0 {

		cookie := http.Cookie{
			Name:     cookieName,
			Path:     "/",
			Value:    token,
			HttpOnly: true,
			MaxAge:   0,
		}

		http.SetCookie(w, &cookie)
	}
}

func removeToken(url string) (string, string) {
	idx := strings.LastIndex(url, "?access_token")

	if idx == -1 {
		return url, ""
	}

	tokenIdx := strings.LastIndex(url, "=") + 1

	cleanURL := url[:idx]
	token := url[tokenIdx:]

	return cleanURL, token
}
