package middleware

import "net/http"

// SocketAuthMiddleware Auth
func SocketAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r * http.Request) {
		tokenStr := r.URL.Query().Get("token")
		if (tokenStr == "") {
			http.Error(w, "Token is Required", http.StatusUnauthorized)
			return
		}

		claims, err := ValidateJWT(tokenStr)
		if err != nil || claims == nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		// add user id to context if needed
		r.Header.Set("UserID", string(claims.UserID))
		next(w,r)
	}
}