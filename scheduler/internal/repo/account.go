package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ideal-rucksack/workflow-scheduler/scheduler/internal/repo/entities"
	"github.com/jmoiron/sqlx"
)

type AccountRepo struct {
	db *sqlx.DB
}

func (a AccountRepo) TableName() string {
	return "account"
}

func (a AccountRepo) Columns() []string {
	return []string{
		"id",
		"nickname",
		"code",
		"email",
		"username",
		"password",
		"secret",
		"status",
		"refresh_token",
		"create_at",
		"modify_at",
		"deleted",
	}
}

func (a AccountRepo) Columns2Query() string {
	columns := a.Columns()
	var columns2Query string
	for i, column := range columns {
		if i == 0 {
			columns2Query = column
		} else {
			columns2Query = fmt.Sprintf("%s, %s", columns2Query, column)
		}
	}
	return columns2Query
}

func (a AccountRepo) Insert(account entities.Account) error {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (:id, :nickname, :code, :email, :username, :password, :secret, :status, :refresh_token, :create_at, :modify_at, :deleted)", a.TableName(), a.Columns2Query())
	_, err := a.db.NamedExec(query, account)
	return err
}

func (a AccountRepo) QueryByUsername(username string, deleted bool) (*entities.Account, error) {
	var account entities.Account
	query := fmt.Sprintf("SELECT %s FROM %s WHERE username = ? and deleted = ?", a.Columns2Query(), a.TableName())
	err := a.db.Get(&account, query, username, deleted)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

func (a AccountRepo) QueryByEmail(email string, deleted bool) (*entities.Account, error) {
	var account entities.Account
	query := fmt.Sprintf("SELECT %s FROM %s WHERE email = ? and deleted = ?", a.Columns2Query(), a.TableName())
	err := a.db.Get(&account, query, email, deleted)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

func (a AccountRepo) Update(account entities.Account) error {
	query := fmt.Sprintf("UPDATE %s SET nickname = :nickname, code = :code, email = :email, username = :username, password = :password, secret = :secret, status = :status, refresh_token = :refresh_token, create_at = :create_at, modify_at = :modify_at, deleted = :deleted WHERE id = :id", a.TableName())
	_, err := a.db.NamedExec(query, account)
	return err
}

func (a AccountRepo) QuerySignIn(username, password string, deleted bool) (*entities.Account, error) {
	var account entities.Account
	query := fmt.Sprintf("SELECT %s FROM %s WHERE username = ? AND password = ? and deleted = ?", a.Columns2Query(), a.TableName())
	err := a.db.Get(&account, query, username, password, deleted)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

func (a AccountRepo) QueryById(id int64, deleted bool) (*entities.Account, error) {
	var account entities.Account
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = ? and deleted = ?", a.Columns2Query(), a.TableName())
	err := a.db.Get(&account, query, id, deleted)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

func NewAccountRepo(db *sqlx.DB) *AccountRepo {
	return &AccountRepo{db: db}
}
