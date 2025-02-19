package bmm

import (
	apiv1 "bcc-media-tools/api/v1"
	"encoding/json"
	"github.com/bcc-code/bmm-sdk-golang/bmm"
	"os"
	"testing"

	"connectrpc.com/connect"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_UnmarshallAlbumWithTracks(t *testing.T) {
	data, err := os.ReadFile("testdata/album_with_tracks.json")
	assert.NoError(t, err)

	album := &BMMItem{}
	err = json.Unmarshal(data, album)
	assert.NoError(t, err)
	assert.NotEmpty(t, album)

	assert.Equal(t, "album", album.Type)
	assert.Equal(t, 6, len(album.Tracks))
}

func Test_GetTranscriptions(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Skip(t, err)
	}

	if os.Getenv("BMM_AUTH0_BASE_URL") == "" {
		t.Skip("Required ENV variables not set for getting token")
	}

	if os.Getenv("BMM_BASE_URL") == "" {
		t.Skip("Required BMM_BASE_URL not set")
	}

	tokenBaseURL := os.Getenv("BMM_AUTH0_BASE_URL")
	clientID := os.Getenv("BMM_CLIENT_ID")
	clientSecret := os.Getenv("BMM_CLIENT_SECRET")
	audience := os.Getenv("BMM_AUDIENCE")

	token, err := bmm.GetToken(tokenBaseURL, clientID, clientSecret, audience)
	assert.NoError(t, err)

	baseURL := os.Getenv("BMM_BASE_URL")

	api := NewBMMApi(baseURL, token)

	req := connect.NewRequest(&apiv1.GetBMMTranscriptionRequest{
		Language:    "no",
		BmmId:       "115012",
		Environment: apiv1.BmmEnvironment_Production,
	})

	res, err := api.GetBMMTranscription(nil, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
