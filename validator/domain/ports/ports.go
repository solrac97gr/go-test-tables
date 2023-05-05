package ports

import "github.com/solrac97gr/go-test-tables/validator/domain/models"

type Validator interface {
	Struct(s models.EvaluableStruct) error
}
