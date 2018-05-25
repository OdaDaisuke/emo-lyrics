package gateways

import (
  "net/http"
  "fmt"
  "os"
  "io/ioutil"
  "github.com/mrjones/oauth"
)

type TwitterGateway struct {
  r *http.Request
  consumer *oauth.Consumer
}

func NewTwitterGW(r *http.Request) *TwitterGateway {
  consumer := oauth.NewConsumer(
    os.Getenv("CONSUMER_KEY"),
    os.Getenv("CONSUMER_SECRET"),
    oauth.ServiceProvider{
      RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
      AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
      AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
    })
  return &TwitterGateway{r, consumer}
}

func (c *TwitterGateway) GetAuthUrl() string {
  _, u, err := c.consumer.GetRequestTokenAndUrl("oob")
  if err != nil {
    return ""
  }
  return u
}

func (c *TwitterGateway) SetVerificationCode() bool {
  requestToken, _, _ := c.consumer.GetRequestTokenAndUrl("oob")
  accessToken, err := c.consumer.AuthorizeToken(requestToken, c.r.FormValue("verification_code"))
  if err != nil {
    return false
  }

  client, err := c.consumer.MakeHttpClient(accessToken)
  if err != nil {
    return false
  }

  response, err := client.Get("https://api.twitter.com/1.1/account/verify_credentials")
  defer response.Body.Close()
  if err != nil {
    return false
  }

  bits, err := ioutil.ReadAll(response.Body)
  fmt.Println(bits)
  return true
}
