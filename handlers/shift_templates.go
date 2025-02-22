package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ptmmeiningen/schichtplaner/database"
	"github.com/ptmmeiningen/schichtplaner/models"
	"github.com/ptmmeiningen/schichtplaner/pkg/responses"
	"gorm.io/gorm"
)

// @Summary Erstellt ein neues Schicht-Template
// @Description Erstellt ein neues Template für Schichtpläne
// @Tags ShiftTemplates
// @Accept json
// @Produce json
// @Param template body models.ShiftTemplate true "Template Details"
// @Success 200 {object} responses.APIResponse{data=models.ShiftTemplate}
// @Failure 400,500 {object} responses.APIResponse
// @Router /shifttemplates [post]
func HandleCreateShiftTemplate(c *fiber.Ctx) error {
	var template models.ShiftTemplate

	if err := c.BodyParser(&template); err != nil {
		return c.Status(400).JSON(responses.ErrorResponse(responses.ErrInvalidInput))
	}

	if !template.Validate() {
		return c.Status(400).JSON(responses.ErrorResponse("Ungültige Datumsangaben"))
	}

	if err := database.GetDB().Create(&template).Error; err != nil {
		return c.Status(500).JSON(responses.ErrorResponse(err.Error()))
	}

	return c.JSON(responses.SuccessResponse(responses.MsgSuccessCreate, template))
}

// @Summary Listet alle Schicht-Templates
// @Description Gibt alle verfügbaren Schicht-Templates zurück
// @Tags ShiftTemplates
// @Produce json
// @Success 200 {object} responses.APIResponse{data=[]models.ShiftTemplate}
// @Failure 500 {object} responses.APIResponse
// @Router /shifttemplates [get]
func HandleAllShiftTemplates(c *fiber.Ctx) error {
	var templates []models.ShiftTemplate

	result := database.GetDB().
		Preload("Department").
		Preload("ShiftDays", func(db *gorm.DB) *gorm.DB {
			return db.Order("week_day ASC")
		}).
		Preload("ShiftDays.ShiftType").
		Order("created_at DESC").
		Find(&templates)

	if result.Error != nil {
		return c.Status(500).JSON(responses.ErrorResponse(result.Error.Error()))
	}

	return c.JSON(responses.SuccessResponse(responses.MsgSuccessGet, templates))
}

// @Summary Gibt ein einzelnes Template zurück
// @Description Zeigt detaillierte Informationen zu einem spezifischen Schicht-Template
// @Tags ShiftTemplates
// @Param id path int true "Template ID"
// @Produce json
// @Success 200 {object} responses.APIResponse{data=models.ShiftTemplate}
// @Failure 404 {object} responses.APIResponse
// @Router /shifttemplates/{id} [get]
func HandleGetOneShiftTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	var template models.ShiftTemplate

	result := database.GetDB().
		Preload("Department").
		Preload("ShiftDays.ShiftType").
		First(&template, id)

	if result.Error != nil {
		return c.Status(404).JSON(responses.ErrorResponse(responses.ErrNotFound))
	}

	return c.JSON(responses.SuccessResponse(responses.MsgSuccessGet, template))
}

// @Summary Gibt Templates einer Abteilung zurück
// @Description Listet alle Schicht-Templates für eine bestimmte Abteilung
// @Tags ShiftTemplates
// @Param id path int true "Abteilungs-ID"
// @Produce json
// @Success 200 {object} responses.APIResponse{data=[]models.ShiftTemplate}
// @Failure 500 {object} responses.APIResponse
// @Router /shifttemplates/department/{id} [get]
func HandleGetDepartmentShiftTemplates(c *fiber.Ctx) error {
	departmentID := c.Params("id")
	var templates []models.ShiftTemplate

	result := database.GetDB().
		Where("department_id = ?", departmentID).
		Preload("ShiftDays.ShiftType").
		Find(&templates)

	if result.Error != nil {
		return c.Status(500).JSON(responses.ErrorResponse(result.Error.Error()))
	}

	return c.JSON(responses.SuccessResponse(responses.MsgSuccessGet, templates))
}

