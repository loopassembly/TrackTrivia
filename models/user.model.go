package models

import (
    "time"

    "github.com/go-playground/validator/v10"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// User Model
// type User struct {
//     ID                   string              `gorm:"type:uuid;primary_key"`
//     Name                 string              `gorm:"type:varchar(100);not null"`
//     Email                string              `gorm:"type:varchar(100);unique;not null"`
//     Password             string              `gorm:"type:varchar(100);not null"`
//     Role                 *string             `gorm:"type:varchar(50);default:'user';not null"`
//     Provider             *string             `gorm:"type:varchar(50);default:'local';not null"`
//     Photo                *string             `gorm:"not null;default:'default.png'"`
//     Verified             *bool               `gorm:"not null;default:false"`
//     VerificationCode     string              `gorm:"type:varchar(100);"`
//     PasswordResetToken   string              `gorm:"type:varchar(100);"`
//     PasswordResetAt      time.Time           `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
//     CreatedAt            time.Time           `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
//     UpdatedAt            time.Time           `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
//     NonTechnical []NonTechnical `gorm:"foreignKey:UserID"`
//     Technical    []Technical    `gorm:"foreignKey:UserID"`
//     GK           []GK           `gorm:"foreignKey:UserID"`
// }

func (u *User) BeforeCreate(tx *gorm.DB) error {
    uuid, err := uuid.NewUUID()
    if err != nil {
        return err
    }
    u.ID = uuid.String()
    return nil
}

// SignUpInput struct
type SignUpInput struct {
    Name            string `json:"name" validate:"required"`
    Email           string `json:"email" validate:"required"`
    Password        string `json:"password" validate:"required,min=8"`
    PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
    Photo           string `json:"photo"`
}

// SignInInput struct
type SignInInput struct {
    Email    string `json:"email" validate:"required"`
    Password string `json:"password" validate:"required"`
}

// UserResponse struct
type UserResponse struct {
    ID        uuid.UUID `json:"id,omitempty"`
    Name      string    `json:"name,omitempty"`
    Email     string    `json:"email,omitempty"`
    Role      string    `json:"role,omitempty"`
    Photo     string    `json:"photo,omitempty"`
    Provider  string    `json:"provider"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// FilterUserRecord function
func FilterUserRecord(user *User) UserResponse {
    id := user.ID
    return UserResponse{
        ID:        uuid.MustParse(id),
        Name:      user.Name,
        Email:     user.Email,
        Role:      *user.Role,
        Photo:     *user.Photo,
        Provider:  *user.Provider,
        CreatedAt: user.CreatedAt,
        UpdatedAt: user.UpdatedAt,
    }
}

// ErrorResponse struct
type ErrorResponse struct {
    Field string `json:"field"`
    Tag   string `json:"tag"`
    Value string `json:"value,omitempty"`
}

// Validator
var validate = validator.New()

// ValidateStruct function
func ValidateStruct[T any](payload T) []*ErrorResponse {
    var errors []*ErrorResponse
    err := validate.Struct(payload)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            var element ErrorResponse
            element.Field = err.StructNamespace()
            element.Tag = err.Tag()
            element.Value = err.Param()
            errors = append(errors, &element)
        }
    }
    return errors
}

// ForgotPasswordInput struct
type ForgotPasswordInput struct {
    Email string `json:"email" binding:"required"`
}

// ResetPasswordInput struct
type ResetPasswordInput struct {
    Password        string `json:"password" binding:"required"`
    PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

// // Exam Model
// type Exam struct {
//     gorm.Model
//     Name           string           `gorm:"not null"`
//     UserID         string
//     User           User
//     StudyMaterials []StudyMaterial  `gorm:"foreignKey:ExamID"`
// }

// // StudyMaterial Model
// type StudyMaterial struct {
//     gorm.Model
//     Name           string     `gorm:"not null"`
//     ExamID         uint
//     Exam           Exam
//     CategoryID     uint
//     Category       Category
//     SubCategoryID  uint
//     SubCategory    SubCategory
//     TopicID        uint
//     Topic          Topic
// }

// // MockExam Model
// type MockExam struct {
//     gorm.Model
//     Name           string     `gorm:"not null"`
//     UserID         string
//     User           User
//     CategoryID     uint
//     Category       Category
//     SubCategoryID  uint
//     SubCategory    SubCategory
//     TopicID        uint
//     Topic          Topic
// }

// // Category Model
// type Category struct {
//     gorm.Model
//     Name           string        `gorm:"unique;not null"`
//     SubCategories  []SubCategory `gorm:"foreignKey:CategoryID"`
// }

// // SubCategory Model
// type SubCategory struct {
//     gorm.Model
//     Name           string     `gorm:"not null"`
//     CategoryID     uint
//     Category       Category
//     Topics         []Topic    `gorm:"foreignKey:SubCategoryID"`
// }

// // Topic Model
// type Topic struct {
//     gorm.Model
//     Name           string     `gorm:"not null"`
//     SubCategoryID  uint
//     SubCategory    SubCategory
// }
