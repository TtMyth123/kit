package TmpFileKit

import (
	"fmt"
	"os"
	"time"
	"ttmyth123/kit/fileKit"
)

var aTmpFileKit *TmpFileKit
func init() {
	aTmpFileKit = new(TmpFileKit)
	aTmpFileKit.mpFile = make(map[string]string)
}
type TmpFileKit struct {
	mpFile map[string]string
}

func AddFile(key,filePath string, t int) {
	aTmpFileKit.mpFile[key] = filePath
	go func(key string) {
		t := time.NewTimer(time.Second * time.Duration(t))
		for{
			<-t.C
			t.Stop()
			aFilePath := aTmpFileKit.mpFile[key]
			os.Remove(aFilePath)
		}
	}(key)
}
func ToNewPath(key, path,name string) (string,error) {
	filePath := aTmpFileKit.mpFile[key]
	fileInfo, e := os.Stat(filePath)
	if e!= nil {
		return filePath,e
	}

	if name=="" {
		fileKit.CreateMutiDir(path)
		a :=fileInfo.Name()
		newFilePath := fmt.Sprintf("%s/%s",path,a)
		return newFilePath,os.Rename(filePath,newFilePath)

	} else {
		fileKit.CreateMutiDir(path)
		newFilePath := fmt.Sprintf("%s/%s",path,name)
		return newFilePath, os.Rename(filePath,newFilePath)
	}
}

func GetFilePath(key string)string  {
	return aTmpFileKit.mpFile[key]
}


