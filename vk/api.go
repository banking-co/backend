package vk

import (
	"github.com/SevereCloud/vksdk/v3/api"
	"os"
)

var Api *api.VK

func Init() {
	token, tokenExist := os.LookupEnv("APP_SERVICE_KEY")

	if !tokenExist {
		panic("VKApp service key is nil")
		return
	}

	Api = api.NewVK(token)
}
