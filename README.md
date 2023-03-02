# Monitor

This project check the web services http status.

## Config file

```yaml
title: My Monitor

endpoints:
  - name: Google
    description: "Check Google web page"
    enable: true
    url: "https://www.google.com"
    interval: "1m"
    condition:
      status: 200

  - name: My Site
    description: "Check my web page"
    enable: false
    url: "https://mysite.com"
    interval: "1m"
    condition:
      status: 200
```

# Get Started

To run the project execute:

```console
go run cmd/main.go
```

If you want to notify by slack when a website goes down,
add `SLACK_WEBHOOK_URL` as environment variable:

```
SLACK_WEBHOOK_URL = https://hooks.slack.com/services/AAAA/BBBBB/CCCCC go run cmd/main.go
```
