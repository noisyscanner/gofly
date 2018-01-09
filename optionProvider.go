package main

import "flag"

type Options struct {
	Code string
	Conf string
	Since int
	Full bool
	OutDir string
	ShouldGzip bool
}

func (o *Options) ShouldWriteFile() bool {
	return len(o.OutDir) > 0
}

func getOpts() *Options {
	code := flag.String("code", "", "The language to do")
	conf := flag.String("conf", "", "DB config file")
	since := flag.Int("since", 0, "Only get changes since the given UNIX timestamp")
	fullPtr := flag.Bool("full", false, "Output the full language: verbs, tenses and pronouns")
	gzipPtr := flag.Bool("gz", false, "GZIP the output")
	outDir := flag.String("out", "", "Directory to write output to")

	flag.Parse()

	return &Options{
		Code: *code,
		Conf: *conf,
		Since: *since,
		Full: *fullPtr,
		OutDir: *outDir,
		ShouldGzip: *gzipPtr,
	}
}
