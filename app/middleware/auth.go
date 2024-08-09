package middleware

import (
	"github.com/revel/revel"
)

func AuthenticateUser(c *revel.Controller) revel.Result {
	if _, ok := c.Session["user_id"]; !ok {
		c.Flash.Error("Please log in to access this page")
		return c.Redirect(revel.MainRouter.Reverse("Auth.Login", nil).URL)
	}
	return nil
}

func AuthorizeRole(requiredRole string) func(*revel.Controller) revel.Result {
	return func(c *revel.Controller) revel.Result {
		if userRole, ok := c.Session["user_role"]; !ok || userRole != requiredRole {
			c.Flash.Error("You don't have permission to access this page")
			return c.Redirect(revel.MainRouter.Reverse("App.Index", nil).URL)
		}
		return nil
	}
}
