# gobet
A golang betfair API implementation. Also my first golang attempt.

## Example Usage
```
conf := gobet.Configuration{Username: "username", Password: "pass", AppKey: "appKey", Locale: "en", LogDest: os.Stdout}
s, err := gobet.NewSession(&conf)
if err != nil {
	log.Fatal("Invalid Configuration! Exiting...")
}

err = s.Login()
if err != nil {
	log.Fatal("Could not login! Exiting...")
}
defer s.Logout()

log.Println("Setting Periodic Keep Alive")
s.DoKeepAliveEvery(time.Hour * 4)

input := gobet.BettingInputParams{Filter: &gobet.MarketFilter{TextQuery: "Horse Racing"}}
eventTypes, err := s.ListEventTypes(input)
log.Println(eventTypes)
```
