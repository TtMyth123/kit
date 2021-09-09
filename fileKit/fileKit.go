package fileKit

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetTextFile(fileName string) (string ,error)  {
	fileT,e:= ioutil.ReadFile(fileName)
	return string(fileT),e
}

func CreateFileEx(c []byte,path, fileName string) error {
	e:= CreateMutiDir(path)
	if e!=nil {
		return e
	}

	localpath := fmt.Sprintf("%s/%s", path,fileName)

	file, err := os.OpenFile(
		localpath,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(c)
	return err
}

func  CreateFile(c []byte, fileName string) error {
	file, err := os.OpenFile(
		fileName,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(c)
	return err
}

func IsFileExist(path string) (bool, error) {
	fileInfo, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}
	//我这里判断了如果是0也算不存在
	if fileInfo.Size() == 0 {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, err
}
func CreateMutiDir(filePath string) error{
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			//fmt.Println("创建文件夹失败,error info:", err)
			return err
		}
		return err
	}
	return nil
}

func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}