package postgre

import (
	"context"

	"github.com/Mitra-Apps/be-user-service/domain/user/entity"

	"gorm.io/gorm"
)

type Postgre struct {
	db *gorm.DB
}

func NewPostgre(db *gorm.DB) *Postgre {
	return &Postgre{db}
}

func (p *Postgre) GetAll(ctx context.Context) ([]*entity.User, error) {
	var accounts []*entity.User
	res := p.db.Order("created_at DESC").Find(&accounts)
	if res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return accounts, nil
}

func (p *Postgre) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user *entity.User
	res := p.db.Where("email = ?", email).First(&user)
	if res.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}
