package main

import (
	"log"
	"medical_master/internal/batch"
	"medical_master/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// DB接続設定
	dsn := "host=localhost user=postgres password=postgres dbname=medical_master port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// マイグレーション
	if err := db.AutoMigrate(
		&model.SpecialMedicalDevice{},
		&model.Comment{},
		&model.Disease{},
		&model.Medication{},
		&model.Medicine{},
		&model.HotCode{},
		&model.CommentRelation{},
		&model.Ward{},
		&model.Tooth{},
		&model.MedicalPractice{},
		&model.MedicalPracticeInclusion{},
		&model.MedicalPracticeConflict{},
		&model.MedicalPracticeSupport{},
		&model.InpatientBasicFee{},
		&model.CalculationCount{},
		&model.DentalPractice{},
		&model.DentalPracticeAdditionRelation{},
		&model.DentalPracticeCalculationCount{},
		&model.DentalPracticeStep{},
		&model.DentalPracticeAgeConstraint{},
		&model.DentalPracticeConflict{},
		&model.DentalPracticeActualDays{},
		&model.DentalPracticeSupport{},
		&model.DentalPracticeInclusion{},
		&model.DentalPracticeConflictDetail{},
		&model.DentalPracticeCalculationCountLimit{},
		&model.VisitingNursingFee{},
		&model.VisitingNursingAddition{},
		&model.VisitingNursingCalculationCount{},
		&model.VisitingNursingConflict{},
		&model.VisitingNursingFacilityStandard{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// 訪問看護療養費マスターのインポート
	batch.ImportVisitingNursingFees(db)
	batch.ImportVisitingNursingAdditions(db)
	batch.ImportVisitingNursingCalculationCounts(db)
	batch.ImportVisitingNursingConflicts(db)
	batch.ImportVisitingNursingFacilityStandards(db)

	// 病棟マスターのインポート
	batch.ImportWards(db)

	// 歯式マスターのインポート
	batch.ImportTeeth(db)

	// 特定器材マスターのインポート
	batch.ImportDevices(db)

	// コメントマスターのインポート
	batch.ImportComments(db)

	// コメント関連テーブルのインポート
	batch.ImportCommentRelations(db)

	// 傷病名マスターのインポート
	batch.ImportDiseases(db)

	// 調剤行為マスターのインポート
	batch.ImportMedications(db)

	// 医薬品マスターのインポート
	batch.ImportMedicines(db)

	// 医薬品HOTコードマスターのインポート
	batch.ImportHotCodes(db)

	// 医科診療行為マスターのインポート
	batch.ImportMedicalPractices(db)

	// 包括テーブルのインポート
	batch.ImportMedicalPracticeInclusions(db)

	// 背反テーブルのインポート
	batch.ImportMedicalPracticeConflicts(db)

	// 補助マスターテーブルのインポート
	batch.ImportMedicalPracticeSupports(db)

	// 入院基本料テーブルのインポート
	batch.ImportInpatientBasicFees(db)

	// 算定回数テーブルのインポート
	batch.ImportCalculationCounts(db)

	// 歯科診療行為マスターのインポート
	batch.ImportDentalPractices(db)

	// 歯科電子点数表（補助マスター）のインポート
	batch.ImportDentalPracticeSupports(db)

	// 歯科電子点数表（包括）のインポート
	batch.ImportDentalPracticeInclusions(db)

	// 歯科電子点数表（背反）のインポート
	batch.ImportDentalPracticeConflicts(db)

	// 歯科電子点数表（算定回数）のインポート
	batch.ImportDentalPracticeCalculationCounts(db)
}
