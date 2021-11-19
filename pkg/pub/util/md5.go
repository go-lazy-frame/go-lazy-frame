package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

var (
	Md5Util = new(md5Util)
)

type md5Util struct {

}

func (*md5Util) Md5(raw string) string {
	h := md5.New()
	_, _ = io.WriteString(h, raw)
	return fmt.Sprintf("%x", h.Sum(nil))
}
