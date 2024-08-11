package business

import (
	"encoding/json"
	"fmt"
	"github.com/SevereCloud/vksdk/v3/vkapps"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
	"rabotyaga-go-backend/dto"
	"rabotyaga-go-backend/models"
	"rabotyaga-go-backend/mysqldb"
	"rabotyaga-go-backend/types"
	"rabotyaga-go-backend/utils"
)

func GetStaff(e types.EventType, conn net.Conn, code ws.OpCode, vkParams *vkapps.Params, data json.RawMessage) {
	var db = mysqldb.DB

	rD, err := utils.UnmarshalData[dto.RequestBusinessStaffGet](data)
	if err != nil {
		fmt.Println(err)
		return
	}

	staff, err := models.GetBusinessStaffByBusinessId(db, rD.BusinessId)

	resData, err := utils.MarshalData[dto.ResponseBusinessStaffGet](e, &dto.ResponseBusinessStaffGet{
		BusinessID:    rD.BusinessId,
		BusinessStaff: staff,
	})
	if err != nil {
		return
	}

	err = wsutil.WriteServerMessage(conn, code, resData)
	if err != nil {
		return
	}
}
