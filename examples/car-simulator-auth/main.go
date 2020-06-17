package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/adaptant-labs/mercedes-connectedvehicle-go"
	"golang.org/x/oauth2"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	oauth2Config = oauth2.Config{
		ClientID:     os.Getenv("MERCEDES_CLIENT_ID"),
		ClientSecret: os.Getenv("MERCEDES_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       mercedes.DefaultScopes,
		Endpoint:     mercedes.Endpoint,
	}

	oauthStateString string
)

func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	var html = `<html>
<body>
	<a href="/login">Mercedes Log in</a>
</body>
</html>
`
	fmt.Fprintf(w, html)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	err := toggleDoorLocks(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func toggleDoorLocks(state string, code string) error {
	// Make sure we have a valid request
	if state != oauthStateString {
		return errors.New("oauth state strings do not match")
	}

	// Exchange code for access token
	token, err := oauth2Config.Exchange(context.TODO(), code)
	if err != nil {
		return err
	}

	// Use the new access token for subsequent API calls
	client := mercedes.NewClient(token.AccessToken, true)
	vehicles, err := client.GetVehicles(context.TODO())
	if err != nil {
		panic(err)
	}

	// Make sure there are vehicles available to the user
	if len(vehicles) == 0 {
		return errors.New("no vehicles discovered")
	}

	// Fetch a vehicle to operate on
	vehicle := client.NewVehicle(vehicles[0].Id)

	// Obtain more detailed information about the vehicle
	detail, err := vehicle.GetVehicleDetail(context.TODO())
	if err != nil {
		return err
	}

	go func() {
		// Lock doors
		fmt.Printf("Locking doors for vehicle %s (VIN %s)...\n", vehicle.VehicleID, detail.VIN)
		_ = vehicle.LockDoors(context.TODO())

		time.Sleep(5 * time.Second)

		// Fetch door lock status
		doors, _ := vehicle.GetDoorStatus(context.TODO())

		println("Front left door is now", doors.DoorLockStatusFrontLeft.Value)
		time.Sleep(5 * time.Second)

		// Unlock doors
		println("Unlocking...")
		_ = vehicle.UnlockDoors(context.TODO())
	}()

	return nil
}

func main() {
	// Generate a random state string for validation
	rand.Seed(time.Now().UnixNano())
	oauthStateString = randomString(16)

	// Handlers for Oauth2 login and consent flows
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)

	println("Awaiting completion of login and consent flow at http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
