package Number_Line_Jumps

// https://www.hackerrank.com/challenges/kangaroo/problem

func Kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if (x1 >= x2 && v1 >= v2) || (x1 <= x2 && v1 <= v2) {
		return "NO"
	} else {
		if ((x1 - x2) % (v1 - v2)) == 0 {
			return "YES"
		}
	}
	return "NO"
}
