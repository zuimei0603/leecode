## 直线相关的题目
### 1. 直线上最多的点数
给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上。

输入: [[1,1],[2,2],[3,3]]
输出: 3

输入: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
输出: 4

**解题四路：**
暴力法解题，需要注意起点重复的情况，起点重复也算在同一条直线上

**代码**
```
func maxPoints(points [][]int) int {
	length := len(points)
	if length <= 2 {
		return length
	}
	max := 0
	for i := 0; i < length-2; i++ {
		same := 1
		for j := i + 1; j < length; j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				same++
				continue
			}
			cnt := same + 1
			for k := j + 1; k < length; k++ {
				if (points[j][1]-points[i][1])*(points[k][0]-points[i][0]) == (points[k][1]-points[i][1])*(points[j][0]-points[i][0]) {
					cnt++
				}
			}
			max = int(math.Max(float64(max), math.Max(float64(same), float64(cnt))))
		}
		if max <= 0 {
			max = same
		}
	}
	return max
}
```
