package client_one

import "github/alvaroenriqueds/Golang_RoadToSenior/05_patrones/creacionales/01-singleton/singleton"

func IncrementAge()  {
	p := singleton.GetInstance()
	p.IncrementAge()
}
