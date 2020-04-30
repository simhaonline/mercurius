module github.com/worldiety/mercurius

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.1.1
	github.com/lib/pq v1.4.0
	github.com/worldiety/sqlm v0.0.0
)

replace github.com/worldiety/sqlm => ../sqlm

