## 1. 定义DP数组含义
```text
dp[i] 代表以s[i]结尾的最长递增子序列长度
```

## 2. 定义状态转移方程
```text
首先根据定义, dp[i-1] 代表以s[i-1]结尾的最长递增子序列长度
那么我们以 [4,10,4,3,8,9] 为例, 比如我想求8所在的位置的dp值, 则我们需要找到上一个比8小的数的dp值然后+1


则伪代码为:
        max := -1
        for i :=0; i < j;i++ {
            if num[i] > num[j+1] {
                if dp[i] > max {
                    max = dp[i]
                }
            }
        }
```

## 3.dp数组的初始化
```text
    此时看状态转移方程, dp[i]只与dp[i-1]相关, 那么我们只需要初始化第一个数就行
    如果字符串只有一个字符, 则最长的递增子序列为1, dp[0]=1
```

## 4. 数组的遍历顺序
```text
由于我们的dp数组是一维的, 要么从前往后, 要么从后往前, 此处我们选择从前往后
代码如下:

//求传入参数的最大值
func max(x ...int) int {
	sort.Ints(x)
	return x[len(x)-1]
}

func dp(nums []int) []int {
	dpTable := make([]int, len(nums))
	dpTable[0] = 1

	for i := 1; i < len(nums); i++ {
		v := -1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] { //获取比当前这个数小的数
				if dpTable[j] > v { //获取其中的最大值
					v = dpTable[j]
				}
			}
		}

		if v == -1 {  //v == -1 代表没有比num[i] 最小的数了, 比如 [2,3,4,5,6,1] 遍历到1时, 就找不到比1小的数,则dp值为1
			dpTable[i] = 1
		} else {      //说明找了, +1
			dpTable[i] = v + 1 
		}
	}

	fmt.Println(dpTable)
	return dpTable
}

func lengthOfLIS(nums []int) int {
	t := dp(nums)
	return max(t...)
}

```