package token

import (
	"blog/models"
	"blog/util"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var admin_key string = "6eb702a91df711e9b303021be74f3a70"

func CreateToken(admin *models.Admin) (string,error){

	claims  :=  make(jwt.MapClaims)
	claims["username"] =  admin.Username
	claims["password"] = admin.Password
	claims["exp"] = time.Now().Add(time.Hour*1).Unix()    //1小时有效期

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	tokenStr ,err := token.SignedString([]byte(admin_key))

	if err!=nil {
		log.Fatalf("token.SignedString error : %v ",err)
		return "",err
	}

	err = util.SetRedisKey("BLOG-TOKEN",tokenStr)

	if err !=nil{
		log.Fatalf("redis set usertoken err： %v ",err)
		return "", err
	}

	userToken,err := util.GetRedisKey("BLOG-TOKEN")

	if  err!=nil{
		log.Fatalf("get redisKey BLOG-TOKEN ERR: %v  ",err)
		return "",err
	}

	return userToken,nil
}


func AuthToken(tokenStr string) (string,string,error){
	token,err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(admin_key),nil
	})

	if err !=nil{
		log.Fatalf("token parsing error : %v",err)
		return "","",err
	}

	if !token.Valid{
		if ve,ok := err.(*jwt.ValidationError);ok{
			if ve.Errors&jwt.ValidationErrorExpired!=0{
				log.Fatalf("token expired :%v  ",token)
				return "","",err
			}
		}
		return "","",nil
	}

	username := fmt.Sprintf("%s",token.Claims.(jwt.MapClaims)["username"])
	password:= fmt.Sprintf("%s",token.Claims.(jwt.MapClaims)["password"])

	return username,password,nil
}


