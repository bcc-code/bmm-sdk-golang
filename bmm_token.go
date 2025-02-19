package bmm

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log/slog"
	"os"
	"time"
)

// NewToken creates a new token for M2M communication with BMM API
func NewToken(tokenBaseURL, clientID, clientSecret, audience string) (*Token, error) {
	t := &Token{
		tokenBaseURL: tokenBaseURL,
		clientID:     clientID,
		clientSecret: clientSecret,
		audience:     audience,
	}

	return t, t.refresh()
}

type Token struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	CreatedAt   time.Time

	tokenBaseURL string
	clientID     string
	clientSecret string
	audience     string
}

// GetAccessToken returns the access token, if it is expired, it will be refreshed automatically
func (t *Token) GetAccessToken() (string, error) {
	if t.expired() {
		err := t.refresh()
		if err != nil {
			return "", err
		}
	}

	return t.AccessToken, nil
}

func (t *Token) expired() bool {
	return time.Since(t.CreatedAt)+10*time.Second > time.Duration(int64(t.ExpiresIn)*int64(time.Second))
}

func (t *Token) refresh() error {
	l := slog.With("component", "bmm_token", "action", "refresh")
	if os.Getenv("BMM_DEBUG_TOKEN") != "" {
		l.Debug("Using DEBUG token. Expired token will not be automatically refreshed")
		t.AccessToken = os.Getenv("BMM_DEBUG_TOKEN")
		t.ExpiresIn = 24 * 60 * 60
		t.CreatedAt = time.Now()
		return nil
	}

	r := resty.New()
	r.SetBaseURL(t.tokenBaseURL)
	res, err := r.R().SetBody(map[string]string{
		"client_id":     t.clientID,
		"client_secret": t.clientSecret,
		"audience":      t.audience,
		"grant_type":    "client_credentials",
	}).SetResult(&Token{}).Post("oauth/token")

	if err != nil {
		l.Error("Failed to refresh token", "error", err)
		return err
	}

	newToken := res.Result().(*Token)

	if newToken.AccessToken == "" {
		l.Error("Failed to get token", "response", res.String())
		return fmt.Errorf("faulure to get token: %s", res.String())
	}

	t.AccessToken = newToken.AccessToken
	t.Scope = newToken.Scope
	t.ExpiresIn = newToken.ExpiresIn
	t.TokenType = newToken.TokenType
	t.CreatedAt = time.Now()

	return nil
}
