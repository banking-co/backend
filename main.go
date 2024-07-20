package main

import (
	"os"
	"rabotyaga-go-backend/api/realtime/balance"
	"rabotyaga-go-backend/api/realtime/base"
	"rabotyaga-go-backend/api/realtime/user"
	"rabotyaga-go-backend/mysql/database"
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
		dbErr := database.New(database.Options{
			Database:       "banking",
			Username:       "app",
			Host:           "147.45.184.220",
			MaxConnections: 10,
			Port:           3306,
			Password:       dbPassword,
		})

		if dbErr != nil {
			panic(dbErr)
		}
	}
}

func main() {
	s := server.Init()

	s.OnSocket(types.EventPing, base.Ping)
	s.OnSocket(types.EventStartApp, base.StartApp)
	s.OnSocket(types.EventUserGet, user.Get)
	s.OnSocket(types.EventBalanceGet, balance.Get)

	s.Listen()
}
