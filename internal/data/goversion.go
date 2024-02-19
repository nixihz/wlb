package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/gorm"
	"wlb/internal/biz"
)

type goVersionRepo struct {
	log *log.Helper
}

// NewGoVersionRepo .
func NewGoVersionRepo(client *gorm.DB, logger log.Logger) biz.GoVersionRepo {
	return &goVersionRepo{
		log: log.NewHelper(logger),
	}
}

func (r *goVersionRepo) Save(ctx context.Context, g *biz.GoVersion) (*biz.GoVersion, error) {
	return g, nil
}

func (r *goVersionRepo) Update(ctx context.Context, g *biz.GoVersion) (*biz.GoVersion, error) {
	return g, nil
}

func (r *goVersionRepo) FindByID(context.Context, int64) (*biz.GoVersion, error) {
	return nil, nil
}

func (r *goVersionRepo) ListByHello(context.Context, string) ([]*biz.GoVersion, error) {
	return nil, nil
}

func (r *goVersionRepo) ListAll(context.Context) ([]*biz.GoVersion, error) {
	return nil, nil
}
