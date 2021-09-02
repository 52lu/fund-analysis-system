/**
 * @Description JWT中的声明
 **/
package request

import "github.com/golang-jwt/jwt"

// 设置声明
type UserClaims struct {
	*jwt.StandardClaims
	Uid uint // 不建议放入会被修改的字段，推荐放入用户ID。
}