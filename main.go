package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"strconv"
	"os"
	"time"
)

func main() {

	code := flag.String("code", "", "The language to do")
	conf := flag.String("conf", "", "DB config file")
	since := flag.Int("since", 0, "Only get changes since the given UNIX timestamp")
	fullPtr := flag.Bool("full", false, "Output the full language: verbs, tenses and pronouns")
	gzipPtr := flag.Bool("gz", false, "GZIP the output")
	outDir := flag.String("out", "", "Directory to write output to")

	flag.Parse()

	shouldWriteFile := len(*outDir) > 0
	shouldGzip := *gzipPtr

	if *code == "" {
		fmt.Println("Please specify a language code.")
	} else {
		var service RealLanguageService
		if *conf != "" {
			configService := FileConfigService{File: *conf}
			service = RealLanguageService{configService: configService}
		} else {
			service = RealLanguageService{}
		}


		var (
			output []byte
			filename string
			err error
		)

		if shouldWriteFile {
			if *outDir != "." {
				if _, err := os.Stat(*outDir); os.IsNotExist(err) {
					os.Mkdir(*outDir, 0755)
				}
			}

			filename = *outDir + "/"
		}

		if *fullPtr {
			var lang Language
			lang, err = service.GetLang(*code)
			if err == nil {
				output, err = lang.MarshalJSON()
			}
			if shouldWriteFile {
				filename += strconv.Itoa(lang.Id) + *code + ".json.full"
			}
		} else {
			var (
				langID int
				verbConf VerbContainer
			)

			if *since > 0 {
				langID, verbConf, err = service.GetVerbsSince(*code, *since)
			} else {
				langID, verbConf, err = service.GetVerbsOnly(*code)
			}

			if err == nil {
				output, err = verbConf.MarshalJSON()
			}

			if shouldWriteFile && *since > 0 {
				filename += strconv.Itoa(langID) + *code + strconv.FormatInt(time.Now().Unix(), 10) + ".json"
			} else {
				filename += strconv.Itoa(langID) + *code + ".json"
			}
		}

		if err == nil {
			var zippedOutput []byte

			if shouldGzip {
				zippedOutput, err = zipBytes(output)
			}
			
			if shouldWriteFile {
				err = ioutil.WriteFile(filename, output, 0644)

				if shouldGzip && len(zippedOutput) > 0 {
					err = ioutil.WriteFile(filename + ".gz", zippedOutput, 0644)
				}

				if err == nil {
					fmt.Println("SUCCESS: Written " + filename)
					if shouldGzip {
						fmt.Println("Written " + filename + ".gz")
					}
				}
			} else {
				if shouldGzip {
					fmt.Println(string(zippedOutput))
				} else {
					fmt.Println(string(output))
				}
			}
		} else {
			fmt.Println(err)
		}

		if err != nil {
			fmt.Printf("ERROR: %s", err.Error())
		}

	}
}
