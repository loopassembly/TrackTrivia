package initializers

import (
    "log"
    "railway/models"
)

func SeedCategories() {
    categories := []models.NonTechnical{
        {Name: "Financial"},
        {Name: "Establishment"},
        {Name: "Rajbhasa"},
        {Name: "DAR/Leave rules"},
    }

    for _, category := range categories {
        if err := DB.Where(models.NonTechnical{Name: category.Name}).FirstOrCreate(&category).Error; err != nil {
            log.Fatalf("Could not seed category %v: %v", category.Name, err)
        }
    }

    // Add similar logic for Technical and GK categories
    // For example:
    technicalCategories := []models.Technical{
        {Title: "C&W"},
        {Title: "Workshop"},
    }

    for _, category := range technicalCategories {
        if err := DB.Where(models.Technical{Title: category.Title}).FirstOrCreate(&category).Error; err != nil {
            log.Fatalf("Could not seed technical category %v: %v", category.Title, err)
        }
    }

    gkCategories := []models.GK{
        {SubCategory: "General Knowledge"},
        {SubCategory: "Aptitude"},
        {SubCategory: "Reasoning"},
    }

    for _, category := range gkCategories {
        if err := DB.Where(models.GK{SubCategory: category.SubCategory}).FirstOrCreate(&category).Error; err != nil {
            log.Fatalf("Could not seed GK category %v: %v", category.SubCategory, err)
        }
    }

    log.Println("Categories seeded successfully!")
}
