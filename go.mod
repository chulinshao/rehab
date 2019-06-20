module github.com/chulinshao/rehab

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190325154230-a5d413f7728c
	golang.org/x/net => github.com/golang/net v0.0.0-20181023162649-9b4f9f5ad519
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190329044733-9eb1bfa1ce65
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20181221001348-537d06c36207
)

require (
	github.com/kataras/iris v11.1.1+incompatible // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)
