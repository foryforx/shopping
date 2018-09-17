package middleware

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	loginRepo "github.com/karuppaiah/shopping/login/repository"
	loginUcase "github.com/karuppaiah/shopping/login/usecase"
	"github.com/karuppaiah/shopping/model"
)

var identityKey = "id"

//// Basic Setup and routers config
type goMiddleware struct {
	// another stuff , may be needed by middleware
}

func (m *goMiddleware) CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
func InitMiddleware() *goMiddleware {
	return &goMiddleware{}
}

func (m *goMiddleware) AuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"), // Change the secret key before going to staging/production
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// Called during login after Authenticator
			fmt.Println("Payloadfunc")
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}

			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			// called for refresh token here and forwarded to authotizator
			// Called when someone asks for an auth API with token and then moves to authorizator
			claims := jwt.ExtractClaims(c)
			fmt.Println("IdentityHandler")
			return &model.User{
				UserName: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			// Called when we ask for new token with username and pwd
			var loginVals model.Login
			fmt.Println("Authenticator")
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password
			timeoutContext := time.Duration(10) * time.Second
			db := model.GetDBInstance()
			loR := loginRepo.NewERepository(db.SDB)
			loU := loginUcase.NewEUsecase(loR, timeoutContext)
			mLogin, err := loU.Fetch(c, userID)
			if len(mLogin) != 1 || err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			if mLogin[0].Username == userID && mLogin[0].Password == password {
				return &model.User{
					UserName:  userID,
					LastName:  "KAL",
					FirstName: "KAL",
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
			//Send the usename to User Business logic and verify if its available in DB
			// call Business logic here
			// If Business logic return true
			//		Allow create User object and return
			// Else
			//      return nil
			//Authenticate only for specific users as of now. User management later to be done
			// old code start: hardcoded JWT
			// if (userID == "admin" && password == "admin") || (userID == "kal" && password == "kal") || (userID == "james" && password == "james") {
			// 	return &model.User{
			// 		UserName:  userID,
			// 		LastName:  "KAL",
			// 		FirstName: "KAL",
			// 	}, nil
			// }

			// return nil, jwt.ErrFailedAuthentication
			// old code end:hardcoded JWT
		},
		//If username is admin/kal/james allow them to proceed
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// Called when someone asks for an auth API with token| called from IdentityHandler
			fmt.Println("Authorizator")
			// Call User Business layer to verify if username exists
			// If user name exists
			//  return true
			// Else
			// return false
			v, ok := data.(*model.User)
			if !ok {
				return false
			}
			db := model.GetDBInstance()
			loR := loginRepo.NewERepository(db.SDB)
			timeoutContext := time.Duration(10) * time.Second
			loU := loginUcase.NewEUsecase(loR, timeoutContext)
			fmt.Println("Reqested user:" + v.UserName)
			mLogin, err := loU.Fetch(c, v.UserName)
			if len(mLogin) != 1 || err != nil {
				return false
			}
			if v.UserName == mLogin[0].Username {
				fmt.Println("username mismatch")
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			// Called when we have invalid token
			fmt.Println("Unauthorized")
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}

// the jwt middleware
