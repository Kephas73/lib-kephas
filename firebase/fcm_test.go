package firebase

import (
	"context"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/spf13/viper"
	_ "golang.org/x/oauth2"
	"log"
	"sync"
	"testing"
	"time"
)

const TOKEN_ANDROID = "ca0Un66nCrk:APA91bFGMkdysvY_rLVTc16A_r_OoJjQTPltEYt1oaamtv_2GYhH0yLtgPANJxuaFt9gNHwdNej1124MFl5tJMIshCSO6pYUYenRaVL13pWmuYQWV5Kx7p7gUzEj0HMzMCk0DfYVlIc2"
const TOKEN_ANDROID2 = "ds3XuMbwoeg:APA91bHPpjomIpp6yetGLTjz_gnHtV0c5LeBCTwim0zM_BBjxKFwYcGm8L8BAReUp6vgLrsshdc0C1ekxlAuQMjZN7_L8gMRs8WI5Cs_4cS-caybZJzfq82iYqvchgAen_lrqovYdlSi"

type NotificationFcm struct {
	ID                 int    `db:"id" json:"id,omitempty"`
	Title              string `db:"title" json:"title,omitempty"`
	Content            string `db:"content" json:"content,omitempty"`
	ImageUrl           string `json:"image_url" json:"image_url,omitempty"`
	Action             string `json:"action" json:"action,omitempty"`
	Payload            string `json:"payload" json:"payload,omitempty"`
	ClickActionAndroid string `json:"click_action_android" json:"click_action_android,omitempty"`
	ClickActionIOS     string `json:"click_action_ios" json:"click_action_ios,omitempty"`

	OffPush      int        `db:"off_push" json:"off_push"`
	StartAfter   string     `db:"start_after" json:"start_after,omitempty"`
	IntervalLoop int64      `db:"interval_loop" json:"interval_loop"`
	LastPush     int64      `db:"last_push" json:"last_push"`
	CreatedTime  int32      `db:"created_time" json:"created_time,omitempty"`
	UpdatedTime  int32      `db:"updated_time" json:"updated_time,omitempty"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeleteAt     *time.Time `db:"delete_at" json:"delete_at,omitempty"`
}

func TestFirebaseFcm_SendMsgToTopic(t *testing.T) {
	viper.SetConfigFile(`config.json`)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fcm := InstallFirebaseFcm()
	GetFirebaseFcmInstance()
	fmt.Println(fcm.SubscribeToTopic(context.TODO(), []string{
		TOKEN_ANDROID2,
	}, "dev_fcm_all"))

	//fmt.Println(fcm.UnsubscribeFromTopic(context.TODO(), []string{
	//	TOKEN_ANDROID2,
	//}, "fcm_all"))

	notify := NotificationFcm{
		OffPush: 3,
	}

	for {
		time.Sleep(time.Second * 10)
		res, err := fcm.SendMsgToTopic(&messaging.Message{
			Topic: "dev_fcm_all",
			Notification: &messaging.Notification{
				Title:    "Hay quá haha",
				Body:     "Xin chào tất cả mọi người ahihi",
				ImageURL: "https://cdn.gametv.vn/news_media/image/900x700_0x0_1650427181.png",
			},
			Android: &messaging.AndroidConfig{
				Notification: &messaging.AndroidNotification{
					DefaultSound: true,
					ClickAction:  "FCM_ALLL",
				},
			},
			APNS: &messaging.APNSConfig{
				Payload: &messaging.APNSPayload{
					Aps: &messaging.Aps{
						Sound:    "default",
						Category: "FCM_ALLL",
					},
				},
			},
			Data: map[string]string{"raw_data": fmt.Sprint(notify)},
		})
		fmt.Println(err)
		fmt.Println(res)
	}
}

type BufferFcm struct {
	fcmClient *messaging.Client

	dispatchInterval time.Duration
	batchCh          chan *messaging.Message
	wg               sync.WaitGroup
}

func (b *BufferFcm) SendPush(msg *messaging.Message) {
	b.batchCh <- msg
}

func (b *BufferFcm) sender() {
	defer b.wg.Done()

	t := time.NewTicker(b.dispatchInterval)

	messages := make([]*messaging.Message, 0, 500)

	defer func() {
		t.Stop()

		b.sendMessages(messages)

		log.Println("batch sender finished")
	}()

	for {
		select {
		case m, ok := <-b.batchCh:
			if !ok {
				return
			}

			messages = append(messages, m)
		case <-t.C:
			b.sendMessages(messages)
			messages = messages[:0]
		}
	}
}

func (b *BufferFcm) Run() {
	b.wg.Add(1)
	go b.sender()
}

func (b *BufferFcm) Stop() {
	close(b.batchCh)
	b.wg.Wait()
}

func (b *BufferFcm) sendMessages(messages []*messaging.Message) {
	if len(messages) == 0 {
		return
	}

	batchResp, err := b.fcmClient.SendAll(context.TODO(), messages)

	log.Printf("batch response: %+v, err: %s \n", batchResp, err)
}
