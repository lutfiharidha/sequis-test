package app

import (
	"time"
)

type DBConfig struct {
	DBDriver string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
}

// type User struct {
// 	ID        string    `gorm:"uniqueIndex" json:"id"`
// 	Email     string    `json:"email" gorm:"unique:not null"  form:"email"`
// 	Name      string    `json:"name" gorm:"not null"  form:"name"`
// 	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
// 	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
// }

// func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	uuid := uuid.New()
// 	tx.Statement.SetColumn("ID", uuid)
// 	return
// }

type Friend struct {
	Requestor string `json:"requestor" gorm:"not null;size:191;primaryKey;autoIncrement:false" form:"requestor" binding:"required"`
	To        string `json:"to" gorm:"not null;size:191;primaryKey;autoIncrement:false" form:"to" binding:"required"`
	Status    string `json:"status"`
	// UserRequestor User   `json:"user_requestor" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:Requestor;references:ID"`
	// UserTo        User   `json:"user_to" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:To;references:ID"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// func (friend *Friend) BeforeCreate(tx *gorm.DB) (err error) {
// 	uuid := uuid.New()
// 	tx.Statement.SetColumn("ID", uuid)
// 	return
// }

type FriendRequest struct {
	Requestor string `json:"requestor" form:"requestor" binding:"required"`
	To        string `json:"to" form:"to" binding:"required"`
}

type FriendResponse struct {
	Success bool `json:"success"`
}

type BlockRequest struct {
	Requestor string `json:"requestor" form:"requestor" binding:"required"`
	Block     string `json:"block" form:"block" binding:"required"`
}

type ListRequest struct {
	Email string `json:"email" form:"email" binding:"required"`
}

type ListResponse struct {
	Friends []string `json:"friends"`
}

type ListFriendsRequestResponse struct {
	Request []FriendsRequestResponse `json:"requests"`
}

type FriendsRequestResponse struct {
	Requestor string `json:"requestator"`
	Status    string `json:"status"`
}

type CommonRequest struct {
	Friends []string `json:"friends" form:"friends" binding:"required"`
}

type CommonResponse struct {
	Success bool         `json:"success"`
	Friends ListResponse `json:"friends"`
	Count   int          `json:"count"`
}
