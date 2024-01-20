package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	config "studioj/boilerplate_go/configs"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Auth는 JWT 기반의 인증 미들웨어 함수를 반환합니다.
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// HTTP 요청에서 사용자 ID 추출
		sub, err := UserID(c.Request)
		if err != nil {
			// 인증 에러가 발생하면 401 Unauthorized 응답을 보내고 요청 처리 중단
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// 토큰과 사용자 ID를 출력 (디버깅용)
		fmt.Println("token", extract(c.Request))
		fmt.Println("sub", *sub)

		// 컨텍스트에 토큰과 사용자 ID 저장
		c.Set("token", extract(c.Request))
		c.Set("sub", *sub)
		c.Next() // 다음 미들웨어나 핸들러로 이동
	}
}

// extract는 HTTP 요청 헤더에서 JWT 토큰을 추출합니다.
func extract(r *http.Request) string {
	authorization := r.Header.Get("Authorization")
	strArr := strings.Split(authorization, " ")
	if len(strArr) == 2 {
		return strArr[1] // "Bearer [토큰]" 형식에서 토큰 부분 반환
	}
	return "" // 토큰이 없는 경우 빈 문자열 반환
}

// verify는 JWT 토큰의 유효성을 검증합니다.
func verify(r *http.Request) (*jwt.Token, error) {
	tokenString := extract(r) // 요청에서 토큰 추출
	jwtToken, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		// 토큰의 서명 방식이 HMAC이 아닌 경우 에러 반환
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}
		return []byte(config.SECRET_KEY), nil // 비밀키로 서명 검증
	})

	return jwtToken, err
}

// UserID는 HTTP 요청에서 JWT 토큰을 검증하여 사용자 ID ('sub' 클레임)를 반환합니다.
func UserID(r *http.Request) (*string, error) {
	jwtToken, err := verify(r) // 토큰 검증
	if err != nil {
		return nil, err // 토큰 검증 실패 시 에러 반환
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return nil, errors.New("refresh token is invalid") // 유효하지 않은 토큰 처리
	}

	sub := claims["sub"].(string) // 'sub' 클레임 추출

	return &sub, nil
}