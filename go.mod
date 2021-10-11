module github.com/el10savio/goSetReconciliation

go 1.16

require (
	github.com/bits-and-blooms/bloom/v3 v3.1.0
	github.com/gorilla/mux v1.8.0
	github.com/mitchellh/hashstructure v1.1.0
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
)

replace github.com/el10savio/goSetReconciliation/set => ../set

replace github.com/el10savio/goSetReconciliation/handlers => ../handlers

replace github.com/el10savio/goSetReconciliation/sync => ../sync
