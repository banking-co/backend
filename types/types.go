package types

type VkParams struct {
	VKUserId                  uint64 `json:"vk_user_id"`
	VKAppId                   uint64 `json:"vk_app_id"`
	VKIsAppUser               bool   `json:"vk_is_app_user"`
	VKAreNotificationsEnabled bool   `json:"vk_are_notifications_enabled"`
	VKGroupId                 uint64 `json:"vk_group_id"`
	VKIsFavorite              bool   `json:"vk_is_favorite"`
	VKTs                      uint64 `json:"vk_ts"`
	VKIsRecommended           bool   `json:"vk_is_recommended"`
	VKProfileId               uint64 `json:"vk_profile_id"`
	VKHasProfileButton        bool   `json:"vk_has_profile_button"`
	VKTestingGroupId          uint64 `json:"vk_testing_group_id"`
	VKViewerGroupRole         string `json:"vk_viewer_group_role"`
	VKPlatform                string `json:"vk_platform"`
	VKLanguage                string `json:"vk_language"`
	VKRef                     string `json:"vk_ref"`
	VKAccessTokenSettings     string `json:"vk_access_token_settings"`
	VKChatId                  uint64 `json:"vk_chat_id"`
}

type EventType = string
type ErrorMessage = string

const (
	EventPing               EventType = "ping"
	EventPong               EventType = "pong"
	EventStartApp           EventType = "start_app"
	EventGetBusiness        EventType = "get_business"
	EventGetPrimaryBusiness EventType = "get_pr_business"
	EventBalanceGet         EventType = "balance_get"
	EventUserGet            EventType = "user_get"
	EventError              EventType = "error"
)

const (
	//ErrorMessageNil          ErrorMessage = "data is nil"
	ErrorMessageParseData ErrorMessage = "data is nil"
	ErrorMessageMsgLength ErrorMessage = "data is nil"
	//ErrorMessageUidUndefined ErrorMessage = "the user ID is not defined or a conversion error has occurred"
	ErrorMessageMissingEvent ErrorMessage = "event is missing"
	ErrorMessageMissingData  ErrorMessage = "data is missing"
)

const (
	BusinessRoleBot = 0
)
