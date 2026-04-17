// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Mysql struct {
		User         string
		Host         string
		Port         int
		DatabaseName string
		Charset      string
		ParseTime    bool
		Locale       string
	}
	EncodingBaseString string `json:",optional"`
	Blacklist          string `json:",optional"`
	ShortDomain        string
	CacheRedis         cache.CacheConf
}
