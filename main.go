package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"os"
	"time"
)

func main() {
	opts := getOpts()

	if opts.Code == "" {
		fmt.Println("Please specify a language code.")
		return
	}

	var (
		configService ConfigService
		langID int
		output []byte
		err error
	)

	if opts.Conf != "" {
		configService = FileConfigService{File: opts.Conf}
	}

	service := &RealLanguageService{configService: configService}

	if opts.Full {
		langID, output, err = performFull(opts, service)
	} else {
		if opts.Since > 0 {
			langID, output, err = performSince(opts, service)
		} else {
			langID, output, err = performVerbs(opts, service)
		}
	}

	if err == nil {
		err = writeOut(opts, langID, output)
	}

	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}

func performFull(opts *Options, service LanguageService) (int, []byte, error) {
	var (
		output []byte
		err error
	)

	lang, err := service.GetLang(opts.Code)

	if err == nil {
		output, err = lang.MarshalJSON()
	}

	return lang.Id, output, err
}

func performVerbs(opts *Options, service LanguageService) (int, []byte, error) {
	var (
		langID int
		verbConf VerbContainer
		output []byte
		err error
	)

	langID, verbConf, err = service.GetVerbsOnly(opts.Code)

	if err == nil {
		output, err = verbConf.MarshalJSON()
	}

	return langID, output, err
}

func performSince(opts *Options, service LanguageService) (int, []byte, error) {
	var (
		langID int
		verbConf VerbContainer
		output []byte
		err error
	)

	langID, verbConf, err = service.GetVerbsSince(opts.Code, opts.Since)

	if err == nil {
		output, err = verbConf.MarshalJSON()
	}

	return langID, output, err
}

func getFileName(opts *Options, langId int) string {
	if !opts.ShouldWriteFile() {
		return ""
	}

	var prefix string

	outDir := opts.OutDir
	if outDir != "." {
		if _, err := os.Stat(outDir); os.IsNotExist(err) {
			os.Mkdir(outDir, 0755)
		}
	}

	prefix = outDir + "/" + strconv.Itoa(langId) + opts.Code

	if opts.Full {
		return prefix + ".json.full"
	}

	if opts.Since > 0 {
		return prefix + strconv.FormatInt(time.Now().Unix(), 10) + ".json"
	}

	return prefix + ".json"

}

func writeOut(opts *Options, langID int, output []byte) error {
	filename := getFileName(opts, langID)
	var (
		zippedOutput []byte
		err error
	)

	if opts.ShouldGzip {
		zippedOutput, err = zipBytes(output)
	}

	if filename != "" {
		err = ioutil.WriteFile(filename, output, 0644)
		if err == nil {
			fmt.Println("SUCCESS: Written " + filename)
		}

		if opts.ShouldGzip && len(zippedOutput) > 0 {
			err = ioutil.WriteFile(filename + ".gz", zippedOutput, 0644)
			if err == nil {
				fmt.Println("Written " + filename + ".gz")
			}
		}
	} else {
		if opts.ShouldGzip {
			fmt.Println(string(zippedOutput))
		} else {
			fmt.Println(string(output))
		}
	}

	return err
}