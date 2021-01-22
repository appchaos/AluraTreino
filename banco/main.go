package main

import (
	"banco/contas"
)



type verificarConta interface {
	Sacar(valor float64) bool
}
strconv.FormatFloat
func pagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

func main() {

	contaDoGuilherme := contas.ContaPoupanca{
		Titular:       "Gulherme",
		NumeroAgencia: 589,
		NumeroConta:   123456,
		Saldo:         623.45,
	}

	pagarBoleto(&contaDoGuilherme, 100)


	fmt.Println(typeOf(contaDoGuilherme))
	// sacar(contaDoGuilherme, 500.00)
	// deposita(contaDoGuilherme, 1000.00)

	// contaDoGuilherme.sacar(100)
	// contaDoGuilherme.deposita(200)

}
