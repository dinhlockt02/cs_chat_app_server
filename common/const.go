package common

const (
	Male   = "male"
	Female = "female"
	Other  = "other"
)

const CurrentUser = "CurrentUser"
const CurrentFriendId = "CurrentFriendId"

var AppDatabase string
var AccessTokenExpiry int

// GroupCollectionName declares names of collections in the mongodb database
const (
	GroupCollectionName = "groups"
	UserCollectionName  = "users"
)
