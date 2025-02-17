// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package g

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/frame/gins"
	"github.com/gogf/gf/i18n/gi18n"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/net/gtcp"
	"github.com/gogf/gf/net/gudp"
	"github.com/gogf/gf/os/gcfg"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gres"
	"github.com/gogf/gf/os/gview"
)

// Client is a convenience function, that creates and returns a new HTTP client.
func Client() *ghttp.Client {
	return ghttp.NewClient()
}

// Server returns an instance of http server with specified name.
func Server(name ...interface{}) *ghttp.Server {
	return gins.Server(name...)
}

// TCPServer returns an instance of tcp server with specified name.
func TCPServer(name ...interface{}) *gtcp.Server {
	return gtcp.GetServer(name...)
}

// UDPServer returns an instance of udp server with specified name.
func UDPServer(name ...interface{}) *gudp.Server {
	return gudp.GetServer(name...)
}

// View returns an instance of template engine object with specified name.
func View(name ...string) *gview.View {
	return gins.View(name...)
}

// Config returns an instance of config object with specified name.
func Config(name ...string) *gcfg.Config {
	return gins.Config(name...)
}

// Cfg is alias of Config.
// See Config.
func Cfg(name ...string) *gcfg.Config {
	return Config(name...)
}

// Resource returns an instance of Resource.
// The parameter <name> is the name for the instance.
func Resource(name ...string) *gres.Resource {
	return gins.Resource(name...)
}

// I18n returns an instance of gi18n.Manager.
// The parameter <name> is the name for the instance.
func I18n(name ...string) *gi18n.Manager {
	return gins.I18n(name...)
}

// I18n returns an instance of gi18n.Manager.
// The parameter <name> is the name for the instance.
func I18nOptions(o gi18n.Options) *gi18n.Manager {
	return gins.I18nOptions(o)
}

// Res is alias of Resource.
// See Resource.
func Res(name ...string) *gres.Resource {
	return Resource(name...)
}

// Log returns an instance of glog.Logger.
// The parameter <name> is the name for the instance.
func Log(name ...string) *glog.Logger {
	return gins.Log(name...)
}

// Database is alias of DB.
// See DB.
// Deprecated, use DB instead.
func Database(name ...string) gdb.DB {
	return gins.Database(name...)
}

// DB returns an instance of database ORM object with specified configuration group name.
func DB(name ...string) gdb.DB {
	return gins.Database(name...)
}

// Table is alias of Model.
// The database component is designed not only for
// relational databases but also for NoSQL databases in the future. The name
// "Table" is not proper for that purpose any more.
// Deprecated, use Model instead.
func Table(tables ...string) *gdb.Model {
	return DB().Model(tables...)
}

// Model creates and returns a model based on configuration of default database group.
func Model(tables ...string) *gdb.Model {
	return DB().Model(tables...)
}

// With creates and returns an ORM model based on meta data of given object.
func With(object interface{}) *gdb.Model {
	return DB().With(object)
}

// Redis returns an instance of redis client with specified configuration group name.
func Redis(name ...string) *gredis.Redis {
	panic("forbidden to use this redis")
	//return gins.Redis(name...)
}
