package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"interview/pkg/utils"
	"net/http"
	"time"
)

const (
	addItemsFormSource = "static/add_item_form.html"
	sessionIDKey       = "ice_session_id"
)

type TaxController struct {
	CalculatorHandler
}

func NewTaxController(calculatorHandler CalculatorHandler) *TaxController {
	return &TaxController{CalculatorHandler: calculatorHandler}
}

func (t *TaxController) ShowAddItemForm(c *gin.Context) {
	var sessionID string
	cookie, err := c.Request.Cookie(sessionIDKey)
	if cookie != nil {
		sessionID = cookie.Value
	}
	if errors.Is(err, http.ErrNoCookie) {
		sessionID = time.Now().String()
		c.SetCookie(sessionIDKey, sessionID, 3600, "/", "localhost", false, true)
	}

	data := map[string]interface{}{
		DataErrorKey: c.Query("error"),
	}

	cartItems, err := t.CalculatorHandler.GetCartData(c, sessionID)
	if cartItems != nil && err == nil {
		data[DataCartItemsKey] = cartItems
	}

	html, err := utils.RenderTemplate(data, addItemsFormSource)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	c.Header("Content-Type", "text/html")
	c.String(200, html)
}

func (t *TaxController) AddItem(c *gin.Context) {
	cookie, err := c.Request.Cookie(sessionIDKey)

	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		redirect(c, 302, "/", nil)
		return
	}
	// read addItem form
	form := CartItemForm{}
	if err := utils.GetRequestForm(c, &form); err != nil {
		redirect(c, 302, "/", err)
		return
	}
	// add the item
	code, err := t.CalculatorHandler.AddItem(c, form, cookie.Value)
	redirect(c, code, "/", err)
}

func (t *TaxController) DeleteCartItem(c *gin.Context) {
	cookie, err := c.Request.Cookie(sessionIDKey)

	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		redirect(c, 302, "/", nil)
		return
	}

	cartItemIDString := c.Query("cart_item_id")
	if cartItemIDString == "" {
		redirect(c, 302, "/", nil)
		return
	}

	code, err := t.CalculatorHandler.DeleteItemFromCart(c, cartItemIDString, cookie.Value)
	redirect(c, code, "/", err)
}

func redirect(c *gin.Context, code int, dest string, err error) {
	if err != nil {
		dest = fmt.Sprintf("%s?error=%s", dest, err.Error())
	}
	c.Redirect(code, dest)
}
