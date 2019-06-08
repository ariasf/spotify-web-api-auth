package spotify

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"main/appsettings"
	"main/structures"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	accountSpotifyAPIURL = "https://accounts.spotify.com"
	authorizeEndpoint    = "authorize"
	spotifyTokenEndpoint = "api/token"
)

var (
	spotifyRedirectUri  = appsettings.GetFromEnvironment("SPOTIFY_APPLE_CALLBACK_URL", "")
	spotifyClientId     = appsettings.GetFromEnvironment("SPOTIFY_CLIENT_ID", "")
	spotifyClientSecret = appsettings.GetFromEnvironment("SPOTIFY_CLIENT_SECRET", "")
	spotifyScopes       = []string{
		"user-library-read",
		"playlist-read-private",
		"playlist-read-collaborative",
		"user-read-playback-state",
		"user-modify-playback-state",
		"user-read-currently-playing",
		"app-remote-control",
		"streaming",
	}
)

func LoginRedirect() string {
	authUrl := accountSpotifyAPIURL + "/" + authorizeEndpoint + "?client_id=" + spotifyClientId + "&response_type=code&redirect_uri=" + encodeParam(spotifyRedirectUri) + "&scope=" + getScopesEncoded()
	return authUrl
}

func SwapToken(code, state string) (structures.TokenResponse, error) {
	var result structures.TokenResponse

	if code == "" && state == "" {
		errr := errors.New(" code, error and state are all empty, unknown error occurred")
		log.Printf("%+v", errr)
		return result, errr
	}
	form := url.Values{}
	form.Add("code", code)
	form.Add("grant_type", "authorization_code")
	form.Add("redirect_uri", spotifyRedirectUri)

	return getAccessToken(form)
}


func RefreshToken(refreshToken string) (structures.TokenResponse, error) {
	form := url.Values{}
	form.Add("grant_type", "refresh_token")
	form.Add("refresh_token", refreshToken)

	return getAccessToken(form)
}

func getAccessToken(form url.Values) (structures.TokenResponse, error) {
	var result structures.TokenResponse

	client := getHttpClient()
	request, err := http.NewRequest(http.MethodPost, accountSpotifyAPIURL+"/"+spotifyTokenEndpoint, strings.NewReader(form.Encode()))
	if err != nil {
		log.Printf("error creating request for token %+v", err)
	}
	request.Header.Add("Authorization", "Basic "+encodeStringBase64(spotifyClientId+":"+spotifyClientSecret))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	response, err := client.Do(request)
	if err != nil {
		log.Printf("error requesting token %+v", err)
		return result, err
	}
	jsonResult, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("error reading response body for token %+v", err)
		return result, err
	}
	err = json.Unmarshal(jsonResult, &result)
	if err != nil {
		log.Printf("error deserializing token response %+v", err)
		return result, err
	}

	//added to help the spotify sdk
	if result.RefreshToken == "" && form.Get("refresh_token") != "" {
		result.RefreshToken = form.Get("refresh_token")
	}

	result.ExpirationTimeStampInMs = time.Now().Add(time.Duration(result.ExpiresIn) * time.Second).UnixNano() / 1000000
	return result, nil
}

func getHttpClient() *http.Client {
	tr := &http.Transport{
		IdleConnTimeout: 120 * time.Second,
	}
	return &http.Client{Transport: tr}
}

func encodeParam(s string) string {
	return url.QueryEscape(s)
}

func encodeStringBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func getScopesEncoded() string {
	var scopesSpaceSeparated string
	for _, scope := range spotifyScopes {
		scopesSpaceSeparated = scopesSpaceSeparated + " " + scope
	}
	return encodeParam(strings.Trim(scopesSpaceSeparated, " "))
}
