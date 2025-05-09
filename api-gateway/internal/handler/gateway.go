package handler

import (
	"context"
	"fmt"
	"github.com/TeslaMode1X/advProg2Final/proto/gen/recipe"
	"github.com/TeslaMode1X/advProg2Final/proto/gen/review"
	"github.com/TeslaMode1X/advProg2Final/proto/gen/statistics"
	"github.com/TeslaMode1X/advProg2Final/proto/gen/user"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type GatewayHandler struct {
	userClient       user.UserServiceClient
	recipeClient     recipe.RecipeServiceClient
	reviewClient     review.ReviewServiceClient
	statisticsClient statistics.StatisticsServiceClient
}

func NewGatewayHandler(userConn, recipeConn *grpc.ClientConn, reviewConn *grpc.ClientConn, statisticsConn *grpc.ClientConn) *GatewayHandler {
	return &GatewayHandler{
		userClient:       user.NewUserServiceClient(userConn),
		recipeClient:     recipe.NewRecipeServiceClient(recipeConn),
		reviewClient:     review.NewReviewServiceClient(reviewConn),
		statisticsClient: statistics.NewStatisticsServiceClient(statisticsConn),
	}
}

func (g *GatewayHandler) UserLogin(c *gin.Context) {
	_, err := c.Cookie("auth_token")
	if err == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "already logged in"})
		return
	}

	var req user.RequestUserLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	resp, err := g.userClient.UserLogin(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie(
		"auth_token",
		resp.Token,
		3600*1,
		"/",
		"",
		false,
		false,
	)

	c.JSON(http.StatusOK, gin.H{
		"id":    resp.Id,
		"token": resp.Token,
	})
}

func (g *GatewayHandler) UserRegistration(c *gin.Context) {
	var req user.RequestUserRegistration

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Email == "" || req.Password == "" || req.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username, email and password are required"})
		return
	}

	resp, err := g.userClient.UserRegistration(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": resp.Id})
}

func (g *GatewayHandler) UserGetById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	req := &user.RequestUserGetById{
		Id: id,
	}

	user, err := g.userClient.UserGetById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (g *GatewayHandler) UserDeleteById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	req := &user.RequestUserGetById{
		Id: id,
	}

	_, err := g.userClient.UserDeleteById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": "was deleted"})
}

func (g *GatewayHandler) UserExists(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	req := &user.RequestUserGetById{
		Id: id,
	}

	exists, err := g.userClient.UserExists(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": exists})
}

func (g *GatewayHandler) UserChangePassword(c *gin.Context) {
	var req user.RequestUserChangePassword
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	if req.OldPassword == "" || req.NewPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Old and new passwords are required"})
		return
	}

	userIDStr := userID.(uuid.UUID).String()

	ctx := context.Background()
	_, err := g.userClient.UserChangePassword(ctx, &user.RequestUserChangePassword{
		Id:          userIDStr,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

func (g *GatewayHandler) RecipeList(c *gin.Context) {
	r := &recipe.Empty{}
	recipeObjects, err := g.recipeClient.RecipeList(context.Background(), r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"recipes": recipeObjects})
}

func (g *GatewayHandler) RecipeCreate(c *gin.Context) {
	const op = "handler.gateway.RecipeCreate"

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not read multipart form"})
		return
	}

	var title, description string
	if len(form.Value["title"]) > 0 {
		title = form.Value["title"][0]
	}
	if len(form.Value["description"]) > 0 {
		description = form.Value["description"][0]
	}

	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	userID, ok := userIDValue.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID format"})
		return
	}

	recipeID, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	files := form.File["photos"]
	var photoPaths []string
	for _, file := range files {
		filename := fmt.Sprintf("%s-%s", recipeID.String(), file.Filename)

		dir := "/app/photo"
		filepath := fmt.Sprintf("%s/%s", dir, filename)

		cwd, _ := os.Getwd()
		log.Printf("Current working directory: %s", cwd)

		if err = os.MkdirAll(dir, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "directory creation error"})
			return
		}

		if err = c.SaveUploadedFile(file, filepath); err != nil {
			log.Printf("Error saving file: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "photo save error"})
			return
		}

		photoPaths = append(photoPaths, fmt.Sprintf("photo/%s", filename))
	}

	req := &recipe.RecipeCreateRequest{
		Title:       title,
		Description: description,
		Photos:      photoPaths,
		AuthorId:    userID.String(),
	}

	id, err := g.recipeClient.RecipeCreate(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (g *GatewayHandler) RecipeByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	req := &recipe.RecipeByIDRequest{
		Id: id,
	}

	recipeObject, err := g.recipeClient.RecipeByID(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"recipe": recipeObject})
}

