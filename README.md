# Events pusher to GA4
### Pushed event with uah/usd ratio to GA4

Steps:
- run `cp config.sample.json config.json`
- write your credentials in `config.json` -> `measurementId`, `appId`, `event`, `apiSecret`
- `docket build -t ga4 .`
- `docker run ga4`

Worker start process at every 15th minute that will push event about uah/usd ratio.

[Report](https://analytics.google.com/analytics/web/#/p354942564/reports/dashboard?params=_u..nav%3Dmaui%26_r..dimension-value%3D%7B%22dimension%22:%22eventName%22,%22value%22:%22uah_usd_ratio%22%7D&r=events-overview)

![image](https://user-images.githubusercontent.com/51129612/221361449-76c23d6d-b620-4db0-8d4e-38c2b67f2c60.png)
![image](https://user-images.githubusercontent.com/51129612/221410336-eca9b7c1-121f-454e-aafd-9cd19213de3c.png)
