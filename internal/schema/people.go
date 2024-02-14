package schema

import "fmt"

type People struct {
}

func (p People) Validate() error {
	return fmt.Errorf("%s", "erro")
}
