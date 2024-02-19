package po

type GoVersion struct {
	ID               int    `gorm:"primary_key"`
	Version          string `gorm:"type:varchar(100);not null;unique"`
	PrimaryVersion   string `gorm:"type:varchar(100);not null;unique"`
	SecondaryVersion string `gorm:"type:varchar(100);not null;unique"`
	ThirdVersion     string `gorm:"type:varchar(100);not null;unique"`
	CreateTime       int64  `gorm:"type:bigint;not null"`
	UpdateTime       int64  `gorm:"type:bigint;not null"`
}
