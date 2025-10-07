package main

func FindUniq(arr []float32) float32 {
	trueTabel := make(map[float32]int)
	for _, value := range arr {
		_, ok := trueTabel[value]
		if ok {
			trueTabel[value]++

		} else {
			trueTabel[value] = 1
		}
	}
	for element, count := range trueTabel {
		// Проверяем, равно ли количество повторений единице
		if count == 1 {
			// Если да, мы нашли уникальный элемент, и сразу его возвращаем
			return element
		}
	}
	return 0
}

func main() {
	FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0})
}
