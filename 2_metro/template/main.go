package main

import (
	"bufio"
	"fmt"
	"os"
)

type time struct {
	start, interval int
}

type query struct {
	n, time int
}

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Run(in, out)

}

func Run(in *bufio.Reader, out *bufio.Writer) {
	var n int // количество веток
	var q int // количество запросов
	fmt.Fscanln(in, &n)
	// fmt.Println("n", n)
	timeTable := make(map[int]time, 0) // карта в которую будем складывать расписание поездов для каждой ветки
	for i := 0; i < n; i++ {           // Заполним карту расписаний
		var a, b int
		fmt.Fscanln(in, &a, &b)
		timeTable[i] = time{start: a, interval: b}
	}
	// fmt.Println(timeTable)
	fmt.Fscanln(in, &q) // Получим число запросов
	// fmt.Println("Q", q)
	qyeryList := make(map[int]query, 0) // карта запросов
	for i := 0; i < q; i++ {            // заполним карту
		var t, d int
		fmt.Fscanln(in, &t, &d)
		qyeryList[i] = query{n: t, time: d}
	}

	// fmt.Println(qyeryList)
	for key, value := range qyeryList {
		timeOfArrival := getTimeOfArrival(timeTable[value.n-1], value.time) // получаем время для конкретной ветки
		fmt.Fprint(out, timeOfArrival)
		if key < q-1 {
			fmt.Fprint(out, "\r\n")
		}
	}
}

func getTimeOfArrival(timetable time, waitTime int) int {
	sum := timetable.start
	for sum < waitTime {
		sum += timetable.interval
	}
	return sum
}
