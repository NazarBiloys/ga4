# Events pusher to GA4
### Pushed event with uah/usd ratio to GA4

Steps:
- run `cp config.sample.json config.json`
- write your credentials in `config.json` -> `measurementId`, `appId`, `event`, `apiSecret`
- `docket build -t ga4 .`
- `docker run ga4`

Worker start process at every 15th minute that will push event about uah/usd ratio.

[Report](https://analytics.google.com/analytics/web/#/p354942564/reports/reportinghub)