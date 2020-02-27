package main

import (
    "github.com/HugoSmits86/dreamcast-vms-icondata-tool/icondata"
    "os"
    "image/png"
    "strings"
    "flag"
    "fmt"
    "log"
    "errors"
)

func showHelpText() {
    fmt.Println("Icontool version 1.0 (c) 2020 by Hugo Smits.")
    fmt.Println("Usage:")
    fmt.Println("    when vms file is presented as input, tool will decode into png.")
    fmt.Println("    when png file is presented as input, tool will encode into vms.")
    fmt.Println("    for more info type -help.")
}

func main() {
    outputFileName := ""
    inputFileName := ""
    
    flag.StringVar(&inputFileName, "i", "", "input file name.")
    flag.StringVar(&outputFileName, "o", "", "output file name")
    flag.Parse()

    tmp := strings.Split(inputFileName, ".")
    if len(tmp) < 2 {
        showHelpText()
        return
    } 

    inputFileType := strings.ToLower(tmp[1])
    if inputFileType != "vms" && inputFileType != "png" {
      log.Fatal(errors.New("input file must be a VMS or PNG file."))
    }

    tmp = strings.Split(outputFileName, ".")
    if len(tmp) < 2 {
        showHelpText()
        return
    } 

    outputFileType := strings.ToLower(tmp[1])
    if outputFileType != "vms" && outputFileType != "png" {
      log.Fatal(errors.New("output file must be a VMS or PNG file."))
    }

    if outputFileType == inputFileType {
        log.Fatal(errors.New("input file type cannot be the same as output file type."))
    }


    file, err := os.Open(inputFileName)
    if err != nil {
        log.Fatal(err)
    }

    if inputFileType == "vms" {
        _, img, err := icondata.Decode(file)
        if err != nil {
            log.Fatal(err)
        }

        file, _ = os.OpenFile(outputFileName, os.O_WRONLY|os.O_CREATE, 0600)
        defer file.Close()

        png.Encode(file, img)
    } else if inputFileType == "png" {
        img, err := png.Decode(file)
        if err != nil {
            log.Fatal(err)
        }

        file, _ = os.OpenFile(outputFileName, os.O_WRONLY|os.O_CREATE, 0600)
        defer file.Close()

        err = icondata.Encode(file, img)
        if err != nil {
            log.Fatal(err)
        }
    }
}
