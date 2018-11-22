package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gen2brain/beeep"
)

var (
	reyhoonHotspotIP               string
	reyhoonHotspotUsername         string
	reyhoonHotspotPassword         string
	reyhoonHotspotLoginFormAddress string
)

func init() {
	reyhoonHotspotIP = os.Getenv("RHN_HS_IP")
	reyhoonHotspotUsername = os.Getenv("RHN_HS_USERNAME")
	reyhoonHotspotPassword = os.Getenv("RHN_HS_PASSWORD")
	if reyhoonHotspotIP == "" || reyhoonHotspotUsername == "" || reyhoonHotspotPassword == "" {
		beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		beeep.Notify("Oops!", "Please set all RHN_HS environment variables!", "")
		os.Exit(1)
	}
	reyhoonHotspotLoginFormAddress = fmt.Sprintf("http://%s/login", reyhoonHotspotIP)
}

func main() {
	response, err := http.PostForm(reyhoonHotspotLoginFormAddress, url.Values{
		"username": {reyhoonHotspotUsername},
		"password": {reyhoonHotspotPassword},
	})
	if err != nil {
		beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		beeep.Notify("Oops!", "Error sending post request to login form!", "")
		os.Exit(1)
	}
	defer response.Body.Close()
	beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	beeep.Notify("Done!", "Successfully logged in to Reyhoon hotspot!", "")
}
