package config

import "os"

var (
	PASSWORD_SALT = os.Getenv("PASSWORD_SALT")
	JWT_SECRET    = os.Getenv("JWT_SECRET")

	POSTGRES_HOST     = os.Getenv("POSTGRES_HOST")
	POSTGRES_USER     = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB       = os.Getenv("POSTGRES_DB")
)
