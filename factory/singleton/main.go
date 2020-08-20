package singleton

import (
	"net"
	"sync"
)

// single thread singleton

type dbConnection struct {
	mx *sync.RWMutex
	// unexported fields here
}

func (db *dbConnection) Connect() (net.Conn, error) {
	// do connection here
	// code here is not concurrent if called from multiple goroutines
	// either use channels or lock
	db.mx.Lock()
	// do write
	db.mx.Unlock()

}

var dbConn *dbConnection

func GetConnection() *dbConnection {
	if dbConn == nil {
		return new(dbConnection)
	}
	return dbConn
}
