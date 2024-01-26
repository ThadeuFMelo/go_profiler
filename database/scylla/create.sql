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