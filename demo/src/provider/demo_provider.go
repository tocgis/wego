package provider

import (
	"github.com/9299381/wego"
	"github.com/9299381/wego/demo/src/controller"
	"github.com/9299381/wego/filters"
)

type DemoProvider struct {
}

func (it *DemoProvider) Boot() {
}

func (it *DemoProvider) Register() {

	//限速
	wego.Handler("one", filters.Limit(&controller.OneController{}))
	wego.Handler("two", filters.New(&controller.TwoController{}))
	wego.Handler("auth", filters.Chain(
		&filters.ResponseEndpoint{},
		&filters.JwtEndpoint{},
		&filters.LimitEndpoint{},
		&filters.CommEndpoint{
			Controller: &controller.AuthController{},
		}))

	wego.Handler("post", filters.New(&controller.PostController{}))
	wego.Handler("sql", filters.New(&controller.SqlController{}))
	wego.Handler("redis", filters.New(&controller.RedisController{}))
	wego.Handler("queue", filters.New(&controller.QueueController{}))

	wego.Handler("cache_set", filters.New(&controller.CacheSetController{}))
	wego.Handler("cache_get", filters.New(&controller.CacheGetController{}))
	//
	wego.Handler("valid", filters.New(&controller.ValidController{}))
	//
	wego.Handler("consul", filters.New(&controller.ConsulController{}))

	wego.Handler("event", filters.New(&controller.EventController{}))
	//
	wego.Handler("publish", filters.New(&controller.PublishController{}))
	wego.Handler("sleep", filters.New(&controller.SleepController{}))

}
