package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"
	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	err := s.db.Create(&session).Error//menambahkan session ke table sessions
	if err != nil {
		return err
	}
	return nil  // TODO: replace this
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	//menghapus session berdasarkan token
	err := s.db.Where("token = ?", token).Delete(&model.Session{}).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	//mengupdate session ke table sessions
	err := s.db.Model(&model.Session{}).Where("username = ?", session.Username).Updates(model.Session{
		Token:  session.Token,
		Expiry: session.Expiry,
	}).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
	//mencari session berdasarkan username
	var session model.Session
	err := s.db.Where("username = ?", name).First(&session).Error
	if err == gorm.ErrRecordNotFound {
		return errors.New("session not found")
	} else if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	//mencari session berdasarkan token di table sessions
	var session model.Session
	err := s.db.Where("token = ?", token).First(&session).Error
	if err == gorm.ErrRecordNotFound {
		return model.Session{}, errors.New("session not found")
	} else if err != nil {
		return model.Session{}, err
	}
	return session, nil// TODO: replace this
}
