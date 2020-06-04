package core

//go:generate mockgen -destination ../mock/context.go -package mock . Context

type Context interface {
	Config() Config
	Series() Series
}
