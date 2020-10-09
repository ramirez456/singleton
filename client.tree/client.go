package client_tree

import (
	singleton "singleton/single"
)

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
