module task5

go 1.17

require utils v0.0.0-00010101000000-000000000000

require (
	golang.org/x/exp v0.0.0-20211210185655-e05463a05a18 // indirect
	golang.org/x/tools v0.1.8 // indirect
	gonum.org/v1/gonum v0.9.3 // indirect
)

replace utils v0.0.0-00010101000000-000000000000 => ../utils
