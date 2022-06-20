package main

import "time"

type User struct {
	Id        int
	Name      string
	Email     string
	Age       int
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Posts     []*Post
}

type Post struct {
	Id      int
	Title   string
	Content string
	Tag     string
	Author  *User
}
