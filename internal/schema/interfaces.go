package schema

type Entity interface {
	Validate() error
}
