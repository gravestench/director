package access_token

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
)

var (
	oauth2Config *clientcredentials.Config
)

func Get(clientID, clientSecret string) string {
	oauth2Config = &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     twitch.Endpoint.TokenURL,
	}

	token, err := oauth2Config.Token(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return token.AccessToken
}
