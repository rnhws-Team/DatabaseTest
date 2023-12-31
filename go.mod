module DatabaseTest

go 1.20

require (
	github.com/google/uuid v1.3.0
	go.etcd.io/bbolt v1.3.7
)

require golang.org/x/sys v0.10.0 // indirect

replace go.etcd.io/bbolt v1.3.7 => ./depends/bbolt
