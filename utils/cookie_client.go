package util

import(
  "strings"
  "net/http"
)

type CookieClient struct {
}

func NewCookieClient() *CookieClient {
  return &CookieClient{}
}

func (c *CookieClient) SetCookie(w http.ResponseWriter) {
  cookie := &http.Cookie {
    ScreenName: "aaaa",
  }
  http.SetCookie(w, cookie)
}

func (c *CookieClient) LoadCookie(r *http.Request) {
  cookie, err := r.Cookie("ScreenName")
  if err != nil {
    fmt.Println("")
  }
  fmt.Println(cookie)
}
