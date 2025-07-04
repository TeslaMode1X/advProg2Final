package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/TeslaMode1X/advProg2Final/user/internal/interfaces"
	"github.com/TeslaMode1X/advProg2Final/user/internal/model"
	"github.com/gofrs/uuid"
	"github.com/mailgun/mailgun-go/v4"
	"time"
)

type UserService struct {
	userRepo   interfaces.UserRepo
	redisCache interfaces.RedisCache
}

func NewUserService(userRepo interfaces.UserRepo, redisCache interfaces.RedisCache) *UserService {
	return &UserService{
		userRepo:   userRepo,
		redisCache: redisCache,
	}
}

func (us *UserService) UserRegisterService(user model.User) (uuid.UUID, error) {
	const op = "user.service.UserRegisterService"

	if err := user.Validate(); err != nil {
		return uuid.Nil, err
	}

	if user.ID == uuid.Nil {
		newID, err := uuid.NewV4()
		if err != nil {
			return uuid.Nil, errors.New(op + ": failed to generate UUID: " + err.Error())
		}
		user.ID = newID
	}

	id, err := us.userRepo.UserRegisterRepo(user)
	if err != nil {
		return uuid.Nil, errors.New(op + ": " + err.Error())
	}

	err = us.redisCache.Set(context.Background(), user)
	if err != nil {
		fmt.Printf("[REDIS DEBUG] Failed to cache user in Redis: %v\n", err)
	} else {
		fmt.Printf("[REDIS DEBUG] User successfully cached in Redis with key pattern: user:%s\n", id)
	}

	apiKey := "9291a387d961180fbedfdd6c186e8261-f3238714-cc4fbb37"
	domain := "sandbox8cdacc3f5b044882b390cabbe20ef6ee.mailgun.org"
	notifyEmail := "anuar.anuar222444@gmail.com"

	message := fmt.Sprintf("%s зарегистрировался в вашем приложении", user.Email)

	idS, err := SendRegistrationNotification(domain, apiKey, notifyEmail, message)
	fmt.Println(idS)
	fmt.Println(err)

	return id, nil
}

func (us *UserService) UserLoginService(login, password string) (uuid.UUID, error) {
	const op = "user.service.UserLoginService"

	id, err := us.userRepo.UserLoginRepo(login, password)
	if err != nil {
		return uuid.Nil, errors.New(op + ": " + err.Error())
	}

	return id, nil
}

func (us *UserService) UserGetByIdService(id string) (*model.User, error) {
	const op = "service.user.UserGetByIdService"

	uid, err := uuid.FromString(id)
	if err != nil {
		return nil, errors.New(op + ": invalid UUID format")
	}

	fmt.Printf("[REDIS DEBUG] Checking Redis cache for user ID: %s\n", id)
	cachedUser, err := us.redisCache.Get(context.Background(), uid)
	if err == nil && cachedUser.ID != uuid.Nil {
		fmt.Printf("[REDIS DEBUG] CACHE HIT! User found in Redis cache: %s (%s)\n",
			cachedUser.Username, cachedUser.Email)
		return &cachedUser, nil
	}

	if err != nil {
		fmt.Printf("[REDIS DEBUG] Redis cache error: %v\n", err)
	} else {
		fmt.Printf("[REDIS DEBUG] CACHE MISS. User not found in Redis cache.\n")
	}

	user, err := us.userRepo.UserGetByIdRepo(id)
	if err != nil {
		return nil, errors.New(op + ": " + err.Error())
	}

	return user, nil
}

func (us *UserService) UserDeleteByIdService(id string) error {
	const op = "service.user.UserDeleteByIdService"

	if err := us.userRepo.UserDeleteByIdRepo(id); err != nil {
		return errors.New(op + ": " + err.Error())
	}

	return nil
}

func (us *UserService) UserExistsService(id string) (bool, error) {
	const op = "service.user.UserExistsService"

	exists, err := us.userRepo.UserExistsRepo(id)
	if err != nil {
		return false, errors.New(op + ": " + err.Error())
	}

	return exists, nil
}

func (us *UserService) UserUpdatePasswordService(id, oldPassword, newPassword string) error {
	const op = "service.user.UserUpdatePasswordService"

	if err := us.userRepo.UserUpdatePasswordRepo(id, oldPassword, newPassword); err != nil {
		return errors.New(op + ": " + err.Error())
	}

	return nil
}

func (us *UserService) RefreshCache() error {
	users, err := us.userRepo.UserGetAllRepo()
	if err != nil {
		return err
	}

	err = us.redisCache.SetMany(context.Background(), users)
	if err != nil {
		return err
	}

	return nil
}

func SendRegistrationNotification(domain, apiKey, recipient, messageBody string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)

	sender := fmt.Sprintf("Mailgun Sandbox <postmaster@%s>", domain)
	subject := "Новый пользователь зарегистрировался"

	message := mg.NewMessage(sender, subject, messageBody, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, id, err := mg.Send(ctx, message)
	return id, err
}
