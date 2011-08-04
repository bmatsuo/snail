package main
/*
 *  Filename:    options.go
 *  Author:      Bryan Matsuo <bmatsuo@soe.ucsc.edu>
 *  Created:     Wed Aug  3 22:17:18 PDT 2011
 *  Description: Parse arguments and options from the command line.
 */
import (
    "flag"
    "fmt"
    "os"
)

type options struct {
    n       int
}

var opt = options{}

func setupFlags() *flag.FlagSet {
    var fs = flag.NewFlagSet("snail", flag.ExitOnError)
    fs.IntVar(&(opt.n), "n", 5, "Specify snail dimensions (n x n).")
    return fs
}
func verifyFlags(fs *flag.FlagSet) {
    if opt.n < 0 {
        panic(fmt.Errorf("Negative matrix dimension %d", opt.n))
    }
}
func parseFlags() {
    var fs = setupFlags()
    fs.Parse(os.Args[1:])
    verifyFlags(fs)
}
