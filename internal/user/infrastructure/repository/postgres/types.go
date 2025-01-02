package postgres

import (
	"proj/internal/domain/entity"
)

// PostgresID преобразует entity.ID в uint
func PostgresID(id entity.ID) uint {
	switch v := id.(type) {
	case uint:
		return v
	case int:
		return uint(v)
	default:
		return 0
	}
}

// ToEntityID преобразует uint в entity.ID
func ToEntityID(id uint) entity.ID {
	return id
}
