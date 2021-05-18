package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/guozhe001/etherscan-api-practice/contant"
	"github.com/guozhe001/etherscan-api-practice/model"
	"github.com/guozhe001/etherscan-api-practice/util"
	"github.com/nanmu42/etherscan-api"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	client := getClient()
	for name, address := range getContractAddress() {
		log.Printf("name=%s, address=%s start get source code\n", name, address)
		if util.Exits(util.GetCoinDir(name)) {
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
	wantedPath := fmt.Sprintf("%s%s%s", contant.PkgContract, string(os.PathSeparator), contant.FileWanted)
	wantedFile, err := os.Open(wantedPath)
	if err != nil {
		panic(err)
	}
	defer wantedFile.Close()
	scanner := bufio.NewScanner(wantedFile)
	m := make(map[string]string)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		m[strings.TrimSpace(split[0])] = strings.TrimSpace(split[1])
	}
	return m
}

// writeSourceCode 把获取到的合约源码写入文件
func writeSourceCode(name, address string, s etherscan.ContractSource) error {
	return writeSourceCodeStr(name, address, s.SourceCode)
}

// writeSourceCode 把获取到的合约源码写入文件
func writeSourceCodeStr(name, address, s string) error {
	if isJson(s) {
		log.Println("is a json!")
		s = strings.TrimPrefix(s, "{")
		s = strings.TrimSuffix(s, "}")
		// 如果code中的是json，则分开处理
		if json.Valid([]byte(s)) {
			source := model.SourceCodeModel{}
			err := json.Unmarshal([]byte(s), &source)
			if err != nil {
				return err
			}
			log.Printf("source.Language=%s", source.Language)
			log.Printf("source.Settings=%s", source.Settings)

			for fileName, contextMap := range source.Sources {
				context := contextMap[model.KeyContent].(string)
				if err := util.WriteFile(name, fileName, context); err != nil {
					return err
				}
			}
		} else {
			return fmt.Errorf("name=%s is not a Valid json", name)
		}
		return nil
	} else {
		return util.WriteFile(name, util.GetFileName(name), s)
	}
}

func isJson(s string) bool {
	return strings.HasPrefix(s, "{{")
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
	file, err := ioutil.ReadFile(contant.FileApiKey)
	if err != nil {
		return "", err
	}
	return string(file), nil
}
