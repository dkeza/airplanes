# airplanes
Script in Go to show how to use SQLBoiler and Postgres database.

First install [SQLBoiler](https://github.com/volatiletech/sqlboiler)

Create new database airplanes in Postgres.

With psql import data

```
psql -d airplanes -f /home/keza/code/mygo/airplanes/db/import.sql
```

Create sqlboiler.toml file in your Go project

```
[psql]
dbname = "airplanes"
user="postgres"
pass="postgres"
host="localhost"
sslmode="disable"
```

Now cd in project directory, and generate models

```
sqlboiler psql --wipe
```

And run main.go