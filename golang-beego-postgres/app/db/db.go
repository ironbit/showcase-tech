package db

import (
	"net/url"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"sketch-tech/golang-beego-postgres/app/tool"
)

const (
	defaultName = "default"
	scheme      = "postgresql"
	sslmode     = "sslmode"
	disable     = "disable"
)

func init() {
	// url connection
	conn := getDatabaseURL()

	// driver env
	driver := tool.GetEnv("DATABASE_DRIVER")

	// register database
	orm.RegisterDriver(driver, orm.DRPostgres)
	err := orm.RegisterDataBase(defaultName, driver, conn)
	if err != nil {
		logs.Error("database error: %v", err)
	} else {
		u, _ := url.Parse(conn)
		logs.Info("database connected: user: %s | host: %s | database: %s", u.User.Username(), u.Host, u.Path[1:])
	}
}

func getDatabaseURL() string {
	// environmental variables
	user := tool.GetEnv("DATABASE_USER")
	pass := tool.GetEnv("DATABASE_PASS")
	host := tool.GetEnv("DATABASE_HOST")
	port := tool.GetEnv("DATABASE_PORT")
	name := tool.GetEnv("DATABASE_NAME")

	// create url
	u := url.URL{}

	// populate url
	u.Scheme = scheme
	u.User = url.UserPassword(user, pass)
	u.Host = host + ":" + port
	u.Path = name

	// populate query
	q := u.Query()
	q.Set(sslmode, disable)
	u.RawQuery = q.Encode()

	// url build
	return u.String()
}
