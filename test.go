package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/auth0/go-auth0/v2/authentication"
	"github.com/auth0/go-auth0/v2/authentication/database"
	"github.com/auth0/go-auth0/v2/authentication/oauth"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/oauth2"
)

func main() {
	// Get these from your Auth0 Application Dashboard.
	domain := os.Getenv("AUTH0_DOMAIN")
	clientID := os.Getenv("AUTH0_CLIENT_ID")
	clientSecret := os.Getenv("AUTH0_CLIENT_SECRET")
	redirectUri := os.Getenv("AUTH0_CALLBACK_URL")

	// Initialize a new client using a domain, client ID and client secret.
	authAPI, err := authentication.New(
		context.TODO(), // Replace with a Context that better suits your usage
		domain,
		authentication.WithClientID(clientID),
		authentication.WithClientSecret(clientSecret), // Optional depending on the grants used
	)
	if err != nil {
		log.Fatalf("failed to initialize the auth0 authentication API client: %+v", err)
	}

	// Now we can interact with the Auth0 Authentication API.
	// Sign up a user
	userData := database.SignupRequest{
		Connection: "Username-Password-Authentication",
		Username:   "mytestaccount",
		Password:   "mypassword sovfuhe9p84t7zq957tgzrehi ÜÜ+äö",
		Email:      "mytestaccount@example.com",
		UserMetadata: &map[string]any{
			"preferred_color": "blue",
		},
	}

	createdUser, err := authAPI.Database.Signup(context.Background(), userData)
	if err != nil {
		log.Fatalf("failed to sign user up: %v", err)
	}

	fmt.Println(createdUser)

	cv := oauth2.GenerateVerifier()

	// Login using OAuth grants
	tokenSet, err := authAPI.OAuth.LoginWithAuthCodeWithPKCE(context.Background(), oauth.LoginWithAuthCodeWithPKCERequest{
		Code:         oauth2.S256ChallengeFromVerifier(cv),
		CodeVerifier: cv,
		RedirectURI:  redirectUri,
	}, oauth.IDTokenValidationOptions{})
	if err != nil {
		log.Fatalf("failed to retrieve tokens: %v", err)
	}

	fmt.Println(tokenSet)
}
