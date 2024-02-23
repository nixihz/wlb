package po

import (
	"time"
)

type GoVersion struct {
	ID             int64     `gorm:"primary_key"`
	Version        string    `gorm:"type:varchar(100);not null;unique"`
	IsMajorVersion int64     `gorm:"type:bigint;not null"`
	VersionDate    time.Time `gorm:"type:datetime;not null"`
	CreateTime     time.Time `gorm:"type:bigint;not null"`
	UpdateTime     time.Time `gorm:"type:bigint;not null"`
}

func (g *GoVersion) TableName() string {
	return "go_version"
}
