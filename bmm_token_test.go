package bmm

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getToken(t *testing.T) {
	if os.Getenv("BMM_AUTH0_BASE_URL") == "" {
		t.Skip("Required ENV variables not set for getting token")
	}

	tokenBaseURL := os.Getenv("BMM_AUTH0_BASE_URL")
	clientID := os.Getenv("BMM_CLIENT_ID")
	clientSecret := os.Getenv("BMM_CLIENT_SECRET")
	audience := os.Getenv("BMM_AUDIENCE")

	res, err := NewToken(tokenBaseURL, clientID, clientSecret, audience)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}
