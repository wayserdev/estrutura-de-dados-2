package main

import "fmt"

type No struct {
	Valor              int
	Esquerdo, Direito *No
}

type Arvore struct {
	Raiz *No
}

func NovaArvore() *Arvore {
	return &Arvore{}
}

func NovoNo(v int) *No {
	return &No{Valor: v}
}

func inserir(no *No, v int) *No {
	if no == nil {
		return NovoNo(v)
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

func emOrdem(no *No) {
	if no == nil {
		return
	}
	emOrdem(no.Esquerdo)
	fmt.Printf("%d ", no.Valor)
	emOrdem(no.Direito)
}

func main() {
	arvore := NovaArvore()
	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}

	for _, valor := range valores {
		arvore.Inserir(valor)
	}

	arvore.Inserir(50)

	fmt.Print("Árvore em ordem: ")
	emOrdem(arvore.Raiz)
	fmt.Println()

	vazia := NovaArvore()
	vazia.Inserir(10)
	fmt.Printf("Raiz após inserir em árvore vazia: %d\n", vazia.Raiz.Valor)
}
