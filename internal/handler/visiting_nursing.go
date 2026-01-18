package handler

import (
	"medical_master/internal/api"
	"medical_master/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- 訪問看護療養費マスター API ---

// 訪問看護療養費マスター (ア 基本テーブル) 名称検索API
func (s *ServerImpl) GetVisitingNursingFeesSearchName(c *gin.Context, params api.GetVisitingNursingFeesSearchNameParams) {
	var fees []model.VisitingNursingFee
	searchTerm := "%" + params.Q + "%"
	if err := s.DB.Where("basic_name LIKE ? OR abbreviated_name LIKE ? OR abbreviated_kana_name LIKE ?", searchTerm, searchTerm, searchTerm).
		Limit(100).Find(&fees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mapVisitingNursingFeesToAPI(fees))
}

// 訪問看護療養費マスター (ア 基本テーブル) コード検索API
func (s *ServerImpl) GetVisitingNursingFeesSearchCode(c *gin.Context, params api.GetVisitingNursingFeesSearchCodeParams) {
	var fees []model.VisitingNursingFee
	if err := s.DB.Where("visiting_nursing_fee_code = ?", params.Q).Find(&fees).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mapVisitingNursingFeesToAPI(fees))
}

// 訪問看護療養費マスター (イ 加算テーブル) コード検索API
func (s *ServerImpl) GetVisitingNursingFeesCodeAdditions(c *gin.Context, code string) {
	var additions []model.VisitingNursingAddition
	// 基本テーブルから加算グループを取得して検索する
	var fee model.VisitingNursingFee
	if err := s.DB.Where("visiting_nursing_fee_code = ?", code).First(&fee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := s.DB.Where("group_number = ?", fee.AdditionGroup).Find(&additions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, additions)
}

// 訪問看護療養費マスター (ウ 算定回数テーブル) コード検索API
func (s *ServerImpl) GetVisitingNursingFeesCodeCalculationCounts(c *gin.Context, code string) {
	var counts []model.VisitingNursingCalculationCount
	if err := s.DB.Where("visiting_nursing_fee_code = ?", code).Find(&counts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, counts)
}

// 訪問看護療養費マスター (エ 相互関係テーブル) コード検索API
func (s *ServerImpl) GetVisitingNursingFeesCodeConflicts(c *gin.Context, code string) {
	var conflicts []model.VisitingNursingConflict
	if err := s.DB.Where("visiting_nursing_fee_code = ?", code).Find(&conflicts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, conflicts)
}

// 訪問看護療養費マスター (オ 施設基準テーブル) コード検索API
func (s *ServerImpl) GetVisitingNursingFeesCodeFacilityStandards(c *gin.Context, code string) {
	var standards []model.VisitingNursingFacilityStandard
	// 基本テーブルから施設基準グループを取得
	var fee model.VisitingNursingFee
	if err := s.DB.Where("visiting_nursing_fee_code = ?", code).First(&fee).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := s.DB.Where("group_number = ?", fee.FacilityStandardGroup).Find(&standards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, standards)
}

// モデルからAPI用構造体へのマッピング
func mapVisitingNursingFeesToAPI(fees []model.VisitingNursingFee) []api.VisitingNursingFee {
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
