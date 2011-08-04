package main
/*
 *  Filename:    options.go
 *  Author:      Bryan Matsuo <bmatsuo@soe.ucsc.edu>
 *  Created:     Wed Aug  3 22:17:18 PDT 2011
 *  Description: Parse arguments and options from the command line.
 */
import (
    "os"
    "flag"
)

type Options struct {
    verbose bool
}
var opt = Options{ }

func SetupFlags() *flag.FlagSet {
    var fs = flag.NewFlagSet("snail", flag.ExitOnError)
    fs.BoolVar(&(opt.verbose), "v", false, "Verbose program output.")
    return fs
}
func VerifyFlags(fs *flag.FlagSet) {
}
func ParseFlags() {
    var fs = SetupFlags()
    fs.Parse(os.Args[1:])
    VerifyFlags(fs)
    // Process the verified options...
}