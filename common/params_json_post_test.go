package common_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"loveraw.club/hsipeng/ginapi/common"
	"loveraw.club/hsipeng/ginapi/test"
)

type exampleJson struct {
	User string `form:"user" json:"user" binding:"required"`
	Age  int    `form:"age" json:"age" binding:"required"`
}

/// 一共通过 application/json 取值了三次 然后再次 bindJSON
func bindingPostJson(c *gin.Context) {
	// first
	user := common.GetString(c, "user")
	// second
	ageAsString := common.GetString(c, "age")
	// third
	ageAsInt := common.GetInt64(c, "age")
	var shouldBind exampleJson
	// 再次bind json
	err := c.BindJSON(&shouldBind)
	if err != nil {
		fmt.Println(c, http.StatusBadRequest, http.StatusBadRequest, "ShouldBind failed, err: %+v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user":          user,
			"age_as_string": ageAsString,
			"age_as_int":    ageAsInt,
			"should_bind":   shouldBind,
		},
		"code": 0,
	})
}

func TestPOSTJSONBinding(t *testing.T) {
	engine := test.NewEngine()

	// router
	engine.Handle("POST", "/testjson", bindingPostJson)

	// request
	jsonStr := `{"age":10,"user":"foo"}`
	req := test.NewPostRequest("/testjson", "application/json", strings.NewReader(string(jsonStr)))
	w := test.Do(engine, req)
	// response
	assert.Equal(t, http.StatusOK, w.Code)

	respBody, err := ioutil.ReadAll(w.Result().Body)
	assert.Nil(t, err)

	var resp map[string]interface{}
	err = jsoniter.Unmarshal(respBody, &resp)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	data, ok := resp["data"].(map[string]interface{})
	assert.True(t, ok)
	assert.NotNil(t, data)
	assert.Equal(t, "foo", data["user"])
	assert.Equal(t, "10", data["age_as_string"])
	assert.Equal(t, float64(10), data["age_as_int"])
	assert.Equal(t, map[string]interface{}{"user": "foo", "age": float64(10)}, data["should_bind"])
}
