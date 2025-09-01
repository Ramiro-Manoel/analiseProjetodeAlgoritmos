package main

import (
	"analiseProjetodeAlgoritmos/internal/prime"
	"analiseProjetodeAlgoritmos/internal/random"
	"analiseProjetodeAlgoritmos/internal/sort"
	"fmt"
	"log"
	"time"
)

type timer struct {
	generate time.Duration
	sort     time.Duration
	prime    time.Duration
}

func (t timer) calculateTotal() time.Duration {
	return t.generate + t.sort + t.prime
}

func (t timer) calculateSortPorcentage() float64 {
	return float64(t.sort) / float64(t.calculateTotal()*100)
}

func (t timer) calculatePrimePorcentage() float64 {
	return float64(t.prime) / float64(t.calculateTotal()*100)
}

type Timers struct {
	timers []timer
}

func (t Timers) calculateAverage(f func(timer) float64) float64 {

	var summ float64

	for _, timer := range t.timers {
		summ += f(timer)
	}

	return float64(summ) / float64(len(t.timers))
}

func (t Timers) calculateTotalAverage() float64 {
	return t.calculateAverage(
		func(tr timer) float64 {
			return float64(tr.calculateTotal())
		})
}

func (t Timers) calculateSortPorcentageAverage() float64 {
	return t.calculateAverage(
		func(tr timer) float64 {
			return tr.calculateSortPorcentage()
		})
}

func (t Timers) calculatePrimePorcentageAverage() float64 {

	return t.calculateAverage(
		func(tr timer) float64 {
			return tr.calculatePrimePorcentage()
		})
}

type ResultsLog struct {
	scenario int
	size     int
	Timers
}

func (r ResultsLog) logScenarios() {
	fmt.Println("| Cenário | Execução | TGerar (ms) | TOrdernar (ms) | TPrimos (ms) | TTotal (ms) |")
	fmt.Println("|---------|----------|-------------|----------------|--------------|-------------|")

	for i, timer := range r.timers {
		fmt.Printf("|%d|%d|%d|%d|%d|%d|\n",
			r.scenario,
			i+1,
			timer.generate.Milliseconds(),
			timer.sort.Milliseconds(),
			timer.prime.Milliseconds(),
			timer.calculateTotal().Milliseconds())
	}
}

func (r ResultsLog) logAverages() {
	fmt.Println("| Cenário | n | %% TOrdenar | %% TPrimos | TMedia |")
	fmt.Println("|---------|---|-------------|------------|--------|")

	fmt.Printf("|%d|%d|%.2f|%.2f|%.2f|\n",
		r.scenario,
		r.size,
		r.calculateSortPorcentageAverage(),
		r.calculatePrimePorcentageAverage(),
		r.calculateTotalAverage())

}

func main() {

	var size int
	var timer timer
	var scenario int
	var scenariosLog ResultsLog

	fmt.Print("Scenario number: ")
	if _, err := fmt.Scan(&scenario); err != nil {
		log.Fatal("The scenario number must be a number")
	}
	scenariosLog.scenario = scenario

	fmt.Print("Array size: ")
	if _, err := fmt.Scan(&size); err != nil {
		log.Fatal("The array size must be a number")
	}
	scenariosLog.size = size

	for i := 0; i < 10; i++ {
		fmt.Println("execution ", i)

		start := time.Now()
		array, err := random.RandomArrayGenerator(size, 1500)
		if err != nil {
			log.Fatal(err)
		}
		timer.generate = time.Since(start)
		fmt.Println(time.Since(start))

		start = time.Now()
		array = sort.BubbleSort(array)
		timer.sort = time.Since(start)

		start = time.Now()
		_ = prime.GetPrimeNumbers(array)
		timer.prime = time.Since(start)

		scenariosLog.timers = append(scenariosLog.timers, timer)
	}

	scenariosLog.logScenarios()
	scenariosLog.logAverages()
}
