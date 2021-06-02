package utils

import (
	"github.com/gin-gonic/gin"
	"strings"
)


func GetMapFromContext(context *gin.Context) (map[string]interface{},error) {
	result := make(map[string]interface{})
	contentType := strings.ToLower(context.Request.Header.Get("content-type"))
	if strings.Contains(contentType,"multipart/form-data"){
		err := context.Request.ParseMultipartForm(128)
		if err != nil {
			return map[string]interface{}{},err
		}
		form := context.Request.Form
		for k,v :=range form{
			if len(v) == 1 {
				result[k] = v[0]
			}else{
				result[k] = strings.Join(v,";")
			}
		}
	}else if  strings.Contains(contentType,"x-www-form-urlencoded") {
		err := context.Request.ParseForm()
		if err != nil {
			return map[string]interface{}{},err
		}
		form := context.Request.Form
		for k,v :=range form{
			if len(v) == 1 {
				result[k] = v[0]
			}else{
				result[k] = strings.Join(v,";")
			}
		}
	}else{
		_ = context.BindJSON(&result)
	}
	return result,nil
}