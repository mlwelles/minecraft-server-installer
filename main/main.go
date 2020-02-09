package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var usage = `mcinstall {SERVER FOLDER} {SERVER SOFTWARE} {VERSION}
{SERVER FOLDER} -- Directory where server should be stored 
{SERVER SOFTWARE} -- spigot/paper 
{VERSION} -- Paper supports 1.15.2, 1.14.4, 1.12.2, and 1.8.8/1.8.9, spigot support all versions at https://www.spigotmc.org/wiki/buildtools/`

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
		flag.PrintDefaults()
		return
	}
	flag.Parse()
	if flag.NArg() != 3 {
		log.Fatal("error: wrong number of arguments")
	}
	path := flag.Arg(0)
	rev := flag.Arg(2)
	if flag.Arg(1) == "spigot" {
		InstallSpigot(path, rev)
	} else if flag.Arg(1) == "paper" {
		InstallPaper(path, rev)
	}

}
