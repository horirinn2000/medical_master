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

// --- 歯科診療行為マスター API ---

func (s *ServerImpl) GetDentalPracticesSearchName(c *gin.Context, params api.GetDentalPracticesSearchNameParams) {
	var practices []model.DentalPractice
	searchTerm := "%" + params.Q + "%"
	if err := db.Where("abbreviated_kanji_name LIKE ?", searchTerm).
		Limit(100).Find(&practices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, practices)
}

func (s *ServerImpl) GetDentalPracticesSearchCode(c *gin.Context, params api.GetDentalPracticesSearchCodeParams) {
	var practices []model.DentalPractice
	if err := db.Where("medical_practice_code = ?", params.Q).Find(&practices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, practices)
}

func (s *ServerImpl) GetDentalPracticesCodeInclusions(c *gin.Context, code string) {
	var inclusions []model.DentalPracticeInclusion
	if err := db.Where("comprehensive_practice_code = ?", code).Find(&inclusions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, inclusions)
}

func (s *ServerImpl) GetDentalPracticesCodeConflicts(c *gin.Context, code string) {
	var conflicts []model.DentalPracticeConflictDetail
	if err := db.Where("medical_practice_code = ?", code).Find(&conflicts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, conflicts)
}

func (s *ServerImpl) GetDentalPracticesCodeSupports(c *gin.Context, code string) {
	var supports []model.DentalPracticeSupport
	if err := db.Where("medical_practice_code = ?", code).Find(&supports).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, supports)
}

