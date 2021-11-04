package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/ppcamp/go-graphql-with-auth/internal/models/usermodels"
	"github.com/ppcamp/go-graphql-with-auth/internal/utils"
)

func (t *UserTransaction) CreateUser(payload *usermodels.UserMutationPayload) (user usermodels.UserEntity, err error) {
	user = usermodels.UserEntity{}
	user.Email = payload.Password
	user.Nick = payload.Nick
	user.Password = payload.Password

	sql := `
	INSERT INTO users(password,email,nick)
	VALUES (:password,:email,:nick)
	RETURNING id,updated_at
	`
	rows, err := t.tx.NamedQuery(sql, payload)
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.UpdatedAt)
	}

	return
}

// TODO: improve updated_at filter
// See https://www.postgresql.org/docs/8.1/queries-limit.html
func (t *UserTransaction) FindUsers(filter *usermodels.UserQueryPayload) (users []usermodels.UserEntity, err error) {
	users = []usermodels.UserEntity{}

	sql := `
	SELECT id, password, email, nick, updated_at
	FROM users
	WHERE id = COALESCE(:id, id)
		AND email = COALESCE(:email, email)
		AND nick = COALESCE(:nick, nick)
		AND updated_at = COALESCE(:updated_at, updated_at)
	OFFSET :skip
	LIMIT :take
	`
	stmt := utils.Must(t.tx.PrepareNamed(sql)).(*sqlx.NamedStmt)
	err = stmt.Select(&users, filter)
	return
}
