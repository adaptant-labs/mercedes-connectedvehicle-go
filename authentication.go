package mercedes

import "golang.org/x/oauth2"

var (
	DefaultScopes = []string{"mb:user:pool:reader", "mb:vehicle:status:general"}
	Endpoint      = oauth2.Endpoint{
		AuthURL:  "https://api.secure.mercedes-benz.com/oidc10/auth/oauth/v2/authorize",
		TokenURL: "https://api.secure.mercedes-benz.com/oidc10/auth/oauth/v2/token",
	}
)
