package db

import (
	"net/url"
	"os"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
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
	driver := getEnv("DATABASE_DRIVER")

	// register database
	orm.RegisterDriver(driver, orm.DRPostgres)
	orm.RegisterDataBase(defaultName, driver, conn)

	u, _ := url.Parse(conn)
	logs.Info("database connected: user: %s | host: %s | database: %s", u.User.Username(), u.Host, u.Path[1:])
}

func getDatabaseURL() string {
	// environmental variables
	user := getEnv("DATABASE_USER")
	pass := getEnv("DATABASE_PASS")
	host := getEnv("DATABASE_HOST")
	port := getEnv("DATABASE_PORT")
	name := getEnv("DATABASE_NAME")

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

func getEnv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		logs.Critical("database: %s is not set", name)
		os.Exit(1)
	}
	return v
}
