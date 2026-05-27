package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/araddon/qlbridge/datasource"
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/expr/builtins"
	"github.com/araddon/qlbridge/value"
)

/*
go build && time ./_bm --command=parse --cpuprofile=cpu.prof
go tool pprof _bm cpu.prof

go build && time ./_bm --command=vm --cpuprofile=cpu.prof
go tool pprof _bm cpu.prof
*/
var (
	cpuProfileFile string
	memProfileFile string
	logging        = "info"
	command        = "parse"

	msg = datasource.NewContextSimpleTs(
		map[string]value.Value{
			"int5":       value.NewIntValue(5),
			"item_count": value.NewStringValue("5"),
			"reg_date":   value.NewStringValue("2014/11/01"),
			"user_id":    value.NewStringValue("abc")},
		time.Now(),
	)
)

func init() {

	flag.StringVar(&logging, "logging", "info", "logging [ debug,info ]")
	flag.StringVar(&cpuProfileFile, "cpuprofile", "", "cpuprofile")
	flag.StringVar(&memProfileFile, "memprofile", "", "memProfileFile")
	flag.StringVar(&command, "command", "parse", "command to run [parse,vm]")
	flag.Parse()

	builtins.LoadAllBuiltins()

}

func main() {

	if cpuProfileFile != "" {
		f, err := os.Create(cpuProfileFile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	switch command {
	case "parse":
		runParse(100000, `select user_id, item_count * 2 as itemsx2, yy(reg_date) > 10 as regyy FROM stdio`, msg)

	case "vm":
		runVm(100000, `select user_id, item_count * 2 as itemsx2, yy(reg_date) > 10 as regyy FROM stdio`, msg)
	}
}

func runParse(repeat int, sql string, readContext expr.ContextReader) {
	_ = "STUB: not implemented"
	return
}

func runVm(repeat int, sql string, readContext expr.ContextReader) {
	_ = "STUB: not implemented"
	return
}

//log.Println(writeContext.All())
