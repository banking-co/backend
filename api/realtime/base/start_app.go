package base

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SevereCloud/vksdk/v3/vkapps"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"gorm.io/gorm"
	"log"
	"net"
	"rabotyaga-go-backend/database"
	"rabotyaga-go-backend/models"
	"rabotyaga-go-backend/responseData"
	"rabotyaga-go-backend/types"
	"rabotyaga-go-backend/utils"
	"strconv"
)

func StartApp(conn net.Conn, code ws.OpCode, vkParams *vkapps.Params, data json.RawMessage) {
	var user models.User
	var db = database.DB
	var uid = "id" + strconv.Itoa(vkParams.VkUserID)

	if err := db.Where("username = ?", uid).First(&user).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			user := models.User{
				Username: uid,
			}
			if err := db.Create(&user).Error; err != nil {
				log.Fatal("Failed to create user:", err)
			}
			fmt.Println("User created successfully:", user)
		} else {
			log.Fatal("Failed to query user:", err)
		}
	} else {
		fmt.Println("User already exists:", user)
	}

	resData, err := utils.MarshalData[responseData.ResponseStartApp](types.EventStartApp, &responseData.ResponseStartApp{
		User:     responseData.UserWrap(user),
		IsLogged: true,
	})

	err = wsutil.WriteServerMessage(conn, code, resData)
	if err != nil {
		return
	}
}
