package config

import (
	"os"
	"strings"
)

var defaultAllowedOrigins = []string{
	"http://localhost:5173",
	"http://127.0.0.1:5173",
	"https://frontend-q3mk.vercel.app",
	"https://backend-kssq.onrender.com",
}

func GetAllowedOrigins() []string {
	origins := defaultAllowedOrigins

	if envOrigins := os.Getenv("ALLOWED_ORIGINS"); envOrigins != "" {
		origins = strings.Split(envOrigins, ",")
	}

	allowedOrigins := make([]string, 0, len(origins))
	for _, origin := range origins {
		origin = strings.TrimSpace(origin)
		if origin != "" {
			allowedOrigins = append(allowedOrigins, origin)
		}
	}

	return allowedOrigins
}
