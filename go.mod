module github.com/404notfoundhard/http-metric.git

go 1.19

require (
	github.com/caarlos0/env/v6 v6.10.1
	github.com/go-chi/chi v1.5.4
	github.com/go-chi/chi/v5 v5.0.8
	
)
require internal/myFlags v1.0.0
replace internal/myFlags => ./internal/myFlags

require internal/myMetrics v1.0.0
replace internal/myMetrics => ./internal/myMetrics

