package controllers

import (
	"github.com/revel/revel"
	"DineFlowPro/app"
	"DineFlowPro/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	if userId, ok := c.Session["user_id"]; ok {
		var user models.User
		app.DB.First(&user, userId)
		return c.Render(user)
	}
	return c.Render()
}

func (c App) TestDB() revel.Result {
	var result int
	if err := app.DB.Raw("SELECT 1").Scan(&result).Error; err != nil {
		return c.RenderJSON(map[string]string{"status": "error", "message": err.Error()})
	}
	return c.RenderJSON(map[string]string{"status": "success", "message": "Database connection successful"})
}

func (c App) Dashboard() revel.Result {
	var orderCount int64
	app.DB.Model(&models.Order{}).Count(&orderCount)

	var userCount int64
	app.DB.Model(&models.User{}).Count(&userCount)

	return c.Render(orderCount, userCount)
}

// API Endpoints

// GetMenuItems returns all menu items
func (c App) GetMenuItems() revel.Result {
	var menuItems []models.MenuItem
	app.DB.Find(&menuItems)
	return c.RenderJSON(menuItems)
}

// GetMenuItem returns a specific menu item by ID
func (c App) GetMenuItem(id int) revel.Result {
	var menuItem models.MenuItem
	if err := app.DB.First(&menuItem, id).Error; err != nil {
		return c.RenderJSON(map[string]string{"error": "Menu item not found"})
	}
	return c.RenderJSON(menuItem)
}

// CreateMenuItem creates a new menu item
func (c App) CreateMenuItem() revel.Result {
	var menuItem models.MenuItem
	c.Params.BindJSON(&menuItem)
	if err := app.DB.Create(&menuItem).Error; err != nil {
		return c.RenderJSON(map[string]string{"error": "Failed to create menu item"})
	}
	return c.RenderJSON(menuItem)
}

// UpdateMenuItem updates an existing menu item
func (c App) UpdateMenuItem(id int) revel.Result {
	var menuItem models.MenuItem
	if err := app.DB.First(&menuItem, id).Error; err != nil {
		return c.RenderJSON(map[string]string{"error": "Menu item not found"})
	}
	c.Params.BindJSON(&menuItem)
	app.DB.Save(&menuItem)
	return c.RenderJSON(menuItem)
}

// DeleteMenuItem deletes a menu item
func (c App) DeleteMenuItem(id int) revel.Result {
	var menuItem models.MenuItem
	if err := app.DB.Delete(&menuItem, id).Error; err != nil {
		return c.RenderJSON(map[string]string{"error": "Failed to delete menu item"})
	}
	return c.RenderJSON(map[string]string{"status": "Menu item deleted"})
}

// GetOrders returns all orders
func (c App) GetOrders() revel.Result {
	var orders []models.Order
	app.DB.Preload("Items").Find(&orders)
	return c.RenderJSON(orders)
}

// GetOrder returns a specific order by ID
func (c App) GetOrder(id int) revel.Result {
	var order models.Order
	if err := app.DB.Preload("Items").First(&order, id).Error; err != nil {
		return c.RenderJSON(map[string]string{"error": "Order not found"})
	}
	return c.RenderJSON(order)
}

// CreateOrder creates a new order
func (c App) CreateOrder() revel.Result {
	var order models.Order
	c.Params.BindJSON(&order)
	if err := app.DB.Create(&order).Error; err != nil {
		return c.RenderJSON(map[string]string{"error": "Failed to create order"})
	}
	return c.RenderJSON(order)
}

// UpdateOrder updates an existing order
func (c App) UpdateOrder(id int) revel.Result {
	var order models.Order
	if err := app.DB.First(&order, id).Error; err != nil {
		return c.RenderJSON(map[string]string{"error": "Order not found"})
	}
	c.Params.BindJSON(&order)
	app.DB.Save(&order)
	return c.RenderJSON(order)
}

// DeleteOrder deletes an order
func (c App) DeleteOrder(id int) revel.Result {
	var order models.Order
	if err := app.DB.Delete(&order, id).Error; err != nil {
		return c.RenderJSON(map[string]string{"error": "Failed to delete order"})
	}
	return c.RenderJSON(map[string]string{"status": "Order deleted"})
}

// GetUsers returns all users
func (c App) GetUsers() revel.Result {
	var users []models.User
	app.DB.Find(&users)
	return c.RenderJSON(users)
}

// GetUser returns a specific user by ID
func (c App) GetUser(id int) revel.Result {
	var user models.User
	if err := app.DB.First(&user, id).Error; err != nil {
		return c.RenderJSON(map[string]string{"error": "User not found"})
	}
	return c.RenderJSON(user)
}

// CreateUser creates a new user
func (c App) CreateUser() revel.Result {
	var user models.User
	c.Params.BindJSON(&user)
	if err := app.DB.Create(&user).Error; err != nil {
		return c.RenderJSON(map[string]string{"error": "Failed to create user"})
	}
	return c.RenderJSON(user)
}

// UpdateUser updates an existing user
func (c App) UpdateUser(id int) revel.Result {
	var user models.User
	if err := app.DB.First(&user, id).Error; err != nil {
		return c.RenderJSON(map[string]string{"error": "User not found"})
	}
	c.Params.BindJSON(&user)
	app.DB.Save(&user)
	return c.RenderJSON(user)
}

// DeleteUser deletes a user
func (c App) DeleteUser(id int) revel.Result {
	var user models.User
	if err := app.DB.Delete(&user, id).Error; err != nil {
		return c.RenderJSON(map[string]string{"error": "Failed to delete user"})
	}
	return c.RenderJSON(map[string]string{"status": "User deleted"})
}
