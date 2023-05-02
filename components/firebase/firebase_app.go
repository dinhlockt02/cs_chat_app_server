package firebase

import (
	"context"
	"cs_chat_app_server/common"
	firebase "firebase.google.com/go/v4"
)

type FirebaseApp interface {
	VerifyToken(ctx context.Context, idToken string) (*string, error)
	ExtractEmailFromUID(ctx context.Context, uid string) (*string, error)
}

type firebaseApp struct {
	app *firebase.App
}

func NewFirebaseApp(app *firebase.App) *firebaseApp {
	return &firebaseApp{app: app}
}

func (fa *firebaseApp) VerifyToken(ctx context.Context, idToken string) (*string, error) {
	client, err := fa.app.Auth(ctx)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}
	return &token.UID, nil
}

func (fa *firebaseApp) ExtractEmailFromUID(ctx context.Context, uid string) (*string, error) {
	client, err := fa.app.Auth(ctx)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	firebaseUser, err := client.GetUser(ctx, uid)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return &firebaseUser.Email, nil
}
