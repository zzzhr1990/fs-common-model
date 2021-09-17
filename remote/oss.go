package remote

import "time"

type OssInfo struct {
}

type OdCookie struct {
	//logger *zap.Logger

	ID        uint64    `gorm:"primarykey" json:"ID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Cookie    string    `json:"cookie"`
	PickCode  string    `json:"pickCode"`
	Name      string    `json:"name"`
	User      uint64    `gorm:"index" json:"user"`
}

type OdPreviewResponse struct {
	State    bool   `json:"state"`
	FileName string `json:"file_name"`
	FileSize string `json:"file_size"`
	Pickcode string `json:"pickcode"`
	FileURL  string `json:"file_url"`
	IsHTML   int    `json:"is_html"`
	Page     int    `json:"page"`
	URL      string `json:"url"`
	Tab      string `json:"tab"`
	Type     string `json:"type"`
	Msg      string `json:"msg"`
	Errtype  string `json:"errtype"`
	MsgCode  int    `json:"msg_code"`
}

type OdTokenResponse struct {
	StatusCode      string    `json:"StatusCode"`
	AccessKeySecret string    `json:"AccessKeySecret"`
	SecurityToken   string    `json:"SecurityToken"`
	Expiration      time.Time `json:"Expiration"`
	AccessKeyID     string    `json:"AccessKeyId"`
}

type OdToken struct {
	StatusCode      string    `json:"statusCode"`
	AccessKeySecret string    `json:"accessKeySecret"`
	SecurityToken   string    `json:"securityToken"`
	Expiration      time.Time `json:"expiration"`
	AccessKeyID     string    `json:"accessKeyId"`
}

type OdUserResponse struct {
	State bool   `json:"state"`
	Error string `json:"error"`
	Data  struct {
		Device      int  `json:"device"`
		Rank        int  `json:"rank"`
		Liang       int  `json:"liang"`
		Mark        int  `json:"mark"`
		Mark1       int  `json:"mark1"`
		Vip         int  `json:"vip"`
		Expire      int  `json:"expire"`
		Global      int  `json:"global"`
		Forever     int  `json:"forever"`
		IsPrivilege bool `json:"is_privilege"`
		// "error":"\u8bf7\u5148\u767b\u5f55,\u540e\u64cd\u4f5c\uff01"
		Privilege struct {
			Start  int  `json:"start"`
			Expire int  `json:"expire"`
			State  bool `json:"state"`
		//	Mark   int  `json:"mark"`
		} `json:"privilege"`
		UserName string `json:"user_name"`
		Face     string `json:"face"`
		UserID   uint64 `json:"user_id"`
	} `json:"data"`
}
