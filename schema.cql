CREATE KEYSPACE cardhub WITH replication = {'class':'SimpleStrategy', 'replication_factor': 1};
USE cardhub;
CREATE TABLE cardhub.deck (id uuid, name text, PRIMARY KEY (id, name));
CREATE TABLE cardhub.word (id uuid, front text, back text, reversable boolean, info text, PRIMARY KEY (id));