package session

import (
	"os"
	"time"

	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

var Sess *sessions.Sessions

func ConnectRedis() {
	db := redis.New(redis.Config{
		Network:   "tcp",
		Addr:      os.Getenv("REDIS_ADDR"),
		Timeout:   time.Duration(30) * time.Second,
		MaxActive: 10,
		// Use REDIS_USERNAME only if redis version >= 6.0
		// Username:  os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: os.Getenv("REDIS_DATABASE"),
		Prefix:   "",
		Driver:   redis.GoRedis(), // defautls.
	})

	Sess = sessions.New(sessions.Config{
		Cookie:          "_session_id",
		Expires:         0,
		AllowReclaim:    true,
		CookieSecureTLS: true,
	})

	Sess.UseDatabase(db)
}
