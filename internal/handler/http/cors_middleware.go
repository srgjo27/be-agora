package http

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupCORS mengonfigurasi CORS middleware untuk aplikasi
func SetupCORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		// Daftar origin yang diizinkan
		AllowOrigins: []string{
			"http://localhost:3000",  // React development server
			"http://localhost:5173",  // Vite development server
			"http://localhost:8080",  // Local backend
			"http://localhost:4200",  // Angular development server
			"https://yourdomain.com", // Production domain
		},

		// Method HTTP yang diizinkan
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},

		// Header yang diizinkan
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-Requested-With",
			"X-CSRF-Token",
			"Content-Length",
		},

		// Header yang di-expose ke client
		ExposeHeaders: []string{
			"Content-Length",
			"X-Total-Count",
			"X-Page",
			"X-Per-Page",
		},

		// Mengizinkan credentials (cookies, authorization headers)
		AllowCredentials: true,

		// Cache preflight request selama 12 jam
		MaxAge: 12 * time.Hour,
	})
}

// SetupCORSForDevelopment mengonfigurasi CORS dengan pengaturan lebih permisif untuk development
func SetupCORSForDevelopment() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true, // HANYA untuk development!
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"}, // Allow all headers
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
