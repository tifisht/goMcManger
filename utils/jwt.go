package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var Session map[string]string

func JwtSign(user string) string {
	// 创建一个新的 token 对象
	token := jwt.New(jwt.SigningMethodHS256)

	// 创建声明，Claims 是一个 map[string]interface{}
	claims := token.Claims.(jwt.MapClaims)

	// 设置声明的内容
	claims["authorized"] = true
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 设置过期时间为 24 小时之后

	// 签署 token 并获取完整的编码后的字符串
	tokenString, err := token.SignedString([]byte(JwtKey))
	if err != nil {
		fmt.Println("Error signing token:", err)
		return ""
	}

	return tokenString
}

// func Auth(tokenString string) (string, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// 检查 token 的签名方法
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		// 返回用于验证签名的 key
// 		return []byte(JwtKey), nil
// 	})

// 	if err != nil {
// 		fmt.Println("Error parsing token:", err)
// 		return "", fmt.Errorf("token解析出错: %v", err)
// 	}

// 	// 检查 token 是否有效
// 	if token.Valid {
// 		fmt.Println("Token is valid")
// 		// 获取声明的内容
// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok {
// 			fmt.Println("Couldn't parse claims")
// 			return "", fmt.Errorf("无法解析token声明")
// 		}
// 		user, ok := claims["user"].(string)
// 		if !ok {
// 			fmt.Println("User claim not found or not a string")
// 			return "", fmt.Errorf("token中缺少用户声明或者声明类型不正确")
// 		}
// 		return user, nil
// 	}

// 	if ve, ok := err.(*jwt.ValidationError); ok {
// 		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
// 			fmt.Println("Token is not valid JWT")
// 			return "", fmt.Errorf("这个token非法啦")
// 		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
// 			fmt.Println("Token is expired or not active yet")
// 			return "", fmt.Errorf("这个token已过期或尚未生效")
// 		} else {
// 			fmt.Println("Couldn't handle this token:", err)
// 			return "", fmt.Errorf("无法处理此token: %v", err)
// 		}
// 	}

// 	fmt.Println("Couldn't handle this token:", err)
// 	return "", fmt.Errorf("无法处理此token: %v", err)
// }
