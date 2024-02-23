package factory

import (
	"wlb/internal/biz"
	"wlb/internal/data/po"
)

func CreateGoVersion(goVersionPO po.GoVersion) biz.GoVersion {
	boolMap := map[int64]bool{
		0: false,
		1: true,
	}

	return biz.GoVersion{
		ID:             goVersionPO.ID,
		Version:        goVersionPO.Version,
		IsMajorVersion: boolMap[goVersionPO.IsMajorVersion],
		VersionDate:    goVersionPO.VersionDate,
		CreateTime:     goVersionPO.CreateTime,
		UpdateTime:     goVersionPO.UpdateTime,
	}
}
func CreateGoVersionPO(goVersion biz.GoVersion) po.GoVersion {
	boolMap := map[bool]int64{
		false: 0,
		true:  1,
	}
	return po.GoVersion{
		ID:             goVersion.ID,
		Version:        goVersion.Version,
		IsMajorVersion: boolMap[goVersion.IsMajorVersion],
		VersionDate:    goVersion.VersionDate,
		CreateTime:     goVersion.CreateTime,
		UpdateTime:     goVersion.UpdateTime,
	}
}
