package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"runtime"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage（Unix/Linux）:  Go_Data_Summary_Tool <filepath> \n Usage（Windows）:  Go_Data_Summary_Tool.exe <filepath>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	fileName := filePath[strings.LastIndex(filePath, "/")+1:] // 适用于Unix/Linux系统
	if runtime.GOOS == "windows" {
		fileName = filePath[strings.LastIndex(filePath, "\\")+1:] // 适用于Windows系统
	}

	md5Hash := ComputeHash(filePath, md5.New())
	sha1Hash := ComputeHash(filePath, sha1.New())
	sha256Hash := ComputeHash(filePath, sha256.New())
	sha512Hash := ComputeHash(filePath, sha512.New())

	fmt.Printf("文件名: %s\n", fileName)
	fmt.Printf("MD5: %s\n", md5Hash)
	fmt.Printf("SHA-1: %s\n", sha1Hash)
	fmt.Printf("SHA-256: %s\n", sha256Hash)
	fmt.Printf("SHA-512: %s\n", sha512Hash)
}

func ComputeHash(filePath string, hashFunction io.Writer) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	hashFunction.Write(make([]byte, 0)) // 清空hashFunction（如果必要）

	if _, err := io.Copy(hashFunction, file); err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}

	sum := hashFunction.(hash.Hash).Sum(nil)
	return hex.EncodeToString(sum)
}

// 注意：这里缺少strings和runtime的导入
// 需要在文件顶部添加：
// import (
//     "runtime"
//     "strings"
//     "hash" // 还需要hash接口
// )
