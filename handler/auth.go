package handler

import (
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
	"oauth2-google/util/logger"
	"strings"
)

func HandleLogin(w http.ResponseWriter, r *http.Request, oauthConf *oauth2.Config, oauthStateString string) {
	URL, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		logger.Log.Error("Parse: " + err.Error())
	}
	logger.Log.Info(URL.String())
	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	logger.Log.Info(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
