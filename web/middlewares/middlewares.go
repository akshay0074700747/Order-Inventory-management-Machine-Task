package middlewares

import (
	"net/http"

	jwttoken "github.com/akshay0074700747/order-inventory_management/web/jwt_token"
	"github.com/gofiber/fiber"
)

type Middleware struct {
	secret string
}

func NewMiddleware(secret string) *Middleware {
	return &Middleware{
		secret: secret,
	}
}

// global middleware set to all the routes
func (middleware *Middleware) GlobalMiddleware(c *fiber.Ctx) {

	//fetching cookies
	cookie := c.Cookies("Token")
	if cookie == "" {
		c.Status(http.StatusUnauthorized).Write("Please Login")
		return
	}

	values, err := jwttoken.ValidateToken(cookie, []byte(middleware.secret))
	if err != nil {
		c.Status(http.StatusUnauthorized).Write(err.Error())
		return
	}

	c.Locals("values", values)

	c.Next()
}

// middleware to validate only admins access specific routes
func (middleware *Middleware) AdminMiddleware(c *fiber.Ctx) {

	//getting passed values from the request context
	value := c.Locals("values")
	if value == nil {
		c.Status(http.StatusInternalServerError).Write("user credentials are not available")
		return
	}

	valueMap, _ := value.(map[string]interface{})

	isAdmin := valueMap["isAdmin"].(bool)

	if !isAdmin {
		c.Status(http.StatusUnauthorized).Write("this route is only accessible to admins")
		return
	}

	c.Next()
}
