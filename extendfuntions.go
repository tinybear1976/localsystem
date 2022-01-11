package localsystem

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true

}

//仅返回主文件名（不包含扩展名）
func FileNameOnly(fullFilename string) string {
	//var filenameWithSuffix string
	filenameWithSuffix := filepath.Base(fullFilename)
	//var fileSuffix string
	fileSuffix := filepath.Ext(filenameWithSuffix)
	//var filenameOnly string
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly
}

// 返回主文件名(不包含路径) 与 扩展名。第一个返回是主文件名,第二返回是扩展名
func FileNameSplit(fullFilename string) (string, string) {
	//var filenameWithSuffix string
	filenameWithSuffix := filepath.Base(fullFilename)
	//var fileSuffix string
	fileSuffix := filepath.Ext(filenameWithSuffix)
	//var filenameOnly string
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly, fileSuffix
}

//返回文件名及扩展名，去除路径
func FileNameWithoutPath(fullFilename string) string {
	filenameWithSuffix := filepath.Base(fullFilename)
	return filenameWithSuffix
}

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()
	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func CreateMutiDir(filePath string) error {
	if !Exists(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func CurrentDirectory() (currentpath string, err error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return
	}
	currentpath = strings.Replace(dir, "\\", "/", -1)
	return
}
