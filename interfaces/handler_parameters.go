package interfaces

/*------
 * parameters. エンドポイントの数とほぼ1:1
 --------*/

type SignupParams struct {
	ProviderId           string
	TwitterId            string
	Lang                 string
	Location             string
	Name                 string
	ProfileBannerUrl     string
	ProfileImageUrlHttps string
	Protected            string
	ScreenName           string
	Url                  string
}

type GetMeParams struct {
	TwitterId string
}

type CreateLyricParams struct {
	Lyric  string
	Title  string
	Singer string
	Url    string
}

type PostFavParams struct {
	UserId  string
	LyricId string
}

type UnFavParams struct {
	UserId  string
	LyricId string
}

type GetFavListParams struct {
	UserId string
}
