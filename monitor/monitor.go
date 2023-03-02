package monitor

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"monitor/config"

	"github.com/rs/zerolog"
)

// Monitor represets the monitor object.
type Monitor struct {
	Context context.Context
	Log     *zerolog.Logger
	Notify  *SlackNotify
}

// NewMonitor returns a new Monitor instance.
func NewMonitor(ctx context.Context, log *zerolog.Logger,
	notify *SlackNotify) *Monitor {
	return &Monitor{
		Context: ctx,
		Log:     log,
		Notify:  notify,
	}
}

// Run loops over each endpoint and starts a goroutine to monitor each endpoint separately.
func (m *Monitor) Run(cfg *config.Config) {
	m.Log.Info().Msgf("running monitor with config: %v", cfg)
	err := m.Notify.SendMessage("api monitor started")
	if err != nil {
		m.Log.Error().Msgf("error on send slack notification %q", err)
	}

	for _, endpoint := range cfg.Endpoints {
		m.Log.Debug().Msgf("running monitor endpoint: %v", endpoint)
		if endpoint.Enable {
			go m.MonitorSleep(endpoint)
		}
	}
}

// Monitor the endpoint forever.
func (m *Monitor) MonitorSleep(endpoint *config.Endpoint) {
	for {
		err := m.Monitor(endpoint)
		if err != nil {
			m.Log.Error().Msgf("error on monitor %q", err)
		}

		time.Sleep(endpoint.Interval)
	}
}

// Monitor the endpoint and verify conditions.
func (m *Monitor) Monitor(endpoint *config.Endpoint) error {
	m.Log.Debug().Msgf("monitor endpoint %v", endpoint)

	if endpoint.Method == "" {
		endpoint.Method = "GET"
	}

	req, err := http.NewRequest(endpoint.Method, endpoint.URL, nil)
	if err != nil {
		return err
	}

	client := http.Client{Timeout: time.Duration(5) * time.Second}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if err = VerifyStatusCode(res, endpoint.Condition); err != nil {
		m.Notify.SendMessage(fmt.Sprintf("monitor %s failed", endpoint.Name))
	}

	return nil
}
