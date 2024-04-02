package controller

import (
	"backend/pkg"

	"github.com/gin-gonic/gin"
)

func InitLedger(c *gin.Context) {
	res, err := pkg.InitLedger()
	if err != nil {
		c.JSON(200, gin.H{
			"message": "initLedger failed" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "initLedger success",
		"txid":    res,
	})
}

func CreateCreditFile(c *gin.Context) {
	res, err := pkg.CreateCreditFile(c.PostForm("id"), c.PostForm("owner"), c.PostForm("academy"), c.PostForm("disciplinarySituation"), c.PostForm("gradePoint"), c.PostForm("awardSituation"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "createCreditFile failed" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "createCreditFile success",
		"txid":    res,
	})
}

func UpdateCreditFile(c *gin.Context) {
	res, err := pkg.UpdateCreditFile(c.PostForm("id"), c.PostForm("owner"), c.PostForm("academy"), c.PostForm("disciplinarySituation"), c.PostForm("gradePoint"), c.PostForm("awardSituation"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "updateCreditFile failed" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "updateCreditFile success",
		"txid":    res,
	})
}

func DeleteCreditFile(c *gin.Context) {
	res, err := pkg.DeleteCreditFile(c.PostForm("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "deleteCreditFile failed" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "deleteCreditFile success",
		"txid":    res,
	})
}

func GetCreditFile(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetCreditFile", c.PostForm("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "getCreditFile failed" + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "getCreditFile success",
		"data":    res,
	})
}
