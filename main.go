package main

import (
	"rabotyaga-go-backend/api/realtime/balance"
	"rabotyaga-go-backend/api/realtime/base"
	"rabotyaga-go-backend/api/realtime/user"
	"rabotyaga-go-backend/mysql/database"
	"rabotyaga-go-backend/server"
	"rabotyaga-go-backend/types"
)

func main() {
	database.New(database.Options{
		Database:       "dev",
		Username:       "root",
		Host:           "localhost",
		MaxConnections: 10,
		Port:           3306,
		Password:       "admin",
	})

	s := server.Init()

	s.On(types.RequestPing, base.Ping)
	s.On(types.RequestStartApp, base.StartApp)

	s.On(types.RequestUserGet, user.Get)
	s.On(types.RequestUserGet, user.Get)

	s.On(types.RequestBalanceGet, balance.Get)

	s.Listen()
}
