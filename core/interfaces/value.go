package interfaces

type Number interface {
    ~int
}

type String interface {
    ~string
}

type Value interface {
	Number | String
}

