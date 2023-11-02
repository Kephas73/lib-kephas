package firebase

import (
	"context"
	"fmt"
	"github.com/Kephas73/lib-kephas/base"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type FirebaseAuth struct {
	*auth.Client
}

func InstallFirebaseAuth(configKeys ...string) *FirebaseAuth {
	if firebaseConfig == nil {
		getFirebaseConfigFromEnv(configKeys...)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	opt := option.WithCredentialsJSON(base.JSONDebugData(firebaseConfig))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}

	client, errAuth := app.Auth(ctx)
	if errAuth != nil {
		panic(errAuth)
	}

	return &FirebaseAuth{Client: client}
}

func (f *FirebaseAuth) GetUser(jwtToken string) *auth.UserRecord {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	token, err := f.VerifyIDToken(ctx, jwtToken)
	if err != nil {
		fmt.Println(fmt.Sprintf("FirebaseAuth::GetUser - VerifyIDToken Error: %v", err))
		return nil
	}

	user, err := f.Client.GetUser(ctx, token.UID)
	if err != nil {
		fmt.Println(fmt.Sprintf("FirebaseAuth::GetUser - GetUser Error: %v", err))
		return nil
	}

	return user
}
