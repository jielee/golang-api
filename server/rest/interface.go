package rest

type Rester interface {
	Poster
	Getter
}

type Getter interface {
	Get()
}

type Putter interface {
	Put()
}

type Poster interface {
	Post()
}

type Deleter interface {
	Delete()
}
