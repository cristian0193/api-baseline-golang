package middleware

import (
	"net/http"
	"os"
	"ws-baseline-golang-v2/domain/dto"
	"ws-baseline-golang-v2/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	var lenguage string
	var responseDto = dto.Response{}

	return func(c *gin.Context) {
		lenguage = c.Request.Header.Get(os.Getenv("LENGUAGE_HEADER"))
		token := utils.ExtractToken(c.Request)
		descripcion, code := utils.VerifyToken(token, lenguage)
		if code != http.StatusOK {
			responseDto.Status = code
			responseDto.Description = utils.StatusText(code)
			responseDto.Message = descripcion
			c.JSON(code, responseDto)
			c.AbortWithStatus(code)
			return
		}
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {

	var allowOrigin = "*"
	var allowCredential = "true"
	var allowHeaders = "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, " +
		"Cache-Control, X-Requested-With, X-Lang, X-Tenant"
	var exposeHeaders = ""
	var allowMethods = "POST, OPTIONS, GET, PUT, PATCH, DELETE"

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", allowCredential)
		c.Writer.Header().Set("Access-Control-Allow-Headers", allowHeaders)
		c.Writer.Header().Set("Access-Control-Expose-Headers", exposeHeaders)
		c.Writer.Header().Set("Access-Control-Allow-Methods", allowMethods)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
