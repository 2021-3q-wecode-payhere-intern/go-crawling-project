module main

go 1.16

require (
	codef v0.0.0
	db v0.0.0

)

replace (
	codef v0.0.0 => ./codef
	db v0.0.0 => ./db
	key v0.0.0 => ./key
)
