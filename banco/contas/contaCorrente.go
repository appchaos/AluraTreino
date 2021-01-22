package contas

import "fmt"

type ContaCorrente struct {
	//parece classe
	Titular string
	// quando variavel começa com minusculo é similiar propriedade private
	// quando variavel começa com maiusculo é similiar propriedade public
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

const limiteSaque = 700.00

func (c *ContaCorrente) Transferencia(valor float64, contaDestino *ContaCorrente) {

	fmt.Println("--------------------- DETALHE das CONTA ------------------------- ")
	fmt.Println(*c)
	fmt.Println(*contaDestino)
	fmt.Println("--------------------- Transferencia ------------------------- ")
	if c.Sacar(valor) {
		if contaDestino.Deposita(valor) {
			fmt.Println("Operacao de transferência realizada com sucesso!")
		}

	}
	fmt.Println("--------------------- DETALHE das CONTA ------------------------- ")
	fmt.Println(*c)
	fmt.Println(*contaDestino)
}

func (c *ContaCorrente) Sacar(valor float64) bool {
	// metodo similar com metodos da classe
	if valor >= 0 {

		if valor <= limiteSaque {
			if valor <= c.Saldo {
				c.Saldo -= valor
				c.contaDetalhada()
				fmt.Println("Saque realizado com sucesso!")
				return true
			} else {
				fmt.Println("Saldo na conta do ", c.Titular, " [", c.Saldo, "] é inferiodo do que valor do saque [", valor, "]!")
				return false
			}

		} else {
			fmt.Println("Valor do saque [", valor, "]maior que limite maximo permitido [", limiteSaque, "] ")
			return false
		}

	} else {
		fmt.Println("Por favor , informe um valor válido [", valor, "]!")
		return false
	}
}

func (c *ContaCorrente) Deposita(valor float64) bool {
	if valor >= 0 {
		c.Saldo += valor
		c.contaDetalhada()
		return true
	} else {
		fmt.Println("Por favor , informe um valor válido [", valor, "]!")
		return false
	}

}

func (c *ContaCorrente) contaDetalhada() {
	fmt.Println(*c)
}

//----------------------------- metodo normal do C#  também funciona no Go------------------

// func sacar(c ContaCorrente, valor float64) {
// 	if valor >= 0 {

// 		if valor <= limiteSaque {
// 			if valor <= c.Saldo {
// 				c.Saldo -= valor
// 				contaDetalhada(c)
// 				fmt.Println("Saque realizado com sucesso!")
// 			} else {
// 				fmt.Println("Saldo na conta do ", c.Titular, " [", c.Saldo, "] é inferiodo do que valor do saque [", valor, "]!")
// 			}

// 		} else {
// 			fmt.Println("Valor do saque [", valor, "]maior que limite maximo permitido [", limiteSaque, "] ")
// 		}

// 	} else {
// 		fmt.Println("Por favor , informe um valor válido!")
// 	}
// }

// func deposita(c ContaCorrente, valor float64) {
// 	if valor >= 0 {
// 		c.Saldo += valor
// 		contaDetalhada(c)
// 	} else {
// 		fmt.Println("Por favor , informe um valor válido!")
// 	}

// }

// func contaDetalhada(c ContaCorrente) {
// 	fmt.Println(c)
// }
