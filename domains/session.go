package domains

import (
	"fmt"
	"net/http"
	"strings"
)

const cookieName = "avosession"

func HandleSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer next.ServeHTTP(w, r)

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

			if !strings.Contains(r.Host, "localhost") {
				cookie.Domain = fmt.Sprintf(".%s", r.Host)
				cookie.Secure = true
			}

			http.SetCookie(w, &cookie)
		}
	})
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
