package controllers

import (
	"net/http"
	"strconv"

	"github.com/akhill4054/room-backend/models"
	"github.com/akhill4054/room-backend/schemas"
	"github.com/gin-gonic/gin"
)

func CreateProfile(c *gin.Context) {
	user := c.Keys["user"].(*models.User)

	var reqBody schemas.CreateUserProfileSchema

	if err := c.BindJSON(&reqBody); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()},
		)
		return
	}

	profile, err := models.CreateUserProfile(
		user.ID,
		reqBody.Name,
		reqBody.Bio,
		reqBody.PicUrl,
	)

	if err == models.ErrUserProfileAlreadyExists {
		c.IndentedJSON(
			http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()},
		)
		return
	}

	c.IndentedJSON(http.StatusCreated, asUserProfileSchema(profile))
}

func GetUserProfile(c *gin.Context) {
	queryUid := c.Query("uid")
	if queryUid == "" {
		c.IndentedJSON(
			http.StatusBadRequest,
			schemas.ErrorResponse{Message: "Missing required query param 'uid'"},
		)
		return
	}

	uid, err := strconv.Atoi(c.Query("uid"))
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest, schemas.ErrorResponse{Message: "Invalid user id"},
		)
		return
	}

	profile, err := models.GetUserProfile(uid)

	if err == models.ErrUserProfileNotFound {
		c.IndentedJSON(
			http.StatusNotFound, schemas.ErrorResponse{Message: err.Error()},
		)
		return
	}

	c.IndentedJSON(http.StatusOK, asUserProfileSchema(profile))
}

func UpdateUserProfile(c *gin.Context) {
	user := c.Keys["user"].(*models.User)

	var reqBody schemas.UpdateUserProfileSchema

	if err := c.BindJSON(&reqBody); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()},
		)
		return
	}

	profile, err := models.UpdateUserProfile(
		user.ID,
		reqBody.Name,
		reqBody.Bio,
		reqBody.PicUrl,
	)

	if err == models.ErrUserProfileDoesNotExist {
		c.IndentedJSON(
			http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()},
		)
		return
	}

	c.IndentedJSON(http.StatusOK, asUserProfileSchema(profile))
}

func asUserProfileSchema(profile *models.UserProfile) *schemas.UserProfileSchema {
	return &schemas.UserProfileSchema{
		ID:       profile.ID,
		Username: profile.User.Username,
		Email:    profile.User.Email,
		UID:      profile.UID,
		Name:     profile.Name,
		Bio:      profile.Bio,
		PicUrl:   profile.PicUrl,
	}
}
