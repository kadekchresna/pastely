## Exec

```
cqlsh localhost -u cassandra -p cassandra

--
CREATE ROLE admin WITH PASSWORD = 'adminpass' AND SUPERUSER = true AND LOGIN = true;

--
CREATE KEYSPACE pastely
   WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};
```