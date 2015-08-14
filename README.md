# gobet
A golang betfair API implementation. Supports the full API (account/betting).

**WARNING: Use it in production at your own risk!**

**<sub><sup>I ve been coding in Go for only a week :-p<sub><sup>**

## Example Usage
You can see some sample code in the examples folder.
```
conf := gobet.Configuration{Username: "username", Password: "pass", AppKey: "appKey", Locale: "en", LogDest: os.Stdout}
s, _ := gobet.NewSession(&conf)

err = s.Login()
defer s.Logout()

s.DoKeepAliveEvery(time.Hour * 4)

input := gobet.BettingInputParams{Filter: &gobet.MarketFilter{TextQuery: "Horse Racing"}}
eventTypes, err := s.ListEventTypes(input)
```