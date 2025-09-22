package strings

//Given two non-negative integers, num1 and num2 represented as string,
// return the sum of num1 and num2 as a string.

// You must solve the problem without using any built-in library for
// handling large integers (such as BigInteger). You must also not
// convert the inputs to integers directly.

// O(1)
func AddStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	result := ""

	for i >= 0 || j >= 0 || carry > 0 {
		var digit1, digit2 int

		if i >= 0 {
			digit1 = int(num1[i] - '0')
			i--
		}

		if j >= 0 {
			digit2 = int(num2[j] - '0')
			j--
		}

		sum := digit1 + digit2 + carry
		carry = sum / 10
		result = string(rune(sum%10+'0')) + result
	}

	return result
}
