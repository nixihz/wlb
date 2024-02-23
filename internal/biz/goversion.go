package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
)

// GoVersionis a GoVersionmodel.
type GoVersion struct {
	ID             int64
	Version        string
	IsMajorVersion bool
	VersionDate    time.Time
	CreateTime     time.Time
	UpdateTime     time.Time
}

// GoVersionRepo is a Greater repo.
type GoVersionRepo interface {
	Save(context.Context, *GoVersion) (*GoVersion, error)
	Update(context.Context, *GoVersion) (*GoVersion, error)
	FindByID(context.Context, int64) (*GoVersion, error)
	ListByHello(context.Context, string) ([]*GoVersion, error)
	ListAll(context.Context) ([]*GoVersion, error)
}

// GoVersionUsecase is a GoVersionusecase.
type GoVersionUsecase struct {
	repo GoVersionRepo
	log  *log.Helper
}

// NewGoVersionUsecase new a GoVersionUsecase.
func NewGoVersionUsecase(repo GoVersionRepo, logger log.Logger) *GoVersionUsecase {
	return &GoVersionUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGoVersion create a GoVersion, and returns the new GoVersion.
func (uc *GoVersionUsecase) CreateGoVersion(ctx context.Context, g *GoVersion) (*GoVersion, error) {
	uc.log.WithContext(ctx).Infof("CreateGoVersion: %v", g.Version)
	return uc.repo.Save(ctx, g)
}
