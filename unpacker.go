package unpacker

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

type Compressed struct {
	pos     int    //позиция
	integer string // десятичное число
	letter  string // символ который нужно повторить (перед числом)
}

func FileData(name string) string {
	data, err := ioutil.ReadFile(name) // data считывает файл
	if err != nil {
		fmt.Println("ошибка файла\n", err)
		fmt.Println(data)
	}

	return string(data)
}

func StringToPos(x string) []Compressed { // функция возвращает поля структуры "Compressed" в виде массива
	var integer string         // целое число
	var positions []Compressed // переменная с типом массив структуры
	lastpos := 0
	for pos, char := range x { // идем циклом по ячейкам и их содержимому(байты) строки "х"

		if unicode.IsDigit(char) { // проверка: если в ячейке десятичное число
			integer += string(char) // десятичные числа в виде байтов помещаются в новую строку "integer" друг за другом, в том количестве, сколько содержится в строке(массиве) "х"
			lastpos = pos
			// и переменной "lastpos" типа int присваивается номер ячейки в которой было записано десятичное число
		} else { // во всех ячейках кроме тех, в которых записаны десятичные числа

			if char != 7 { // если в ячейке не ?7?   не понимаю условие  (7 это 55)

				if pos == lastpos+1 { // и если номер ячейки равен следующей после той в которой было десятичное число

					if len(integer) > 0 { // и если строка "integer" в которой в виде байтов помещены десятичные числа больше 0
						positions = append(positions, Compressed{pos - len(integer), integer, string(char)}) /* в структуру, в виде одной ячейки массива записываются:
						{(индекс ячейки в которой было десятич.число) (само число) и (символ перед которым было это число)}
						*/
						// "обнуление" переменной для того, чтобы предыдущее десятичное число не добавлялось в следущую ячейку
						integer = ""
					}
				}
			}
		}
	}
	return positions
}

func FileUnpack(x string) string {
	var positions []Compressed
	positions = StringToPos(x) // массив который получился на выходе в функции "StringToPos(x string)"
	var needle string          // новая переменная стринг
	times := 0
	if len(positions) > 0 {
		for i := range positions { // циклом проходимся по массиву - ({ячейка} {число} {символ перед которым это число})
			needle = positions[i].integer // к переменной присваивается значение {число} из массива в виде строки (байтов)
			needle += positions[i].letter // к этой же переменной прибавляется строка - {символ перед которым это число}, получается запись "2a3b4d" и т.д
			//fmt.Println(string(needle))
			times, _ = strconv.Atoi(positions[i].integer) // перевод из байт в integer ячеек {число} и запись в переменную
			//fmt.Println(times)
			x = strings.Replace(x, needle, strings.Repeat(positions[i].letter, times), 1) /* Replace меняет, например запись "2a" на "аа"
			с помощью Repeat в котором в качестве аргументов служат 3 аргумента:
			1) строка в которой надо произвести повтор "positions[i]", 2) что нужно повторить ".letter", 3) сколько раз повторить "times"
			*/
		}
		return x
	}
	fmt.Println(positions)
	return x
}
