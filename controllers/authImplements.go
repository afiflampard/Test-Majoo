package controllers

import (
	"fmt"
	"majoo/helpers"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authClaims struct {
	UserId int64 `json:"user_id"`
	User   bool  `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

type AccessDetails struct {
	UserId int64
}

func JWTAuthService() JWTServices {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Bikash",
	}
}

var db *gorm.DB

func InitiateDB(dataDB *gorm.DB) {
	db = dataDB
}

func GetDB() *gorm.DB {
	return db
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	fmt.Println(secret)
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(c *gin.Context, ID int, isUser bool) (string, error) {
	td := &authClaims{
		int64(ID),
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	claims := jwt.MapClaims{}
	claims["user_id"] = td.UserId
	claims["exp"] = td.ExpiresAt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		helpers.Error(c, 404, "Login dulu")
	}
	return t, nil

}
func (service *jwtServices) ExtractToken(c *gin.Context) string {
	bearToken := c.GetHeader("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}

func (service *jwtServices) ExtractTokenMetadata(c *gin.Context) (*AccessDetails, error) {
	stringToken := service.ExtractToken(c)
	token, err := service.ValidateToken(stringToken)
	fmt.Println("Error", err)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			UserId: userID,
		}, nil
	}
	return nil, err
}
