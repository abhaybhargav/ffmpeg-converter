package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	ffmpeg "github.com/floostack/transcoder/ffmpeg"
)

func main() {
	var inputPath string
	var outPath string
	var outFormat string
	flag.StringVar(&inputPath, "input", "", "Input Path to read from")
	flag.StringVar(&outPath, "out", "", "Output Path to write to")
	flag.StringVar(&outFormat, "format", "", "Format for exported file")

	flag.Parse()

	overwrite := true

	opts := ffmpeg.Options{
		OutputFormat: &outFormat,
		Overwrite:    &overwrite,
	}

	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath:   "/usr/local/bin/ffmpeg",
		FfprobeBinPath:  "/usr/local/bin/ffprobe",
		ProgressEnabled: true,
	}

	fileList := []string{}
	err := filepath.Walk(inputPath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("Unable to parse files")
		}

		if f.IsDir() {
		}

		if matched, err := filepath.Match("*.mp4", filepath.Base(path)); err != nil {
		} else if matched {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range fileList {
		ext := path.Ext(file)
		outFile := file[0:len(file)-len(ext)] + "." + outFormat
		baseFile := filepath.Base(outFile)
		finalFile := fmt.Sprintf("%s/%s", outPath, baseFile)
		progress, err := ffmpeg.
			New(ffmpegConf).
			Input(file).
			Output(finalFile).
			WithOptions(opts).
			Start(opts)

		if err != nil {
			log.Fatal(err)
		}

		for msg := range progress {
			log.Printf("Processing file: %s %+v", baseFile, msg)
		}
	}
}
