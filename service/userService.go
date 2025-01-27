package service

import (
	"errors"
	"example/web-service-gin/database"
	"example/web-service-gin/dto"
	"example/web-service-gin/models"
	"example/web-service-gin/repository"
	"example/web-service-gin/utils"
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
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.UserRepo.CreateUser(user)
}

// func (s *UserService) Login(userLogin *dto.LoginUser) error {
// 	user, err := s.UserRepo.GetUserByEmail(userLogin.Email)
// 	if err != nil {
// 		return err // Restituisci l'errore se l'utente non viene trovato
// 	}
// 	match := utils.VerifyPassword(userLogin.Password, user.Password)
// 	if !match {
// 		return errors.New("invalid credentials") // Restituisci un errore se la password è errata
// 	}
// 	return nil // Login riuscito, nessun errore
// }

func (s *UserService) Login(userLogin *dto.LoginUser) (string, error) {
	user, err := s.UserRepo.GetUserByEmail(userLogin.Email)
	if err != nil {
		return "", err // Restituisci l'errore se l'utente non viene trovato
	}

	match := utils.VerifyPassword(userLogin.Password, user.Password)
	if !match {
		return "", errors.New("invalid credentials") // Restituisci un errore se la password è errata
	}

	token, err := utils.CreateToken(userLogin.Email) // Funzione per generare un JWT
	if err != nil {
		return "", err
	}

	return token, nil // Restituisce il token se il login è corretto
}
