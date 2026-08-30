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

func minimo(no *No) *No {
	for no != nil && no.Esquerdo != nil {
		no = no.Esquerdo
	}
	return no
}

func Remover(no *No, v int) *No {
	if no == nil {
		return nil
	}

	if v < no.Valor {
		no.Esquerdo = Remover(no.Esquerdo, v)
	} else if v > no.Valor {
		no.Direito = Remover(no.Direito, v)
	} else {
		if no.Esquerdo == nil && no.Direito == nil {
			return nil
		}

		if no.Esquerdo == nil {
			return no.Direito
		}
		if no.Direito == nil {
			return no.Esquerdo
		}

		sucessor := minimo(no.Direito)
		no.Valor = sucessor.Valor
		no.Direito = Remover(no.Direito, sucessor.Valor)
	}
	return no
}

func (a *Arvore) Remover(v int) {
	a.Raiz = Remover(a.Raiz, v)
}

func EmOrdem(no *No) {
	if no == nil {
		return
	}
	EmOrdem(no.Esquerdo)
	fmt.Printf("%d ", no.Valor)
	EmOrdem(no.Direito)
}

func novaArvore(valores []int) *Arvore {
	a := &Arvore{}
	for _, valor := range valores {
		a.Inserir(valor)
	}
	return a
}

func testarRemocao(descricao string, arvore *Arvore, valor int) {
	fmt.Printf("%s - removendo %d\n", descricao, valor)
	fmt.Print("Antes:  ")
	EmOrdem(arvore.Raiz)
	fmt.Println()
	arvore.Remover(valor)
	fmt.Print("Depois: ")
	EmOrdem(arvore.Raiz)
	fmt.Println("\n")
}

func main() {
	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}

	testarRemocao("Caso folha", novaArvore(valores), 20)
	testarRemocao("Caso com um filho", novaArvore(valores), 40)
	testarRemocao("Caso com dois filhos", novaArvore(valores), 50)
}
