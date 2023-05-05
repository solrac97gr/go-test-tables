package models

type EvaluableStruct interface {
	Validate() error
}
