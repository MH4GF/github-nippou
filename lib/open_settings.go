package lib

import (
	"context"
	"fmt"

	"github.com/skratchdot/open-golang/open"
)

// OpenSettings opens settings url with web browser
func OpenSettings() error {
	var settings Settings

	accessToken, err := getAccessToken()
	if err != nil {
		return err
	}

	ctx := context.Background()
	client := getClient(ctx, accessToken)
	gistID := getGistID()

	if err := settings.Init(ctx, client, gistID); err != nil {
		return nil
	}

	fmt.Printf("Open %s\n", settings.URL)
	return open.Run(settings.URL)
}
