module github.com/S035779/yahoo-finance-openapi/go/test

replace github.com/S035779/yahoo-finance-openapi => ../../

go 1.13

require (
	github.com/pasdam/yahoo-finance-openapi/gen/go/yq1 v0.0.0-20231005014554-ffccf256a4c7
	github.com/stretchr/testify v1.7.0
)
