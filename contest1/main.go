package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var houses = []int{3, 5, 0, 0, 6, 21, 3, 0, 2, 90}
// var houses = []int{0, 7, 9, 4, 8, 20}
// var houses = []int{0, 7, 9, 4, 8, 20}
// var houses = []int{98, 0, 10, 77, 0, 59, 28, 0, 94}
var houses = []int{5, 8, 9, 12, 15, 26, 30, 0, 0, 55, 0, 0, 67, 0, 76, 80, 82, 0, 0, 98}

func getCountHouse() (int, error) {

	fmt.Println("input count houses:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	count, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, err
	}
	if count <= 0 {
		return 0, errors.New("must be zero")
	}

	return count, nil
}

func getHouses(l int) ([]int, error) {

	houses := make([]int, 0)

	fmt.Println("input houses:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return houses, err
	}

	housesIndex := strings.Split(scanner.Text(), " ")

	if len(housesIndex) != l {
		return houses, errors.New("Должны совпадать")
	}

	for _, idx := range housesIndex {
		i, err := strconv.Atoi(idx)
		if err != nil {
			return houses, err
		}
		houses = append(houses, i)
	}

	return houses, nil
}

func getInputData() ([]int, error) {

	countHouses, err := getCountHouse()
	if err != nil {
		return make([]int, 0), err
	}

	return getHouses(countHouses)
}

func min(arr []int) int {
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

func getEmptyHouses(h []int) map[int]struct{} {
	emptyHouse := make(map[int]struct{})
	for i, v := range houses {
		if v == 0 {
			emptyHouse[i] = struct{}{}
		}
	}
	return emptyHouse
}

func getMinimalDistance(emptyHouses map[int]struct{}, indexHouse int) int {

	distances := make([]int, 0)

	for emptyIdx := range emptyHouses {
		distance := emptyIdx - indexHouse
		if distance < 0 {
			distance = -distance // значение по модулю
		}
		distances = append(distances, distance)
	}

	return min(distances)
}

func main() {

	houses, err := getInputData()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Houses: %+v\n", houses)

	// Получаем индексы всех пустых домов
	emptyHouses := getEmptyHouses(houses)

	out := make([]int, len(houses))
	for houseIndex, houseNumber := range houses {
		if houseNumber != 0 {
			out[houseIndex] = getMinimalDistance(emptyHouses, houseIndex)
		} else {
			out[houseIndex] = 0 // участок пустой
		}
	}
	fmt.Print(out)
}
