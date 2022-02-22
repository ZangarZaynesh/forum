package module

import (
	"net/http"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserDTO struct {
	Login           string `json: "login"`
	Password        string `json: "password"`
	ConfirmPassword string `json: "confirmpassword"`
	Email           string `json: "email"`
}

func (c *CreateUserDTO) Add(r *http.Request) {
	c.Login = r.FormValue("login")
	c.Email = r.FormValue("email")
	c.Password = r.FormValue("password")
	c.ConfirmPassword = r.FormValue("confirm")
}

func (c *CreateUserDTO) IsEmailValid() bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(c.Email)
}

func (c *CreateUserDTO) CheckPassConfirm() bool {
	return c.Password == c.ConfirmPassword
}

func (c *CreateUserDTO) GeneratePassword() bool {
	password, err := bcrypt.GenerateFromPassword([]byte(c.Password), 8)
	c.Password = string(password)
	if err != nil {
		return false
	}
	return true
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
	CategoryId int       `json: "categoryid"`
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

type CreateCommentDTO struct {
	Comment string    `json: "comment"`
	Date    time.Time `json: "date"`
	UserId  int       `json: "userid"`
	PostId  int       `json: "postid"`
}

type UpdateCommentDTO struct {
	Comment string    `json: "comment"`
	Date    time.Time `json: "date"`
}

type DeleteCommentDTO struct {
	Id     int `json: "id"`
	PostId int `json: "PostId"`
}

type CreatePostLikesDTO struct {
	Value  byte `json: "value"`
	UserId int  `json: "userid"`
	PostId int  `json: "postid"`
}

type UpdatePostLikesDTO struct {
	Value byte `json: "value"`
}

type CreateCommentLikesDTO struct {
	Value     string `json: "value"`
	UserId    int    `json: "userid"`
	CommentId int    `json: "commentid"`
}

type UpdateCommentLikesDTO struct {
	Value string `json: "value"`
}

type CreateSessionDTO struct {
	Key    string    `json: "key"`
	Date   time.Time `json: "date"`
	UserId int       `json: "userid"`
}
