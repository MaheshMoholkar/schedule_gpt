package handlers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MaheshMoholkar/schedule_gpt/internal/database"
	"github.com/MaheshMoholkar/schedule_gpt/internal/utils"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	System string `json:"system"`
	Stream bool   `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	System   string    `json:"system"`
	Stream   bool      `json:"stream"`
}

type GenerateResponse struct {
	Model              string `json:"model"`
	CreatedAt          string `json:"created_at"`
	Response           string `json:"response"`
	Done               bool   `json:"done"`
	Context            []int  `json:"context"`
	TotalDuration      int64  `json:"total_duration"`
	LoadDuration       int64  `json:"load_duration"`
	PromptEvalCount    int    `json:"prompt_eval_count"`
	PromptEvalDuration int64  `json:"prompt_eval_duration"`
	EvalCount          int    `json:"eval_count"`
	EvalDuration       int64  `json:"eval_duration"`
}

type ChatResponse struct {
	Model              string  `json:"model"`
	CreatedAt          string  `json:"created_at"`
	Message            Message `json:"message"`
	Done               bool    `json:"done"`
	Context            []int   `json:"context"`
	TotalDuration      int64   `json:"total_duration"`
	LoadDuration       int64   `json:"load_duration"`
	PromptEvalCount    int     `json:"prompt_eval_count"`
	PromptEvalDuration int64   `json:"prompt_eval_duration"`
	EvalCount          int     `json:"eval_count"`
	EvalDuration       int64   `json:"eval_duration"`
}

type OllamaHandler struct {
	store *database.Store
}

func NewOllamaHandler(store *database.Store) *OllamaHandler {
	return &OllamaHandler{
		store: store,
	}
}

func (h *OllamaHandler) HandleGenerateRequest(ctx *fiber.Ctx) error {
	var params GenerateRequest
	err := ctx.BodyParser(&params)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request payload")
	}

	params.Model = "llama3"
	params.Stream = false
	endpoint := "/api/generate"

	response, err := h.interactWithOllama(params, endpoint)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(response)
}

func (h *OllamaHandler) HandleChatRequest(ctx *fiber.Ctx) error {
	var params ChatRequest
	err := ctx.BodyParser(&params)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request payload")
	}

	params.Model = "llama3"
	params.Stream = false
	endpoint := "/api/chat"
	params.System = utils.CHAT_SYSTEM_PROMPT

	response, err := h.chatWithOllama(params, endpoint)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(response)
}

func (h *OllamaHandler) HandleGenerateSchedule(ctx *fiber.Ctx) error {
	var params GenerateRequest
	err := ctx.BodyParser(&params)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request payload")
	}

	params.Model = "llama3"
	params.Stream = false
	params.System = utils.SCHEDULE_SYSTEM_PROMPT
	endpoint := "/api/generate"

	response, err := h.interactWithOllama(params, endpoint)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(response)
}

func (h *OllamaHandler) interactWithOllama(params interface{}, endpoint string) (*GenerateResponse, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		Post(os.Getenv("OLLAMA_HOST") + endpoint)

	if err != nil {
		return nil, fmt.Errorf("failed to interact with Ollama: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("ollama responded with an error: %s", resp.Status())
	}

	var result GenerateResponse
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Ollama response: %w", err)
	}

	return &result, nil
}

func (h *OllamaHandler) chatWithOllama(params interface{}, endpoint string) (*ChatResponse, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(params).
		Post(os.Getenv("OLLAMA_HOST") + endpoint)

	if err != nil {
		return nil, fmt.Errorf("failed to interact with Ollama: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("ollama responded with an error: %s", resp.Status())
	}

	var result ChatResponse
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Ollama response: %w", err)
	}

	return &result, nil
}
