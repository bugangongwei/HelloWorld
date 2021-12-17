package code

/*
位运算
*/

/*
1720. 解码异或后的数组
encoded = [1,2,3], first = 1

假如 z = x^y
x = 0 ^ x = (y ^ y) ^ x = y ^ (y ^ x) = y ^ z
推出: x = y^z

encoded[i] = ori[i] ^ ori[i+1]
ori[i+1] = encoded[i] ^ ori[i]
*/
func Decode(encoded []int, first int) []int {
	var ori = make([]int, len(encoded)+1)
	ori[0] = first

	for i := 0; i < len(encoded); i++ {
		ori[i+1] = encoded[i] ^ ori[i]
	}

	return ori
}

/*
1734. 解码异或后的排列
ori = [A,B,C,D,E]
encoded = [AB,BC,CD,DE]
ABCDE=A(BC^DE)
A = ABCDE(BC^DE)
*/
func Decode2(encoded []int) []int {
	// 计算 first
	var (
		all  = 0
		nall = 0
		ori  = make([]int, len(encoded)+1)
	)
	for i := 1; i <= len(encoded)+1; i++ {
		// fmt.Println(all)
		all ^= i
	}

	for i := 1; i < len(encoded); i = i + 2 {
		nall ^= encoded[i]
	}

	// 计算 ori
	ori[0] = all ^ nall
	for i := 0; i < len(encoded); i++ {
		ori[i+1] = encoded[i] ^ ori[i]
	}

	return ori
}
