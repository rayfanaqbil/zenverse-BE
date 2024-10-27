package module

import(
	"golang.org/x/oauth2"
	"os"
	"github.com/rayfanaqbil/zenverse-BE/v2/model"
	"golang.org/x/oauth2/google"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
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

func isAdmin(db *mongo.Database, col string, email string) bool {
	var admin model.Admin
	err := db.Collection(col).FindOne(context.Background(), bson.M{"email": email}).Decode(&admin)
	return err == nil
}