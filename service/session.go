package service

import (
	"cygo_iris/session"

	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"
)

func SetLoginSession(ctx iris.Context, uid bson.ObjectId) {
	sess := session.Sess.Start(ctx)
	sess.Set("auth", true)
	sess.Set("uid", uid.Hex())
}

func SetLogoutSession(ctx iris.Context) {
	sess := session.Sess.Start(ctx)
	sess.Destroy()
}