func (g *GatewayHandler) RecipeUpdate(c *gin.Context) {
	const op = "handler.gateway.RecipeUpdate"

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not read multipart form"})
		return
	}

	if len(form.Value["id"]) == 0 || form.Value["id"][0] == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "recipe ID is required"})
		return
	}
	recipeID := form.Value["id"][0]

	idUUID := uuid.FromStringOrNil(recipeID)
	if idUUID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid recipe ID"})
		return
	}

	var title, description string
	if len(form.Value["title"]) > 0 {
		title = form.Value["title"][0]
	}
	if len(form.Value["description"]) > 0 {
		description = form.Value["description"][0]
	}

	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}

	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	userID, ok := userIDValue.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID format"})
		return
	}

	var photoPaths []string
	files := form.File["photos"]
	if len(files) > 0 {
		for _, file := range files {
			filename := fmt.Sprintf("%s-%s", idUUID.String(), file.Filename)

			dir := "/app/photo"
			filepath := fmt.Sprintf("%s/%s", dir, filename)

			if err = os.MkdirAll(dir, 0755); err != nil {
				log.Printf("Error creating directory: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "directory creation error"})
				return
			}

			if err = c.SaveUploadedFile(file, filepath); err != nil {
				log.Printf("Error saving file: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "photo save error"})
				return
			}

			photoPaths = append(photoPaths, fmt.Sprintf("photo/%s", filename))
		}
	}

	req := &recipe.RecipeUpdateRequest{
		Id:          recipeID,
		Title:       title,
		Description: description,
		AuthorId:    userID.String(),
	}

	if len(photoPaths) > 0 {
		req.Photos = photoPaths
	} else {
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = g.recipeClient.RecipeUpdate(ctx, req)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
			return
		} else if status.Code(err) == codes.PermissionDenied {
			c.JSON(http.StatusForbidden, gin.H{"error": "you don't have permission to update this recipe"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": recipeID})
}

func (g *GatewayHandler) RecipeDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	userID, ok := userIDValue.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID format"})
		return
	}

	req := &recipe.RecipeDeleteRequest{Id: id, AuthorId: userID.String()}

	_, err := g.recipeClient.RecipeDelete(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "recipe deleted"})
}

func (g *GatewayHandler) ReviewCreate(c *gin.Context) {
	var req review.ReviewCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	userID, ok := userIDValue.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID format"})
		return
	}

	if req.RecipeId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "recipe_id is required"})
		return
	}

	reqForExist := &recipe.RecipeExistsRequest{
		RecipeId: req.RecipeId,
	}

	exist, err := g.recipeClient.RecipeExists(context.Background(), reqForExist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !exist.Check {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
		return
	}

	if req.Rating < 1 || req.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rating must be between 1 and 5"})
		return
	}

	if req.Comment == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "comment is required"})
		return
	}

	req.UserId = userID.String()

	id, err := g.reviewClient.ReviewCreate(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (g *GatewayHandler) ReviewList(c *gin.Context) {
	var req review.Empty

	objectList, err := g.reviewClient.ReviewGetList(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, objectList)
}

func (g *GatewayHandler) ReviewById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	object, err := g.reviewClient.ReviewGetById(context.Background(), &review.ReviewGetByIdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, object)
}

func (g *GatewayHandler) ReviewUpdate(c *gin.Context) {
	var req review.ReviewUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	userID, ok := userIDValue.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID format"})
		return
	}

	req.UserId = userID.String()

	id, err := g.reviewClient.ReviewUpdate(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (g *GatewayHandler) ReviewDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	req := &review.ReviewDeleteRequest{Id: id}

	_, err := g.reviewClient.ReviewDelete(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "review deleted"})
}

func (g *GatewayHandler) GetUserRegisteredStatistics(c *gin.Context) {
	r := &statistics.Empty{}
	userStats, err := g.statisticsClient.StatisticsUser(context.Background(), r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userStats)
}
