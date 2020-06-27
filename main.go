package main

import (
	"database/sql"
	"fmt"
	gofly "github.com/noisyscanner/gofly"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {
	err := errorWrapper(getOpts())

	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}

func errorWrapper(opts *Options) error {
	var (
		configService gofly.ConfigService
		langID        int
		output        []byte
	)

	if opts.Conf != "" {
		configService = gofly.FileConfigService{File: opts.Conf}
	}

	dbs := gofly.DatabaseService{ConfigService: configService}

	db, err := dbs.GetDb()

	if err != nil {
		return err
	}

	if opts.ImportFile != "" {
		err = performImport(opts, db)
	} else {
		if opts.Code == "" {
			return fmt.Errorf("please specify a language code")
		}

		service := &gofly.Fetcher{Db: db}

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
	}

	return err
}

func performImport(opts *Options, db *sql.DB) error {
	data, err := ioutil.ReadFile(opts.ImportFile)

	language := &gofly.Language{}

	if err == nil {
		err = language.UnmarshalJSON(data)
	}

	if language.Id == 0 || language.Code == "" {
		err = fmt.Errorf("malformed language data")
	}

	if err == nil {
		inserter := &gofly.Inserter{Db: db}
		err = inserter.InsertLanguage(language)
	}

	return err
}

func performFull(opts *Options, service gofly.LanguageService) (int, []byte, error) {
	var (
		output []byte
		err    error
	)

	lang, err := service.GetLang(opts.Code)

	if err == nil {
		output, err = lang.MarshalJSON()
	}

	return lang.Id, output, err
}

func performVerbs(opts *Options, service gofly.LanguageService) (int, []byte, error) {
	var (
		langID   int
		verbConf gofly.VerbContainer
		output   []byte
		err      error
	)

	langID, verbConf, err = service.GetVerbsOnly(opts.Code)

	if err == nil {
		output, err = verbConf.MarshalJSON()
	}

	return langID, output, err
}

func performSince(opts *Options, service gofly.LanguageService) (int, []byte, error) {
	var (
		langID   int
		verbConf gofly.VerbContainer
		output   []byte
		err      error
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

	prefix = outDir + "/" + opts.Code

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
		err          error
	)

	if opts.ShouldGzip {
		zippedOutput, err = gofly.ZipBytes(output)
	}

	if filename != "" {
		err = ioutil.WriteFile(filename, output, 0644)
		if err == nil {
			fmt.Println("SUCCESS: Written " + filename)
		}

		if opts.ShouldGzip && len(zippedOutput) > 0 {
			err = ioutil.WriteFile(filename+".gz", zippedOutput, 0644)
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
