package controllers

import (
	"rhiona-api/config"
	"rhiona-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create
func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&customer)
	c.JSON(http.StatusOK, customer)
}

// Read All
func GetCustomers(c *gin.Context) {
	var customers []models.Customer
	config.DB.Find(&customers)
	c.JSON(http.StatusOK, customers)
}

// Read by ID
func GetCustomerByID(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	c.JSON(http.StatusOK, customer)
}

// Update
func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	var input models.Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.Name = input.Name
	customer.Phone = input.Phone
	customer.Address = input.Address

	config.DB.Save(&customer)
	c.JSON(http.StatusOK, customer)
}

// Delete
func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := config.DB.Delete(&customer, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
}
