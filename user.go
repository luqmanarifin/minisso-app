package main

import "time"

type User struct {
	Id          int64     `json:"id"`
	UserId      string    `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Picture     string    `xorm:"text" json:"picture"`
	Gender      string    `json:"gender"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	LatestLogin time.Time `json:"latest_login"`
	LastIp      string    `json:"last_ip"`
	Connection  string    `json:"connection"`
	CreatedAt   time.Time `xorm:"created" json:"created_at"`
	UpdatedAt   time.Time `xorm:"updated" json:"updated_at"`
}
