package main

import (
	"app/database"
	"app/middleware"
	"fmt"
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/joho/godotenv"
)

// func healthCheck(c echo.Context) error {
// 	database.Connect()
// 	sqlDB, _ := database.DB.DB()
// 	defer sqlDB.Close()
// 	err := sqlDB.Ping()
// 	if err != nil {
// 		return c.String(http.StatusInternalServerError, "データベース接続失敗")
// 	} else {
// 		return c.String(http.StatusOK, "200, OK")
// 	}
// }

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	database.Connect()
	sqlDB, _ := database.DB.DB()

	// router := http.NewServeMux()
	http.Handle("/api/public", String("hoge"))
	http.Handle("/api/private", middleware.EnsureValidToken()(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

			claims := token.CustomClaims.(*middleware.CustomClaims)
			if !claims.HasScope("read:messages") {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte(`{"message":"Insufficient scope."}`))
				return
			}
		}),
	))

	http.ListenAndServe(":8080", nil)

	defer sqlDB.Close()
}
