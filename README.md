# Monitor

This project check the web services http status.

## Config file

```yaml
title: My Monitor
slack_webhook_url: "https://hooks.slack.com/services/ABABABABA/BABABABABAB/BLABLABLABLABLABLABLABLA"

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

If you don't want send notifications by slack when a website goes down,
remove `slack_webhook_url` from config file.
