/* Funcionalidades:

O programa deve mostrar um título ao iniciar

Deve pedir o nome do usuário

Deve pedir uma senha numérica

Se a senha estiver correta:

mostrar uma mensagem de acesso liberado

Se a senha estiver errada:

mostrar acesso negado

O usuário pode tentar novamente quantas vezes quiser

Deve existir uma opção para sair do programa */

package main

import "fmt"

func cabecalho() {
	fmt.Println("=== CONTROLE DE ACESSO ~SIMPLES~ ===")
}

func menu() {
	fmt.Println("1. acessar")
	fmt.Println("0. sair")
	fmt.Print("escolha: ")
}

func acessar() (string, int) {
	fmt.Print(">> NOME DE USUARIO ")
	var nome string
	fmt.Scanln(&nome)

	fmt.Printf(">> SENHA DE %s ", nome)
	var senha int
	fmt.Scanln(&senha)

	fmt.Println("verificando permissão...")

	return nome, senha
}

func verificar(nome string, senha int) {
	if nome == "Admin" && senha == 1234 {
		fmt.Println("acesso liberado!")
	} else {
		fmt.Println("acesso negado!")
	}
}

func main() {
	var opcao int
	for {

		cabecalho()
		menu()
		fmt.Scanln(&opcao)

		switch opcao {
		case 0:
			return
		case 1:
			nome, senha := acessar()			
			verificar(nome, senha)
		default:
			fmt.Println("escolha invalida")
		}
	}
}