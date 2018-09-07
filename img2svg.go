package main

import (
	"os"
	"os/exec"
	"flag"
	"log"
	"path/filepath"
	"io/ioutil"
	"strings"
)

func main() {
	srcPtr := flag.String("src", "", "source image path")
	convertPtr := flag.String("convert", "/usr/local/bin/convert", "imagemagick convert command path.")
	potracePtr := flag.String("potrace", "/usr/local/bin/potrace", "potrace command path.")
	flag.Parse()
	if _, err := os.Stat(*convertPtr); os.IsNotExist(err) {
		log.Fatal("\"brew install imagemagick\" 명령을 이용해서 imagemagick 명령어를 설치해주세요.")
	}
	if _, err := os.Stat(*potracePtr); os.IsNotExist(err) {
		log.Fatal("\"brew install potrace\" 명령을 이용해서 potrace 명령어를 설치해주세요.")
	}
	if *srcPtr == "" {
		flag.PrintDefaults()
		log.Fatal("-src 옵션을 이용해서 이미지 경로를 설정해 주세요.")
	}
	src, err := filepath.Abs(*srcPtr)
	if err != nil {
		log.Fatal(err)
	}
	tempDir, err := ioutil.TempDir("", "bmp")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir) //  clean up
	ext := filepath.Ext(src)
	bmpPath := tempDir + "/" + strings.Replace(filepath.Base(src), ext, ".bmp", -1)
	svgPath := strings.Replace(src, ext, ".svg", -1)
	err = exec.Command(*convertPtr, *srcPtr, bmpPath).Run()
	if err != nil {
		log.Fatal(err)
	}
	err = exec.Command(*potracePtr, "--svg", bmpPath, "-o", svgPath, "-k", "0.8").Run()
	if err != nil {
		log.Fatal(err)
	}
}
