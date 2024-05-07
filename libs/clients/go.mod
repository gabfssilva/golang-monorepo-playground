module github.com/gabfssilva/gomonorepo/libs/clients

go 1.21

require github.com/gabfssilva/gomonorepo/libs/result v0.0.0
replace github.com/gabfssilva/gomonorepo/libs/result v0.0.0 => ./../result

require github.com/gabfssilva/gomonorepo/libs/async v0.0.0
replace github.com/gabfssilva/gomonorepo/libs/async v0.0.0 => ./../async

require github.com/gabfssilva/gomonorepo/libs/observability v0.0.0
replace github.com/gabfssilva/gomonorepo/libs/observability v0.0.0 => ./../observability

require github.com/go-resty/resty/v2 v2.12.0

require golang.org/x/net v0.22.0 // indirect
