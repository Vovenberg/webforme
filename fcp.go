package main;

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
	"bufio"
	"golang.design/x/clipboard"
	"time"
)

func main() {	
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	err = clipboard.Init()
	if err != nil {
		panic(err)
	}
	
	scanner := bufio.NewScanner(file) 
    scanner.Split(bufio.ScanLines)
    var text []string
    for scanner.Scan() {
        text = append(text, scanner.Text())
    }
	for _, each_ln := range text {
        fmt.Println(each_ln)
    }

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()


	pointer := 0
	auto := false
	
	copy(text[0], 0)

	go func(msg string) {
		for {
			if auto {
				size := len(text)
				if pointer < size-1 {
					pointer++
				}
				copy(text[pointer], pointer)

				fmt.Println("Сплю 5 секунд и копирую следующую строчку")
				fmt.Println()
				time.Sleep(5 * time.Second)
			}
		}
    }("going")
	
	fmt.Println("Нажмите ESC чтобы выйти")
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
        if key == keyboard.KeyEsc {
			break
		}
        if key == keyboard.KeyEnter {
			if auto {
				auto = false			
				fmt.Println("--------Выключено автокопирование-------")
			} else {
				auto = true
				fmt.Println("--------Включено автокопирование-------")
			}
		}
		if key == keyboard.KeyArrowRight {
			size := len(text)
			if pointer < size-1 {
				pointer++
			}
			var current = text[pointer]
			copy(current, pointer)
		}
		if key == keyboard.KeyArrowLeft {
			if pointer > 0 {
				pointer--
			}
			var current = text[pointer]
			copy(current, pointer)
		}
		
	}	
}

func copy(text string, line int){
	clipboard.Write(clipboard.FmtText, []byte(text))
	fmt.Println()
	fmt.Printf("Скопирована строчка %d: %s \n", line + 1, text)
	fmt.Println()
}