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
		user.Id,
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
		c.IndentedJSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Missing required query param 'uid'"})
		return
	}

	uid, err := strconv.Atoi(c.Query("uid"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Invalid user id"})
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

	type Uri struct {
		ProfileId int `uri:"profileId" binding:"required"`
	}

	var uri Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	var reqBody schemas.UpdateUserProfileSchema

	if err := c.BindJSON(&reqBody); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest, schemas.ErrorResponse{Message: err.Error()},
		)
		return
	}

	profile, err := models.UpdateUserProfile(
		user,
		uri.ProfileId,
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
		Id:       profile.Id,
		Username: profile.User.Username,
		Email:    profile.User.Email,
		Uid:      profile.Uid,
		Name:     profile.Name,
		Bio:      profile.Bio,
		PicUrl:   profile.PicUrl,
	}
}
