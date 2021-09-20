package remote

import "time"

type FileInfo struct {
	//logger *zap.Logger

	ID             uint64    `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Bucket         string    `json:"bucket" gorm:"type:varchar(80)"`
	Parent         uint64    `gorm:"parent" json:"parent"`
	Key            string    `json:"key"`
	FileSize       int64     `json:"fileSize"`
	UploadTime     int64     `json:"uploadTime"`
	Type           int       `json:"type"`
	Source         string    `json:"source"`
	Name           string    `json:"name"`
	FakeHeaderHash string    `json:"fakeHeaderHash" gorm:"type:varchar(64)"`
	FakeHash       string    `json:"fakeHash" gorm:"type:varchar(64)"`
	FakeSize       int64     `json:"fakeSize"`
}

type FakeInfo struct {
	//logger *zap.Logger

	ID             uint64    `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	FakeHeaderHash string    `json:"fakeHeaderHash" gorm:"type:varchar(64)"`
	FakeHash       string    `json:"fakeHash" gorm:"type:varchar(64)"`
	FakeSize       int64     `json:"fakeSize"`
	Bucket         string    `json:"bucket" gorm:"type:varchar(80)"`
	Key            string    `json:"key"`
}
