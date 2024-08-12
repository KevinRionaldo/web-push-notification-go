package main

import (
	"github.com/rs/zerolog/log"

	webpush "github.com/SherClockHolmes/webpush-go"
)

func main() {
	// Generate VAPID keys
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		log.Err(err).Msg("Error generating VAPID keys:")
	}
	log.Info().Any("Private Key:", privateKey).Msg("log private key")
	log.Info().Any("Public Key:", publicKey).Msg("log public key")
}
