package handler

import (
	"medical_master/internal/api"
	"medical_master/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// --- 医薬品マスター API ---

// 医薬品マスター 検索API
func (s *ServerImpl) GetMedicines(c *gin.Context, params api.GetMedicinesParams) {
	var results []api.MedicineSearchResult

	query := s.DB.Model(&model.Medicine{})

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
		subQuery := s.DB.Model(&model.HotCode{})
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
		s.DB.Where("receipt_code_1 = ?", m.Code).Find(&products)

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

// 医薬品マスター HOTコード検索API
func (s *ServerImpl) GetMedicinesProductsHotCode(c *gin.Context, hotCode string) {
	var product model.HotCode
	if err := s.DB.Where("hot_code = ?", hotCode).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, product)
}
