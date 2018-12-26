package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetFileContent(file string) string {
	b, e := ioutil.ReadFile(file)
	if e != nil {

		fmt.Println("GetFileContent", e)
		return ""
	}
	return string(b)
}
func GetFileContentbytes(file string) []byte {
	b, e := ioutil.ReadFile(file)
	if e != nil {

		fmt.Println("GetFileContent", e)
		return []byte("")
	}
	return b
}
func SaveToFile(filePath, str string) int64 {
	ioutil.WriteFile(filePath, []byte(str), 0644)
	return 0
}

func Exist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func CopyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println("copyFile", src, dst, err.Error())
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}
func CreateDir(dir string) bool {
	err := os.MkdirAll(dir, 0777)
	return err == nil
}
func RemoveFile(file string) error {
	err := os.Remove(file)
	return err
}
func RemoveAll(path string) error {
	err := os.RemoveAll(path)
	return err
}
func CheckAndCreateDir(path string) {
	arr := strings.Split(path, "/")
	var b string = ""
	for _, s := range arr {
		b += s + "/"
		if _, err := os.Stat(b); err != nil {
			if os.IsNotExist(err) {
				fmt.Println("not exist")
				os.Mkdir(b, os.ModeDir)
				os.Chmod(b, 0777)
			}
			//fmt.Println(err)
		}
	}

}

func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}
