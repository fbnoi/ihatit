package main

import (
	"context"
	"database/sql"
	"strings"
)

const sql_get_user_by_name_and_age = "SELECT * FROM user WHERE name = ? AND age = ?"
const sql_add_user = `INSERT INTO user (id, name, gender, email) VALUES (?, ?, ?, ?)`
const sql_add_users = `INSERT INTO user (id, name, gender, email) VALUES`

func (q *Queries) GetUserByNameAndAge(ctx context.Context, name string, age int) (*User, error) {
	row := q.db.QueryRowContext(ctx, sql_get_user_by_name_and_age)
	u := &User{}
	err := row.Scan(&u.Id, &u.Name, &u.Age)

	return u, err
}

func (q *Queries) AddUser(ctx context.Context, user *User) (sql.Result, error) {
	return q.db.ExecContext(ctx, sql_add_user, user.Name, user.Age)
}

func (q *Queries) AddUsers(ctx context.Context, users []*User) (sql.Result, error) {
	sqlBuilder := &strings.Builder{}
	sqlBuilder.WriteString(sql_add_users)
	vals := []interface{}{}
	for i, user := range users {
		sqlBuilder.WriteString("(?,?)")
		vals = append(vals, user.Name, user.Age)
		if i != len(users) {
			sqlBuilder.WriteString(",")
		}
	}

	return q.db.ExecContext(ctx, sqlBuilder.String(), vals...)
}
