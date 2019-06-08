package spotify

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"main/structures"
	"net/http"
)

/*
Redirects to Spotify login, hitting this endpoint might not be necessary if you are using  a mobile sdk like iOSon the
client side. Use this endpoint otherwise
*/
func Login(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, LoginRedirect())
}

/*
After the user successfully accepts and authorizes our app's permissions you will get a authorization code, we will use
that to get an access token and a refresh token.
*/
func Swap(ctx *gin.Context) {
	code := ctx.PostForm("code")
	errCode := ctx.PostForm("error")
	state := ctx.PostForm("state") //will implement maybe later
	if errCode != "" {

		errorResponse := structures.TokenResponse{ErrorCode: errCode}
		jsonErrorResponse, _ := json.Marshal(errorResponse)
		ctx.Data(http.StatusBadRequest, "text/json", jsonErrorResponse)

	} else if code != "" {
		tokenResponse, err := SwapToken(code, state)
		if err != nil {
			jsonErr, _ := json.Marshal(err)
			ctx.Data(http.StatusBadRequest, "text/json", jsonErr)
		} else {
			jsonResponse, _ := json.Marshal(tokenResponse)
			ctx.Data(http.StatusOK, "text/json", jsonResponse)
		}
	} else {
		errorResponse := structures.TokenResponse{ErrorCode: "missing parameters or unknown error"}
		jsonErrorResponse, _ := json.Marshal(errorResponse)
		ctx.Data(http.StatusBadRequest, "text/json", jsonErrorResponse)
	}
}

/*
After getting the initial access and refresh token from the swap endpoint using the authorization code, the refresh
token endpoint allows you to get a new access token when the one issued previously has expired, it will also return a
new refresh token occasionally, it will send back the send refresh token we sent otherwise.
*/
func Refresh(ctx *gin.Context) {

	refreshToken := ctx.PostForm("refresh_token")
	if refreshToken != "" {
		tokenInfo, err := RefreshToken(refreshToken)
		if err != nil {
			jsonErr, _ := json.Marshal(err)
			ctx.Data(http.StatusBadRequest, "text/json", jsonErr)
		} else {
			jsonResponse, _ := json.Marshal(tokenInfo)
			ctx.Data(http.StatusOK, "text/json", jsonResponse)
		}

	} else {
		errorResponse := structures.TokenResponse{ErrorCode: "missing parameters or unknown error"}
		jsonErrorResponse, _ := json.Marshal(errorResponse)
		ctx.Data(http.StatusBadRequest, "text/json", jsonErrorResponse)
	}
}
