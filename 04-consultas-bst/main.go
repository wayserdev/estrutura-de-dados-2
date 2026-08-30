package main

import "fmt"

type No struct {
	Valor              int
	Esquerdo, Direito *No
}

type Arvore struct {
	Raiz *No
}

func inserir(no *No, v int) *No {
	if no == nil {
		return &No{Valor: v}
	}
	if v < no.Valor {
		no.Esquerdo = inserir(no.Esquerdo, v)
	} else if v > no.Valor {
		no.Direito = inserir(no.Direito, v)
	}
	return no
}

func (a *Arvore) Inserir(v int) {
	a.Raiz = inserir(a.Raiz, v)
}

func Altura(no *No) int {
	if no == nil {
		return -1
	}
	esquerda := Altura(no.Esquerdo)
	direita := Altura(no.Direito)
	if esquerda > direita {
		return esquerda + 1
	}
	return direita + 1
}

func Contar(no *No) int {
	if no == nil {
		return 0
	}
	return 1 + Contar(no.Esquerdo) + Contar(no.Direito)
}

func ContarFolhas(no *No) int {
	if no == nil {
		return 0
	}
	if no.Esquerdo == nil && no.Direito == nil {
		return 1
	}
	return ContarFolhas(no.Esquerdo) + ContarFolhas(no.Direito)
}

func Minimo(no *No) *No {
	if no == nil {
		return nil
	}
	for no.Esquerdo != nil {
		no = no.Esquerdo
	}
	return no
}

func Maximo(no *No) *No {
	if no == nil {
		return nil
	}
	for no.Direito != nil {
		no = no.Direito
	}
	return no
}

func mostrarConsultas(nome string, arvore *Arvore) {
	fmt.Println(nome)
	fmt.Printf("Altura: %d\n", Altura(arvore.Raiz))
	fmt.Printf("Quantidade de nós: %d\n", Contar(arvore.Raiz))
	fmt.Printf("Quantidade de folhas: %d\n", ContarFolhas(arvore.Raiz))

	minimo := Minimo(arvore.Raiz)
	maximo := Maximo(arvore.Raiz)
	if minimo == nil {
		fmt.Println("Mínimo: árvore vazia")
		fmt.Println("Máximo: árvore vazia")
	} else {
		fmt.Printf("Mínimo: %d\n", minimo.Valor)
		fmt.Printf("Máximo: %d\n", maximo.Valor)
	}
	fmt.Println()
}

func main() {
	vazia := &Arvore{}
	mostrarConsultas("Árvore vazia:", vazia)

	unicoNo := &Arvore{}
	unicoNo.Inserir(10)
	mostrarConsultas("Árvore com um nó:", unicoNo)

	completa := &Arvore{}
	for _, valor := range []int{50, 30, 70, 20, 40, 60, 80, 35, 65} {
		completa.Inserir(valor)
	}
	mostrarConsultas("Árvore do exercício:", completa)
}

