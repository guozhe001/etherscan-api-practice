package main

import (
	"fmt"
	"github.com/guozhe001/etherscan-api-practice/contant"
	"github.com/nanmu42/etherscan-api"
	"io/fs"
	"io/ioutil"
	"log"
)

func main() {
	client := getClient()
	for name, address := range getContractAddress() {
		source, err := client.ContractSource(address)
		if err != nil {
			panic(err)
		}
		log.Printf("len(source)=%d\n", len(source))
		for i, s := range source {
			log.Printf("index=%d, s=【\n%s\n】", i, s.SourceCode)
			err := writeSourceCode(name, address, s)
			if err != nil {
				panic(err)
			}
		}
	}
}

// getContractAddress 获取需要获取源码的合约地址
// 返回结果是一个map，key：contractName，value：contractAddress
func getContractAddress() map[string]string {
	m := make(map[string]string)
	m["BNB"] = "0xb8c77482e45f1f44de1745f52c74426c631bdd52"
	return m
}

// writeSourceCode 把获取到的合约源码写入文件
func writeSourceCode(name, address string, s etherscan.ContractSource) error {
	return writeSourceCodeStr(name, address, s.SourceCode)
}

// writeSourceCode 把获取到的合约源码写入文件
func writeSourceCodeStr(name, address, s string) error {
	return ioutil.WriteFile(fmt.Sprintf("%s.%s", name, contant.SoliditySuffix), []byte(s), fs.ModePerm)
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
