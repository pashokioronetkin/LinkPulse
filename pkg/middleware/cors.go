package middleware

import "net/http"

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			next.ServeHTTP(w, r)
			return
		}
		header := w.Header()
		header.Set("Access-Control-Allow-Origin", origin)      // Allow all origins
		header.Set("Access-Control-Allow-Credentials", "true") // Allow cookies to be sent

		if r.Method == http.MethodOptions {
			header.Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, HEAD, PATCH")
			header.Set("Access-Control-Allow-Headers", "authorization, content-type, content-length")
			header.Set("Access-Control-Max-Age", "86400")
		} // Add more headers as needed

		next.ServeHTTP(w, r)
	})
}
