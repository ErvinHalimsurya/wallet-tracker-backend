package service

func (u *userService) GetUser() string {
	user := u.userRepository.GetUser()
	return user
}
