package client_two

import (
	singleton "singleton/single"
)

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
