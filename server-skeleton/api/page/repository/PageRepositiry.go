package repository

import (
	"context"
	"server-skeleton/api/page/model"
)

type PageRepositoryInterface interface {
	GetById(ctx context.Context, id string)
	Create(ctx context.Context, page *model.Page)
}
