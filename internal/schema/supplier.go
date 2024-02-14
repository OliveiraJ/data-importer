package schema

import "fmt"

type Supplier struct {
}

func (s Supplier) Validate() error {
	return fmt.Errorf("%s", "erro")
}
