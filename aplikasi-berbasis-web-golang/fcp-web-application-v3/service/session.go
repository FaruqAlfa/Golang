package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type SessionService interface {
	GetSessionByEmail(email string) (model.Session, error)
}

type sessionService struct {
	sessionRepo repo.SessionRepository
}

func NewSessionService(sessionRepo repo.SessionRepository) *sessionService {
	return &sessionService{sessionRepo}
}

func (c *sessionService) GetSessionByEmail(email string) (model.Session, error) {
	var sessionRepo model.Session //membuat variabel penampung unntuk menyimpan session
	sessionRepo, err := c.sessionRepo.SessionAvailEmail(email)//mengambil session berdasarkan email
	if err != nil {	//pengecekan error
		return model.Session{}, err
	}

	return sessionRepo, nil//mengembalikan session
	// return model.Session{}, nil // TODO: replace this
}
