CREATE KEYSPACE test WITH replication = {'class': 'NetworkTopologyStrategy', 'replication_factor' : 1};

use test;

CREATE TABLE users (
    id uuid PRIMARY KEY,
    first_name text,
    last_name text,
    email text,
    picture_location text,
);

INSERT INTO users (id, first_name, last_name, email, picture_location) VALUES (uuid(), 'John', 'Doe', 'john.doe@test.com', 'http://www.example.com/pictures/john.doe.jpg');

type ProcessMessage struct {
	Pid       uint32    `json:"pid"`
	Cpu       float64   `json:"cpu"`
	Mem       float32   `json:"mem"`
	Name      string    `json:"name"`
	TimeStamp time.Time `json:"time"`
	Ctime     int64     `json:"ctime"`
}

CREATE TABLE process (
    pid int,
    cpu float,
    mem float,
    name text,
    time timestamp,
    ctime bigint,
    PRIMARY KEY (pid, ctime)
) WITH CLUSTERING ORDER BY (ctime DESC);