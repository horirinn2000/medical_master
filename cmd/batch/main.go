package main

import (
	"fmt"
	"log"
	"medical_master/internal/batch"
	"medical_master/internal/model"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// DB接続設定
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "medical_master"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_SSLMODE", "disable"),
		getEnv("DB_TIMEZONE", "Asia/Tokyo"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("failed to get db from gorm: %v", err)
		}
		sqlDB.Close()
	}()

	// マイグレーション
	if err := db.AutoMigrate(
		&model.MedicalPractice{},
		&model.MedicalPracticeSupport{},
		&model.MedicalPracticeInclusion{},
		&model.MedicalPracticeConflict{},
		&model.InpatientBasicFee{},
		&model.CalculationCount{},
		&model.DentalPractice{},
		&model.DentalPracticeSupport{},
		&model.DentalPracticeInclusion{},
		&model.DentalPracticeConflict{},
		&model.DentalPracticeCalculationCount{},
		&model.DentalPracticeAdditionRelation{},
		&model.DentalPracticeProcedureMaterial{},
		&model.DentalPracticeCalculationCountLimit{},
		&model.DentalPracticeStep{},
		&model.DentalPracticeAgeConstraint{},
		&model.DentalPracticeConflictDetail{},
		&model.DentalPracticeActualDays{},
		&model.SpecialMedicalDevice{},
		&model.Comment{},
		&model.Medicine{},
		&model.HotCode{},
		&model.Disease{},
		&model.Medication{},
		&model.VisitingNursingFee{},
		&model.VisitingNursingAddition{},
		&model.VisitingNursingCalculationCount{},
		&model.VisitingNursingConflict{},
		&model.VisitingNursingFacilityStandard{},
		&model.Ward{},
		&model.Tooth{},
		&model.CommentRelation{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// インポート処理の実行
	batch.ImportMedicalPractices(db)
	batch.ImportMedicalPracticeSupports(db)
	batch.ImportMedicalPracticeInclusions(db)
	batch.ImportMedicalPracticeConflicts(db)
	batch.ImportInpatientBasicFees(db)
	batch.ImportCalculationCounts(db)

	batch.ImportDentalPractices(db)
	batch.ImportDentalPracticeSupports(db)
	batch.ImportDentalPracticeInclusions(db)
	batch.ImportDentalPracticeConflicts(db)
	batch.ImportDentalPracticeCalculationCounts(db)

	batch.ImportDevices(db)
	batch.ImportComments(db)
	batch.ImportMedicines(db)
	batch.ImportHotCodes(db)
	batch.ImportDiseases(db)
	batch.ImportMedications(db)

	batch.ImportVisitingNursingFees(db)
	batch.ImportVisitingNursingAdditions(db)
	batch.ImportVisitingNursingCalculationCounts(db)
	batch.ImportVisitingNursingConflicts(db)
	batch.ImportVisitingNursingFacilityStandards(db)

	batch.ImportWards(db)
	batch.ImportTeeth(db)
	batch.ImportCommentRelations(db)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
