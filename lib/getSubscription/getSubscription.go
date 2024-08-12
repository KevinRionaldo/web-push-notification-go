package getSubscription

import (
	"github.com/KevinRionaldo/web-push-notification-go/config"
	"github.com/KevinRionaldo/web-push-notification-go/lib/models"
	"github.com/rs/zerolog/log"
	db "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/cockroachdb"
)

func Do() ([]models.Push_notification, error) {
	// open the session
	session, err := cockroachdb.Open(config.CurrentDBConnectionURL())
	if err != nil {
		log.Err(err).Msg("cockroachdb.Open")
		return nil, err
	}
	defer session.Close()

	//init table if you dont have create the push_notification table before
	// err = config.InitNotifTable(session)
	// if err != nil {
	// 	log.Error().Err(err).Msg("error init push notif table")
	// 	return nil, err
	// }

	// set the settings table
	notifTable := session.Collection(config.GetTableNameOnCurrentSchema("push_notification"))

	if config.IsInDevelopmentStage() {
		db.LC().SetLevel(db.LogLevelTrace)
	}

	subscribtionList := []models.Push_notification{}
	err = notifTable.Find().All(&subscribtionList)
	if err != nil {
		log.Err(err).Msg("error fetching subscribtion data")
		return nil, err
	}

	return subscribtionList, nil
}
