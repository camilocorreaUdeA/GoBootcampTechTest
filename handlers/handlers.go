package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/camilocorreaUdeA/GoBootcampTechTest/models"

	"github.com/gin-gonic/gin"
)

const ApiUrl = "https://reqres.in/api/users?page=2"

func GetCustomData(c *gin.Context) (interface{}, error) {

	req, err := http.NewRequest(http.MethodGet, ApiUrl, nil)
	if err != nil {
		fmt.Println("Error building request object")
		return models.ApiData{}, err
	}

	client := &http.Client{}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request")
		return models.ApiData{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error response from Api")
		return models.ApiData{}, err
	}
	defer resp.Body.Close()

	apiData := models.ApiData{}
	fmt.Println(string(body))

	err = json.Unmarshal(body, &apiData)
	if err != nil {
		fmt.Println("Error unmarshaling data")
		return models.ApiData{}, fmt.Errorf("Error unmarshaling data: %s", err.Error())
	}

	return apiData, nil
}

func SayHello(c *gin.Context) (interface{}, error) {
	return "Hello World!", nil
}

func RequestWrapper(f func(c *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := f(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal Error %s", err.Error())})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "OK", "response": resp})
	}
}
