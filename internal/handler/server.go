package handler

import "gorm.io/gorm"

// ServerImpl OpenAPIから生成されたServerInterfaceを実装する構造体
type ServerImpl struct {
	DB *gorm.DB
}
