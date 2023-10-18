package schemas

type UserScheama struct {
	Id       int    `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	IsAdmin  bool   `json:"is_amdin" binding:"required"`
}

type CreateUserProfileSchema struct {
	Name   string `json:"name" binding:"required"`
	Bio    string `json:"bio"`
	PicUrl string `json:"pic_url"`
}

type UpdateUserProfileSchema struct {
	Name   string `json:"name" binding:"required"`
	Bio    string `json:"bio"`
	PicUrl string `json:"pic_url"`
}

type UserProfileSchema struct {
	Id       int    `json:"id" binding:"required"`
	Uid      int    `json:"uid" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Bio      string `json:"bio,omitempty"`
	PicUrl   string `json:"pic_url,omitempty"`
}

type DeleteUserResponseSchema struct {
	Message string `json:"message" binding:"required"`
}
