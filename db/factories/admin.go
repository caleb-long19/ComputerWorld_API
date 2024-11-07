package factories

import (
	m "ComputerWorld_API/db/models"
	"ComputerWorld_API/pkg/rand"
	"github.com/jinzhu/gorm"
	"log"
)

func NewAdmin(db *gorm.DB, admin *m.Admin) {
	fillAdminDefaults(admin)

	err := db.Create(admin).Error
	if err != nil {
		log.Printf("Error creating admin in factory: %v", err)
	}
}

func NewAdmins(db *gorm.DB, defaultAdmin *m.Admin, total int) []*m.Admin {
	if total < 1 || total > 100 {
		log.Fatal("Count can only be in the range 1 - 100")
	}
	admins := make([]*m.Admin, total)
	for i := 0; i < total; i++ {
		admins[i] = &m.Admin{
			UID:      "",
			Email:    defaultAdmin.Email,
			Name:     defaultAdmin.Name,
			Password: defaultAdmin.Password,
		}
		NewAdmin(db, admins[i])
	}

	return admins
}

func fillAdminDefaults(admin *m.Admin) {
	if admin.Email == "" {
		admin.Email = rand.Email()
	}
	if admin.Name == "" {
		admin.Name = rand.String()
	}
	if admin.Password == "" {
		admin.Password = rand.String()
	}
	if admin.UID == "" {
		admin.UID = rand.String()
	}
}
