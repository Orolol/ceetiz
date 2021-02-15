package main

import (
	"math"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func initRoutes() {

	r := gin.Default()

	r.Use(LiberalCORS)

	r.NoRoute(GlobalHandler)
	r.GET("/", GlobalHandler)
	r.POST("/estimate", Estimate)

	r.Run(":3010")

}

func GlobalHandler(c *gin.Context) {
}

func Estimate(c *gin.Context) {
	var ent Entreprise
	c.ShouldBindJSON(&ent)
	var isValid bool = false
	var taxRate float64 = 0.0
	var errorMsg = "Data invalid"
	for _, typeE := range AllTypes {
		if typeE.ID == ent.TypeId {
			isValid = true
			taxRate = typeE.TaxRate
			for _, field := range typeE.Mandatory {
				if getField(&ent, field) == "" {
					errorMsg = errorMsg + " " + field + " missing "
					isValid = false
				}
			}
		}
	}

	if isValid == false {
		c.JSON(http.StatusBadRequest, errorMsg)
	} else {
		var tax float64 = ent.Revenue * taxRate
		c.JSON(http.StatusOK, math.Ceil(tax*100)/100)
	}

}

func getField(v *Entreprise, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}

func LiberalCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
			c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
		}
		c.AbortWithStatus(http.StatusOK)
	}
}
