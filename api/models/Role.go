package models

import (
	"errors"
	"html"
	"strings"
	"time"
	"github.com/jinzhu/gorm"
)

type Role struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name	  string    `gorm:"size:255;not null;unique" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (r *Role) Prepare() {
	r.ID 		= 0
	r.Name 		= html.EscapeString(strings.TrimSpace(r.Name))
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

func (r *Role) Validate() error {
	if r.Name == "" {
		return errors.New("Required Name")
	}
	return nil
}
func (r *Role) SaveRole(db *gorm.DB) (*Role, error) {

	var err error
	err = db.Debug().Create(&r).Error
	if err != nil {
		return &Role{}, err
	}
	return r, nil
}

func (r *Role) FindAllRoles(db *gorm.DB) (*[]Role, error) {
	var err error
	roles := []Role{}
	err = db.Debug().Model(&Role{}).Limit(100).Find(&roles).Error
	if err != nil {
		return &[]Role{}, err
	}
	return &roles, err
}

func (r *Role) FindRoleByID(db *gorm.DB, rid uint32) (*Role, error) {
	var err error
	err = db.Debug().Model(Role{}).Where("id = ?", rid).Take(&r).Error
	if err != nil {
		return &Role{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Role{}, errors.New("Role Not Found")
	}
	return r, err
}

func (r *Role) UpdateARole(db *gorm.DB, rid uint32) (*Role, error) {

	db = db.Debug().Model(&Role{}).Where("id = ?", rid).Take(&Role{}).UpdateColumns(
		map[string]interface{}{
			"name":  r.Name,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Role{}, db.Error
	}
	return r, nil
}

func (r *Role) DeleteARole(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Role{}).Where("id = ?", uid).Take(&Role{}).Delete(&Role{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}