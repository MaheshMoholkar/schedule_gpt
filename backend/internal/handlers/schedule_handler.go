package handlers

import (
	"database/sql"

	"github.com/MaheshMoholkar/schedule_gpt/internal/database"
	"github.com/MaheshMoholkar/schedule_gpt/internal/types"
	"github.com/gofiber/fiber/v2"
)

type ScheduleHandler struct {
	store *database.Store
}

func NewSchduleHandler(store *database.Store) *ScheduleHandler {
	return &ScheduleHandler{
		store: store,
	}
}

func (h *ScheduleHandler) HandleStoreSubjects(ctx *fiber.Ctx) error {
	var data types.TeacherSubject
	err := ctx.BodyParser(&data)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request payload")
	}

	for _, subject := range data.Subjects {
		query := `INSERT INTO subjects (teacher_name, subject_name, Occurrence) VALUES (@teacher_name, @subject_name, @occurrence)`
		_, err := h.store.DB.Exec(query, sql.Named("teacher_name", subject.TeacherName), sql.Named("subject_name", subject.SubjectName), sql.Named("occurrence", subject.Occurrence))
		if err != nil {
			return err
		}
	}

	return ctx.SendStatus(200)
}

func (h *ScheduleHandler) HandleGetSubjects(ctx *fiber.Ctx) error {
	rows, err := h.store.DB.Query("SELECT teacher_name, subject_name, occurrence FROM subjects")
	if err != nil {
		return err
	}
	defer rows.Close()

	var subjects []types.Subject
	for rows.Next() {
		var subject types.Subject
		if err := rows.Scan(&subject.TeacherName, &subject.SubjectName, &subject.Occurrence); err != nil {
			return err
		}
		subjects = append(subjects, subject)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return ctx.JSON(subjects)
}
