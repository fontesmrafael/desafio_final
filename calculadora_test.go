package main
import "testing"
func TestSoma(t *testing.T) { 
(assinatura do método)

teste := soma(36, 24, 12)

resultado := 72

if teste != resultado { 
t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
}
}
func TestSoma2(t *testing.T) { 
teste := soma(36, 24, 12) 
resultado := 100 
if teste != resultado { 
t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
}
}
func TestMult(t *testing.T) {
teste := multiplica(100, 100)
resultado := 1000
if teste != resultado {
t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
}
}
func TestMult2(t *testing.T) {
teste := multiplica(100, 100)
resultado := 25604
if teste != resultado {
t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
}
}
func TestSub(t *testing.T) {
	teste := subtrai(20, 5)
	resultado := 15
if teste != resultado {
	t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
	func TestSub2(t *testing.T) {
	teste := subtrai(20, 5)
	resultado := 5

	if teste != resultado {
	t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
	func TestDiv(t *testing.T) {
	teste := divide(200)
	resultado := 1
	if teste != resultado {
	t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
	}
	func TestDiv2(t *testing.T) {
	teste := divide(200)
	resultado := 5
	if teste != resultado {
	t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}	