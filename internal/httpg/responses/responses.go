package responses

import "github.com/gin-gonic/gin"

func ResponseOkData(ctx *gin.Context, data any) {
	ctx.JSON(200, gin.H{
		"data": data,
		"ok":   true,
	})
}
func ResponseOkCreated(ctx *gin.Context, data any) {
	if data != nil {
		ctx.JSON(201, gin.H{
			"data": data,
			"ok":   true,
		})
		ctx.Abort()
		return
	}
	ctx.JSON(201, gin.H{
		"msg": "created",
		"ok":  true,
	})
}

func ReponseInternalError(ctx *gin.Context, err error) {
	ctx.JSON(500, gin.H{
		"error": err.Error(),
		"ok":    false,
	})
	ctx.Abort()

}

func ResponseClientError(ctx *gin.Context, err error) {
	ctx.JSON(400, gin.H{
		"error": err.Error(),
		"ok":    false,
	})
	ctx.Abort()
}

func ResponseNotFound(ctx *gin.Context) {
	ctx.JSON(404, gin.H{
		"error": "not found",
		"ok":    false,
	})
	ctx.Abort()
}
