package entity

type Dapp struct {
	Id          int64  `json:"id"`
	AccountId   int64  `json:"accountId"`
	DappName    string `json:"dappName"`
	Developer   string `json:"developer"`
	Catagory    int8   `json:"catagory"`
	Tagline     string `json:"tagline"`
	Version     string `json:"version"`
	Description string `json:"description"`
	WebsiteUrl  string `json:"websiteUrl"`
	Status      int8   `json:"status"`
	CreateTime  string `json:"createTime"`
	OnsaleTime  string `json:"onsaleTime"`
	OffsaleTime string `json:"offsaleTime"`
	UpdateTime  string `json:"updateTime"`
}

type DappContract struct {
	Id      int64  `json:"id"`
	DappId  int64  `json:"dappId"`
	Address string `json:"address"`
	Name    string `json:"name"`
	Type    int8   `json:"type"` //1 - main ;2 - test
}

type DappExt struct {
	Id         int64  `json:"id"`
	DappId     int64  `json:"dappId"`
	Github     string `json:"github"`
	Twitter    string `json:"twitter"`
	Facebook   string `json:"facebook"`
	Telegram   string `json:"telegram"`
	Reddit     string `json:"reddit"`
	Medium     string `json:"medium"`
	Webchat    string `json:"webchat"`
	Weibo      string `json:"weibo"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type DappPreview struct {
	Id                 int64  `json:"id"`
	AappId             int64  `json:"dappId"`
	LogoImageUrl       string `json:"logoImageUrl"`
	ScreenshotImageUrl string `json:"screenshotImageUrl"`
	Status             int8   `json:"status"`
}
