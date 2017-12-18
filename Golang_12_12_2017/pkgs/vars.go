package pkgs

var int1 int = 30

// zero value initialization
var b1, b2, b3 bool

// value initialization
var b4, b6, b7 bool = true, false, true

// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var c, python, java = true, false, "no!"

var (
	toBe   bool   = false
	maxInt uint64 = 1<<64 - 1
)

const t int = 30

func test() {
	h := 45 // is int
}
