package demo

type QueryReq struct {
	Limit     	int64  		`json:"limit,omitempty"`
	Start     	int64  		`json:"start,omitempty"`
	Username	string 		`json:"username,omitempty"`
	Password	string		`json:"password,omitempty"`
}

type QueryResp struct {
	Total 		int64		`json:"total"`
	Data 		[]*Demo		`json:"data"`
}

type Demo struct {
	Id 			int64 		`json:"id"`
	Username	string 		`json:"username"`
	Password	string		`json:"password"`
	CreateTime	string 		`json:"createTime"`
	UpdateTime	string		`json:"updateTime"`
}

