package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/KevinRionaldo/myGoLibrary/responseLib"
	"github.com/KevinRionaldo/web-push-notification-go/lib/getSubscription"
	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog/log"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	type bodyType struct {
		Message map[string]interface{} `json:"message"`
		// Rfid    string                 `json:"rfid"`
	}
	var body bodyType
	err := json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		log.Err(err).Msg("invalid request body")
		return responseLib.Generate(event, 400, "Invalid Request Body")
	}

	subscribtionList, err := getSubscription.Do()
	if err != nil {
		return responseLib.Generate(event, 400, err.Error())
	}

	for _, item := range subscribtionList {
		s := &webpush.Subscription{
			Endpoint: item.Endpoint,
			Keys: webpush.Keys{
				Auth:   item.Keys_Auth,
				P256dh: item.Keys_p256dh,
			},
		}
		// Marshal data to JSON
		msgStr, err := json.Marshal(body.Message)
		if err != nil {
			log.Err(err).Msg("error marshal message notif")
		}

		// Send notification
		resp, err := webpush.SendNotification([]byte(string(msgStr)), s, &webpush.Options{
			Subscriber:      "evgate-support@stroomer.id",
			VAPIDPublicKey:  os.Getenv("PUSH_NOTIF_PUBLIC_KEY"),
			VAPIDPrivateKey: os.Getenv("PUSH_NOTIF_PRIVATE_KEY"),
			TTL:             30,
		})
		if err != nil {
			log.Err(err).Msg(fmt.Sprintf("Error sending notification with id %s", item.Id))
			continue
		}
		defer resp.Body.Close()

	}
	return responseLib.Generate(event, 200, "OK")
}

func main() {
	lambda.Start(Handler)
}
