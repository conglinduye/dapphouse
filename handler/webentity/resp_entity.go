package webentity

//vo for Dapp query list
type DappListInfo struct {
	PageInfo      *PageInfo        `json:"pageiInfo"`
	DappInfo4List *[]DappInfo4List `json:"dappInfos"`
}

//vo for Dapp in list
type DappInfo4List struct {
	Name      string `json:"name"`      //dapp nmae
	Logo      string `json:"logo"`      //dapp logo url,show only for dappList API
	Category  string `json:"category"`  //dapp category
	Tagline   string `json:"tagline"`   //dapp tagline
	Desc      string `json:"desc"`      //dapp desc
	Ver       string `json:"ver"`       //dapp ver
	Status    string `json:"status"`    //dapp status,show only for myDapps API
	Developer string `json:"developer"` //dapp developer
}

//vo for dapp detail
type DappDetailInfo struct {
	Name               string          `json:"name"`      //dapp nmae
	Logo               string          `json:"logo"`      //dapp logo url
	Category           string          `json:"category"`  //dapp category
	Tagline            string          `json:"tagline"`   //dapp tagline
	Desc               string          `json:"desc"`      //dapp desc
	Ver                string          `json:"ver"`       //dapp ver
	Developer          string          `json:"developer"` //dapp developer
	WebsiteUrl         string          `json:"websiteUrl"`
	Contracts          *[]Contract     `json:"contracts"`
	SocialAccounts     *SocialAccounts `json:"SocialAccounts"`
	ScreenshotImageUrl string          `json:"ScreenshotImageUrl"`
	Status             string          `json:"status"` //dapp status
}

//social accounts
type SocialAccounts struct {
	Github   string `json:"github"`
	Twitter  string `json:"twitter"`
	Facebook string `json:"facebook"`
	Telegram string `json:"telegram"`
	Reddit   string `json:"reddit"`
	Medium   string `json:"medium"`
	Webchat  string `json:"webchat"`
	Weibo    string `json:"weibo"`
}

type Contract struct {
	Address string `json:"Address"`
	Name    string `json:"name"`
	Type    int8   `json:"type"` //1 - main ;2 - test
}
