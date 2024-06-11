package internal

type Usecase interface {
	CheckIsOnline(addr string) error
}
