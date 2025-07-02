package postgre

import (
	"github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq"
)

type DataSource struct {
	dsType        string
	connectionURL string
	connection    *dbr.Connection
}

func (d *DataSource) SetConnectionURL(connectionUrl string) {
	d.connectionURL = connectionUrl
}

func (d *DataSource) SetDataSourceType(dsType string) {
	d.dsType = dsType
}

func (d *DataSource) GetConnection() (*dbr.Connection, error) {
	if d.connection == nil {
		var err error
		d.connection, err = dbr.Open(d.dsType, d.connectionURL, nil)
		if err != nil {
			return nil, err
		}
	}
	return d.connection, nil
}
