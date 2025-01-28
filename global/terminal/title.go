//go:build !windows

package terminal

import (
	"fmt"
	"time"

	"github.com/Fripine/go-cqhttp-apad/internal/base"
)

// SetTitle 设置标题为 go-cqhttp `版本` `版权`
func SetTitle() {
	fmt.Printf("\033]0;go-cqhttp "+base.Version+" © 2020 - %d Mrs4s"+"\007", time.Now().Year())
}

// SetTitleExtra 设置标题为 go-cqhttp `版本` `Uin`
func SetTitleExtra(uin int64) {
	fmt.Printf("\033]0;go-cqhttp "+base.Version+" Uin: %v"+"\007", uin)
}
