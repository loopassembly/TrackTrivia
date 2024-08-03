package controllers

import (
	"railway/initializers"
	"railway/models"

	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user": user,
		},
	})
}

func Dashboard(c *fiber.Ctx) error {
	var nonTechnical []models.NonTechnical
	var technical []models.Technical
	var gk []models.GK

	
	if err := initializers.DB.Find(&nonTechnical).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Error fetching NonTechnical data"})
	}

	if err := initializers.DB.Find(&technical).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Error fetching Technical data"})
	}

	if err := initializers.DB.Find(&gk).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Error fetching GK data"})
	}

	
	data := fiber.Map{
		"nonTechnical": nonTechnical,
		"technical":    technical,
		"gk":           gk,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": data})
}






func AddBookToCategory(c *fiber.Ctx) error {
	type BookInput struct {
		Category    string `json:"category" validate:"required"`
		SubCategory string `json:"sub_category"`
		Title       string `json:"title" validate:"required"`
		BookContent string `json:"book_content" validate:"required"`
	}

	var input BookInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var err error

	switch input.Category {
	case "NonTechnical":
		err = addBookToNonTechnical(input.SubCategory, input.Title, input.BookContent)
	case "Technical":
		err = addBookToTechnical(input.SubCategory, input.Title, input.BookContent)
	case "GeneralKnowledge":
		err = addBookToGK(input.SubCategory, input.Title, input.BookContent)
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid category"})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Book added successfully"})
}

func addBookToNonTechnical(subCategory, title, content string) error {
    var category models.NonTechnical
    if err := initializers.DB.Where("name = ?", subCategory).First(&category).Error; err != nil {
        return err
    }

    book := models.Book{
        Title:          title,
        Content:        content,
        NonTechnicalID: category.ID,
    }

    return initializers.DB.Create(&book).Error
}


func addBookToTechnical(subCategory, title, content string) error {
	var err error

	if subCategory == "C&W" {
		var category models.CAndW
		if err = initializers.DB.Where("sub_category = ?", subCategory).First(&category).Error; err != nil {
			return err
		}

		book := models.Book{
			Title:   title,
			Content: content,
			CAndWID: category.ID,
		}
		err = initializers.DB.Create(&book).Error

	} else if subCategory == "Workshop" {
		var category models.Workshop
		if err = initializers.DB.Where("sub_category = ?", subCategory).First(&category).Error; err != nil {
			return err
		}

		book := models.Book{
			Title:      title,
			Content:    content,
			WorkshopID: category.ID,
		}
		err = initializers.DB.Create(&book).Error
	} else {
		err = fiber.NewError(fiber.StatusBadRequest, "Invalid subcategory")
	}

	return err
}

func addBookToGK(subCategory, title, content string) error {
	var category models.GK
	if err := initializers.DB.Where("sub_category = ?", subCategory).First(&category).Error; err != nil {
		return err
	}

	book := models.Book{
		Title:   title,
		Content: content,
		GKID:    category.ID,
	}

	return initializers.DB.Create(&book).Error
}



func ViewBooks(c *fiber.Ctx) error {
	category := c.Query("category")
	subCategory := c.Query("sub_category")

	if category == "" || subCategory == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Category and sub_category are required"})
	}

	var books []models.Book
	var err error

	switch category {
	case "NonTechnical":
		err = initializers.DB.Joins("JOIN non_technicals ON non_technicals.id = books.non_technical_id").
			Where("non_technicals.name = ?", subCategory).
			Find(&books).Error
	case "Technical":
		if subCategory == "C&W" {
			err = initializers.DB.Joins("JOIN c_and_ws ON c_and_ws.id = books.c_and_w_id").
				Where("c_and_ws.sub_category = ?", subCategory).
				Find(&books).Error
		} else if subCategory == "Workshop" {
			err = initializers.DB.Joins("JOIN workshops ON workshops.id = books.workshop_id").
				Where("workshops.sub_category = ?", subCategory).
				Find(&books).Error
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid subcategory"})
		}
	case "GeneralKnowledge":
		err = initializers.DB.Joins("JOIN gks ON gks.id = books.gk_id").
			Where("gks.sub_category = ?", subCategory).
			Find(&books).Error
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid category"})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	if len(books) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No books found in the specified category and subcategory"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": books})
}