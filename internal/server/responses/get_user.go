package responses

import "github.com/alishchenko/discountaria/internal/data"

type GetUserResponse struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToGetUserResponse(user data.User) GetUserResponse {
	return GetUserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToListUserResponse(users []data.User) []GetUserResponse {
	var resp []GetUserResponse
	for _, user := range users {
		resp = append(resp, ToGetUserResponse(user))
	}
	return resp
}
