package monitor

import (
	"fmt"
	"net/http"

	"monitor/config"
)

// VerifyStatusCode verify status code.
func VerifyStatusCode(res *http.Response, cfg *config.Condition) error {
	if cfg.StatusCode == 0 {
		return nil
	}

	if cfg.StatusCode != res.StatusCode {
		return fmt.Errorf("status is not equals - expected: %d actual: %d",
			cfg.StatusCode, res.StatusCode)
	}

	return nil
}
