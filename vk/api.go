package vk

import (
	"github.com/SevereCloud/vksdk/v3/api"
	"os"
)

var Api *api.VK

func Init() {
	token, tokenExist := os.LookupEnv("APP_SERVICE_KEY")

	if !tokenExist {
		panic("VK App service key is nil")
	}

	Api = api.NewVK(token)
}