func (s *ServerImpl) GetDentalPracticesCodeCalculationCounts(c *gin.Context, code string) {
	var counts []model.DentalPracticeCalculationCountLimit
	if err := db.Where("medical_practice_code = ?", code).Find(&counts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, counts)
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

// --- 訪問看護療養費マスター API ---

func (s *ServerImpl) GetVisitingNursingFeesSearchName(c *gin.Context, params api.GetVisitingNursingFeesSearchNameParams) {
	var fees []model.VisitingNursingFee
	searchTerm := "%" + params.Q + "%"
	if err := db.Where("basic_name LIKE ? OR abbreviated_name LIKE ? OR abbreviated_kana_name LIKE ?", searchTerm, searchTerm, searchTerm).
		Limit(100).Find(&fees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, s.mapVisitingNursingFeesToAPI(fees))
}

func (s *ServerImpl) GetVisitingNursingFeesSearchCode(c *gin.Context, params api.GetVisitingNursingFeesSearchCodeParams) {
	var fees []model.VisitingNursingFee
	if err := db.Where("visiting_nursing_fee_code = ?", params.Q).Find(&fees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, s.mapVisitingNursingFeesToAPI(fees))
}

func (s *ServerImpl) GetVisitingNursingFeesCodeAdditions(c *gin.Context, code string) {
	var additions []model.VisitingNursingAddition
	// 基本テーブルから加算グループを取得して検索する
	var fee model.VisitingNursingFee
	if err := db.Where("visiting_nursing_fee_code = ?", code).First(&fee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := db.Where("group_number = ?", fee.AdditionGroup).Find(&additions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, additions)
}

func (s *ServerImpl) GetVisitingNursingFeesCodeCalculationCounts(c *gin.Context, code string) {
	var counts []model.VisitingNursingCalculationCount
	if err := db.Where("visiting_nursing_fee_code = ?", code).Find(&counts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, counts)
}

func (s *ServerImpl) GetVisitingNursingFeesCodeConflicts(c *gin.Context, code string) {
	var conflicts []model.VisitingNursingConflict
	if err := db.Where("visiting_nursing_fee_code = ?", code).Find(&conflicts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, conflicts)
}

func (s *ServerImpl) GetVisitingNursingFeesCodeFacilityStandards(c *gin.Context, code string) {
	var standards []model.VisitingNursingFacilityStandard
	// 基本テーブルから施設基準グループを取得
	var fee model.VisitingNursingFee
	if err := db.Where("visiting_nursing_fee_code = ?", code).First(&fee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := db.Where("group_number = ?", fee.FacilityStandardGroup).Find(&standards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, standards)
}

func (s *ServerImpl) mapVisitingNursingFeesToAPI(fees []model.VisitingNursingFee) []api.VisitingNursingFee {
	res := make([]api.VisitingNursingFee, len(fees))
	for i, f := range fees {
		price := float32(f.Price)
		stepPrice := float32(f.StepPrice)
		lowerLimit := int(f.StepLowerLimit)
		upperLimit := int(f.StepUpperLimit)
		stepValue := int(f.StepValue)

		jobCodes := []string{
			f.JobCategory1, f.JobCategory2, f.JobCategory3, f.JobCategory4, f.JobCategory5,
			f.JobCategory6, f.JobCategory7, f.JobCategory8, f.JobCategory9, f.JobCategory10,
			f.JobCategory11, f.JobCategory12, f.JobCategory13, f.JobCategory14, f.JobCategory15,
		}
		// 空のコードを除去
		filteredJobCodes := []string{}
		for _, jc := range jobCodes {
			if jc != "00" && jc != "" {
				filteredJobCodes = append(filteredJobCodes, jc)
			}
		}

		symbols := []string{
			f.ReceiptSymbol1, f.ReceiptSymbol2, f.ReceiptSymbol3, f.ReceiptSymbol4, f.ReceiptSymbol5,
			f.ReceiptSymbol6, f.ReceiptSymbol7, f.ReceiptSymbol8, f.ReceiptSymbol9,
		}

		res[i] = api.VisitingNursingFee{
			AbbreviatedKanaName:        &f.AbbreviatedKanaName,
			AbbreviatedName:            &f.AbbreviatedName,
			AdditionGroup:              &f.AdditionGroup,
			BasicName:                  &f.BasicName,
			ChangeCategory:             &f.ChangeCategory,
			DataStandardCode:           &f.DataStandardCode,
			ElderlyMedicalCategory:     &f.ElderlyMedicalCategory,
			ExpiryDate:                 &f.ExpiryDate,
			FacilityStandardGroup:      &f.FacilityStandardGroup,
			JobCategoryCodes:           &filteredJobCodes,
			LowerAge:                   &f.LowerAge,
			MasterType:                 &f.MasterType,
			MedicalObservationCategory: &f.MedicalObservationCategory,
			NotificationBranch:         &f.NotificationBranch,
			NotificationItem:           &f.NotificationItem,
			NotificationSection:        &f.NotificationSection,
			NursingInstructionCategory: &f.NursingInstructionCategory,
			Price:                      &price,
			PriceCategory:              &f.PriceCategory,
			ReceiptDisplayItem:         &f.ReceiptDisplayItem,
			ReceiptDisplaySection:      &f.ReceiptDisplaySection,
			ReceiptDisplaySerial:       &f.ReceiptDisplaySerial,
			ReceiptSymbols:             &symbols,
			SoloAdditionCategory:       &f.SoloAdditionCategory,
			SpecialInstructionCategory: &f.SpecialInstructionCategory,
			StepCalculationCategory:    &f.StepCalculationCategory,
			StepLowerLimit:             &lowerLimit,
			StepPrice:                  &stepPrice,
			StepUpperLimit:             &upperLimit,
			StepValue:                  &stepValue,
			UpdateDate:                 &f.UpdateDate,
			UpperAge:                   &f.UpperAge,
			VisitTimesCategory:         &f.VisitTimesCategory,
			VisitingNursingFeeCode:     &f.VisitingNursingFeeCode,
			VisitingNursingType:        &f.VisitingNursingType,
		}
	}
	return res
}

// --- 医薬品マスター API ---

func (s *ServerImpl) GetMedicines(c *gin.Context, params api.GetMedicinesParams) {
	var results []api.MedicineSearchResult

	query := db.Model(&model.Medicine{})

	if params.ReceptCode != nil {
		query = query.Where("code = ?", *params.ReceptCode)
	}
	if params.Name != nil {
		searchTerm := "%" + *params.Name + "%"
		query = query.Where("name_kanji LIKE ? OR basic_name LIKE ?", searchTerm, searchTerm)
	}

	// HOTコードやJANコードによる検索の場合はMedicineとの紐付けが必要
	if params.JanCode != nil || params.HotCode != nil {
		var hotCodes []model.HotCode
		subQuery := db.Model(&model.HotCode{})
		if params.JanCode != nil {
			subQuery = subQuery.Where("jan_code = ?", *params.JanCode)
		}
		if params.HotCode != nil {
			subQuery = subQuery.Where("hot_code = ?", *params.HotCode)
		}
		subQuery.Select("receipt_code_1").Find(&hotCodes)

		var codes []string
		for _, h := range hotCodes {
			codes = append(codes, h.ReceiptCode1)
		}
		query = query.Where("code IN ?", codes)
	}

	var medicines []model.Medicine
	if err := query.Limit(100).Find(&medicines).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, m := range medicines {
		var products []model.HotCode
		db.Where("receipt_code_1 = ?", m.Code).Find(&products)

		price := float32(m.Price)
		// APIのレスポンス形式に変換
		apiMedicine := api.Medicine{
			Code:             &m.Code,
			NameKanji:        &m.NameKanji,
			UnitNameKanji:    &m.UnitNameKanji,
			Price:            &price,
			DosageForm:       &m.DosageForm,
			NationalDrugCode: &m.NationalDrugCode,
			BasicName:        &m.BasicName,
		}

		var apiProducts []api.HotCode
		for _, p := range products {
			pCopy := p
			optQty := float32(pCopy.OptionPackageQuantity)
			apiProducts = append(apiProducts, api.HotCode{
				HotCode:               &pCopy.HotCode,
				JanCode:               &pCopy.JanCode,
				ReceiptCode1:          &pCopy.ReceiptCode1,
				SalesName:             &pCopy.SalesName,
				Manufacturer:          &pCopy.Manufacturer,
				OptionPackageQuantity: &optQty,
			})
		}

		results = append(results, api.MedicineSearchResult{
			Medicine: &apiMedicine,
			Products: &apiProducts,
		})
	}

	c.JSON(http.StatusOK, results)
}

func (s *ServerImpl) GetMedicinesProductsHotCode(c *gin.Context, hotCode string) {
	var product model.HotCode
	if err := db.Where("hot_code = ?", hotCode).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, product)
}

// --- 公費マスター API ---

func (s *ServerImpl) GetPublicFundsNationalSearchCode(c *gin.Context, params api.GetPublicFundsNationalSearchCodeParams) {
	var funds []model.NationalPublicFund
	if err := db.Where("law_number = ?", params.Q).Find(&funds).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, funds)
}

func (s *ServerImpl) GetPublicFundsNationalSearchName(c *gin.Context, params api.GetPublicFundsNationalSearchNameParams) {
	var funds []model.NationalPublicFund
	searchTerm := "%" + params.Q + "%"
	if err := db.Where("project_name LIKE ?", searchTerm).Limit(100).Find(&funds).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, funds)
}

