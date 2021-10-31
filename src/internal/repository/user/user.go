package user

import (
	"github.com/ppcamp/go-graphql-with-auth/internal/models"
)

func (t *UserTransaction) CreateUser(payload *models.User) (user models.User, err error) {
	user = models.User{}
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
