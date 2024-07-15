package main

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

var token = "2zeCr3Tco8J8q8s30GjfphpZwBHMFiGz"

func main() {
	router := gin.Default()

	router.GET("/verify", func(ctx *gin.Context) {
		signature := ctx.Query("signature")
		timestamp := ctx.Query("timestamp")
		nonce := ctx.Query("nonce")
		echostr := ctx.Query("echostr")
		tmpArr := []string{token, timestamp, nonce}
		sort.Strings(tmpArr)
		tmpStr := strings.Join(tmpArr, "")
		hasher := sha1.New()
		hasher.Write([]byte(tmpStr))
		hexStr := hex.EncodeToString(hasher.Sum(nil))
		if hexStr == signature {
			log.Println("[/]:  the signature is valid")
			ctx.String(http.StatusOK, echostr)
			return
		}
		ctx.String(http.StatusOK, "Bad signature")
		log.Println("[/]: Bad signature")
	})

	log.Fatal(router.Run(":8080"))
}
