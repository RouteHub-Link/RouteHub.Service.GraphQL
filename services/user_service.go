package services

import (
	"errors"

	database_models "github.com/RouteHub-Link/routehub-service-graphql/database/models"
	database_relations "github.com/RouteHub-Link/routehub-service-graphql/database/relations"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (u UserService) Users() (users []*database_models.User, err error) {
	err = u.DB.Find(&users).Error
	return
}

func (u UserService) UserCompany(userId uuid.UUID) (company []*database_models.Company, err error) {
	userCompanies := []database_relations.UserCompany{}
	err = u.DB.Where("user_id = ?", userId).Find(&userCompanies).Error
	if err != nil {
		return
	}

	companyIDs := []uuid.UUID{}
	for _, userCompany := range userCompanies {
		companyIDs = append(companyIDs, userCompany.CompanyID)
	}

	err = u.DB.Where("id IN ?", companyIDs).Find(&company).Error
	return
}

func (u UserService) Login(email string, password string) (user *database_models.User, err error) {
	err = u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return
	}

	if !CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid password")
	}

	return
}

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u UserService) User(id uuid.UUID) (user *database_models.User, err error) {
	err = u.DB.Where("id = ?", id).First(&user).Error
	return
}
