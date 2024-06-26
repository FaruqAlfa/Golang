package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
	"errors"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)

	FetchByID(id int) (*model.Session, error)
}

type sessionsRepoImpl struct {
	db *sql.DB
}

func NewSessionRepo(db *sql.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (u *sessionsRepoImpl) AddSessions(session model.Session) error {
	_, err := u.db.Exec("INSERT INTO sessions (token, username, expiry) VALUES ($1, $2, $3)", session.Token, session.Username, session.Expiry)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *sessionsRepoImpl) DeleteSession(token string) error {
	_, err := u.db.Exec("DELETE FROM sessions WHERE token = $1", token)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	_, err := u.db.Exec("UPDATE sessions SET token = $1, expiry = $2 WHERE username = $3", session.Token, session.Expiry, session.Username)
	if err != nil {
		return err
	}
	return nil
	 // TODO: replace this
}

func (u *sessionsRepoImpl) SessionAvailName(name string) error {
	row := u.db.QueryRow("SELECT id FROM sessions WHERE username = $1", name)//mengambil data dari database berdasarkan username
	var id int//membuat variabel baru
	err := row.Scan(&id)//mengambil data dari database berdasarkan id
	if err == sql.ErrNoRows {//kondisi jika user tidak ditemukan
		return errors.New("session not found")
	} else if err != nil {
		return err
	}
	return nil
	// TODO: replace this
}

func (u *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	
	row := u.db.QueryRow("SELECT id, token, username, expiry FROM sessions WHERE token = $1", token)//mengambil data dari database berdasarkan token

	var session model.Session//membuat variabel baru untuk menyimpan session
	err := row.Scan(&session.ID, &session.Token, &session.Username, &session.Expiry)
	if err == sql.ErrNoRows {//kondisi jika user tidak ditemukan
		return model.Session{}, errors.New("session not found")
	} else if err != nil {
		return model.Session{}, err
	}

	return session, nil // TODO: replace this
}

func (u *sessionsRepoImpl) FetchByID(id int) (*model.Session, error) {
	row := u.db.QueryRow("SELECT id, token, username, expiry FROM sessions WHERE id = $1", id)//mengambil data dari database berdasarkan id

	var session model.Session//membuat variabel baru untuk menyimpan session
	err := row.Scan(&session.ID, &session.Token, &session.Username, &session.Expiry)
	if err != nil {
		return nil, err
	}

	return &session, nil
}
