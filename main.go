// Package main
package main

import (
	"github.com/Fripine/go-cqhttp-apad/cmd/gocq"
	"github.com/Fripine/go-cqhttp-apad/global/terminal"

	_ "github.com/Fripine/go-cqhttp-apad/db/leveldb"   // leveldb 数据库支持
	_ "github.com/Fripine/go-cqhttp-apad/modules/silk" // silk编码模块
	// 其他模块
	// _ "github.com/Fripine/go-cqhttp-apad/db/sqlite3"   // sqlite3 数据库支持
	// _ "github.com/Fripine/go-cqhttp-apad/db/mongodb"    // mongodb 数据库支持
	// _ "github.com/Fripine/go-cqhttp-apad/modules/pprof" // pprof 性能分析
)

func main() {
	terminal.SetTitle()
	gocq.InitBase()
	gocq.PrepareData()
	gocq.LoginInteract()
	_ = terminal.DisableQuickEdit()
	_ = terminal.EnableVT100()
	gocq.WaitSignal()
	_ = terminal.RestoreInputMode()
}
