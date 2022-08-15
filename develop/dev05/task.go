package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===
Реализовать утилиту фильтрации (man grep)
Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	after := flag.Int("A", 0, "печатать N строк после совпадения")
	before := flag.Int("B", 0, "печатать N строк до совпадения")
	contextText := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	countBool := flag.Bool("c", false, "количество строк")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение co строкой, не паттерн")
	lineNum := flag.Bool("n", false, "печатать номер строки")
	filePath := flag.String("fp", "", "указать абсолютный путь до файла")
	flag.Parse()

	findPhrase := os.Args[len(os.Args)-1]

	data, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Fatal(err)
	}
	tempStr := strings.Split(string(data), "\n")
	text := make([]string, 0, len(tempStr))

	if *ignoreCase {
		str := strings.ToLower(string(data))
		text = strings.Split(str, "\n")
		findPhrase = strings.ToLower(findPhrase)
	} else {
		text = strings.Split(string(data), "\n")
	}

	arr := []int{}
	if *fixed {
		for ind, val := range text {
			if findPhrase == val {
				arr = append(arr, ind)
			}
		}
	} else {
		for ind, val := range text {
			check, err := regexp.MatchString(findPhrase, val)
			if err != nil {
				log.Fatal(err)
			}
			if check {
				arr = append(arr, ind)
			}
		}
	}

	if *before > 0 {
		for _, v := range arr {
			if *before > v {
				fmt.Println(strings.Join(text[:v+1], "\n")) //саму строку печатаем
				fmt.Println()
				continue
			}
			fmt.Println(strings.Join(text[v-*before:v+1], "\n")) // печатаем до строки + саму строку
			fmt.Println()
		}
	}

	if *after > 0 {
		for _, v := range arr {
			if len(text[v:])-1 < *after {
				fmt.Println(strings.Join(text[v:], "\n")) // печатаем саму строку
				fmt.Println()
				continue
			}
			fmt.Println(strings.Join(text[v:v+*after+1], "\n")) // печатаем сама + после
			fmt.Println()
		}
	}
	// after + before
	if *contextText > 0 {
		for _, v := range arr {
			if *contextText > v {
				fmt.Println(strings.Join(text[:v+1], "\n"))
			} else {
				fmt.Println(strings.Join(text[v-*contextText:v+1], "\n"))
			}
			if len(text)-1 < *contextText {
				fmt.Println(strings.Join(text[v+1:], "\n"))
				fmt.Println()
				continue
			}
			fmt.Println(strings.Join(text[v+1:v+*contextText+1], "\n"))
			fmt.Println()
		}

	}
	if *countBool {
		fmt.Println("Найдено строк - ", len(arr))
	}

	if *invert {
		// выводим строки до первого найденного элемента
		if arr[0] != 0 {
			fmt.Println(strings.Join(text[:arr[0]], "\n"))
		}
		for i, v := range arr {
			// пропуск найденного элекмента
			if arr[0] == v {
				continue
			}
			// если разница между индексам 1, то тоже пропускаем
			if arr[i-1]-arr[i] == 1 {
				continue
			}
			// выводим строки между найденными
			fmt.Println(strings.Join(text[arr[i-1]+1:arr[i]], "\n"))
		}
		// если последний элемент не равен, количеству всех строк, выводим остальное
		if arr[len(arr)-1] != len(text)-1 {
			fmt.Println(strings.Join(text[arr[len(arr)-1]+1:], "\n"))
		}
	}
	if *lineNum {
		for _, val := range arr {
			fmt.Println(val)
		}
	}
	//Чтобы не печатло 2-й раз, если есть совпадения
	if *after > 0 || *before > 0 || *contextText > 0 || *countBool || *lineNum {
		return
	}
	for _, val := range arr {
		fmt.Println(text[val])
	}

}
