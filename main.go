package main

import (
	"flag"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/simonfuhrer/gangdert/conf"
	"github.com/simonfuhrer/gangdert/logger"
	"github.com/simonfuhrer/gangdert/model"
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Set("v", "3")
	flag.Set("mode", "release")
	flag.Parse()
	conf.Load()

}

func ceateQRCode(c *gin.Context) {
	name := c.Params.ByName("code")
	qrobj := model.CreateQR(name)
	c.Data(200, "image/png", qrobj.Bytes())
}

func main() {
	runtime.GOMAXPROCS(conf.GangDert.MaxProcs)
	defer glog.Flush()
	glog.Infof(conf.GangDert.Listen)

	engine := gin.New()
	gin.SetMode(conf.GangDert.RuntimeMode)

	engine.Use(gin.Recovery(), logger.Logger())
	engine.GET("/:code", ceateQRCode)

	engine.Run(conf.GangDert.Listen)
}
