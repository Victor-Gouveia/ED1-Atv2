package main

import "fmt"

type ITree interface {
	Add(value int)
	Search(value int) bool
	Min() int
	Max() int
	PrintPre()
	PrintIn()
	PrintPos()
	PrintLevels()
	Height() int
	Remove(value int) *BstNode
}

type BstNode struct {
	left  *BstNode
	value int
	right *BstNode
}

// insere um novo valor na arvore
func (n *BstNode) Add(value int) {
	if value < n.value {
		// se o valor for menor, insere na esquerda
		if n.left == nil {
			n.left = &BstNode{value: value}
		} else {
			n.left.Add(value)
		}
	} else if value >= n.value {
		// se o valor for maior ou igual, insere na direita
		if n.right == nil {
			n.right = &BstNode{value: value}
		} else {
			n.right.Add(value)
		}
	}
}

// busca por um valor na arvore
func (n *BstNode) Search(value int) bool {
	if n == nil {
		return false
	}

	if value < n.value {
		return n.left.Search(value)
	} else if value > n.value {
		return n.right.Search(value)
	}
	// se value == n.value, o noh foi encontrado
	return true
}

// retorna o menor valor na arvore
func (n *BstNode) Min() int {
	if n.left == nil {
		return n.value
	}
	return n.left.Min()
}

// retorna o maior valor na arvore
func (n *BstNode) Max() int {
	if n.right == nil {
		return n.value
	}
	return n.right.Max()
}

// implementa o percurso em Pre-Ordem
func (n *BstNode) PrintPre() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.value)
	n.left.PrintPre()
	n.right.PrintPre()
}

// implementa o percurso Em-Ordem
func (n *BstNode) PrintIn() {
	if n == nil {
		return
	}
	n.left.PrintIn()
	fmt.Printf("%d ", n.value)
	n.right.PrintIn()
}

// implementa o percurso em Pos-Ordem
func (n *BstNode) PrintPos() {
	if n == nil {
		return
	}
	n.left.PrintPos()
	n.right.PrintPos()
	fmt.Printf("%d ", n.value)
}

// implementa o percurso por Niveis
func (n *BstNode) PrintLevels() {
	if n == nil {
		return
	}

	// usa um slice como uma fila
	queue := []*BstNode{n}

	for len(queue) > 0 {
		// Dequeue
		currentNode := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", currentNode.value)

		// Enqueue filhos
		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}
		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}
	}
}

// calcula a altura da arvore
func (n *BstNode) Height() int {
	if n == nil {
		// altura de uma arvore vazia eh -1
		return -1
	}

	leftHeight := n.left.Height()
	rightHeight := n.right.Height()

	// retorna o maior entre as duas subarvores (+1 para contar a aresta ate o noh atual)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// remove um valor da arvore e retorna a nova raiz da subarvore
func (n *BstNode) Remove(value int) *BstNode {
	if n == nil {
		return nil
	}

	// vai ate o noh a ser removido
	if value < n.value {
		n.left = n.left.Remove(value)
	} else if value > n.value {
		n.right = n.right.Remove(value)
	} else {
		// encontrando o noh, segue os casos

		// caso 1: noh sem filhos ou com um filho
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}

		// caso 2: noh tem dois filhos
		// encontra o sucessor em-ordem
		minRightSubtree := n.right.Min()
		// copia o valor do sucessor para o noh atual
		n.value = minRightSubtree
		// Remove o noh sucessor da subarvore
		n.right = n.right.Remove(minRightSubtree)
	}
	return n
}

// funcao de verificar se eh bst valida
func (bstNode *BstNode) isBst() bool {
	// se nao tem filhos, essa subarvore eh uma bst
	if bstNode.left == nil && bstNode.right == nil {
		return true
	}

	// inicia os boolean como verdadeiros, pois se um filho estiver nulo ele esta ok
	leftOK := true
	rightOK := true

	// esquerda esta ok se o valor for menor q atual e ele for uma bst
	if bstNode.left != nil {
		leftOK = (bstNode.left.value < bstNode.value) && bstNode.left.isBst()
	}
	// esquerda esta ok se o valor for maior ou igual ao atual e ele for uma bst
	if bstNode.right != nil {
		rightOK = (bstNode.right.value >= bstNode.value) && bstNode.right.isBst()
	}

	// tanto o filho da esquerda como o da direita devem estar ok
	return (leftOK && rightOK)
}

func (bstNode *BstNode) Size() int {
	if bstNode == nil {
		return 0
	}
	return 1 + bstNode.left.Size() + bstNode.right.Size()
}

func (node *BstNode) Par() int {
	if node == nil {
		return 0
	}

	par := 0
	if node.value%2 == 0 {
		par++
	}
	return par + node.left.Par() + node.right.Par()
}

func main() {
	// Cria a raiz
	root := &BstNode{value: 50}

	// Adiciona elementos
	values := []int{30, 70, 20, 40, 60, 80, 10, 25, 35, 45}
	for _, v := range values {
		root.Add(v)
	}

	fmt.Println("Árvore Inicial:")
	fmt.Println("--------------------------------")

	fmt.Print("Pré-Ordem (Raiz, Esq, Dir):   ")
	root.PrintPre()
	fmt.Println()

	fmt.Print("Em-Ordem (Esq, Raiz, Dir):    ")
	root.PrintIn()
	fmt.Println()

	fmt.Print("Pós-Ordem (Esq, Dir, Raiz):   ")
	root.PrintPos()
	fmt.Println()

	fmt.Print("Por Níveis:                   ")
	root.PrintLevels()
	fmt.Println()
	fmt.Println("--------------------------------")

	fmt.Printf("Altura da árvore: %d\n", root.Height())
	fmt.Printf("Valor mínimo: %d\n", root.Min())
	fmt.Printf("Valor máximo: %d\n", root.Max())
	fmt.Printf("Buscar 40: %t\n", root.Search(40))
	fmt.Printf("Buscar 99: %t\n", root.Search(99))
	fmt.Println("--------------------------------")

	// --- Demo da Remocao ---
	fmt.Println("\nRemovendo 10 (nó folha)...")
	root = root.Remove(10)
	fmt.Print("Em-Ordem após remover 10: ")
	root.PrintIn()
	fmt.Println()

	fmt.Println("\nRemovendo 20 (nó com 1 filho)...")
	root = root.Remove(20)
	fmt.Print("Em-Ordem após remover 20: ")
	root.PrintIn()
	fmt.Println()

	fmt.Println("\nRemovendo 30 (nó com 2 filhos)...")
	root = root.Remove(30)
	fmt.Print("Em-Ordem após remover 30: ")
	root.PrintIn()
	fmt.Println()

	fmt.Println("\nRemovendo 50 (a raiz)...")
	root = root.Remove(50)
	fmt.Print("Em-Ordem após remover 50: ")
	root.PrintIn()
	fmt.Println()
	fmt.Printf("Nova raiz: %d\n", root.value)
	fmt.Println("--------------------------------")
	fmt.Printf("Tamanho da árvore: %d\n", root.Size())
	fmt.Println("--------------------------------")
	fmt.Printf("Quantidade de pares: %d\n", root.Par())
	fmt.Println("--------------------------------")

	isTreeABst := root.isBst()
	fmt.Printf("A árvore é uma BST? %t\n", isTreeABst)
	// Forca a arvore a se tornar invalida para teste
	fmt.Println("Invalidando a árvore: root.left.right.value = 99")
	if root.left != nil && root.left.right != nil {
		root.left.right.value = 99
	}
	isTreeABst = root.isBst()
	fmt.Printf("A árvore modificada é uma BST? %t\n", isTreeABst)
}
