package db

type GormService interface {
	Conn() error
}
