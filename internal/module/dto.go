package module

import "time"

type CreateUserDTO struct {
	Login           string `json: "login"`
	Password        string `json: "password"`
	ConfirmPassword string `json: "confirmpassword"`
	Email           string `json: "email"`
}

type ReadUserDTO struct {
	Id    int    `json: "id"`
	Login string `json: "login"`
}

type UpdateUserDTO struct {
	Login    string `json: "login"`
	Password string `json: "password"`
	Email    string `json: "email"`
}

type DeleteUserDTO struct {
	Id int `json: "id"`
}

type CreatePostDTO struct {
	Title      string    `json: "title"`
	PostText   string    `json: "posttext"`
	Date       time.Time `json: "date"`
	UserId     int       `json: "userid"`
	CategoryId int       `json: "categoryid"`
}

type ReadPostDTO struct {
	Id     int `json: "id"`
	UserId int `json: "userid"`
}

type UpdatePostDTO struct {
	Title      string    `json: "title"`
	PostText   string    `json: "posttext"`
	Date       time.Time `json: "date"`
	CategoryId int       `json: "categoryid"`
}

type DeletePostDTO struct {
	Id     int `json: "id"`
	UserId int `json: "userid"`
}
