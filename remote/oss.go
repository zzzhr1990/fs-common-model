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
}
