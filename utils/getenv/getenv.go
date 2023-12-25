package getenv

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

// 判断当前环境（开发/测试/gray/product）的函数
func GetEnv() string {
	var hostname string
	var env string
	testPrefix := "test_test"
	grayPrefix := "hzv_webgray"
	productPrefix := "hzv_"
	commonPrefix := "common_"
	env = "dev"

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "localhost"
	}
	//qa测试环境
	position := strings.Index(hostname, testPrefix)
	if position != -1 {
		env = "qa"
	}
	//gray环境
	position = strings.Index(hostname, grayPrefix)
	if position != -1 {
		env = "gray"
	}
	//production环境
	position = strings.Index(hostname, productPrefix)
	if position != -1 {
		env = "production"
	}
	position = strings.Index(hostname, commonPrefix)
	if position != -1 {
		env = "production"
	}
	return env
}

// 获取根目录
// 仅仅针对git仓库的项目
func GetRootPath() string {
	cmdOut, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		fmt.Printf(`Error on getting the go-kit base path: %s - %s`, err.Error(), string(cmdOut))
		os.Exit(2)
	}
	return strings.TrimSpace(string(cmdOut))
}

func GetRootDir(file string) string {
	if filepath.IsAbs(file) {
		return path.Dir(file)
	}
	dir := path.Dir(file)
	dirArr := strings.Split(dir, "/")
	result := ""
	for _, s := range dirArr {
		if s != ".." {
			continue
		}
		result += s + "/"
	}
	absPath, err := filepath.Abs(result)
	if err != nil {
		log.Fatalf("get root dir error :%s", err)
	}
	return absPath
}

// GetAbsPath 获取绝对路径
func GetAbsPath(rootBase, file string) string {
	var (
		absPath = ""
		err     error
	)
	if filepath.IsAbs(file) {
		absPath = file
	} else {
		absPath, err = filepath.Abs(strings.TrimRight(rootBase, "/") + "/" + strings.TrimLeft(file, "./"))
		if err != nil {
			log.Fatalf("get abs dir error :%s", err)
		}
	}
	return absPath
}
