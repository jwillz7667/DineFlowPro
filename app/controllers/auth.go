package controllers

import (
	"DineFlowPro/app"
	"DineFlowPro/app/models"

	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	*revel.Controller
}

func (c Auth) Login() revel.Result {
	return c.Render()
}

func (c Auth) Authenticate(email, password string) revel.Result {
	var user models.User
	if err := app.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.Flash.Error("Invalid email or password")
		return c.Redirect(Auth.Login)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		c.Flash.Error("Invalid email or password")
		return c.Redirect(Auth.Login)
	}

	c.Session["user_id"] = string(rune(user.ID))
	c.Session["user_role"] = user.Role

	return c.Redirect(App.Index)
}

func (c Auth) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(App.Index)
}

func AuthenticateUser(c *revel.Controller) revel.Result {
	if _, ok := c.Session["user_id"]; !ok {
		c.Flash.Error("Please log in to access this page")
		return c.Redirect(Auth.Login)
	}
	return nil
}

func AuthorizeRole(requiredRole string) func(*revel.Controller) revel.Result {
	return func(c *revel.Controller) revel.Result {
		if userRole, ok := c.Session["user_role"]; !ok || userRole != requiredRole {
			c.Flash.Error("You don't have permission to access this page")
			return c.Redirect(App.Index)
		}
		return nil
	}
}
