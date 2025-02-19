package persistence

import (
	"log"
	"time"

	"github.com/gocql/gocql"
)

func StartPersistenceService(){

}

func startScylla() *gocql.Session {
	cluster := gocql.NewCluster("127.0.0.1") // Replace with your Scylla/Cassandra host(s)
	cluster.Keyspace = "mandcondor_chat"     // Replace with your keyspace
	cluster.Consistency = gocql.Quorum
	// Optionally, adjust timeouts and other settings:
	cluster.Timeout = 10 * time.Second

	// Create a session. You may wish to handle errors differently in production.
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to ScyllaDB: %v", err)
	}

	return session
}