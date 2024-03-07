package middleware

import (
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {

	return &AuthMiddleware{Handler: handler}
}
func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	authRequiredRoutes := map[string]bool{
		"/list_product":   true,
		"/update_stock":   true,
		"/insert_product": true,
		// Add more routes here as needed
	}

	// Check if the requested route requires authentication
	if _, ok := authRequiredRoutes[request.URL.Path]; ok {
		authToken := request.Header.Get("Authorization")
		if authToken == "" {
			// If no token is provided for an authenticated route, respond with an error
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}

		//if check == true {
		//	middleware.Handler.ServeHTTP(writer, request)
		//} else {
		//	http.Error(writer, "Invalid token", http.StatusUnauthorized)
		//	return
		//}
	}
	middleware.Handler.ServeHTTP(writer, request)

}