func (s *ServerImpl) GetPublicFundsLocalSearchPayerNumber(c *gin.Context, params api.GetPublicFundsLocalSearchPayerNumberParams) {
	var funds []model.LocalPublicFund
	if err := db.Where("payer_number_8 = ? OR payer_number_other = ?", params.Q, params.Q).Find(&funds).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, funds)
}

func (s *ServerImpl) GetPublicFundsLocalSearchName(c *gin.Context, params api.GetPublicFundsLocalSearchNameParams) {
	var funds []model.LocalPublicFund
	searchTerm := "%" + params.Q + "%"
	if err := db.Where("official_name LIKE ? OR abbreviated_name LIKE ?", searchTerm, searchTerm).Limit(100).Find(&funds).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, funds)
}

func (s *ServerImpl) GetPublicFundsLocalSearchRegion(c *gin.Context, params api.GetPublicFundsLocalSearchRegionParams) {
	var funds []model.LocalPublicFund
	query := db.Model(&model.LocalPublicFund{})
	if params.PrefectureCode != nil {
		query = query.Where("prefecture_code = ?", *params.PrefectureCode)
	}
	if params.MunicipalityCode != nil {
		query = query.Where("municipality_code = ?", *params.MunicipalityCode)
	}
	if err := query.Limit(100).Find(&funds).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, funds)
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
