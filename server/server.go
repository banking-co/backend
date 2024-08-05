package server

import (
	"encoding/json"
	"fmt"
	"github.com/SevereCloud/vksdk/v3/vkapps"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"net"
	"net/http"
	"os"
	"rabotyaga-go-backend/database"
	"rabotyaga-go-backend/models"
	"rabotyaga-go-backend/responseData"
	"rabotyaga-go-backend/types"
	"rabotyaga-go-backend/utils"
	"strconv"
	"time"
)

type OnCallbackFunc = func(conn net.Conn, op ws.OpCode, sign *vkapps.Params, data json.RawMessage)

type Server struct {
	events map[types.EventType][]OnCallbackFunc
}

func Init() *Server {
	s := Server{
		events: make(map[types.EventType][]OnCallbackFunc),
	}

	fmt.Println("Server started!")
	return &s
}

func (s *Server) Listen() {
	mode, modeExist := os.LookupEnv("APP_ENV")

	err := http.ListenAndServe(":3001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validate, err := vkapps.ParamsVerify(r.URL.String(), "8PogTmDn5uru9WPdXuup")
		if !validate {
			utils.SendError(w, "Signature verification error", http.StatusForbidden)
			return
		}

		if validate {
			fmt.Printf("Connection at [%s] \r\n", time.Now())
		}

		if err != nil {
			utils.SendError(w, "Signature verification error", http.StatusForbidden)
			return
		}

		vkParams, err := vkapps.NewParams(r.URL)
		if err != nil {
			utils.SendError(w, "Error converting the starting parameters", http.StatusForbidden)
			return
		}

		vkTs, err := strconv.ParseInt(vkParams.VkTs, 10, 64)
		if err != nil {
			utils.SendError(w, "Error getting the launch date", http.StatusForbidden)
			return
		}

		if !modeExist || mode != "development" {
			if time.Now().Sub(time.Unix(vkTs, 0)) >= 10*time.Minute {
				utils.SendError(w, "The authorization period has expired", http.StatusForbidden)
				return
			}
		}

		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Panicln("Upgrade HTTP error")
			return
		}

		go func() {
			defer func() {
				if err := conn.Close(); err != nil {
					logInfo := fmt.Sprintf("[ %d ]: connection closed, address: %s", time.Now().Unix(), conn.RemoteAddr())
					fmt.Println(logInfo)
				}
			}()

			user, err := models.GetUserByUsername(database.DB, vkParams.VkUserID)
			if err != nil {
				utils.SendError(w, "User is nil", http.StatusUnauthorized)
				return
			}

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					break
				}

				message, err := utils.UnmarshalData[responseData.EventParams](msg)
				if err != nil {
					break
				}

				if len(message.Event) >= 20 {
					utils.SendError(w, "Event is very big", http.StatusBadRequest)
					break
				}

				if message.Event != types.EventStartApp && len(user.Bans) >= 1 {
					utils.SendError(w, "User is blocked", http.StatusUnauthorized)
					break
				}

				if cbs, ok := s.events[message.Event]; ok {
					for _, fc := range cbs {
						fc(conn, op, vkParams, message.Data)
					}
				}
			}
		}()
	}))

	if err != nil {
		log.Panicln("Server error starting")
	}
}

func (s *Server) OnSocket(e types.EventType, cb OnCallbackFunc) {
	if _, ok := s.events[e]; !ok {
		s.events[e] = []OnCallbackFunc{}
	}

	s.events[e] = append(s.events[e], cb)
}
