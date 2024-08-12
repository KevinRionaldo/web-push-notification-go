package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/KevinRionaldo/myGoLibrary/responseLib"
	"github.com/KevinRionaldo/web-push-notification-go/config"
	"github.com/KevinRionaldo/web-push-notification-go/lib/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/upper/db/v4/adapter/cockroachdb"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := models.Push_notification{Id: string(uuid.NewString()), Created_at: time.Now(), Updated_at: time.Now()}
	err := json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		return responseLib.Generate(event, 400, "Invalid Request Body")
	}

	// open the session
	session, err := cockroachdb.Open(config.CurrentDBConnectionURL())
	if err != nil {
		log.Error().Err(err).Msg("cockroachdb.Open")
		return responseLib.Generate(event, 500, err.Error())
	}
	defer session.Close()

	//init table if you dont have create the push_notification table before
	err = config.InitNotifTable(session)
	if err != nil {
		log.Error().Err(err).Msg("error init push notif table")
		return responseLib.Generate(event, 500, err.Error())
	}

	// set the settings table
	notifTable := session.Collection(config.GetTableNameOnCurrentSchema("push_notification"))

	_, err = notifTable.Insert(body)
	if err != nil {
		return responseLib.Generate(event, 400, err.Error())
	}

	return responseLib.Generate(event, 200, body)
}

func main() {
	lambda.Start(Handler)
}
