package scyllaDB

import (
	"time"

	"go-profiler/database"
	log "go-profiler/gopsutil"

	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

func CreateCluster(consistency gocql.Consistency, keyspace string, hosts ...string) *gocql.ClusterConfig {
	retryPolicy := &gocql.ExponentialBackoffRetryPolicy{
		Min:        time.Second,
		Max:        10 * time.Second,
		NumRetries: 5,
	}
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace
	cluster.Timeout = 5 * time.Second
	cluster.RetryPolicy = retryPolicy
	cluster.Consistency = consistency
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	return cluster
}

func Connect() *gocql.Session {
	logger := log.CreateLogger("info")
	cluster := CreateCluster(gocql.Quorum, "test", "localhost:9042")
	session, err := gocql.NewSession(*cluster)
	if err != nil {
		logger.Error("connect to scylla", zap.Error(err))
	}

	return session
}

func SelectQuery(session *gocql.Session, table database.IScyllaTable, selectFields []string, logger *zap.Logger) []gocql.RowData {
	logger.Info("Displaying Results:")
	query := table.BuildSelectQuery(selectFields)
	q := session.Query(query)
	it := q.Iter()
	defer func() {
		if err := it.Close(); err != nil {
			logger.Warn("select catalog.mutant", zap.Error(err))
		}
	}()

	var res []map[string]interface{}
	for {
		m := make(map[string]interface{})
		if !it.MapScan(m) {
			break
		}
		res = append(res, m)
	}

	var returndata []gocql.RowData
	for _, row := range res {
		columns := []string{}
		values := []interface{}{}
		for column, value := range row {
			columns = append(columns, column)
			values = append(values, value)
		}
		returndata = append(returndata, gocql.RowData{
			Columns: columns,
			Values:  values,
		})
	}
	return returndata

}

func InsertQuery(session *gocql.Session, logger *zap.Logger) {
	logger.Info("Inserting Mike")
	if err := session.Query("INSERT INTO mutant_data (first_name,last_name,address,picture_location) VALUES ('Mike','Tyson','1515 Main St', 'http://www.facebook.com/mtyson')").Exec(); err != nil {
		logger.Error("insert catalog.mutant_data", zap.Error(err))
	}
}
