package base62

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base = len(alphabet)
)


/* По заданию будем просто переводить циферку в алфавит base62:
Короткие ссылки должны основываться на id записи(sequence) в БД,
переведённой в систему счисления с алфавитом [A-Za-z0-9]
Вообще, можно придумать (взять из интернета) более сложные алгоритмы, но не будем
отклоняться от задания
*/


func reverse(s string) (result string) {
	for _,v := range s {
		result = string(v) + result
	}
	return
}

func ToAlphabet(decimal int, alphabet string) string {
	base := len(alphabet)
	result := ""
	if decimal == 0 {
		result = "a"
	}
	for decimal > 0 {
		pos := decimal % base
		result += string(alphabet[pos])
		decimal /= base
	}
	return reverse(result)
}

//func main() {
//	num := 63
//	fmt.Println(toAlphabet(num, alphabet))
//}