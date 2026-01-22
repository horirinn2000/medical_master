package main

import (
	"fmt"
	"log"
	"medical_master/internal/batch"
	"medical_master/internal/model"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// .envファイルの読み込み
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

	// 対象モデルのリスト
	tables := []interface{}{
		&model.SpecialMedicalDevice{},
		&model.Comment{},
		&model.Disease{},
		&model.Medication{},
		&model.Medicine{},
		&model.MedicinePrice{},
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
	}

	// マイグレーション
	if err := db.AutoMigrate(tables...); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// 全テーブルのTruncate
	truncateTables(db, tables)

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

	// 歯科診療行為マスター（加算）のインポート
	batch.ImportDentalPracticeAdditionRelations(db)

	// 歯科診療行為マスター（算定回数・正規）のインポート
	batch.ImportDentalPracticeCalculationCountsMaster(db)

	// 歯科診療行為マスター（きざみ）のインポート
	batch.ImportDentalPracticeSteps(db)

	// 歯科診療行為マスター（年齢制限）のインポート
	batch.ImportDentalPracticeAgeConstraints(db)

	// 歯科診療行為マスター（背反・正規）のインポート
	batch.ImportDentalPracticeConflictsMaster(db)

	// 歯科診療行為マスター（実日数）のインポート
	batch.ImportDentalPracticeActualDays(db)
}

func truncateTables(db *gorm.DB, tables []interface{}) {
	var tableNames []string
	for _, table := range tables {
		stmt := &gorm.Statement{DB: db}
		if err := stmt.Parse(table); err != nil {
			log.Fatalf("failed to parse table: %v", err)
		}
		tableNames = append(tableNames, stmt.Schema.Table)
	}

	if len(tableNames) == 0 {
		return
	}

	// PostgreSQL特有のTRUNCATE構文: TRUNCATE table1, table2, ... RESTART IDENTITY CASCADE;
	query := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", strings.Join(tableNames, ", "))
	if err := db.Exec(query).Error; err != nil {
		log.Fatalf("failed to truncate tables: %v", err)
	}
	log.Println("All tables truncated successfully.")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
