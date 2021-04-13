package main

import (
	"fmt"
	"github.com/guozhe001/etherscan-api-practice/contant"
	"github.com/nanmu42/etherscan-api"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	client := getClient()
	for name, address := range getContractAddress() {
		if exits(name) {
			log.Printf("name=%s, address=%s source code exits\n", name, address)
			continue
		}
		source, err := client.ContractSource(address)
		if err != nil {
			panic(err)
		}
		log.Printf("len(source)=%d\n", len(source))
		for i, s := range source {
			log.Printf("index=%d, s=【\n%s\n】", i, s.SourceCode)
			//err := writeSourceCode(name, address, s)
			if err != nil {
				panic(err)
			}
		}
	}
}

// exits 检查文件是否存在
func exits(name string) bool {
	file, err := os.Open(getFileName(name))
	if err != nil {
		// 如果不存在，则返回false
		if os.IsNotExist(err) {
			return false
		}
		// 如果是其他错误，则返回true
		return true
	}
	// 如果file不为nil，返回true，否则返回false
	return file != nil
}

// getFileName 根据币的名称获取拼接的文件名称
func getFileName(name string) string {
	return fmt.Sprintf("%s.%s", name, contant.SoliditySuffix)
}

// getContractAddress 获取需要获取源码的合约地址
// 返回结果是一个map，key：contractName，value：contractAddress
func getContractAddress() map[string]string {
	m := make(map[string]string)
	m["BNB"] = "0xb8c77482e45f1f44de1745f52c74426c631bdd52"
	m["Uniswap"] = "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"
	return m
}

// writeSourceCode 把获取到的合约源码写入文件
func writeSourceCode(name, address string, s etherscan.ContractSource) error {
	return writeSourceCodeStr(name, address, s.SourceCode)
}

// writeSourceCode 把获取到的合约源码写入文件
func writeSourceCodeStr(name, address, s string) error {
	return ioutil.WriteFile(getFileName(name), []byte(s), fs.ModePerm)
}

// getClient 获取调用etherscan的client
func getClient() *etherscan.Client {
	key, err := getApiKey()
	if err != nil {
		panic(err)
	}
	log.Printf("apiKey=%s\n", key)
	return etherscan.New(etherscan.Mainnet, key)
}

// getApiKey 获取自己的apikey，会读取当前项目根目录下的api_key.txt文件
func getApiKey() (string, error) {
	file, err := ioutil.ReadFile("api_key.txt")
	if err != nil {
		return "", err
	}
	return string(file), nil
}
