package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// 打开文件
	file, err := os.Open("jwt.txt")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 创建一个扫描器来读取文件
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokenString := scanner.Text()
		tokenString = strings.TrimSpace(tokenString)

		// 如果 tokenString 以 "bearer " 开头，则去掉这个前缀
		if strings.HasPrefix(tokenString, "bearer ") {
			tokenString = strings.TrimPrefix(tokenString, "bearer ")
		}

		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		}

		token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
		if err != nil {
			fmt.Println("解析 JWT 失败:", err)
			continue
		}
		// 输出 token 的 id
		fmt.Println("token 的 id:", token.Claims.(jwt.MapClaims)["id"])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件时出错:", err)
	}
}
