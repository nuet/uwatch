package controllers

import (
	"github.com/astaxie/beego"
	"library/autofill"
	"runtime"
)

type MetricsCallBackController struct {
	beego.Controller
}

func (c *MetricsCallBackController) Get() {
	//获取panic
	defer func() {
		if panic_err := recover(); panic_err != nil {
			var buf []byte = make([]byte, 1024)
			runtimec := runtime.Stack(buf, false)
			beego.Error("控制器错误:", panic_err, string(buf[0:runtimec]))
		}
	}()

	beego.Info("start to get metrics")

	endpoint := c.GetString("endpoint")
	metric := c.GetString("metric")
	status := c.GetString("status")
	step := c.GetString("step")
	priority := c.GetString("priority")
	time := c.GetString("time")
	tplId := c.GetString("tpl_id")
	expId := c.GetString("exp_id")
	straId := c.GetString("stra_id")
	tags := c.GetString("tags")
	leftvalue := c.GetString("leftvalue")
	rightvalue := c.GetString("rightvalue")
	counter := c.GetString("counter")

	beego.Info("endpoint is", endpoint)
	beego.Info("metric is", metric)
	beego.Info("status is", status)
	beego.Info("step is", step)
	beego.Info("priority is", priority)
	beego.Info("time is", time)
	beego.Info("tpl_id is", tplId)
	beego.Info("exp_id is", expId)
	beego.Info("stra_id is", straId)
	beego.Info("tags is", tags)
	beego.Info("leftvalue is", leftvalue)
	beego.Info("rightvalue is", rightvalue)
	beego.Info("counter is", counter)

	fv := autofill.FalconInput{}
	fv.Endpoint = endpoint
	fv.Metric = metric
	fv.Tags = tags
	fv.Status = status
	fv.Time = time
	fv.Leftvalue = leftvalue
	fv.Rightvalue = rightvalue
	fv.Counter = counter


	beego.Info("end of get metrics")


	//2018/03/22 14:54:52 [I] [metrics_callback.go:12] start to get metrics
	//2018/03/22 14:54:52 [I] [metrics_callback.go:28] endpoint is 10.205.130.18
	//2018/03/22 14:54:52 [I] [metrics_callback.go:29] metric is cpu.idle
	//2018/03/22 14:54:52 [I] [metrics_callback.go:30] status is PROBLEM
	//2018/03/22 14:54:52 [I] [metrics_callback.go:31] step is 1
	//2018/03/22 14:54:52 [I] [metrics_callback.go:32] priority is 0
	//2018/03/22 14:54:52 [I] [metrics_callback.go:33] time is 2018-03-22 14:54:00
	//2018/03/22 14:54:52 [I] [metrics_callback.go:34] tpl_id is 296
	//2018/03/22 14:54:52 [I] [metrics_callback.go:35] exp_id is 0
	//2018/03/22 14:54:52 [I] [metrics_callback.go:36] stra_id is 1644
	//2018/03/22 14:54:52 [I] [metrics_callback.go:37] tags is
	//2018/03/22 14:54:52 [I] [metrics_callback.go:38] leftvalue is 94.64117831400814
	//2018/03/22 14:54:52 [I] [metrics_callback.go:39] rightvalue is 1
	//2018/03/22 14:54:52 [I] [metrics_callback.go:40] counter is 10.205.130.18/cpu.idle
	//2018/03/22 14:54:52 [I] [metrics_callback.go:42] end of get metrics


	go func() {
		err := autofill.RecoverMatch(fv)
		if err != nil {
			beego.Info(err.Error())
		}
	}()
}