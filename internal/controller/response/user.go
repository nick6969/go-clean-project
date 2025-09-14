package response

type Register struct {
	AccessToken string `json:"access_token"`
}

type Login struct {
	AccessToken string `json:"access_token"`
}
