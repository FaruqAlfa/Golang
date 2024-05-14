package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserService interface {
	Register(user *model.User) (model.User, error)
	Login(user *model.User) (token *string, err error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepository repo.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Register(user *model.User) (model.User, error) {
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return *user, err
	}

	if dbUser.Email != "" || dbUser.ID != 0 {
		return *user, errors.New("email already exists")
	}

	user.CreatedAt = time.Now()

	newUser, err := s.userRepo.CreateUser(*user)
	if err != nil {
		return *user, err
	}

	return newUser, nil
}

func (s *userService) Login(user *model.User) (token *string, err error) {
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if dbUser.Email != user.Email {
		return nil, errors.New("user not found")
	}

	if dbUser.Email != user.Email || dbUser.Password != user.Password {
		return nil, errors.New("wrong email or password")
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	// Buat claims berisi data username dan role yang akan kita embed ke JWT
	claims := &model.Claims{
		UserID: dbUser.ID,
	StandardClaims: jwt.StandardClaims{
			// expiry time menggunakan time millisecond
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Buat token menggunakan encoded claim dengan salah satu algoritma yang dipakai
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Buat JWT string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan (proses encoding JWT)
	tokenString, err := newToken.SignedString(model.JwtKey)
	if err != nil {
		// return internal error ketika ada kesalahan saat pembuatan JWT string
		return nil, err
	}

	return &tokenString, nil // TODO: replace this
}

func (s *userService) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	return s.userRepo.GetUserTaskCategory()// TODO: replace this
}