// @Summary Aktualisiert ein Template
// @Description Aktualisiert die Daten eines bestehenden Schicht-Templates
// @Tags ShiftTemplates
// @Accept json
// @Produce json
// @Param id path int true "Template ID"
// @Param template body models.ShiftTemplate true "Aktualisierte Template Details"
// @Success 200 {object} responses.APIResponse{data=models.ShiftTemplate}
// @Failure 400 {object} responses.APIResponse
// @Failure 404 {object} responses.APIResponse
// @Failure 500 {object} responses.APIResponse
// @Router /shifttemplates/{id} [put]
func HandleUpdateShiftTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	var template models.ShiftTemplate

	if err := database.GetDB().First(&template, id).Error; err != nil {
		return c.Status(404).JSON(responses.ErrorResponse(responses.ErrNotFound))
	}

	if !template.CanBeModified() {
		return c.Status(400).JSON(responses.ErrorResponse("Template kann nicht mehr bearbeitet werden"))
	}

	if err := c.BodyParser(&template); err != nil {
		return c.Status(400).JSON(responses.ErrorResponse(responses.ErrInvalidInput))
	}

	if err := database.GetDB().Save(&template).Error; err != nil {
		return c.Status(500).JSON(responses.ErrorResponse(err.Error()))
	}

	return c.JSON(responses.SuccessResponse(responses.MsgSuccessUpdate, template))
}

// @Summary Löscht ein Template
// @Description Löscht ein Schicht-Template und zugehörige Tage
// @Tags ShiftTemplates
// @Param id path int true "Template ID"
// @Produce json
// @Success 200 {object} responses.APIResponse
// @Failure 400 {object} responses.APIResponse
// @Failure 404 {object} responses.APIResponse
// @Failure 500 {object} responses.APIResponse
// @Router /shifttemplates/{id} [delete]
func HandleDeleteShiftTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	var template models.ShiftTemplate

	if err := database.GetDB().First(&template, id).Error; err != nil {
		return c.Status(404).JSON(responses.ErrorResponse(responses.ErrNotFound))
	}

	if !template.CanBeModified() {
		return c.Status(400).JSON(responses.ErrorResponse("Template kann nicht gelöscht werden"))
	}

	tx := database.GetDB().Begin()

	if err := tx.Where("shift_template_id = ?", id).Delete(&models.ShiftTemplateDay{}).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(responses.ErrorResponse(err.Error()))
	}

	if err := tx.Delete(&template).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(responses.ErrorResponse(err.Error()))
	}

	tx.Commit()

	return c.JSON(responses.SuccessResponse(responses.MsgSuccessDelete, nil))
}

// @Summary Aktualisiert den Template-Status
// @Description Ändert den Status eines Schicht-Templates
// @Tags ShiftTemplates
// @Accept json
// @Produce json
// @Param id path int true "Template ID"
// @Param status body object{status=string} true "Neuer Status"
// @Success 200 {object} responses.APIResponse{data=models.ShiftTemplate}
// @Failure 400 {object} responses.APIResponse
// @Failure 404 {object} responses.APIResponse
// @Failure 500 {object} responses.APIResponse
// @Router /shifttemplates/{id}/status [put]
func HandleUpdateShiftTemplateStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var template models.ShiftTemplate

	if err := database.GetDB().First(&template, id).Error; err != nil {
		return c.Status(404).JSON(responses.ErrorResponse(responses.ErrNotFound))
	}

	var input struct {
		Status string `json:"status"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(responses.ErrorResponse(responses.ErrInvalidInput))
	}

	template.Status = input.Status
	if err := database.GetDB().Save(&template).Error; err != nil {
		return c.Status(500).JSON(responses.ErrorResponse(err.Error()))
	}

	return c.JSON(responses.SuccessResponse(responses.MsgSuccessUpdate, template))
}
