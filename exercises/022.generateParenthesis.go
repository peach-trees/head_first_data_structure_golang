package exercises

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	dfs022(0, 0, n, "", &ret)
	return ret
}

// 回溯算法，回溯跳出条件就是左右括号都已经排完的情况.
// 括号成对存在，先有左括号再有右括号，所以只有右括号的数量小于左括号才进行右括号的添加.
// 最后如果右括号的数量等于0，表示右括号已经排完了，同时意味着左括号也排完了.
func dfs022(leftParenthesisCount, rightParenthesisCount, n int, path string, ret *[]string) {
	if leftParenthesisCount == n && rightParenthesisCount == n {
		*ret = append(*ret, path)
		return
	}
	// 生成左括号
	if leftParenthesisCount < n {
		dfs022(leftParenthesisCount+1, rightParenthesisCount, n, path+"(", ret)
	}
	// 括号成对存在，有左括号才会有右括号
	if rightParenthesisCount < n && rightParenthesisCount < leftParenthesisCount {
		dfs022(leftParenthesisCount, rightParenthesisCount+1, n, path+")", ret)
	}
}
