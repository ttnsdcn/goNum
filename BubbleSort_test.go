// BubbleSort_test
/*
------------------------------------------------------
作者   : Black Ghost
日期   : 2019-03-05
版本   : 0.0.0
------------------------------------------------------
    冒泡排序法
理论：
    时间复杂度: O(n^2)
    最好情况  : O(n)
    最坏情况  : O(n^2)
    空间复杂度: O(1)
    稳定性    : 稳定
------------------------------------------------------
输入   :
    in      输入矩阵, 1xn
输出   :
    sol     排序结果
    err     解出标志：false-未解出或达到步数上限；
                     true-全部解出
------------------------------------------------------
*/

package goNum_test

import (
	"testing"

	"github.com/chfenger/goNum"
)

// BubbleSort 冒泡排序法
func BubbleSort(in goNum.Matrix) (goNum.Matrix, bool) {
	/*
	      冒泡排序法
	   输入   :
	       in      输入矩阵, 1xn
	   输出   :
	       sol     排序结果
	       err     解出标志：false-未解出或达到步数上限；
	                        true-全部解出
	*/
	//判断初值维数
	if in.Rows != 1 {
		panic("Error in goNum.BubbleSort: Input Matrix error")
	}
	if in.Columns < 1 {
		panic("Error in goNum.BubbleSort: Empty input Matrix")
	} else if in.Columns == 1 {
		return in, true
	}

	n := in.Columns
	sol := goNum.ZeroMatrix(1, n)
	var err bool = false

	//初始化sol
	for i := 0; i < n; i++ {
		sol.Data[i] = in.Data[i]
	}
	//排序开始
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sol.Data[j] > sol.Data[j+1] {
				sol.Data[j], sol.Data[j+1] = sol.Data[j+1], sol.Data[j]
			}
		}
	}

	err = true
	return sol, err
}

func BenchmarkBubbleSort(b *testing.B) {
	x58 := goNum.NewMatrix(1, 5, []float64{5.0, 3.2, 2.4, 1.8, 0.1})
	for i := 0; i < b.N; i++ {
		goNum.BubbleSort(x58)
	}
}
