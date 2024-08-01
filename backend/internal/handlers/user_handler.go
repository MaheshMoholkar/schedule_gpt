package handlers

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/MaheshMoholkar/schedule_gpt/internal/database"
	"github.com/MaheshMoholkar/schedule_gpt/internal/types"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserHandler struct {
	store *database.Store
}

func NewUserHandler(store *database.Store) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (h *UserHandler) HandleCreateUser(ctx *fiber.Ctx) error {
	var params *types.UserParams
	err := ctx.BodyParser(&params)
	if err != nil {
		return err
	}
	email := strings.ToLower(params.Email)
	password := params.Password

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	passwordHash := string(hashedBytes)

	query := `INSERT INTO users (email, password_hash) VALUES (@email, @password_hash )`

	_, err = h.store.DB.Exec(query, sql.Named("email", email), sql.Named("password_hash", passwordHash))
	if err != nil {
		return err
	}

	return ctx.SendStatus(200)
}

func (h *UserHandler) HandleGetUser(ctx *fiber.Ctx) error {
	userID := ctx.Query("id", "0")

	query := fmt.Sprintf(`SELECT id, email, password_hash FROM users WHERE id = %s `, userID)

	var user types.UserParams

	err := h.store.DB.QueryRow(query).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return err
	}

	return ctx.JSON(user)
}
