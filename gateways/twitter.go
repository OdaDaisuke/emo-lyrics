package gateways

import (
  "github.com/garyburd/go-oauth/oauth"
)

type TwitterGateway struct {
}

func NewTwitterGW() *TwitterGateway {
  return &TwitterGateway{}
}

func (c *TwitterGateway) GetConnect() *oauth.Client {
    return &oauth.Client{
        TemporaryCredentialRequestURI: "https://api.twitter.com/oauth/request_token",
        ResourceOwnerAuthorizationURI: "https://api.twitter.com/oauth/authorize",
        TokenRequestURI:               "https://api.twitter.com/oauth/access_token",
        Credentials: oauth.Credentials{
            Token:  "n3oHoTErxwwRl8ktx4H7Ni8zz",
            Secret: "KmD76QhvP2FRs0KpyRUpmhLG3waa1erTImV2xwbCIDLZJsaEkN",
        },
    }
}

/*
// GetAccessToken アクセストークンを取得する
func (c *TwitterGateway) GetAccessToken(rt *oauth.Credentials, oauthVerifier string) (*oauth.Credentials, error) {
    oc := GetConnect()
    at, _, err := oc.RequestToken(nil, rt, oauthVerifier)

    return at, err
}
*/

/*
func GetMe(at *oauth.Credentials, user *Account) error {
    oc := GetConnect()

    v := url.Values{}
    v.Set("include_email", "true")

    resp, err := oc.Get(nil, at, "https://api.twitter.com/1.1/account/verify_credentials.json", v)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 500 {
        return errors.New("Twitter is unavailable")
    }

    if resp.StatusCode >= 400 {
        return errors.New("Twitter request is invalid")
    }

    err = json.NewDecoder(resp.Body).Decode(user)
    if err != nil {
        return err
    }

    return nil
}
*/
