package clase01

import "testing"


//para ejecutar todo -> go test
//para ejecutar test especificos con expresiones regulares-> go test -run AddAcum
//                                                           go test -run Add$

func TestAdd(t *testing.T) {
        want := 5
        got := Add(3,3)
        if got != want {
                t.Log("Se esperaba: ",want)
                t.Log("Se obtuvo: ", got)
                t.Fail()
        }
}

func TestAddAcum(t *testing.T)  {
        want := 14
        got := AddAcum(5,5,5)
        if got != want {
                t.Log("Se esperaba: ",want)
                t.Log("Se obtuvo: ", got)
                t.Fail()
        }
}
