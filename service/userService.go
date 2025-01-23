package service

import (
	"example/web-service-gin/database"
	"example/web-service-gin/models"
	"example/web-service-gin/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

// NewUserService inizializza il repository e restituisce un servizio
func NewUserService() *UserService {
	// Inizializza il database se non è già connesso
	if database.DB == nil {
		database.ConnectDatabase()
	}

	// Crea il repository
	userRepo := repository.NewUserRepository(database.DB)

	// Ritorna il servizio con il repository già istanziato
	return &UserService{UserRepo: userRepo}
}

// GetUsers restituisce tutti gli utenti
func (s *UserService) GetUsers() ([]models.User, error) {
	return s.UserRepo.GetUsers()
}

// GetUserByID restituisce un utente per ID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.UserRepo.GetUserByID(id)
}

// CreateUser crea un nuovo utente
func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepo.CreateUser(user)
}
