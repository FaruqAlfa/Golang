package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"time"
)

type SessionRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailEmail(email string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	TokenExpired(session model.Session) bool
}

type sessionsRepo struct {
	filebasedDb *filebased.Data
}

func NewSessionsRepo(filebasedDb *filebased.Data) *sessionsRepo {
	return &sessionsRepo{filebasedDb}
}

func (u *sessionsRepo) AddSessions(session model.Session) error {
	u.filebasedDb.AddSession(session)
	return nil // TODO: replace this
}

func (u *sessionsRepo) DeleteSession(token string) error {
	session, err := u.SessionAvailToken(token)//mengambil session berdasarkan token
	if err != nil {//pengecekan error
		return err
	}

	u.filebasedDb.DeleteSession(session.Token)//menghapus session
	return nil // TODO: replace this
}

func (u *sessionsRepo) UpdateSessions(session model.Session) error {
	u.filebasedDb.UpdateSession(session)
	return nil // TODO: replace this
}

func (u *sessionsRepo) SessionAvailEmail(email string) (model.Session, error) {
	sessionsRepo, err := u.filebasedDb.SessionAvailEmail(email)//mengambil session berdasarkan email
	if err != nil {//pengecekan error
		return model.Session{}, err
	}
	return sessionsRepo, nil//mengembalikan session
	// return model.Session{}, nil // TODO: replace this
}

func (u *sessionsRepo) SessionAvailToken(token string) (model.Session, error) {
	sessionsRepo, err := u.filebasedDb.SessionAvailToken(token)//mengambil session berdasarkan token
	if err != nil {//pengecekan error
		return model.Session{}, err
	}
	return sessionsRepo, nil//mengembalikan session
	// return model.Session{}, nil // TODO: replace this
}

func (u *sessionsRepo) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err := u.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, err
	}

	return session, nil
}

func (u *sessionsRepo) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
