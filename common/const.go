package common

const (
	Male   = "male"
	Female = "female"
	Other  = "other"
)

const CurrentUser = "CurrentUser"
const CurrentFriendId = "CurrentFriendId"
const CurrentGroupId = "CurrentGroupId"

var AppDatabase string
var AccessTokenExpiry int

// GroupCollectionName declares names of collections in the mongodb database
const (
	GroupCollectionName            = "groups"
	UserCollectionName             = "users"
	RequestCollectionName          = "requests"
	GroupChatHistoryCollectionName = "groupChatHistory"
)
