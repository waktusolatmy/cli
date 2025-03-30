package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://api.waktusolat.app"

type Zone struct {
	JakimCode string `json:"jakimCode"`
	Negeri    string `json:"negeri"`
	Daerah    string `json:"daerah"`
}

type ZonePrayerTimes struct {
	Zone        string    `json:"zone"`
	Year        int       `json:"year"`
	Month       string    `json:"month"`
	LastUpdated time.Time `json:"last_updated"`
	Prayers     []struct {
		Maghrib int    `json:"maghrib"`
		Dhuhr   int    `json:"dhuhr"`
		Fajr    int    `json:"fajr"`
		Hijri   string `json:"hijri"`
		Syuruk  int    `json:"syuruk"`
		Day     int    `json:"day"`
		Asr     int    `json:"asr"`
		Isha    int    `json:"isha"`
	} `json:"prayers"`
}

func GetZones() ([]Zone, error) {
	var zones []Zone

	url := fmt.Sprintf("%s/zones", baseURL)
	resp, err := http.Get(url)
	if err != nil {
		return zones, fmt.Errorf("Error sending GET request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return zones, fmt.Errorf("Request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return zones, fmt.Errorf("Error reading response body: %s", err)
	}

	err = json.Unmarshal(body, &zones)
	if err != nil {
		return zones, fmt.Errorf("Error unmarshalling JSON: %s", err)
	}

	return zones, nil
}

func GetPrayerTimesByZone(zoneCode string) (ZonePrayerTimes, error) {
	var zpt ZonePrayerTimes

	url := fmt.Sprintf("%s/v2/solat/%s", baseURL, zoneCode)
	resp, err := http.Get(url)
	if err != nil {
		return zpt, fmt.Errorf("Error sending GET request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return zpt, fmt.Errorf("Request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return zpt, fmt.Errorf("Error reading response body: %s", err)
	}

	err = json.Unmarshal(body, &zpt)
	if err != nil {
		return zpt, fmt.Errorf("Error unmarshalling JSON: %s", err)
	}

	return zpt, nil
}
