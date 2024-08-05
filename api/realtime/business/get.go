package business

import (
	"encoding/json"
	"fmt"
	"github.com/SevereCloud/vksdk/v3/vkapps"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"net"
	"rabotyaga-go-backend/database"
	"rabotyaga-go-backend/models"
	"rabotyaga-go-backend/responseData"
	"rabotyaga-go-backend/types"
	"rabotyaga-go-backend/utils"
)

func Get(conn net.Conn, code ws.OpCode, vkParams *vkapps.Params, data json.RawMessage) {
	var db = database.DB
	var bussiness *models.Business

	pData, err := utils.UnmarshalData[responseData.RequestBusinessGet](data)
	if err != nil {
		fmt.Println(err)
		return
	}

	if pData.UserId == nil && pData.BusinessId == nil {
		fmt.Println("UserID and Business ID is nil")
		return
	}

	if pData.BusinessId != nil {
		b, err := models.GetBusinessById(db, *pData.BusinessId)
		if err != nil {
			fmt.Println(err)
			return
		}

		bussiness = b
	}

	if pData.UserId != nil {
		b, err := models.GetBusinessByUserId(db, *pData.UserId)
		if err != nil {
			fmt.Println(err)
			return
		}

		bussiness = b
	}

	resData, err := utils.MarshalData[responseData.ResponseBusinessGet](types.EventGetBusiness, &responseData.ResponseBusinessGet{
		Business: responseData.BusinessWrap(bussiness),
	})
	if err != nil {
		return
	}

	err = wsutil.WriteServerMessage(conn, code, resData)
	if err != nil {
		return
	}
}
