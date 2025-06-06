package repository

import (
	"context"
	"server-skeleton/api/page/model"
)

func (r *Repository) GetById(ctx context.Context, id string) (*model.Page, error) {

	var page model.Page

	if err := r.db.WithContext(ctx).First(&page, id).Error; err != nil {
		return nil, err
	}
	return &page, nil
}

func (r *Repository) Create(ctx context.Context, page *model.Page) error {
	return r.db.WithContext(ctx).Create(page).Error
}
