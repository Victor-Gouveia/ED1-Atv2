package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SelectionSort
func selectionSort(arr []int) {
	n := len(arr)
	// loop para todo o array
	for i := 0; i < n-1; i++ {
		// encontra o indice do menor elemento no subarray nao ordenado
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// troca o menor elemento encontrado com o primeiro elemento do subarray nao ordenado
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// bubbleSort implementa o algoritmo de ordenação Bubble Sort de forma in-place e otimizada.
func bubbleSort(arr []int) {
	n := len(arr)
	// O loop externo controla as passagens pelo array.
	for i := 0; i < n-1; i++ {
		// 'swapped' é usado para otimizar. Se nenhuma troca ocorrer em uma passagem,
		// o array já está ordenado e podemos parar.
		swapped := false
		// O loop interno empurra o maior elemento para o final do subarray nao ordenado
		for j := 0; j < n-i-1; j++ {
			// Compara elementos adjacentes.
			if arr[j] > arr[j+1] {
				// troca os elementos se estiverem na ordem errada
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// se nenhuma troca foi feita nesta passagem, o array esta ordenado
		if !swapped {
			break
		}
	}
}

// InsertionSort
func insertionSort(arr []int) {
	n := len(arr)
	// Comeca do segundo elemento, pois consideramos o primeiro
	// elemento como a subarray inicial
	for i := 1; i < n; i++ {
		// 'key' eh o elemento que va na posicao correta na subarray ordenada
		key := arr[i]
		// 'j' eh o índice do ultimo elemento da subarray ordenada
		j := i - 1

		// move os elementos da subarray ordenada que são maiores que a 'key' pra direita
		// para abrir espaço para a 'key'
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		// insere a 'key' na posicao correta
		arr[j+1] = key
	}
}

// MergeSort
func mergeSort(arr []int) []int {
	// se o array tem 0 ou 1 elemento, ele já está ordenado
	if len(arr) <= 1 {
		return arr
	}

	// encontra o meio do array
	meio := len(arr) / 2
	// divide o array em duas metades, esquerda e direita
	esquerda := mergeSort(arr[:meio])
	direita := mergeSort(arr[meio:])

	// combina as duas metades ordenadas
	return merge(esquerda, direita)
}

// merge combina dois slices ordenados
func merge(esquerda, direita []int) []int {
	// cria um slice para armazenar o resultado da mesclagem
	resultado := make([]int, 0, len(esquerda)+len(direita))
	i, j := 0, 0

	// compara os elementos das duas listas e adiciona o menor ao resultado
	for i < len(esquerda) && j < len(direita) {
		if esquerda[i] < direita[j] {
			resultado = append(resultado, esquerda[i])
			i++
		} else {
			resultado = append(resultado, direita[j])
			j++
		}
	}

	// adiciona os elementos restantes de qualquer uma das listas
	resultado = append(resultado, esquerda[i:]...)
	resultado = append(resultado, direita[j:]...)

	return resultado
}

// QuickSort
func quickSort(arr []int) {
	// inicializa a semente do gerador de numeros aleatorios para garantir
	// que a escolha do pivo sera diferente toda execucao
	rand.Seed(time.Now().UnixNano())
	quickSortRecursive(arr, 0, len(arr)-1)
}

// quickSortRecursive implementa a logica de recursao
func quickSortRecursive(arr []int, low, high int) {
	if low < high {
		// Encontra o índice do pivô, particionando o array.
		// arr[pivotIndex] agora está na sua posição final ordenada.
		pivotIndex := partition(arr, low, high)

		// Ordena recursivamente os elementos antes e depois do pivô.
		quickSortRecursive(arr, low, pivotIndex-1)
		quickSortRecursive(arr, pivotIndex+1, high)
	}
}

// partition reorganiza o array com base no pivo escolhido aleatoriamente
func partition(arr []int, low, high int) int {
	// escolhe um pivo aleatorio
	pivotRandomIndex := low + rand.Intn(high-low+1)

	// move o pivo para o final para facilitar a particao
	arr[pivotRandomIndex], arr[high] = arr[high], arr[pivotRandomIndex]
	pivot := arr[high]

	// 'i' eh o índice do ultimo elemento menor que o pivo
	i := low - 1

	// percorre o array (exceto pivo)
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++ // incrementa o índice do "menor" elemento
			// troca o elemento atual com o elemento em 'i'
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// troca o pivo com arr[i+1].
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}

// CountingSort
func countingSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	// encontra o valor maximo no array para determinar o range
	max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	// cria um array de contagem (count) de tamanho 'max + 1'
	count := make([]int, max+1)

	// armazena a contagem de cada elemento.
	for _, value := range arr {
		count[value]++
	}

	// modifica o array de contagem para armazenar a soma cumulativa
	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	// constroi o array output
	output := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		value := arr[i]
		// pega a posicao correta do valor a partir do array de contagem
		position := count[value] - 1
		// coloca o valor na sua posição no array de saida
		output[position] = value
		// diminui a contagem para o próximo elemento igual (se houver)
		count[value]--
	}

	return output
}

func main() {
	// array teste
	data := []int{4, 2, 2, 8, 3, 3, 6, 6, 5, 1}
	fmt.Println("Array desordenado:		", data)
	// CountingSort
	dataCount := countingSort(data)
	fmt.Println("Array ordenado (Counting):	", dataCount)
	// QuickSort
	dataQuick := data
	quickSort(dataQuick)
	fmt.Println("Array ordenado (Quick):		", dataQuick)
	// MergeSort
	dataMerge := mergeSort(data)
	fmt.Println("Array ordenado (Merge):		", dataMerge)
	// InsertionSort
	dataIns := data
	insertionSort(dataIns)
	fmt.Println("Array ordenado (Insertion):	", dataIns)
	// BubbleSort
	dataBubble := data
	bubbleSort(dataBubble)
	fmt.Println("Array ordenado (Bubble):	", dataBubble)
	// SelectionSort
	dataSel := data
	selectionSort(dataSel)
	fmt.Println("Array ordenado (Selection):	", dataSel)
}
