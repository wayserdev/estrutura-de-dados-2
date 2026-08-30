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

func PreOrdem(no *No) {
	if no == nil {
		return
	}
	fmt.Printf("%d ", no.Valor)
	PreOrdem(no.Esquerdo)
	PreOrdem(no.Direito)
}

func EmOrdem(no *No) {
	if no == nil {
		return
	}
	EmOrdem(no.Esquerdo)
	fmt.Printf("%d ", no.Valor)
	EmOrdem(no.Direito)
}

func PosOrdem(no *No) {
	if no == nil {
		return
	}
	PosOrdem(no.Esquerdo)
	PosOrdem(no.Direito)
	fmt.Printf("%d ", no.Valor)
}

func EmLargura(raiz *No) {
	if raiz == nil {
		return
	}
	fila := []*No{raiz}
	for len(fila) > 0 {
		atual := fila[0]
		fila = fila[1:]
		fmt.Printf("%d ", atual.Valor)

		if atual.Esquerdo != nil {
			fila = append(fila, atual.Esquerdo)
		}
		if atual.Direito != nil {
			fila = append(fila, atual.Direito)
		}
	}
}

func main() {
	arvore := &Arvore{}
	for _, valor := range []int{50, 30, 70, 20, 40, 60, 80, 35, 65} {
		arvore.Inserir(valor)
	}

	fmt.Print("Pré-ordem: ")
	PreOrdem(arvore.Raiz)
	fmt.Println()

	fmt.Print("Em ordem: ")
	EmOrdem(arvore.Raiz)
	fmt.Println()

	fmt.Print("Pós-ordem: ")
	PosOrdem(arvore.Raiz)
	fmt.Println()

	fmt.Print("Em largura: ")
	EmLargura(arvore.Raiz)
	fmt.Println()
}

// EmOrdem fica crescente porque visita primeiro os menores valores da BST.
