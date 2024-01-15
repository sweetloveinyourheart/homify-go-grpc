package repositories

import (
	"homify-go-grpc/internal/authentication-service/models"

	"gorm.io/gorm"
)

type IAccountRepository interface {
	CreateAccount(account *models.Account) error
	GetAccountByID(id uint) (*models.Account, error)
	GetAccountByField(fieldName string, value interface{}) (*models.Account, error)
	UpdateAccount(account *models.Account) error
	DeleteAccount(account *models.Account) error
}

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) IAccountRepository {
	return &AccountRepository{db}
}

func (r *AccountRepository) CreateAccount(account *models.Account) error {
	return r.db.Create(account).Error
}

func (r *AccountRepository) GetAccountByID(id uint) (*models.Account, error) {
	account := new(models.Account)
	err := r.db.First(account, id).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r *AccountRepository) GetAccountByField(fieldName string, value interface{}) (*models.Account, error) {
	account := &models.Account{}
	if err := r.db.Where(fieldName+" = ?", value).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (r *AccountRepository) UpdateAccount(account *models.Account) error {
	return r.db.Save(account).Error
}

func (r *AccountRepository) DeleteAccount(account *models.Account) error {
	return r.db.Delete(account).Error
}
