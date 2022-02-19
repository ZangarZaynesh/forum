package module

import "time"

type User struct {
	Id       int    `json: "id"`
	Login    string `json: "login"`
	Password string `json: "password"`
	Email    string `json: "email"`
}

type Category struct {
	Id   int    `json: "id"`
	Name string `json: "name"`
}

type Post struct {
	Id         int       `json: "id"`
	Title      string    `json: "title"`
	PostText   string    `json: "posttext"`
	Date       time.Time `json: "date"`
	UserId     int       `json: "userid"`
	CategoryId int       `json: "categoryid"`
}

type PostLikes struct {
	Id     int  `json: "id"`
	Value  byte `json: "value"`
	UserId int  `json: "userid"`
	PostId int  `json: "postid"`
}

type Comment struct {
	Id      int       `json: "id"`
	Comment string    `json: "comment"`
	Date    time.Time `json: "date"`
	UserId  int       `json: "userid"`
	PostId  int       `json: "postid"`
}

type CommentLikes struct {
	Id        int    `json: "id"`
	Value     string `json: "value"`
	UserId    int    `json: "userid"`
	CommentId int    `json: "commentid"`
}

type Session struct {
	Id     int       `json: "id"`
	Key    string    `json: "key"`
	Date   time.Time `json: "date"`
	UserId int       `json: "userid"`
}
