package module

import(
	"golang.org/x/oauth2"
	"os"
	"golang.org/x/oauth2/google"
)

var (
	googleOAuthConfig = &oauth2.Config{
		RedirectURL:  "https://hrisz.github.io/zenverse_FE/",
		ClientID: os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint:     google.Endpoint,
	}
	allowedAdmins = []string{"rayfana09@gmail.com", "harissaefuloh@gmail.com"}
)