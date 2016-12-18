package mproxy

import (
	"fmt"
	"github.com/siddontang/go-mysql/client"
	"sync"
)

// ConnPoolCache manages a connection pool instance
type ConnPoolCache struct {
	conf  Config
	cache map[string]*sync.Pool
}

// NewConnPoolCache creates and intializes new conn pool
func NewConnPoolCache(conf Config) *ConnPoolCache {
	return &ConnPoolCache{
		conf:  conf,
		cache: make(map[string]*sync.Pool),
	}
}

// GetConn returns a mysql connection from the pool if there is else returns a new connection
func (cpc *ConnPoolCache) GetConn(dbName string) (conn *client.Conn) {

	if cp := cpc.cache[dbName]; cp == nil {
		cpc.cache[dbName] = &sync.Pool{}
	} else {
		if tmpConn := cp.Get(); tmpConn != nil {
			conn = tmpConn.(*client.Conn)
		}
	}

	if conn == nil {
		remoteConf := cpc.conf.Remote
		addr := fmt.Sprintf("%s:%d", remoteConf.Host, remoteConf.Port)
		if newConn, err := client.Connect(addr, remoteConf.Username, remoteConf.Password, dbName); err == nil {
			conn = newConn
		}
	}

	return
}

// ReturnConn to return a connection to the pool
func (cpc *ConnPoolCache) ReturnConn(dbName string, c *client.Conn) {
	if cp := cpc.cache[dbName]; cp != nil {
		cp.Put(c)
	}
}
