package remote

import "time"

type FileInfo struct {
	//logger *zap.Logger

	ID             uint64    `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Bucket         string    `json:"bucket"`
	Parent         uint64    `gorm:"parent" json:"parent"`
	Key            string    `json:"key"`
	FileSize       int64     `json:"fileSize"`
	UploadTime     int64     `json:"uploadTime"`
	Type           int       `json:"type"`
	Source         string    `json:"source"`
	Name           string    `json:"name"`
	FakeHeaderHash string    `json:"fakeHeaderHash"`
	FakeHash       string    `json:"fakeHash" gorm:"type:varchar(64)"`
	FakeSize       int64     `json:"fakeSize" gorm:"type:varchar(64)"`
}
