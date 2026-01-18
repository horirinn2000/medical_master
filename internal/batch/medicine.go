package batch

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"medical_master/internal/model"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"gorm.io/gorm"
)

func ImportMedicines(db *gorm.DB) {
	importCsv(db, 42, "csv/medicine/y_20260116.csv", func(record []string) model.Medicine {
		return model.Medicine{
			UpdateCategory:           record[0],
			MasterType:               record[1],
			Code:                     record[2],
			NameKanjiLen:             parseInt(record[3]),
			NameKanji:                record[4],
			NameKanaLen:              parseInt(record[5]),
			NameKana:                 record[6],
			UnitCode:                 record[7],
			UnitNameKanjiLen:         parseInt(record[8]),
			UnitNameKanji:            record[9],
			PriceType:                record[10],
			Price:                    parseFloat(record[11]),
			Reserved1:                record[12],
			NarcoticsCategory:        record[13],
			NeurolyticFlag:           record[14],
			BiologicFlag:             record[15],
			GenericFlag:              record[16],
			Reserved2:                record[17],
			DentalSpecificFlag:       record[18],
			ContrastAgentCategory:    record[19],
			InjectionVolume:          parseInt(record[20]),
			ListingMethodType:        record[21],
			BrandNameCode:            record[22],
			OldPriceType:             record[23],
			OldPrice:                 parseFloat(record[24]),
			NameKanjiUpdateFlag:      record[25],
			NameKanaUpdateFlag:       record[26],
			DosageForm:               record[27],
			Reserved3:                record[28],
			UpdateDate:               record[29],
			DiscontinuedDate:         record[30],
			NationalDrugCode:         record[31],
			PublishOrder:             parseInt(record[32]),
			TransitionalDate:         record[33],
			BasicName:                record[34],
			ListingDate:              record[35],
			GenericNameCode:          record[36],
			GenericNameStandard:      record[37],
			GenericAdditionType:      record[38],
			AntiHivFlag:              record[39],
			LongTermListingCode:      record[40],
			SelectionMedicalCategory: record[41],
		}
	})
}

func ImportHotCodes(db *gorm.DB) {
	fmt.Println("Starting import from csv/medicine/MEDIS20251130.TXT and MEDIS20251130_OP.TXT...")

	// 1. オプション情報を先に読み込んでマップに保持
	optionsMap := make(map[string]struct {
		PackQty    float64
		PackUnit   string
		PackInQty  float64
		PackInUnit string
	})

	optFile, err := os.Open("csv/medicine/MEDIS20251130_OP.TXT")
	if err != nil {
		log.Printf("warning: failed to open hot code option file: %v", err)
	} else {
		defer optFile.Close()
		optReader := csv.NewReader(transform.NewReader(optFile, japanese.ShiftJIS.NewDecoder()))
		optReader.LazyQuotes = true
		for {
			record, err := optReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				continue
			}
			if len(record) < 7 || record[0] == "" {
				continue
			}
			optionsMap[record[0]] = struct {
				PackQty    float64
				PackUnit   string
				PackInQty  float64
				PackInUnit string
			}{
				PackQty:    parseFloat(record[1]),
				PackUnit:   record[2],
				PackInQty:  parseFloat(record[3]),
				PackInUnit: record[4],
			}
		}
	}

	// 2. 標準セットを読み込んでDB保存
	importCsv(db, 24, "csv/medicine/MEDIS20251130.TXT", func(record []string) model.HotCode {
		hotCode := model.HotCode{
			HotCode:          record[0],
			Hot7:             record[1],
			CompanyCode:      record[2],
			DispensingNo:     record[3],
			LogisticsNo:      record[4],
			JanCode:          record[5],
			NationalCode:     record[6],
			IndividualCode:   record[7],
			ReceiptCode1:     record[8],
			ReceiptCode2:     record[9],
			NotificationName: record[10],
			SalesName:        record[11],
			ReceiptName:      record[12],
			SpecUnit:         record[13],
			PackageForm:      record[14],
			PackageCount:     parseFloat(record[15]),
			PackageUnit:      record[16],
			TotalVolume:      parseFloat(record[17]),
			TotalVolumeUnit:  record[18],
			Category:         record[19],
			Manufacturer:     record[20],
			Distributor:      record[21],
			RecordType:       record[22],
			UpdateDate:       record[23],
		}

		// オプション情報のマージ
		if opt, ok := optionsMap[hotCode.HotCode]; ok {
			hotCode.OptionPackageQuantity = opt.PackQty
			hotCode.OptionPackageQuantityUnit = opt.PackUnit
			hotCode.OptionPackageInQuantity = opt.PackInQty
			hotCode.OptionPackageInQuantityUnit = opt.PackInUnit
		}

		return hotCode
	})

	fmt.Println("Finished import from csv/medicine/MEDIS20251130.TXT and MEDIS20251130_OP.TXT...")
}
