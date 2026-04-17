// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"fmt"
	"os"
	"strings"

	"github.com/cylixlee/shorturl/internal/config"
	"github.com/cylixlee/shorturl/internal/dispencer"
	"github.com/cylixlee/shorturl/internal/model"
	"github.com/cylixlee/shorturl/pkg/blacklist"
	"github.com/jxskiss/base62"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	MapModel      model.MapModel
	SequenceModel model.SequenceModel
	Dispencer     dispencer.Interface
	Encoder       *base62.Encoding
	Blacklist     blacklist.Interface
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		c.Mysql.User,
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		c.Mysql.Host,
		c.Mysql.Port,
		c.Mysql.DatabaseName,
		c.Mysql.Charset,
		c.Mysql.ParseTime,
		c.Mysql.Locale,
	)
	conn := sqlx.NewMysql(dsn)

	encoder := base62.StdEncoding
	if c.EncodingBaseString != "" {
		encoder = base62.NewEncoding(c.EncodingBaseString)
	}

	blacklist := blacklist.NewACBuilder()
	if c.Blacklist != "" {
		b, err := os.ReadFile(c.Blacklist)
		if err != nil {
			panic(err)
		}
		lines := strings.SplitSeq(string(b), "\n")
		for line := range lines {
			line = strings.TrimSpace(line)
			if len(line) == 0 || strings.HasPrefix(line, "#") {
				continue
			}
			blacklist.Add(line)
		}
	}
	return &ServiceContext{
		Config:        c,
		MapModel:      model.NewMapModel(conn, c.CacheRedis),
		SequenceModel: model.NewSequenceModel(conn, c.CacheRedis),
		Dispencer:     dispencer.NewMysql(conn),
		Encoder:       encoder,
		Blacklist:     blacklist.Build(),
	}
}
