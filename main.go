package main

import (
	"PaystackInterviewTest/controllers"
	_ "PaystackInterviewTest/routers"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/beego/i18n"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.InsertFilter("/*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "content-type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"content-length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
}

func main() {
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.AddFuncMap("ExtractReadableDateTime", ExtractReadableDateTime)
	beego.AddFuncMap("GetTestSecretKey", GetTestSecretKey)
	beego.AddFuncMap("GetTestPublicKey", GetTestPublicKey)
	beego.Run(":" + os.Getenv("PORT"))
}

func GetTotalCount(rawData []interface{}) int {
	return len(rawData)
}

func ExtractReadableDateTime(rawDate string) string {
	t, _ := time.Parse(time.RFC3339, rawDate)
	return t.Format(time.RFC850)
}

func GetTestSecretKey(value bool) string {
	// Key can be encrypted and saved in config
	return beego.AppConfig.String("PaystackTestSecretKey")
}

func GetTestPublicKey(value bool) string {
	// Key can be encrypted and saved in config
	return beego.AppConfig.String("PaystackTestPublicKey")
}
