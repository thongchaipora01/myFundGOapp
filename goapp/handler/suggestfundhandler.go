package handler

import (
	"github.com/thonp571/myfundgoapp/goapp/model"
	serv "github.com/thonp571/myfundgoapp/goapp/services"

	"github.com/gin-gonic/gin"
)

func suggestFund() gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		//Bind Json body request
		requestfe := model.RequestFE{}
		if err := ctx.ShouldBindJSON(&requestfe); err != nil {
			ctx.JSON(400, gin.H{
				"errorCode": "40000",
				"Message":   "Bind JSON Request fail",
			})
			return
		}
		fund := serv.GetSuggestFund(requestfe.TimeRange)

		ctx.JSON(200, fund)

	}
	return gin.HandlerFunc(fn)
}

// `` << this special single quote
