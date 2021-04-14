package util

import (
	"fmt"
	"github.com/guozhe001/etherscan-api-practice/contant"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// writeFile 把源码写入文件中
func WriteFile(coinName, fileName, context string) error {
	log.Printf("source.Sources: coinName=%s, fileName=%s, context=%s", coinName, fileName, context)
	absolutePath := GetFileAbsolutePath(coinName, fileName)
	index := strings.LastIndex(absolutePath, string(os.PathSeparator))
	path := absolutePath[0:index]
	if err := os.MkdirAll(path, fs.ModePerm); err != nil {
		return err
	}
	return ioutil.WriteFile(absolutePath, []byte(context), fs.ModePerm)
}

// GetFileAbsolutePath 根据币名称和文件名称获取绝对路径
func GetFileAbsolutePath(coinName, fileName string) string {
	return fmt.Sprintf("%s%s%s", GetCoinDir(coinName), string(os.PathSeparator), fileName)
}

// GetCoinDir 获取币的合约存储文件夹
func GetCoinDir(coinName string) string {
	return fmt.Sprintf("%s%s%s", contant.PkgContract, string(os.PathSeparator), coinName)
}

// GetFileName 根据币的名称获取拼接的文件名称
func GetFileName(coinName string) string {
	return fmt.Sprintf("%s.%s", coinName, contant.SoliditySuffix)
}

// exits 检查文件夹是否存在
func Exits(name string) bool {
	dir, err := ioutil.ReadDir(name)
	if err != nil {
		// 如果不存在，则返回false
		if os.IsNotExist(err) {
			return false
		}
		// 如果是其他错误，则返回true
		return true
	}
	// 如果文件夹下的文件个数大于0，则说明已经存在
	return len(dir) > 0
}
