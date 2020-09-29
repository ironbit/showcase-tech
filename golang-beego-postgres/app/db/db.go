package db

import (
	"net/url"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

	"sketch-tech/golang-beego-postgres/app/tool"
)

const (
	defaultName = "default"
	scheme      = "postgresql"
	sslmode     = "sslmode"
	disable     = "disable"

	connRetry = "connection::retry"
	connDelay = "connection::delay"
)

// Init is used to initialize the database.
func Init() {
	// url connection
	conn := getDatabaseURL()

	// driver env
	driver := tool.GetEnv("DATABASE_DRIVER")

	// register database
	orm.RegisterDriver(driver, orm.DRPostgres)

	// read 'retry' from config
	retry, err := beego.AppConfig.Int(connRetry)
	if err != nil {
		logs.Critical("db | Init | cannot read: %v", err)
		os.Exit(1)
	}

	// read 'delay' from config
	delay, err := beego.AppConfig.Int(connDelay)
	if err != nil {
		logs.Critical("db | Init | cannot read: %v", err)
		os.Exit(1)
	}

	// log information
	u, _ := url.Parse(conn)
	logs.Info("try to connect to database: user: %s | host: %s | database: %s", u.User.Username(), u.Host, u.Path[1:])

	isConnected := false
	for 0 < retry {
		err := orm.RegisterDataBase(defaultName, driver, conn)
		if err != nil {
			logs.Error("connection error: %v", err)
			logs.Error("Retry: %d", retry)
			time.Sleep(time.Duration(delay) * time.Second)
			retry--
		} else {
			isConnected = true
			break
		}
	}

	if isConnected {
		logs.Info("database connected")
	} else {
		logs.Critical("db | Init | cannot connect to database")
		os.Exit(1)
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
