package repository

import (
	"example/web-service-gin/models"

	"gorm.io/gorm"
)

type PetRepository struct {
	DB *gorm.DB
}

func NewPetRepository(db *gorm.DB) *PetRepository {
	return &PetRepository{DB: db}
}

func (r *PetRepository) GetPets() ([]models.Pet, error) {
	var pets []models.Pet
	err := r.DB.Find(&pets).Error
	return pets, err
}

func (r *PetRepository) GetPetByID(id string) (*models.Pet, error) {
	var pet models.Pet
	err := r.DB.Preload("User").First(&pet, id).Error
	return &pet, err
}

func (r *PetRepository) GetPetsByUserID(userID string) ([]models.Pet, error) {
	var pets []models.Pet
	err := r.DB.Where("user_id = ?", userID).Find(&pets).Error
	return pets, err
}

func (r *PetRepository) CreatePet(pet *models.Pet) error {
	return r.DB.Create(pet).Error
}

func (r *PetRepository) UpdatePet(pet *models.Pet) error {
	return r.DB.Save(pet).Error
}

func (r *PetRepository) DeletePet(pet *models.Pet) error {
	return r.DB.Delete(pet).Error
}

func (r *PetRepository) DeletePetByID(id uint) error {
	return r.DB.Delete(&models.Pet{}, id).Error
}

func (r *PetRepository) DeletePetsByUserID(userID string) error {
	return r.DB.Where("user_id = ?", userID).Delete(&models.Pet{}).Error
}

func (r *PetRepository) DeleteAllPets() error {
	return r.DB.Delete(&models.Pet{}).Error
}
