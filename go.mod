module github.com/worldiety/mercurius

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golangee/forms v0.0.0-20200518162453-b89914aa6b82
	github.com/golangee/http v0.0.0-20200513135818-75d24690459f
	github.com/golangee/log v0.0.0-20200520130746-6f54960293eb
	github.com/golangee/log-zap v0.0.0-20200520145355-1ae3f536d1c8
	github.com/golangee/reflectplus v0.0.0
	github.com/golangee/sql v0.0.0-20200513144143-4ddbdfb22669
	github.com/golangee/uuid v0.0.0-20200513144043-882c55e8ee6c
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/golangee/reflectplus => ../../golangee/reflectplus

replace github.com/golangee/sql => ../../golangee/sql

replace github.com/golangee/uuid => ../../golangee/uuid

replace github.com/golangee/http => ../../golangee/http

replace github.com/golangee/log => ../../golangee/log

replace github.com/golangee/log-zap => ../../golangee/log-zap
