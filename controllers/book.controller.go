package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/simple-bookStore-server/config"
	"github.com/simple-bookStore-server/models"
)

func CreateBook(c echo.Context) error {
	b := new(models.Book)
	db := config.DB()

	//binding data
	if err := c.Bind(b); err != nil {
		//we are doing this because the Bind function returns a interface by default so whenever we use bind we have to return the error in the format of these
		data := map[string]interface{}{
			"message": err.Error(),
		} //here we are actually sending the error if incountered
		return c.JSON(http.StatusInternalServerError, data)
	}

	//now here we binded the data to the model
	if err := db.Create(&b).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	//Catching the response here
	res := map[string]interface{}{
		"data": b,
	}
	return c.JSON(http.StatusOK, res)
	//returned the status ok after processing of the data
}

func UpdateBook(c echo.Context) error {
	id := c.Param("id")
	b := new(models.Book)
	db := config.DB()

	//binding data
	if err := c.Bind(&b); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	exBook := new(models.Book)

	if err := db.First(&exBook, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusNotFound, data)
	}

	exBook.Name = b.Name
	exBook.Description = b.Description
	if err := db.Save(&exBook).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)

	}

	//actual response
	res := map[string]interface{}{
		"data": exBook,
	}
	return c.JSON(http.StatusOK, res)

}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	book := new(models.Book)

	err := db.Delete(&book, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "a book has been deleted",
	}

	return c.JSON(http.StatusOK, response)
}

func GetBook(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	var books []*models.Book

	if res := db.Find(&books, id); res.Error != nil {
		data := map[string]interface{}{
			"message": res.Error.Error(),
		}

		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"data": books[0],
	}

	return c.JSON(http.StatusOK, response)
}
