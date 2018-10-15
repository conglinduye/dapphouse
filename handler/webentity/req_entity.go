package webentity

// common pageInfo
type PageInfo struct {
	PageSize  int64 `json:"pageSize,omitempty"`
	CurPage   int64 `json:"curPage,omitempty"`
	TotalPage int64 `json:"totalPage,omitempty"`
}

//request param for Dapp query
type DappListRequest struct {
	Category  int8      `json:"category,omitempty"`
	PageInfo  *PageInfo `json:"pageInfo,omitempty"`
	Developer string    `json:"developer,omitempty"`
}

type DappAddRequest struct {
	DappName           string          `json:"dappName"`
	Developer          string          `json:"developer"`
	Catagory           int8            `json:"catagory"`
	Tagline            string          `json:"tagline"`
	Version            string          `json:"version"`
	Description        string          `json:"description"`
	WebsiteUrl         string          `json:"websiteUrl"`
	Contracts          *[]Contract     `json:"contracts"`
	SocialAccounts     *SocialAccounts `json:"SocialAccounts"`
	ScreenshotImageUrl string          `json:"ScreenshotImageUrl"`
}
