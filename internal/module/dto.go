package module

import (
	"net/http"
	"regexp"
	"time"

	"github.com/satori/uuid"
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
	if err != nil {
		return false
	}
	c.Password = string(password)
	return true
}

type SignUserDTO struct {
	Login          string    `json: "login"`
	Password       string    `json: "password"`
	UserId         int       `json: "user_id"`
	UUID           uuid.UUID `json: "uuid"`
	CreateTimeUUID time.Time `json: "createuuid"`
	Duration       time.Time `json: "duration"`
}

func (c *SignUserDTO) Add(r *http.Request) {
	c.Login = r.FormValue("login")
	c.Password = r.FormValue("password")
}

type HomePageDTO struct {
	UserId   int
	UserName string
	Posts    []ShowPostDTO
}

type ShowPostDTO struct {
	Id     int       `json: "id"`
	Title  string    `json: "title"`
	Post   string    `json: "posttext"`
	Date   time.Time `json: "date"`
	UserId int       `json: "userid"`
}
