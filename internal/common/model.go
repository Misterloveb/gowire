package common

// ModleDataInterface TODO
type ModleDataInterface[T any] interface {
	Set([]*T)
	GetOne(col string, query string, args ...any) T
	Get(col string, query string, args ...any) []*T
}
type ModelData[T any] struct {
}
