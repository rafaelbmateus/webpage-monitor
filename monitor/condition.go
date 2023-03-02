package monitor

import (
	"fmt"
	"net/http"

	"monitor/config"
)

// VerifyStatusCode verify status code.
func VerifyStatusCode(res *http.Response, con *config.Condition) error {
	if con.StatusCode == 0 {
		return nil
	}

	if res.StatusCode != con.StatusCode {
		return fmt.Errorf("status is not equals - expected: %d actual: %d",
			con.StatusCode, res.StatusCode)
	}

	return nil
}
