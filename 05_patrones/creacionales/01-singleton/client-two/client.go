package client_two

import "github/alvaroenriqueds/patrones-go/creacionales/01-singleton/singleton"

func IncrementAge()  {
	p := singleton.GetInstance()
	p.IncrementAge()
}
