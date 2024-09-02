package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos, delay = 2, 5

func main() {

	exibeIntroducao()

	imprimeLogs()

	for {
		exibeMenu()
		comando := leComando()
		executaComando(comando)
	}
}

func executaComando(comando int) {
	switch comando {
	case 1:
		iniciarMonitoramento()
		// verificaSlices()
	case 2:
		fmt.Println("Exibindo Logs...")
		imprimeLogs()
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Tá maluco só pode")
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, error := os.Open("sites.txt")
	// arquivoBytes, error := os.ReadFile("sites.txt")

	if error != nil {
		fmt.Println("Deu treta: ", error)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, error := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)
		sites = append(sites, linha)

		if error == io.EOF {
			fmt.Println("Momento de parar a leitura: ", error)
			break
		}
	}

	// fmt.Println(string(arquivoBytes))
	arquivo.Close()
	return sites
}

func verificaSlices() {
	nomes := []string{"rodrigo", "rafaela"}

	fmt.Println("lista de capacidade de: ", cap(nomes), " e len de: ", len(nomes))

	nomes = append(nomes, "joao", "edson")

	fmt.Println("lista de capacidade de: ", cap(nomes), " e len de: ", len(nomes))

	nomes = append(nomes, "rafael")

	fmt.Println("lista de capacidade de: ", cap(nomes), " e len de: ", len(nomes))

}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func exibeIntroducao() {
	fmt.Println("************************************")
	nome, versao := retornaIntroducaoNomeEVersao()
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("************************************")
}

func retornaIntroducaoNomeEVersao() (string, float64) {
	return "Rodrigo", 1.1
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("Ele escolheu:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
			time.Sleep(delay * time.Second)
		}
	}
}

func testaSite(site string) {
	response, error := http.Get(site)

	// verificar o arquivo de log se existe ou nao e abrir ou criar e abrir
	// com o arquivo aberto
	// verificar teste site falha ou sucesso
	// guardar as infos no arquivo
	// fechar o arquivo

	if error != nil {
		fmt.Println("Deu treta: ", error)
	}

	if response.StatusCode == http.StatusOK {
		fmt.Println("Site:", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", response.StatusCode)
		registraLog(site, false)
	}
}

func registraLog(site string, status bool) {
	arquivo, error := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)

	if error != nil {
		fmt.Println(error)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivoBytes, error := os.ReadFile("log.txt")

	if error != nil {
		fmt.Println("deu treta")
	}

	fmt.Println(string(arquivoBytes))
}
