package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		tokenStr, err := ctx.Cookie("session_token")
        if err != nil {
            if err == http.ErrNoCookie {
                // return unauthorized ketika token kosong
				if ctx.Request.Header.Get("Content-Type") == "application/json" {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse("unauthorized"))
					 return
				}

				ctx.Redirect(303, "/login")
					return
            }
            // return bad request ketika field token tidak ada
            ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewSuccessResponse(err.Error()))
            return
        }

        // Deklarasi variable claims yang akan kita isi dengan data hasil parsing JWT
        claims := &model.Claims{}

        //parse JWT token ke dalam claims
        tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return model.JwtKey, nil
        })
        if err != nil {
            if err == jwt.ErrSignatureInvalid {
                // return unauthorized ketika ada kesalahan ketika parsing token
                ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse(err.Error()))
                return
            }
            // return bad request ketika field token tidak ada
            ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewSuccessResponse(err.Error()))
            return
        }

        //return unauthorized ketika token sudah tidak valid (biasanya karena token expired)
        if !tkn.Valid {
         ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse("token invalid"))
            return
        }
		ctx.Set("email", claims.Email)
		ctx.Next()
	})
}
