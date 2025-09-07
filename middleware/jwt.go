//package middleware
//
//import (
//	"My-Blog/utils"
//	"My-Blog/utils/errmsg"
//	"github.com/dgrijalva/jwt-go"
//	"github.com/gin-gonic/gin"
//	"strings"
//	"time"
//)
//
//var JwtKey = []byte(utils.JwtKey)
//
//type MyClaims struct {
//	Username string `json:"username"`
//	jwt.StandardClaims
//}
//
//var code int
//
//// 生成token
//func SetToken(username string) (string, int) {
//	expireTime := time.Now().Add(time.Hour * 10)
//	SetClaims := MyClaims{
//		Username: username,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: expireTime.Unix(),
//			Issuer:    "My blog",
//		},
//	}
//	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
//	token, err := reqClaim.SignedString(JwtKey)
//	if err != nil {
//		return "", errmsg.ERROR
//	}
//	return token, errmsg.SUCCSE
//}
//
//// 验证token
//func CheckToken(token string) (*MyClaims, int) {
//	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return JwtKey, nil
//	})
//	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
//		return key, errmsg.SUCCSE
//	} else {
//		return nil, errmsg.ERROR
//	}
//
//}
//
//// jwt中间件
//func JwtToken() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		tokenHeader := c.Request.Header.Get("Authorization")
//		if tokenHeader == "" {
//			c.JSON(200, gin.H{
//				"code":    code,
//				"message": errmsg.GetErrMsg(code),
//			})
//			code = errmsg.ERROR_TOKEN_EXIST
//			c.Abort()
//			return
//		}
//		checkToken := strings.SplitN(tokenHeader, " ", 2)
//		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
//			code = errmsg.ERROR_TOKEN_TYPE_WRONG
//			c.JSON(200, gin.H{
//				"code":    code,
//				"message": errmsg.GetErrMsg(code),
//			})
//			c.Abort()
//			return
//		}
//		key, tCode := CheckToken(checkToken[1])
//		if tCode != errmsg.ERROR {
//			code = errmsg.ERROR_TOKEN_WRONG
//			c.JSON(200, gin.H{
//				"code":    code,
//				"message": errmsg.GetErrMsg(code),
//			})
//			c.Abort()
//			return
//		}
//		if time.Now().Unix() > key.ExpiresAt {
//			code = errmsg.ERROR_TOKEN_RUNTIME
//			c.JSON(200, gin.H{
//				"code":    code,
//				"message": errmsg.GetErrMsg(code),
//			})
//			c.Abort()
//			return
//		}
//		c.Set("username", key.Username)
//		c.Next()
//	}
//}

package middleware

import (
	"My-Blog/utils"
	"My-Blog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(time.Hour * 10) // 有效期10小时
	SetClaims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "My blog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCSE
}

// 验证token
func CheckToken(token string) (*MyClaims, int) {
	setToken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, errmsg.ERROR // 解析失败返回错误
	}
	// 验证token有效性
	if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
		return key, errmsg.SUCCSE // 验证成功
	}
	return nil, errmsg.ERROR // 验证失败
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 每次请求都重置错误码，避免全局变量污染
		code := 0

		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST // 未提供Token
			c.JSON(200, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 分割Bearer和Token
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG // Token格式错误
			c.JSON(200, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 验证Token
		key, tCode := CheckToken(checkToken[1])
		if tCode == errmsg.ERROR { // 这里修正为：如果验证失败
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(200, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 检查Token是否过期（先确认key不为空）
		if key != nil && time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME // Token已过期
			c.JSON(200, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		// 将用户名存入上下文，供后续接口使用
		c.Set("username", key.Username)
		c.Next()
	}
}
