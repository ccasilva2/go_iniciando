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

const monitoramento = 5

func main() {

	leSitesDoArquivo()

	cidade, populacao, capital := devolveCidadeEPopulacao()
	if capital {
		fmt.Println("A capital ", cidade, "tem", populacao, "habitantes")
	} else {
		fmt.Println("A cidade ", cidade, "tem", populacao, "habitantes")
	}

	nome, idade := devolveNomeEidade()
	fmt.Println(nome, "e tenho ", idade, "anos.")

	_, idade1 := devolveNomeEidade()
	fmt.Println("e tenho ", idade1, "anos.")

	exibeIntroducao()

	for {

		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa.")
			os.Exit(0) //os.Exit(-1)
		default:
			fmt.Println("Não conheço este comando.")
		}
	}

}

func devolveCidadeEPopulacao() (string, int, bool) {
	return "Vila Sem Nome", 4328, true
}

func exibeIntroducao() {
	//idade := 24
	var nome string = "Carlos"
	var versao float32 = 1.1
	fmt.Println("Ola, sr.", nome)
	fmt.Println("Este programa esta na versao", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comandolido int
	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comandolido)

	fmt.Println("O comando escolhido foi", comandolido)
	return comandolido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	//Slices em GO
	//sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br"}
	//sites = append(sites, "https://www.caelum.com.br")

	/*for i := 0; i < len(sites); i++ {
		fmt.Println(sites[i])
	}*/

	sites := leSitesDoArquivo()

	time.Sleep(monitoramento * time.Second)

	for i, site := range sites {
		fmt.Println("Testando site ", i, " : ", site)
		testaSite(site)
		time.Sleep(monitoramento * time.Second)
	}

}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF { // quando chegar ao final da linha
			break
		}

	}

	arquivo.Close()
	return sites
}

func devolveNomeEidade() (string, int) {
	nome := "Carlos"
	idade := 35
	return nome, idade
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	} else {
		if resp.StatusCode == 200 {
			fmt.Println("Site: ", site, " oi carregado com sucesso!!!")
			registraLog(site, true)
		} else {
			fmt.Println("Site: ", site, " está com problemas. Status Code: ", resp.StatusCode)
			registraLog(site, false)
		}
	}

}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	//https://go.dev/src/time/format.go
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "- " + site + " = online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

/*

### Criado primeiro executavel.

carlossilva@Carlos-Silva hello % ls
hello.go
carlossilva@Carlos-Silva hello % go build helli.go
no required module provides package helli.go: go.mod file not found in current directory or any parent directory; see 'go help modules'
carlossilva@Carlos-Silva hello % go build hello.go
carlossilva@Carlos-Silva hello % ls
hello		hello.go
carlossilva@Carlos-Silva hello % ./hello
Ola Mundo, Douglas!
carlossilva@Carlos-Silva hello %

### Compilando e atualizando o executavel.

carlossilva@Carlos-Silva hello % go run hello.go
Ola Mundo, Douglas!!!
carlossilva@Carlos-Silva hello %

*/

/*

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa.")
	} else {
		fmt.Println("Não conheço este comando.")
	}

*/
