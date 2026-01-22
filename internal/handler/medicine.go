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
		query = query.Where("name_kanji LIKE ? OR basic_name LIKE ? OR name_kana LIKE ?", searchTerm, searchTerm, searchTerm)
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

		// 価格履歴の取得
		var prices []model.MedicinePrice
		s.DB.Where("medicine_code = ?", m.Code).Order("start_date DESC").Find(&prices)

		apiPrices := make([]api.MedicinePrice, len(prices))
		for i, p := range prices {
			pCopy := p
			pr := float32(pCopy.Price)
			apiPrices[i] = api.MedicinePrice{
				Price:     &pr,
				PriceType: &pCopy.PriceType,
				StartDate: &pCopy.StartDate,
			}
		}

		price := float32(m.Price)
		// APIのレスポンス形式に変換
		apiMedicine := api.Medicine{
			UpdateCategory:           &m.UpdateCategory,
			MasterType:               &m.MasterType,
			Code:                     &m.Code,
			NameKanjiLen:             &m.NameKanjiLen,
			NameKanji:                &m.NameKanji,
			NameKanaLen:              &m.NameKanaLen,
			NameKana:                 &m.NameKana,
			UnitCode:                 &m.UnitCode,
			UnitNameKanjiLen:         &m.UnitNameKanjiLen,
			UnitNameKanji:            &m.UnitNameKanji,
			PriceType:                &m.PriceType,
			Price:                    &price,
			Reserved1:                &m.Reserved1,
			NarcoticsCategory:        &m.NarcoticsCategory,
			NeurolyticFlag:           &m.NeurolyticFlag,
			BiologicFlag:             &m.BiologicFlag,
			GenericFlag:              &m.GenericFlag,
			Reserved2:                &m.Reserved2,
			DentalSpecificFlag:       &m.DentalSpecificFlag,
			ContrastAgentCategory:    &m.ContrastAgentCategory,
			InjectionVolume:          &m.InjectionVolume,
			ListingMethodType:        &m.ListingMethodType,
			BrandNameCode:            &m.BrandNameCode,
			OldPriceType:             &m.OldPriceType,
			NameKanjiUpdateFlag:      &m.NameKanjiUpdateFlag,
			NameKanaUpdateFlag:       &m.NameKanaUpdateFlag,
			DosageForm:               &m.DosageForm,
			Reserved3:                &m.Reserved3,
			UpdateDate:               &m.UpdateDate,
			DiscontinuedDate:         &m.DiscontinuedDate,
			NationalDrugCode:         &m.NationalDrugCode,
			PublishOrder:             &m.PublishOrder,
			TransitionalDate:         &m.TransitionalDate,
			BasicName:                &m.BasicName,
			ListingDate:              &m.ListingDate,
			GenericNameCode:          &m.GenericNameCode,
			GenericNameStandard:      &m.GenericNameStandard,
			GenericAdditionType:      &m.GenericAdditionType,
			AntiHivFlag:              &m.AntiHivFlag,
			LongTermListingCode:      &m.LongTermListingCode,
			SelectionMedicalCategory: &m.SelectionMedicalCategory,
			Prices:                   &apiPrices,
		}

		var apiProducts []api.HotCode
		for _, p := range products {
			pCopy := p
			optQty := float32(pCopy.OptionPackageQuantity)
			optInQty := float32(pCopy.OptionPackageInQuantity)
			hot9PkgCount := float32(pCopy.Hot9PackageCount)
			hot9TotalVol := float32(pCopy.Hot9TotalVolume)
			apiProducts = append(apiProducts, api.HotCode{
				HotCode:                     &pCopy.HotCode,
				Hot7:                        &pCopy.Hot7,
				CompanyCode:                 &pCopy.CompanyCode,
				DispensingNo:                &pCopy.DispensingNo,
				LogisticsNo:                 &pCopy.LogisticsNo,
				JanCode:                     &pCopy.JanCode,
				NationalCode:                &pCopy.NationalCode,
				IndividualCode:              &pCopy.IndividualCode,
				ReceiptCode1:                &pCopy.ReceiptCode1,
				ReceiptCode2:                &pCopy.ReceiptCode2,
				NotificationName:            &pCopy.NotificationName,
				SalesName:                   &pCopy.SalesName,
				ReceiptName:                 &pCopy.ReceiptName,
				SpecUnit:                    &pCopy.SpecUnit,
				PackageForm:                 &pCopy.PackageForm,
				PackageUnit:                 &pCopy.PackageUnit,
				TotalVolumeUnit:             &pCopy.TotalVolumeUnit,
				Category:                    &pCopy.Category,
				Manufacturer:                &pCopy.Manufacturer,
				Distributor:                 &pCopy.Distributor,
				RecordType:                  &pCopy.RecordType,
				UpdateDate:                  &pCopy.UpdateDate,
				OptionPackageQuantity:       &optQty,
				OptionPackageQuantityUnit:   &pCopy.OptionPackageQuantityUnit,
				OptionPackageInQuantity:     &optInQty,
				OptionPackageInQuantityUnit: &pCopy.OptionPackageInQuantityUnit,
				OptionRecordType:            &pCopy.OptionRecordType,
				OptionUpdateDate:            &pCopy.OptionUpdateDate,
				Hot9:                        &pCopy.Hot9,
				Hot9Hot7:                    &pCopy.Hot9Hot7,
				Hot9CompanyCode:             &pCopy.Hot9CompanyCode,
				Hot9DispensingNo:            &pCopy.Hot9DispensingNo,
				Hot9LogisticsNo:             &pCopy.Hot9LogisticsNo,
				Hot9JanCode:                 &pCopy.Hot9JanCode,
				Hot9NationalCode:            &pCopy.Hot9NationalCode,
				Hot9IndividualCode:          &pCopy.Hot9IndividualCode,
				Hot9ReceiptCode1:            &pCopy.Hot9ReceiptCode1,
				Hot9ReceiptCode2:            &pCopy.Hot9ReceiptCode2,
				Hot9NotificationName:        &pCopy.Hot9NotificationName,
				Hot9SalesName:               &pCopy.Hot9SalesName,
				Hot9ReceiptName:             &pCopy.Hot9ReceiptName,
				Hot9SpecUnit:                &pCopy.Hot9SpecUnit,
				Hot9PackageForm:             &pCopy.Hot9PackageForm,
				Hot9PackageCount:            &hot9PkgCount,
				Hot9PackageUnit:             &pCopy.Hot9PackageUnit,
				Hot9TotalVolume:             &hot9TotalVol,
				Hot9TotalVolumeUnit:         &pCopy.Hot9TotalVolumeUnit,
				Hot9Category:                &pCopy.Hot9Category,
				Hot9Manufacturer:            &pCopy.Hot9Manufacturer,
				Hot9Distributor:             &pCopy.Hot9Distributor,
				Hot9UpdateCategory:          &pCopy.Hot9UpdateCategory,
				Hot9UpdateDate:              &pCopy.Hot9UpdateDate,
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
	var p model.HotCode
	if err := s.DB.Where("hot_code = ?", hotCode).First(&p).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	optQty := float32(p.OptionPackageQuantity)
	optInQty := float32(p.OptionPackageInQuantity)
	hot9PkgCount := float32(p.Hot9PackageCount)
	hot9TotalVol := float32(p.Hot9TotalVolume)

	// APIのレスポンス形式に変換 (すべての項目を網羅)
	apiProduct := api.HotCode{
		HotCode:                     &p.HotCode,
		Hot7:                        &p.Hot7,
		CompanyCode:                 &p.CompanyCode,
		DispensingNo:                &p.DispensingNo,
		LogisticsNo:                 &p.LogisticsNo,
		JanCode:                     &p.JanCode,
		NationalCode:                &p.NationalCode,
		IndividualCode:              &p.IndividualCode,
		ReceiptCode1:                &p.ReceiptCode1,
		ReceiptCode2:                &p.ReceiptCode2,
		NotificationName:            &p.NotificationName,
		SalesName:                   &p.SalesName,
		ReceiptName:                 &p.ReceiptName,
		SpecUnit:                    &p.SpecUnit,
		PackageForm:                 &p.PackageForm,
		PackageUnit:                 &p.PackageUnit,
		TotalVolumeUnit:             &p.TotalVolumeUnit,
		Category:                    &p.Category,
		Manufacturer:                &p.Manufacturer,
		Distributor:                 &p.Distributor,
		RecordType:                  &p.RecordType,
		UpdateDate:                  &p.UpdateDate,
		OptionPackageQuantity:       &optQty,
		OptionPackageQuantityUnit:   &p.OptionPackageQuantityUnit,
		OptionPackageInQuantity:     &optInQty,
		OptionPackageInQuantityUnit: &p.OptionPackageInQuantityUnit,
		OptionRecordType:            &p.OptionRecordType,
		OptionUpdateDate:            &p.OptionUpdateDate,
		Hot9:                        &p.Hot9,
		Hot9Hot7:                    &p.Hot9Hot7,
		Hot9CompanyCode:             &p.Hot9CompanyCode,
		Hot9DispensingNo:            &p.Hot9DispensingNo,
		Hot9LogisticsNo:             &p.Hot9LogisticsNo,
		Hot9JanCode:                 &p.Hot9JanCode,
		Hot9NationalCode:            &p.Hot9NationalCode,
		Hot9IndividualCode:          &p.Hot9IndividualCode,
		Hot9ReceiptCode1:            &p.Hot9ReceiptCode1,
		Hot9ReceiptCode2:            &p.Hot9ReceiptCode2,
		Hot9NotificationName:        &p.Hot9NotificationName,
		Hot9SalesName:               &p.Hot9SalesName,
		Hot9ReceiptName:             &p.Hot9ReceiptName,
		Hot9SpecUnit:                &p.Hot9SpecUnit,
		Hot9PackageForm:             &p.Hot9PackageForm,
		Hot9PackageCount:            &hot9PkgCount,
		Hot9PackageUnit:             &p.Hot9PackageUnit,
		Hot9TotalVolume:             &hot9TotalVol,
		Hot9TotalVolumeUnit:         &p.Hot9TotalVolumeUnit,
		Hot9Category:                &p.Hot9Category,
		Hot9Manufacturer:            &p.Hot9Manufacturer,
		Hot9Distributor:             &p.Hot9Distributor,
		Hot9UpdateCategory:          &p.Hot9UpdateCategory,
		Hot9UpdateDate:              &p.Hot9UpdateDate,
	}

	c.JSON(http.StatusOK, apiProduct)
}
