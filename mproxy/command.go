package mproxy

import (
	"fmt"
	"github.com/siddontang/go-mysql/client"
	go_mysql "github.com/siddontang/go-mysql/mysql"
)

// ProxyHandler handles all the mysql related commands
type ProxyHandler struct {
	currentDBName string
	remoteConn    *client.Conn
	connPoolCache *ConnPoolCache
}

// UseDB change the currently used database
func (h *ProxyHandler) UseDB(dbName string) error {
	if h.remoteConn != nil {
		h.connPoolCache.ReturnConn(h.currentDBName, h.remoteConn)
	}

	h.remoteConn = h.connPoolCache.GetConn(dbName)
	h.currentDBName = dbName

	if h.remoteConn == nil {
		fmt.Println("Error in establishing a remote connection")
	} else {
		if err := h.remoteConn.Ping(); err != nil {
			fmt.Println("Error in pinging conn")
		}
	}
	return nil
}

// HandleQuery handles sql queries
func (h *ProxyHandler) HandleQuery(query string) (*go_mysql.Result, error) {
	fmt.Println("Exec Q: ", query)
	return h.remoteConn.Execute(query)
}

// HandleFieldList field list queries
func (h ProxyHandler) HandleFieldList(table string, fieldWildcard string) ([]*go_mysql.Field, error) {
	return nil, fmt.Errorf("not supported now")
}

// HandleStmtPrepare handles prepared statements
func (h ProxyHandler) HandleStmtPrepare(query string) (int, int, interface{}, error) {
	fmt.Println("prep: ", query)
	return 0, 0, nil, fmt.Errorf("not supported now")
}

// HandleStmtExecute handles execute statements
func (h ProxyHandler) HandleStmtExecute(context interface{}, query string, args []interface{}) (*go_mysql.Result, error) {
	fmt.Println("context: ", context, " query: ", query, " args:", args)
	return nil, fmt.Errorf("not supported now")
}

// HandleStmtClose handles close statement
func (h *ProxyHandler) HandleStmtClose(context interface{}) error {
	fmt.Println("Close context: ", context)
	return nil
}

// ReturnConn return connection to the ConnPoolCache
func (h *ProxyHandler) ReturnConn() {
	h.connPoolCache.ReturnConn(h.currentDBName, h.remoteConn)
}
