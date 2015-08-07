# gobet
A golang betfair API implementation. Also my first golang attempt.

## Example Usage
### Session
```
s := gobet.NewSession("username", "pass", "appKey")
s.Login()
s.DoKeepAliveEvery(time.Hour * 11)
s.Logout()
```
