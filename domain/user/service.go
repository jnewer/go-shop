package user

import hash "go-shop/utils"

type Service struct {
	r Repository
}

func NewUserService(r Repository) *Service {
	r.Migration()
	r.InsertSampleData()
	return &Service{
		r: r,
	}
}

func (s *Service) Create(user *User) error {
	if user.Password != user.Password2 {
		return ErrMismatchPasswords
	}

	_, err := s.r.GetByUsername(user.Username)
	if err != nil {
		return ErrUserExistWithName
	}

	if ValidateUsername(user.Username) {
		return ErrInvalidUsername
	}

	if ValidatePassword(user.Password) {
		return ErrInvalidPassword
	}

	err = s.r.Create(user)
	return err
}

func (s *Service) GetUser(username string, password string) (User, error) {
	user, err := s.r.GetByUsername(username)

	if err != nil {
		return User{}, ErrUserNotFound
	}

	match := hash.CheckPasswordHash(password+user.Salt, user.Password)

	if !match {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

func (s *Service) UpdateUser(user *User) error {
	return s.r.Update(user)
}
