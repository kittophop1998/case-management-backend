package model

type UserResponse struct {
	Id             string     `json:"id"`
	UserName       string     `json:"userName"`
	Email          string     `json:"email"`
	Role           RoleItem   `json:"role"`
	Team           string     `json:"team"`
	Center         CenterItem `json:"center"`
	IsActive       string     `json:"isActive"`
	CreateDatetime string     `json:"createDatetime"`
	UpdateDatetime string     `json:"updateDatetime"`
}

type RoleItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CenterItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
