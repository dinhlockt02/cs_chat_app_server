package authmodel

import "cs_chat_app_server/components/tokenprovider"

type AuthToken struct {
	AccessToken    *tokenprovider.Token `json:"access_token"`
	RefreshToken   *tokenprovider.Token `json:"refresh_token"`
	EmailVerified  bool                 `json:"email_verified"`
	ProfileUpdated bool                 `json:"profile_updated"`
}
