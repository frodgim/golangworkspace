module example/secondmodule

go 1.23.1

replace example/firstmodule => ../firstmodule

replace example/codeorgmodule => ../codeorgmodule

require (
	example/codeorgmodule v0.0.0-00010101000000-000000000000
	example/firstmodule v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/text v0.18.0 // indirect
	rsc.io/quote v1.5.2 // indirect
	rsc.io/sampler v1.99.99 // indirect
)
