//- apps/handlers/auth_handler.go

package handlers

import (
	"go-fiber-dummyapi-svc/apps/configs"
	"go-fiber-dummyapi-svc/apps/entities"
	"go-fiber-dummyapi-svc/pkgs/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	Cfg *configs.Config
	DB  *gorm.DB
}

func NewAuthHandler(cfg *configs.Config, db *gorm.DB) *AuthHandler {
	return &AuthHandler{Cfg: cfg, DB: db}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Hash password
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := entities.User{
		Email:     data["email"],
		Password:  string(password),
		Firstname: data["firstname"],
		Lastname:  data["lastname"],
	}

	if err := h.DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Email sudah terdaftar"})
	}

	return c.JSON(user)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user entities.User
	h.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == "0" {
		return c.Status(404).JSON(fiber.Map{"message": "User tidak ditemukan"})
	}

	// Bandingkan password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(data["password"])); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Password salah"})
	}

	// Generate JWT Token
	tokenData := utils.TokenData{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Avatar:    user.Avatar,
	}
	token, err := utils.GenerateToken(h.Cfg, tokenData)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"accessToken":  token,
		"refreshToken": token,
		"user": entities.RespAuthUser{
			BaseID: entities.BaseID{
				ID: user.ID,
			},
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
			Avatar:    user.Avatar,
		},
	})
}
