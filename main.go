package main

import (
	"os"
	"rabotyaga-go-backend/api/realtime/balance"
	"rabotyaga-go-backend/api/realtime/base"
	"rabotyaga-go-backend/api/realtime/business"
	"rabotyaga-go-backend/api/realtime/user"
	"rabotyaga-go-backend/database"
	"rabotyaga-go-backend/server"
	"rabotyaga-go-backend/types"

	"github.com/joho/godotenv"
)

func init() {
	if envErr := godotenv.Load(".env.local"); envErr != nil {
		panic(envErr)
	}

	dbPassword, dbPasswordExist := os.LookupEnv("DB_PASSWORD")
	if dbPasswordExist {
		database.Init(database.Options{
			Database:       "banking",
			Username:       "backend",
			Host:           "localhost",
			MaxConnections: "10",
			Port:           "3306",
			Password:       dbPassword,
		})

		database.Migrate()
	} else {
		panic("DB_PASSWORD not set in environment")
	}
}

func main() {
	s := server.Init()

	s.OnSocket(types.EventPing, base.Ping)
	s.OnSocket(types.EventStartApp, base.StartApp)
	s.OnSocket(types.EventGetBusiness, business.Get)
	s.OnSocket(types.EventGetPrimaryBusiness, business.Get)
	s.OnSocket(types.EventUserGet, user.Get)
	s.OnSocket(types.EventBalanceGet, balance.Get)

	s.Listen()
}
