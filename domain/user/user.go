package user

import "errors"

var (
	ErrUserExistWithName = errors.New("用户名已存在")
	ErrUserNotFound      = errors.New("用户未找到")
	
	ErrMismatchPasswords = errors.New("密码不匹配")
	ErrInvalidUsername   = errors.New("无效用户名")
	ErrInvalidPassword   = errors.New("无效密码")
)
