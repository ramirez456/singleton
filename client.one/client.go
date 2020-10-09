package client_one

import (
	singleton "singleton/single"
)

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
