package models

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Book struct {
	ID     uint32 `gorm:"primary_key;auto_increment" json:"id"`
	ISBN   string `gorm:"type:varchar(13);unique" json:"isbn"`
	Titulo string `gorm:"type:varchar(256);unique" json:"titulo"`
	Slug   string `gorm:"type:varchar(256)" json:"Slug"`
	// Autores []*Autor `gorm:"many2many:livro_autor;" json:"autores"`
	Autor     string    `gorm:"type:varchar(256)" json:"autor"`
	Idioma    string    `gorm:"default:'Português'" json:"idioma"`
	Formato   string    `gorm:"default:'Capa comum'" json:"formato"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (b *Book) Prepare() {
	b.Slug = strings.Replace(b.Titulo, " ", "-", -1)
	b.Slug = strings.ToLower(b.Slug)
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
}

func (b *Book) Validate(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if len(b.ISBN) != 13 {
			return errors.New("ISBN inválido")
		}
		if b.Titulo == "" {
			return errors.New("Titulo não pode estar em branco")
		}
		if b.Autor == "" {
			return errors.New("Autor não pode estar em branco")
		}
		return nil
	case "update":
		if len(b.ISBN) != 13 {
			return errors.New("ISBN inválido")
		}
		if b.Titulo == "" {
			return errors.New("Titulo não pode estar em branco")
		}
		if b.Autor == "" {
			return errors.New("Autor não pode estar em branco")
		}
		return nil
	default:
		if len(b.ISBN) != 13 {
			return errors.New("ISBN inválido")
		}
		if b.Titulo == "" {
			return errors.New("Titulo não pode estar em branco")
		}
		if b.Autor == "" {
			return errors.New("Autor não pode estar em branco")
		}
		return nil
	}
}

func (b *Book) SaveBook(db *gorm.DB) (*Book, error) {
	var err error
	err = db.Debug().Create(&b).Error
	if err != nil {
		return &Book{}, err
	}
	return b, nil
}

func (b *Book) FindAllBooks(db *gorm.DB) (*[]Book, error) {
	var err error
	book := []Book{}
	err = db.Debug().Model(&Book{}).Limit(100).Find(&book).Error
	if err != nil {
		return &[]Book{}, err
	}
	return &book, err
}

func (b *Book) FindBookByISBN(db *gorm.DB, isbn string) (*Book, error) {
	var err error
	err = db.Debug().Model(Book{}).Where("isbn = ?", isbn).Take(&b).Error
	if err != nil {
		return &Book{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Book{}, errors.New("Book Not Found")
	}
	return b, err
}

func (b *Book) UpdateABook(db *gorm.DB, isbn string) (*Book, error) {

	db = db.Debug().Model(&Book{}).Where("isbn = ?", isbn).Take(&Book{}).UpdateColumns(
		map[string]interface{}{
			"isbn":       b.ISBN,
			"titulo":     b.Titulo,
			"slug":       b.Slug,
			"autor":      b.Autor,
			"idioma":     b.Idioma,
			"formato":    b.Formato,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Book{}, db.Error
	}
	err := db.Debug().Model(&Book{}).Where("isbn = ?", isbn).Take(&b).Error
	if err != nil {
		return &Book{}, err
	}
	return b, nil
}

func (b *Book) DeleteABook(db *gorm.DB, bid uint32) (int64, error) {
	db = db.Debug().Model(&Book{}).Where("id = ?", bid).Take(&Book{}).Delete(&Book{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
