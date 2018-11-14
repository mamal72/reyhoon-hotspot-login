package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gen2brain/beeep"
)

var (
	reyhoonHotspotIP       string
	reyhoonHotspotUsername string
	reyhoonHotspotPassword string
)

func getLoginFormAddress() string {
	return fmt.Sprintf("http://%s/login", reyhoonHotspotIP)
}

func loadEnvVars() {
	reyhoonHotspotIP = os.Getenv("RHN_HS_IP")
	reyhoonHotspotUsername = os.Getenv("RHN_HS_USERNAME")
	reyhoonHotspotPassword = os.Getenv("RHN_HS_PASSWORD")
	if reyhoonHotspotIP == "" || reyhoonHotspotUsername == "" || reyhoonHotspotPassword == "" {
		beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		beeep.Notify("Oops!", "Please set all RHN_HS environment variables!", "")
		os.Exit(1)
	}
}

func main() {
	loadEnvVars()
	address := getLoginFormAddress()
	response, err := http.PostForm(address, url.Values{
		"username": {reyhoonHotspotUsername},
		"password": {reyhoonHotspotPassword},
	})
	if err != nil {
		beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		beeep.Notify("Oops!", "Error sending post request to login form!", "")
		os.Exit(1)
	}
	beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	beeep.Notify("Done!", "Successfully logged in to Reyhoon hotspot!", "")
	defer response.Body.Close()
}
