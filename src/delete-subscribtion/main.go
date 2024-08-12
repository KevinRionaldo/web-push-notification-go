package main

import (
	"context"

	"github.com/KevinRionaldo/myGoLibrary/responseLib"
	"github.com/KevinRionaldo/web-push-notification-go/config"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	// "github.com/KevinRionaldo/web-push-notification-go/lib/models"
	"github.com/rs/zerolog/log"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/cockroachdb"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// open the session
	session, err := cockroachdb.Open(config.CurrentDBConnectionURL())
	if err != nil {
		log.Error().Err(err).Msg("cockroachdb.Open")
		return responseLib.Generate(event, 500, err.Error())
	}
	defer session.Close()

	if config.IsInDevelopmentStage() {
		db.LC().SetLevel(db.LogLevelTrace)
	}

	// set the settings table
	notifTable := session.Collection(config.GetTableNameOnCurrentSchema("push_notification"))

	res := notifTable.Find(db.Cond{"id": event.PathParameters["id"]})
	err = res.Delete()
	if err != nil {
		log.Err(err).Msg("error delete qry")
		return responseLib.Generate(event, 400, err.Error())
	}

	return responseLib.Generate(event, 200, res)
}

func main() {
	lambda.Start(Handler)
}
