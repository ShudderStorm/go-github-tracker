package ssl

type Mode string

const (
	Disable Mode = "disable"
	Prefer  Mode = "prefer"
	Require Mode = "require"
)
