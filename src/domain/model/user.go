package model

type User struct {
	Id        uint    `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name"`
	Email     string  `json:"email" gorm:"not null"`
	Password  string  `json:"password" gorm:"not null"`
	Token     *string `json:"token"`
	CreatedAt int64   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64   `json:"updated_at" gorm:"autoUpdateTime:milli"`
}
