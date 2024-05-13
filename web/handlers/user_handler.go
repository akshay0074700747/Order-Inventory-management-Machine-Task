package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/akshay0074700747/order-inventory_management/entities"
	"github.com/akshay0074700747/order-inventory_management/helpers"
	usecaseports "github.com/akshay0074700747/order-inventory_management/usecase/usecase_ports"
	jwttoken "github.com/akshay0074700747/order-inventory_management/web/jwt_token"
	"github.com/gofiber/fiber"
)

// user handler consists of all the methods for user related api
type UserHandler struct {
	UserUsecase usecaseports.UserUsecasePort
	secret      string
}

func NewUserHandler(userUsecase usecaseports.UserUsecasePort, secret string) *UserHandler {

	return &UserHandler{
		UserUsecase: userUsecase,
		secret:      secret,
	}
}

func (userHandler *UserHandler) SignUp(c *fiber.Ctx) {

	//checking wheather the user is already logged in
	cookie := c.Cookies("Token")
	if cookie != "" {
		c.Status(http.StatusConflict).Write("the user is already logged in")
		return
	}

	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		fmt.Printf("the body cannot be parsed : %s", err.Error())
		c.Status(http.StatusBadRequest).JSON(entities.Response{
			StatusCode: 400,
			Message:    "can't bind the credentials",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	result, err := userHandler.UserUsecase.Signup(user)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "unable to signUp",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	//generating the jwt Token
	jwt, err := jwttoken.GenerateJwt(result.UserID, false, []byte(userHandler.secret))
	if err != nil {
		fmt.Printf("the jwt Token cannot be generated : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "cannot generate Token",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	//setting the jwt in the cookie
	cookieConfig := fiber.Cookie{
		Name:     "Token",
		Value:    jwt,
		Expires:  time.Now().Add(1 * time.Hour),
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HTTPOnly: true,
	}
	c.Cookie(&cookieConfig)

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "user SignedUp successfully",
		Data:       result,
		Errors:     nil,
	})
}

func (userHandler *UserHandler) Login(c *fiber.Ctx) {

	//checking wheather the user is already logged in
	cookie := c.Cookies("Token")
	if cookie != "" {
		c.Status(http.StatusConflict).Write("the user is already logged in")
		return
	}

	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		fmt.Printf("the body cannot be parsed : %s", err.Error())
		c.Status(http.StatusBadRequest).JSON(entities.Response{
			StatusCode: 400,
			Message:    "can't bind the credentials",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	result, err := userHandler.UserUsecase.Login(user)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "unable to Login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	//generating jwt
	jwt, err := jwttoken.GenerateJwt(result.UserID, result.IsAdmin, []byte(userHandler.secret))
	if err != nil {
		fmt.Printf("the jwt Token cannot be generated : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "cannot generate Token",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	//setting jwt to cookie
	cookieConfig := fiber.Cookie{
		Name:     "Token",
		Value:    jwt,
		Expires:  time.Now().Add(1 * time.Hour),
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HTTPOnly: true,
	}
	c.Cookie(&cookieConfig)

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "user Logged In successfully",
		Data:       result,
		Errors:     nil,
	})
}

func (userHandler *UserHandler) GetMostOrderedUsers(c *fiber.Ctx) {

	//getting the pageNo and limit from query for pagination
	pageNostr := c.Query("pageNo")
	limitStr := c.Query("limit")
	var pageNo, limit int
	var err error

	if pageNostr != "" {
		pageNo, err = helpers.StrtoInt(pageNostr)
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(entities.Response{
				StatusCode: 400,
				Message:    "cannot parse query to int",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	if limitStr != "" {
		limit, err = helpers.StrtoInt(limitStr)
		if err != nil {
			c.Status(http.StatusBadRequest).JSON(entities.Response{
				StatusCode: 400,
				Message:    "cannot parse query to int",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	result, err := userHandler.UserUsecase.GetMostOrderedUsers(pageNo, limit)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
		c.Status(http.StatusInternalServerError).JSON(entities.Response{
			StatusCode: 500,
			Message:    "cannot get most ordered users",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(entities.Response{
		StatusCode: 200,
		Message:    "fetched most ordered users successfully",
		Data:       result,
		Errors:     nil,
	})
}
