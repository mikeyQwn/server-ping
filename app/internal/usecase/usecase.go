package usecase

import (
	"net"
	"time"

	"github.com/mikeyQwn/server-ping/internal"
)

type Usecase struct {
}

func New() internal.Usecase {
	return &Usecase{}
}

func (u *Usecase) CheckIsOnline(addr string) error {
	c, err := net.DialTimeout("tcp", addr, time.Second*3)
	if err != nil {
		return err
	}
	c.Close()
	return nil
}
