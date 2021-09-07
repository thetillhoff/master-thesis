package hardware_inspection

import (
	"net/http"
	"time"
)

var (
	httpClient = http.Client{
		Timeout: 2 * time.Second,
	}
)
