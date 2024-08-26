package main

import (
	"rabotyaga-go-backend/api/realtime/balance"
	"rabotyaga-go-backend/api/realtime/base"
	"rabotyaga-go-backend/api/realtime/bonus"
	"rabotyaga-go-backend/api/realtime/business"
	"rabotyaga-go-backend/api/realtime/user"
	"rabotyaga-go-backend/mysqldb"
	"rabotyaga-go-backend/redisdb"
	"rabotyaga-go-backend/server"
	"rabotyaga-go-backend/types"
	"rabotyaga-go-backend/vk"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env.local"); err != nil {
		panic(err)
	}

	mysqldb.Init()
	redisdb.Init()
	vk.Init()
}

func main() {
	s := server.Init()

	// server
	s.OnSocket(types.EventPing, base.Ping)

	// start
	s.OnSocket(types.EventStartApp, base.StartApp)

	// user
	s.OnSocket(types.EventUserGet, user.Get)

	// bonus
	s.OnSocket(types.EventBonusGet, bonus.Get)
	s.OnSocket(types.EventBonusReceive, bonus.Receive)

	// balance
	s.OnSocket(types.EventBalanceGet, balance.Get)

	// business
	s.OnSocket(types.EventGetBusiness, business.Get)
	s.OnSocket(types.EventGetBusinessStaff, business.GetStaff)
	s.OnSocket(types.EventGetBusinessStaffRecruit, business.RecruitGetStaff)
	s.OnSocket(types.EventBuyBusinessStaffRecruit, business.RecruitBuyStaff)

	s.Listen()
}
