package batch

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"gorm.io/gorm"
)

func importCsv[T any](db *gorm.DB, recordLength int, filePath string, convert func([]string) T) error {
	fmt.Printf("Starting import from %s...\n", filePath)
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open csv file: %w", err)
	}
	defer f.Close()

	reader := csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	reader.LazyQuotes = true

	var models []T
	batchSize := 500
	count := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to read record: %v", err)
			continue
		}

		if len(record) < recordLength {
			log.Printf("invalid record length: expected at least %d, got %d", recordLength, len(record))
			continue
		}

		model := convert(record)

		models = append(models, model)

		if len(models) >= batchSize {
			if err := db.Create(&models).Error; err != nil {
				return fmt.Errorf("failed to bulk insert %s: %w", filePath, err)
			}
			count += len(models)
			fmt.Printf("Inserted %d records from %s...\n", count, filePath)
			models = nil
		}
	}

	if len(models) > 0 {
		if err := db.Create(&models).Error; err != nil {
			return fmt.Errorf("failed to bulk insert %s (last batch): %w", filePath, err)
		}
		count += len(models)
	}

	fmt.Printf("%s Batch finished. Total %d records inserted.\n", filePath, count)

	return nil
}
