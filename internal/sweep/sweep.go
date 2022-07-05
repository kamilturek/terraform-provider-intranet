package sweep

import (
	"fmt"
	"os"

	"github.com/kamilturek/intranet-go"
)

func SharedClient() (*intranet.Client, error) {
	sessionId := os.Getenv("INTRANET_SESSION_ID")
	if sessionId == "" {
		return nil, fmt.Errorf("INTRANET_SESSION_ID must be set for sweepers")
	}

	return intranet.NewClient(sessionId), nil
}
