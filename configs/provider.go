package configs

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
)

func SetupOAuth(clientID, clientSecret, callbackURL string) {
	goth.UseProviders(
		github.New(clientID, clientSecret, callbackURL))
}
