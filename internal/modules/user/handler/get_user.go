package handler

import "fmt"

func (h *userHandler) GetUserHandler() {
	res := h.userService.GetUser()

	fmt.Println(res)
}
