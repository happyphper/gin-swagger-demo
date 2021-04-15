// +build doc

package main

import (
	_ "github.com/happyphper/gin-swagger-demo/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	url := ginSwagger.URL("/swagger/doc.json")
	swagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler, url)
}
