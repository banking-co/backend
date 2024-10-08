package server

import (
	"fmt"
	"github.com/SevereCloud/vksdk/v3/vkapps"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"log"
	"net/http"
	"os"
	"rabotyaga-go-backend/entities"
	"rabotyaga-go-backend/models"
	"rabotyaga-go-backend/mysqldb"
	"rabotyaga-go-backend/types"
	"rabotyaga-go-backend/utils"
	"strconv"
	"time"
)

type OnCallbackFunc = func(req *entities.Request)

type Server struct {
	events           map[types.EventType][]OnCallbackFunc
	biggestEventSize uint8
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

	for e := range s.events {
		eventLength := uint8(len(e))

		if s.biggestEventSize < eventLength {
			s.biggestEventSize = eventLength
		}
	}

	err := http.ListenAndServe(":3001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// routing

		//if rest
		// TODO: add payment routing

		// if ws
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
			if time.Now().Sub(time.Unix(vkTs, 0)) >= 60*time.Minute {
				utils.SendError(w, "The authorization period has expired", http.StatusForbidden)
				return
			}
		}

		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Panicln("Upgrade HTTP error: ", err)
		}

		go func() {
			defer func() {
				if err := conn.Close(); err != nil {
					logInfo := fmt.Sprintf("[ %d ]: connection closed, address: %s", time.Now().Unix(), conn.RemoteAddr())
					fmt.Println(logInfo)
				}
			}()

			user, _, err := models.GetUserByUsername(mysqldb.DB, vkParams.VkUserID)
			if err != nil {
				utils.SendError(w, "GetUser executed with an error", http.StatusUnauthorized)
				return
			}
			if user == nil {
				utils.SendError(w, "User is nil", http.StatusUnauthorized)
				return
			}

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					utils.SendError(w, err.Error(), http.StatusBadRequest)
					break
				}
				if len(msg) >= 100 {
					utils.SendError(w, "Message is very big", http.StatusBadRequest)
					break
				}

				err = conn.SetDeadline(time.Now().Add(20 * time.Second))
				if err != nil {
					utils.SendError(w, "Connection SetDeadline executed with an error", http.StatusBadRequest)
					break
				}

				message, err := utils.UnmarshalData[types.EventParams](msg)
				if err != nil {
					utils.SendError(w, "Data is falsified", http.StatusBadRequest)
					break
				}

				if uint8(len(message.Event)) > s.biggestEventSize {
					utils.SendError(w, "Event is very big", http.StatusBadRequest)
					break
				}

				if message.Event != types.EventStartApp && len(user.Bans) >= 1 {
					utils.SendError(w, "User is blocked", http.StatusUnauthorized)
					break
				}

				if cbs, ok := s.events[message.Event]; ok {
					req := entities.Request{
						Conn: conn,
						Op:   op,

						Event:       message.Event,
						Data:        message.Data,
						StartParams: vkParams,
					}

					for _, fc := range cbs {
						fc(&req)
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
		fmt.Printf("[%v]: launched \r\n", e)
	}

	s.events[e] = append(s.events[e], cb)
}
