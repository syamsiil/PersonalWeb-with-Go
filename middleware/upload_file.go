package middleware

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("input-image") // get image by name
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		src, err := file.Open() //open file from input
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		defer src.Close() 

		tempFile, err := ioutil.TempFile("uploads", "aaa-*.png") //make extension the image became .png and send or save to uploads folder 

		if err !=nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		
		defer tempFile.Close()

		io.Copy(tempFile, src) // copy the file from src to tempFile 

		data := tempFile.Name() //path/to/uploads/image-name.png
		filename := data[8:]

		c.Set("dataFile", filename) 

		return next(c)
	}
}