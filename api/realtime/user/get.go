package user

import (
	"encoding/json"
	"fmt"
	"github.com/SevereCloud/vksdk/v3/vkapps"
	"github.com/gobwas/ws"
	"gorm.io/gorm"
	"log"
	"net"
	"rabotyaga-go-backend/database"
	"rabotyaga-go-backend/models"
	"strconv"
)

func Get(conn net.Conn, code ws.OpCode, startParams *vkapps.Params, data json.RawMessage) {
	var user models.User
	var db = database.DB

	if err := db.Where("username = ?", startParams.VkUserID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			user := models.User{
				Username: "id" + strconv.Itoa(startParams.VkUserID),
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

	fmt.Println("user")
	//reqData, err := utils.UnmarshalData[structures.RequestUserGet](data)

	//if err == nil {
	//user, _ := mysql.USER_GET_BY_UID(reqData.UserId)

	//resData, err := utils.MarshalData[structures.ResponseUserGet](types.EventUserGet, &structures.ResponseUserGet{User: *user})
	//if err != nil {
	//	return
	//}
	//
	//err = wsutil.WriteServerMessage(conn, code, resData)
	//if err != nil {
	//	return
	//}
	//}
}
