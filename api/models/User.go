package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	User      string    `gorm:"size:8;not null;unique" json:"user"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	Admin     bool      `gorm:"default:false;" json:"admin"`
	Active    bool      `gorm:"default:true;" json:"active"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Encriptar a senha
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Confirmar se a senha encriptada é correta
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Antes de salvar encripta e senha
func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.User = strings.Replace(u.User, " ", "", -1)
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// Validar informações de acesso do usuario caso email e senha estejam em branco ou formato do e-mail seja inválidos
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if len(u.User) < 4 || len(u.User) > 8 {
			return errors.New("Usuário precisa ter no minimo 4 e no máximo 8 caracteres")
		}
		if u.Password == "" {
			return errors.New("Senha obrigatória")
		}
		if u.Email == "" {
			return errors.New("Email obrigatório")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email inválido")
		}
		return nil
	case "update":
		if u.User == "" {
			return errors.New("Usuário obrigatório")
		}
		if u.Password == "" {
			return errors.New("Senha obrigatória")
		}
		if u.Email == "" {
			return errors.New("Email obrigatório")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email inválido")
		}
		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Senha obrigatória")
		}
		if u.Email == "" {
			return errors.New("Email obrigatório")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email inválido")
		}
		return nil

	default:
		if u.User == "" {
			return errors.New("Senha obrigatória")
		}
		if u.Password == "" {
			return errors.New("Senha obrigatória")
		}
		if u.Email == "" {
			return errors.New("Email obrigatório")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email inválido")
		}
		return nil
	}
}

// Salvar usuario
func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// Encontrar todos os usuários
func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

// Encontrar usuario pelo ID
func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

// Atualizar Usuario
func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":   u.Password,
			"user":       u.User,
			"email":      u.Email,
			"updated_at": time.Now(),
			"admin":      u.Admin,
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// Apagar usuario
func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
