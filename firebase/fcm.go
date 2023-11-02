package firebase

import (
	"context"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/Kephas73/lib-kephas/base"
	"time"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type FirebaseFcm struct {
	*messaging.Client
}

var firebaseFcm *FirebaseFcm

func InstallFirebaseFcm(configKeys ...string) *FirebaseFcm {

	if firebaseFcm != nil {
		return firebaseFcm
	}

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

	client, errAuth := app.Messaging(ctx)
	if errAuth != nil {
		panic(errAuth)
	}

	firebaseFcm = &FirebaseFcm{client}

	return firebaseFcm
}

func GetFirebaseFcmInstance() *FirebaseFcm {
	if firebaseFcm == nil {
		firebaseFcm = InstallFirebaseFcm()
	}

	return firebaseFcm
}

func (fcm *FirebaseFcm) SendMsgToTopic(message *messaging.Message) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	resp, err := fcm.Send(ctx, message)
	if err != nil {
		fmt.Println(fmt.Sprintf("FirebaseFcm::SendMsgToTopic - SendMsgToTopic Error: %v", err))
		return "", nil
	}

	return resp, nil
}

func (fcm *FirebaseFcm) SendMsgToTopics(message ...*messaging.Message) (int, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	resp, err := fcm.SendAll(ctx, message)
	if err != nil {
		fmt.Println(fmt.Sprintf("FirebaseFcm::SendMsgToTopics - SendMsgToTopics Error: %v", err))
		return 0, 0, nil
	}

	return resp.SuccessCount, resp.FailureCount, nil
}
