package user

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&User{})

	if err != nil {
		log.Print(err)
	}
}

func (r *Repository) Create(u *User) error {
	result := r.db.Create(u)

	return result.Error
}

func (r *Repository) GetByUsername(username string) (User, error) {
	var user User
	err := r.db.Where("Username=?", username).Where("IsDeleted=?", 0).First(&user, "Username=?", username).Error

	if err != nil {
		return User{}, nil
	}

	return user, nil
}

func (r *Repository) InsertSampleData() {
	user := NewUser("admin", "admin", "admin")
	user.IsAdmin = true

	r.db.Where(User{Username: user.Username}).Attrs(
		User{
			Username: user.Username,
			Password: user.Password,
		}).FirstOrCreate(&user)

	user = NewUser("user", "user", "user")
	r.db.Where(User{Username: user.Username}).Attrs(
		User{
			Username: user.Username,
			Password: user.Password,
		}).FirstOrCreate(&user)
}

func (r *Repository) Update(u *User) error {
	return r.db.Save(&u).Error
}
