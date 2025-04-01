package gin

import "github.com/gin-gonic/gin"

var R *gin.Engine

func init() {
	R = gin.Default()
	R.GET(LoginEndpoint, LoginHandler())
	R.GET(CallbackEndpoint, CallbackHandler())
}
