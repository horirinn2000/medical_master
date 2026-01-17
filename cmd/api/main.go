package main

import (
	"log"
	"medical_master/internal/api"
	"medical_master/model"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// ServerImpl OpenAPIから生成されたServerInterfaceを実装する構造体
type ServerImpl struct{}

// 手動定義
type GetMedicationsSearchNameParams struct {
	Q string `form:"q" json:"q"`
}
type GetMedicationsSearchCodeParams struct {
	Q string `form:"q" json:"q"`
}

// --- 傷病名マスター API ---

func (s *ServerImpl) GetDiseases(c *gin.Context) {
	var diseases []model.Disease
	if err := db.Limit(1000).Find(&diseases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.attachMigrationTargets(&diseases)
	c.JSON(http.StatusOK, diseases)
}

func (s *ServerImpl) GetDiseasesSearchName(c *gin.Context, params api.GetDiseasesSearchNameParams) {
	var diseases []model.Disease
	searchTerm := "%" + params.Q + "%"
	if err := db.Where("name_kanji LIKE ? OR name_abbr LIKE ? OR name_kana LIKE ?", searchTerm, searchTerm, searchTerm).
		Limit(1000).Find(&diseases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.attachMigrationTargets(&diseases)
	c.JSON(http.StatusOK, diseases)
}

func (s *ServerImpl) GetDiseasesSearchCode(c *gin.Context, params api.GetDiseasesSearchCodeParams) {
	var diseases []model.Disease
	if err := db.Where("code = ?", params.Q).Limit(1000).Find(&diseases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.attachMigrationTargets(&diseases)
	c.JSON(http.StatusOK, diseases)
}

func (s *ServerImpl) GetDiseasesSearchIcd10(c *gin.Context, params api.GetDiseasesSearchIcd10Params) {
	var diseases []model.Disease
	if err := db.Where("icd10_1 = ? OR icd10_2 = ?", params.Q, params.Q).
		Limit(1000).Find(&diseases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.attachMigrationTargets(&diseases)
	c.JSON(http.StatusOK, diseases)
}

func (s *ServerImpl) attachMigrationTargets(diseases *[]model.Disease) {
	for i := range *diseases {
		d := &(*diseases)[i]
		if d.IsOld && d.MigrationManagementNumber != "" && d.MigrationManagementNumber != "00000000" {
			var target model.Disease
			if err := db.Where("disease_management_number = ? AND is_old = ?", d.MigrationManagementNumber, false).First(&target).Error; err == nil {
				d.MigrationTarget = &target
			}
		}
	}
}

// --- 病棟マスター API ---

func (s *ServerImpl) GetWards(c *gin.Context) {
	var wards []model.Ward
	if err := db.Limit(1000).Find(&wards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wards)
}

func (s *ServerImpl) GetWardsSearchName(c *gin.Context, params api.GetWardsSearchNameParams) {
	var wards []model.Ward
	searchTerm := "%" + params.Q + "%"
	if err := db.Where("name_kanji LIKE ? OR name_kana LIKE ? OR fullname LIKE ?", searchTerm, searchTerm, searchTerm).
		Limit(1000).Find(&wards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wards)
}

func (s *ServerImpl) GetWardsSearchCode(c *gin.Context, params api.GetWardsSearchCodeParams) {
	var wards []model.Ward
	if err := db.Where("code = ?", params.Q).Find(&wards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wards)
}

// --- 特定器材マスター API ---

func (s *ServerImpl) GetDevices(c *gin.Context) {
	var devices []model.SpecialMedicalDevice
	if err := db.Limit(1000).Find(&devices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, devices)
}

func (s *ServerImpl) GetDevicesSearchName(c *gin.Context, params api.GetDevicesSearchNameParams) {
	var devices []model.SpecialMedicalDevice
	searchTerm := "%" + params.Q + "%"
	if err := db.Where("name_kanji LIKE ? OR name_kana LIKE ? OR fullname LIKE ?", searchTerm, searchTerm, searchTerm).
		Limit(100).Find(&devices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, devices)
}

func (s *ServerImpl) GetDevicesSearchCode(c *gin.Context, params api.GetDevicesSearchCodeParams) {
	var devices []model.SpecialMedicalDevice
	if err := db.Where("code = ?", params.Q).Find(&devices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, devices)
}

// --- コメントマスター API ---

func (s *ServerImpl) GetComments(c *gin.Context) {
	var comments []model.Comment
	if err := db.Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (s *ServerImpl) GetCommentsSearchName(c *gin.Context, params api.GetCommentsSearchNameParams) {
	var comments []model.Comment
	searchTerm := "%" + params.Q + "%"
	if err := db.Where("name_kanji LIKE ? OR name_kana LIKE ?", searchTerm, searchTerm).
		Limit(100).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (s *ServerImpl) GetCommentsSearchCode(c *gin.Context, params api.GetCommentsSearchCodeParams) {
	var comments []model.Comment
	if err := db.Where("code = ?", params.Q).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

// --- 調剤行為マスター API ---

func (s *ServerImpl) GetMedicationsSearchName(c *gin.Context, params GetMedicationsSearchNameParams) {
	var medications []model.Medication
	searchTerm := "%" + params.Q + "%"
	if err := db.Where("name_kanji LIKE ? OR name_kana LIKE ?", searchTerm, searchTerm).
		Limit(100).Find(&medications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medications)
}

func (s *ServerImpl) GetMedicationsSearchCode(c *gin.Context, params GetMedicationsSearchCodeParams) {
	var medications []model.Medication
	if err := db.Where("code = ?", params.Q).Find(&medications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medications)
}

// --- 医科診療行為マスター API ---

func (s *ServerImpl) GetMedicalPracticesSearchName(c *gin.Context, params api.GetMedicalPracticesSearchNameParams) {
	var practices []model.MedicalPractice
	searchTerm := "%" + params.Q + "%"
	if err := db.Where("abbreviated_kanji_name LIKE ? OR abbreviated_kana_name LIKE ? OR basic_kanji_name LIKE ?", searchTerm, searchTerm, searchTerm).
		Limit(100).Find(&practices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, practices)
}

func (s *ServerImpl) GetMedicalPracticesSearchCode(c *gin.Context, params api.GetMedicalPracticesSearchCodeParams) {
	var practices []model.MedicalPractice
	if err := db.Where("medical_practice_code = ?", params.Q).Find(&practices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, practices)
}

func (s *ServerImpl) GetMedicalPracticesCodeInclusions(c *gin.Context, code string) {
	var inclusions []model.MedicalPracticeInclusion
	if err := db.Where("comprehensive_practice_code = ?", code).Find(&inclusions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, inclusions)
}

func (s *ServerImpl) GetMedicalPracticesCodeConflicts(c *gin.Context, code string) {
	var conflicts []model.MedicalPracticeConflict
	if err := db.Where("medical_practice_code = ?", code).Find(&conflicts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, conflicts)
}

func (s *ServerImpl) GetMedicalPracticesCodeSupports(c *gin.Context, code string) {
	var supports []model.MedicalPracticeSupport
	if err := db.Where("medical_practice_code = ?", code).Find(&supports).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, supports)
}

func (s *ServerImpl) GetMedicalPracticesCodeInpatientFees(c *gin.Context, code string) {
	var fees []model.InpatientBasicFee
	if err := db.Where("basic_fee_code = ?", code).Find(&fees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fees)
}

func (s *ServerImpl) GetMedicalPracticesCodeCalculationCounts(c *gin.Context, code string) {
	var counts []model.CalculationCount
	if err := db.Where("medical_practice_code = ?", code).Find(&counts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, counts)
}

// --- コメント関連テーブル API ---

func (s *ServerImpl) GetCommentsRelated(c *gin.Context, params api.GetCommentsRelatedParams) {
	var relations []model.CommentRelation
	if err := db.Where("act_code = ?", params.ActCode).Find(&relations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, relations)
}

// --- 歯式マスター API ---

func (s *ServerImpl) GetTeethSearchName(c *gin.Context, params api.GetTeethSearchNameParams) {
	var teeth []model.Tooth
	searchTerm := "%" + params.Q + "%"
	if err := db.Where("name LIKE ?", searchTerm).Limit(100).Find(&teeth).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teeth)
}

func (s *ServerImpl) GetTeethSearchCode(c *gin.Context, params api.GetTeethSearchCodeParams) {
	var teeth []model.Tooth
	if err := db.Where("code = ?", params.Q).Find(&teeth).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teeth)
}

func main() {
	// DB接続設定
	dsn := "host=localhost user=postgres password=postgres dbname=medical_master port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	r := gin.Default()

	// Swagger UI
	r.StaticFile("/doc/openapi.yaml", "./doc/openapi.yaml")
	config := &ginSwagger.Config{
		URL: "/doc/openapi.yaml",
	}
	r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))

	// ルートアクセスを Swagger UI にリダイレクト
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// 自動生成されたハンドラーの登録
	serverImpl := &ServerImpl{}
	api.RegisterHandlers(r, serverImpl)

	log.Println("Server starting on http://localhost:8080")
	log.Println("Swagger UI available on http://localhost:8080/swagger/index.html")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
