//- apps/entities/user_entity.go

package entities

type Gender string

// 2. Define constants
const (
	GenderMale   Gender = "M"
	GenderFemale Gender = "F"
)

type User struct {
	BaseID

	Firstname string `gorm:"column:firstname;type:varchar(255);not null" json:"firstname"`
	Lastname  string `gorm:"column:lastname;type:varchar(255)" json:"lastname"`
	Gender    Gender `gorm:"column:gender;type:varchar(1);default:'M'" json:"gender"`
	Avatar    string `gorm:"column:avatar;type:varchar(255);default:null" json:"avatar"`

	Email        string `gorm:"column:email;unique;type:varchar(255);not null" json:"email"`
	Password     string `gorm:"column:password;type:varchar(255);not null" json:"-"`
	PasswordHash string `gorm:"column:password_hash;type:varchar(255)" json:"-"`
	Phone        string `gorm:"column:phone;type:varchar(30)" json:"phone"`

	BaseTimestamp
}

type RespAuthUser struct {
	BaseID

	Firstname string `gorm:"column:firstname" json:"firstname"`
	Lastname  string `gorm:"column:lastname" json:"lastname"`
	Email     string `gorm:"column:email" json:"email"`
	Avatar    string `gorm:"column:avatar" json:"avatar"`
}

type RespListUser struct {
	BaseID

	Firstname string `gorm:"column:firstname" json:"firstname"`
	Lastname  string `gorm:"column:lastname" json:"lastname"`
	Email     string `gorm:"column:email" json:"email"`
	Phone     string `gorm:"column:phone" json:"phone"`
	Avatar    string `gorm:"column:avatar" json:"avatar"`
	Gender    Gender `gorm:"column:gender" json:"gender"`
}

type RespDetailUser struct {
	RespListUser

	Password string `gorm:"column:password" json:"password"`
}

func (User) TableName() string {
	return GetTableName(TABLE_USER)
}
