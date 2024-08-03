package models

import (
	"time"

	"gorm.io/gorm"
)

// User Model
// type User struct {
//     gorm.Model
//     Name     string
//     Email    string
//     Password string
//     NonTechnical []NonTechnical `gorm:"foreignKey:UserID"`
//     Technical    []Technical    `gorm:"foreignKey:UserID"`
//     GK           []GK           `gorm:"foreignKey:UserID"`
// }

type User struct {
	ID                 string         `gorm:"type:uuid;primary_key"`
	Name               string         `gorm:"type:varchar(100);not null"`
	Email              string         `gorm:"type:varchar(100);unique;not null"`
	Password           string         `gorm:"type:varchar(100);not null"`
	Role               *string        `gorm:"type:varchar(50);default:'user';not null"`
	Provider           *string        `gorm:"type:varchar(50);default:'local';not null"`
	Photo              *string        `gorm:"not null;default:'default.png'"`
	Verified           *bool          `gorm:"not null;default:false"`
	VerificationCode   string         `gorm:"type:varchar(100);"`
	PasswordResetToken string         `gorm:"type:varchar(100);"`
	PasswordResetAt    time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	CreatedAt          time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	NonTechnical       []NonTechnical `gorm:"foreignKey:UserID"`
	Technical          []Technical    `gorm:"foreignKey:UserID"`
	GK                 []GK           `gorm:"foreignKey:UserID"`
}

// NonTechnical represents the non-technical subjects
type NonTechnical struct {
	gorm.Model
	Name   string `gorm:"not null"` // E.g., "Financial", "Establishment", etc.
	Books  []Book `gorm:"foreignKey:NonTechnicalID"`
	UserID uint
}

// Technical represents the technical subjects with titles like C&W, Workshop
type Technical struct {
	gorm.Model
	Title    string     `gorm:"not null"` // E.g., "C&W", "Workshop"
	CAndW    []CAndW    `gorm:"foreignKey:TechnicalID"`
	Workshop []Workshop `gorm:"foreignKey:TechnicalID"`
	UserID   uint
}

// CAndW represents the Carriage & Wagon section
type CAndW struct {
	gorm.Model
	SubCategory string `gorm:"not null"` // E.g., "Airbrake", "Bogie", etc.
	Books       []Book `gorm:"foreignKey:CAndWID"`
	TechnicalID uint
}

// Workshop represents the workshop section
type Workshop struct {
	gorm.Model
	SubCategory string `gorm:"not null"` // E.g., "Incentive", "Workorder", etc.
	Books       []Book `gorm:"foreignKey:WorkshopID"`
	TechnicalID uint
}

// GK represents the General Knowledge section
type GK struct {
	gorm.Model
	SubCategory string `gorm:"not null"` // E.g., "General Knowledge", "Aptitude", etc.
	Books       []Book `gorm:"foreignKey:GKID"`
	UserID      uint
}

// Book represents the books posted in various categories
type Book struct {
	gorm.Model
	Title          string `gorm:"not null"` // Book title
	Content        string `gorm:"not null"` // Book content as string
	NonTechnicalID uint
	CAndWID        uint
	WorkshopID     uint
	GKID           uint
}
