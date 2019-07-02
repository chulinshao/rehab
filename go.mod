module github.com/chulinshao/rehab

go 1.12

replace (
	cloud.google.com/go => github.com/GoogleCloudPlatform/google-cloud-go v0.34.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190325154230-a5d413f7728c
	golang.org/x/net => github.com/golang/net v0.0.0-20181023162649-9b4f9f5ad519
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190329044733-9eb1bfa1ce65
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20181221001348-537d06c36207
	google.golang.org/appengine => github.com/golang/appengine v1.4.0
)

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/core v0.6.2
	github.com/go-xorm/xorm v0.7.3
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	gopkg.in/yaml.v2 v2.2.2
)
