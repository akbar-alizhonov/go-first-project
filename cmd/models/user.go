package models

import "time"

type Users struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Measurements struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
	BodyFat float64 `json:"body_fat"`
	CreatedAt time.Time `json:"created_at"`
}