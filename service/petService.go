package service

import (
	"example/web-service-gin/database"
	"example/web-service-gin/models"
	"example/web-service-gin/repository"
)

type PetService struct {
	PetRepo  *repository.PetRepository
	UserRepo *repository.UserRepository
}

// NewPetService inizializza i repository e restituisce un servizio
func NewPetService() *PetService {
	if database.DB == nil {
		database.ConnectDatabase()
	}

	petRepo := repository.NewPetRepository(database.DB)
	userRepo := repository.NewUserRepository(database.DB) // Aggiunto il repository User

	return &PetService{
		PetRepo:  petRepo,
		UserRepo: userRepo,
	}
}

// GetPets restituisce tutti gli animali
func (s *PetService) GetPets() ([]models.Pet, error) {
	return s.PetRepo.GetPets()
}

// GetPetByID restituisce un animale per ID
func (s *PetService) GetPetByID(id string) (*models.Pet, error) {
	return s.PetRepo.GetPetByID(id)
}

// GetPetsByUserID restituisce tutti gli animali di un utente
func (s *PetService) GetPetsByUserID(userID string) ([]models.Pet, error) {
	return s.PetRepo.GetPetsByUserID(userID)
}

// CreatePet crea un nuovo animale
func (s *PetService) CreatePet(pet *models.Pet) error {
	return s.PetRepo.CreatePet(pet)
}

// UpdatePet aggiorna un animale esistente
func (s *PetService) UpdatePet(pet *models.Pet) error {
	return s.PetRepo.UpdatePet(pet)
}

// DeletePet elimina un animale esistente
func (s *PetService) DeletePet(pet *models.Pet) error {
	return s.PetRepo.DeletePet(pet)
}

// DeletePetByID elimina un animale per ID
func (s *PetService) DeletePetByID(id string) error {
	pet, err := s.GetPetByID(id)
	if err != nil {
		return err
	}
	return s.PetRepo.DeletePet(pet)
}

// DeletePetsByUserID elimina tutti gli animali di un utente
func (s *PetService) DeletePetsByUserID(userID string) error {
	return s.PetRepo.DeletePetsByUserID(userID)
}

// DeleteAllPets elimina tutti gli animali
func (s *PetService) DeleteAllPets() error {
	return s.PetRepo.DeleteAllPets()
}
