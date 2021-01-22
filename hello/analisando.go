package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const nroMonitoramento = 3
const delay = 3
const arquivoLeitura = "teste.txt"
const logFile = "Log.txt"

func main() {

	exibeIntroducao()
	exibirMenu()
	comando := leComando()
	executarTarefa(comando)
}
func sair() {
	println("saindo do sistema....")
	os.Exit(0)
}
func exibirMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}
func exibeIntroducao() {
	nome := "Wang"
	versao := 1.0
	fmt.Println("Olá. sr", nome, "!")
	fmt.Println("Versão do sistema : ", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("Comando lido é : ", comandoLido)
	return comandoLido
}

func executarTarefa(comando int) {
	// if comando < 0 {
	// 	fmt.Println("Comando foi executado ", comando)
	// } else {
	// 	fmt.Println("Comando invalido")
	// }

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		exibirLog()
	case 0:
		sair()
	default:
		os.Exit(-1)
	}
}

func iniciarMonitoramento() {

	fmt.Println("Iniciando monitoramento ....")
	sites := carregarSitesDoArquivo()
	// sites := []string{
	// 	"http://teste.teste.com.br",
	// }
	for {

		verificarStatusSite(sites)

	}

}
func linhaDivisor() {
	println("---------------------------------------------------")
}
func verificarStatusSite(sites []string) {

	for i := 0; i < nroMonitoramento; i++ {

		for i, site := range sites {
			resp, err := http.Get(sites[i])
			if err != nil {
				fmt.Println("Aconteceu um erro : ", err)
			} else if resp.StatusCode == 200 {
				println("Site ", site, " foi carregado com sucesso : ", resp.StatusCode, "!")
				registrarLog(site, true, resp.StatusCode)
			} else {
				println("Site", site, "Está com problemas com codigo : ", resp.StatusCode, "!")
				registrarLog(site, false, resp.StatusCode)
			}

			linhaDivisor()
			num := i + 1
			println("Terminando exercicio de monitoramento de : ", num, "/", len(sites), "[", cap(sites), "].")

			linhaDivisor()

		}
		println("Dormindo por ", delay, " segundos...")
		time.Sleep(delay * time.Second)
	}
	println("Monitoramento ", nroMonitoramento, "x encerrado com sucesso!")
	exibirMenu()
	executarTarefa(leComando())

}

func exibirLog() {
	arquivo, err := ioutil.ReadFile(logFile)

	if err != nil {
		exibirErro(err)
	} else {
		fmt.Println(string(arquivo))
	}

}
func carregarSitesDoArquivo() []string {
	var sites []string
	// arquivo, err := ioutil.ReadFile("sites.txt")
	// ioutil.ReadFile é mais usado para fazer leitura de todo conteudo em uma vez so

	arquivo, err := os.Open(arquivoLeitura)
	// metodo para abrir arquivo pela função do pacote do os

	leitor := bufio.NewReader(arquivo)
	// bufio. ja consegue mais operações depois da leitura com NewReader
	exibirErro(err)

	// ReadString do Bufio permite que a leitura para no deteminado caracter ,
	// caso até final da linha é \n,
	// usando aspa simples (') pois arquivo carregado pelo os.Open é em formato Byte[]

	// fmt.Println(string(arquivo))
	for {

		linha, err := leitor.ReadString('\n')

		if err != io.EOF {
			linha = strings.TrimSpace(linha)
			// pacote strings possui uma função que chama TrinSpace1 que remove caracters vazia e tambem pula linha
			fmt.Println(linha)
			sites = append(sites, linha)
		} else {

			break
			// break serve para encerrar loop infinita do for
		}

	}
	arquivo.Close()
	// sempre manter costume de fechar arquivo depois de utilizar
	return sites
}

func exibirErro(err error) {
	if err != nil {
		fmt.Println("Ocorreu um erro : ", err, "...")
	}
}

func registrarLog(site string, status bool, code int) {

	arquivo, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// o metodo OpenFile permite que voce ter mais opção na hora de abrir aquivo , exemplo criar um arquivo quando nao existe0

	if err != nil {
		exibirErro(err)
	}
	now := time.Now().Format("02/01/2006 15:04:05")

	// formatação de tempo no GO, eles tem uma forma especifica de definição
	// segue documentação https://golang.org/src/time/format.go
	arquivo.WriteString("[ " + now + "] - Site : " + site + " - status : " + strconv.FormatBool(status) + "! \n")

	arquivo.Close()
}
