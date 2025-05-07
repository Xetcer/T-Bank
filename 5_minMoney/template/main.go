package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var n, swapCost, changeCost int
	var str string
	fmt.Fscanln(in, &n, &swapCost, &changeCost)
	strLen := 2 * n  // Длина строки
	if strLen == 0 { // Если пустая строка
		fmt.Fprint(out, 0)
	}
	fmt.Fscanln(in, &str)
	// fmt.Println(str)
	// fmt.Println(strLen, swapCash, changeCost)
	// counterMap := make(map[rune]int) // Карта, в которой будем смотреть сколько у нас открывающих и закрывающих скобок в строке
	// for _, char := range str {
	// 	counterMap[char]++
	// }
	// fmt.Println(counterMap)
	// deltaCount := counterMap['('] - counterMap[')'] // смотрим сколько разница между открывающими и закрывающими скобками по количеству
	// if deltaCount < 0 {
	// 	deltaCount *= -1
	// }
	// changesCount := deltaCount
	// fmt.Println(changesCount)
	// // Смотрим на меньшую сторону: сколько можно сэкономить
	// minSwaps := min(counterMap['('], counterMap[')'])

	// // Если смена мест выгодна, используем её, уменьшая количество изменений
	// if swapCost < changeCost {
	// 	changesCount -= minSwaps
	// }

	// // Финальная стоимость: учитываем изменение и возможное уменьшение за счет перестановок
	// finalCost := changesCount*changeCost + minSwaps*swapCost
	finalCost := BalanceStr(str, swapCost, changeCost)
	fmt.Fprint(out, finalCost)

}

// BalanceBrackets определяет минимальное количество монет для приведения строки к правильной форме
func BalanceStr(s string, swapCost, changeCost int) int {
	n := len(s)

	minCost := 0      // Минимальная стоимость
	balancer := 0     // Баланс открытых и закрытых скобок
	needSwap := false // Нужна ли перестановка

	for i := 0; i < n; i++ {
		ch := rune(s[i])

		if ch == '(' {
			balancer++ // Открывающая скобка увеличивает баланс
		} else if ch == ')' {
			balancer-- // Закрывающая уменьшает баланс

			if balancer < 0 {
				// Обнаружено нарушение, принимаем решение
				if needSwap {
					// Предложена предварительная перестановка
					minCost += swapCost
					needSwap = false
				} else {
					// Нельзя просто поменять местами, нужно или заменить, или дождаться следующего хода
					if i+1 < n && s[i+1] == '(' {
						// Перестановка следующая, пометка установлена
						needSwap = true
					} else {
						// Прямая замена необходима
						minCost += changeCost
					}
				}
			}
		}
	}

	// Заканчиваем начатые процессы
	if needSwap {
		minCost += swapCost
	}

	// Добавляем остаток неотбалансированных скобок
	if balancer > 0 {
		minCost += balancer * changeCost
	}

	return minCost
}
