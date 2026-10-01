package train

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 股票问题1: 只能做一次交易，手上最多持有一只股票，获取的最大收益
// 最低点买入，最高点卖出
func maxProfit1(prices []int) int {
	ans := 0
	minVal := prices[0]
	for i := 0; i < len(prices); i++ {
		minVal = Min(minVal, prices[i])
		ans = Max(ans, prices[i]-minVal)
	}
	return ans
}

// 股票问题2: 可以做无限次交易，手上最多持有一只股票，求能获取的最大收益
func maxProfit2(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans += Max(prices[i]-prices[i-1], 0)
	}
	return ans
}

// 股票问题3: 最多可以左2笔交易，手上最多持有一只股票，求能获取的最大收益
func maxProfit3(prices []int) int {
	// 完成一次交易并买入第二支股票的最大收益
	doneOneMinusBuyMax := -prices[0]
	// 完成一次交易的最大收益
	doneOneMax := 0
	// [0,i-1]的最小值
	minVal := prices[0]
	// 最大收益
	ans := 0
	for i := 0; i < len(prices); i++ {
		ans = Max(ans, doneOneMinusBuyMax+prices[i])
		minVal = Min(minVal, prices[i])
		doneOneMax = Max(doneOneMax, prices[i]-minVal)
		// 1. i位置不买入
		// 2. i位置买入
		doneOneMinusBuyMax = Max(doneOneMinusBuyMax, doneOneMax-prices[i])
	}
	return ans
}

// 股票问题4: 最大可以交易K次，手上最多持有一只股票，求能获取的最大收益
func maxProfit4(k int, prices []int) int {
	// i位置必须是最后一次交易的卖出时机
	N := len(prices)
	// 优化
	if k >= N/2 {
		return maxProfit2(prices)
	}
	dp := make([][]int, N)
	for index := range dp {
		dp[index] = make([]int, k+1)
	}
	// dp[0][k] == 0
	// dp[i][0] == 0
	for i := 1; i < N; i++ {
		for j := 1; j <= k; j++ {
			// i位置不参与交易
			dp[i][j] = dp[i-1][j]
			// i位置参与交易, 枚举在p位置买入，在i位置卖出
			for p := 0; p <= i; p++ {
				dp[i][j] = Max(dp[i][j], dp[p][j-1]+prices[i]-prices[p])
			}
		}
	}
	return dp[N-1][k]
}

// 股票问题4的最优解
func maxProfit4Max(k int, prices []int) int {
	// i位置必须是最后一次交易的卖出时机
	N := len(prices)
	// 优化
	if k >= N/2 {
		return maxProfit2(prices)
	}
	dp := make([][]int, k+1)
	for index := range dp {
		dp[index] = make([]int, N)
	}
	ans := 0
	for j := 1; j <= k; j++ {
		pre := dp[j][0]
		best := pre - prices[0]
		for i := 1; i < N; i++ {
			pre = dp[j-1][i]
			dp[j][i] = Max(dp[j][i-1], prices[i]+best)
			best = Max(best, pre-prices[i])
			ans = Max(ans, dp[j][i])
		}
	}
	return ans
}

// 股票问题5：可以做无限次交易，但是卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)，手上最多持有一只股票，求能获取的最大收益
// 最优尝试如下：
// buy[i] : 在0...i范围上，最后一次操作是buy动作，
// 这最后一次操作有可能发生在i位置，也可能发生在i之前
// buy[i]值的含义是：max{ 所有可能性[之前交易获得的最大收益 - 最后buy动作的收购价格] }
// 比如：arr[0...i]假设为[1,3,4,6,2,7,1...i之后的数字不用管]
// 什么叫，所有可能性[之前交易获得的最大收益 - 最后buy动作的收购价格]？
// 比如其中一种可能性：
// 假设最后一次buy动作发生在2这个数的时候，那么之前的交易只能在[1,3,4]上结束，因为6要cooldown的，
// 此时最大收益是多少呢？是4-1==3。那么，之前交易获得的最大收益 - 最后buy动作的收购价格 = 3 - 2 = 1
// 另一种可能性：
// 再比如最后一次buy动作发生在最后的1这个数的时候，那么之前的交易只能在[1,3,4,6,2]上发生，因为7要cooldown的，
// 此时最大收益是多少呢？是6-1==5。那么，之前交易获得的最大收益 - 最后buy动作的收购价格 = 5 - 1 = 4
// 除了上面两种可能性之外，还有很多可能性，你可以假设每个数字都是最后buy动作的时候，
// 所有可能性中，(之前交易获得的最大收益 - 最后buy动作的收购价格)的最大值，就是buy[i]的含义
// 为啥buy[i]要算之前的最大收益 - 最后一次收购价格？尤其是最后为什么要减这么一下？
// 因为这样一来，当你之后以X价格做成一笔交易的时候，当前最好的总收益直接就是 X + buy[i]了
//
// sell[i] :0...i范围上，最后一次操作是sell动作，这最后一次操作有可能发生在i位置，也可能发生在之前
// sell[i]值的含义：0...i范围上，最后一次动作是sell的情况下，最好的收益
//
// 于是通过分析，能得到以下的转移方程：
// buy[i] = Math.max(buy[i - 1], sell[i - 2] - prices[i])
// 如果i位置没有发生buy行为，说明有没有i位置都一样，那么buy[i] = buy[i-1]，这显而易见
// 如果i位置发生了buy行为, 那么buy[i] = sell[i - 2] - prices[i]，
// 因为你想在i位置买的话，你必须保证之前交易行为发生在0...i-2上，
// 因为如果i-1位置有可能参与交易的话，i位置就要cooldown了，
// 而且之前交易行为必须以sell结束，你才能buy，而且要保证之前交易尽可能得到最好的利润，
// 这正好是sell[i - 2]所代表的含义，并且根据buy[i]的定义，最后一定要 - prices[i]
//
// sell[i] = Math.max(sell[i - 1], buy[i - 1] + prices[i])
// 如果i位置没有发生sell行为，那么sell[i] = sell[i-1]，这显而易见
// 如果i位置发生了sell行为，那么我们一定要找到 {之前获得尽可能好的收益 - 最后一次的收购价格尽可能低}，
// 而这正好是buy[i - 1]的含义！之前所有的"尽可能"中，最好的一个！
func maxProfit5(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	// 在i位置决定买 / 卖 / 不参与
	N := len(prices)
	// 买数组
	buy := make([]int, N)
	// 卖数组
	sell := make([]int, N)
	buy[1] = Max(-prices[0], -prices[1])
	sell[1] = Max(0, prices[1]-prices[0])
	for i := 2; i < N; i++ {
		// 在i位置要买入
		buy[i] = Max(buy[i-1], sell[i-2]-prices[i])
		// 在i位置要卖出
		sell[i] = Max(sell[i-1], buy[i-1]+prices[i])
	}
	return sell[N-1]
}

// 股票问题5最优解(空间上的优化)
func maxProfit5Max(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	// 在i位置决定买 / 卖 / 不参与
	N := len(prices)
	buy1 := Max(-prices[0], -prices[1])
	sell1 := Max(0, prices[1]-prices[0])
	sell2 := 0
	for i := 2; i < N; i++ {
		tmp := sell1
		// 在i位置要买入
		buy1 = Max(buy1, sell2-prices[i])
		// 在i位置要卖出
		sell1 = Max(sell1, buy1+prices[i])
		sell2 = tmp
	}
	return sell1
}

// 股票问题6: 可以做无限次交易，但是每笔交易需要手续费，手上最多持有一只股票，求能获取的最大收益
func maxProfit6(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	N := len(prices)
	// 持有股票
	hold := make([]int, N)
	// 不持有股票
	notHold := make([]int, N)
	hold[0] = -prices[0]
	notHold[0] = 0

	for i := 1; i < N; i++ {
		// hold[i] 持有一只股票
		// 1. i位置不参与
		// 2. 第i天要持有股票，i-1天没有持有股票的最大收益 - 第i天买入价格
		hold[i] = Max(hold[i-1], notHold[i-1]-prices[i])
		// 第i天不持有股票，i-1天持有股票的最大收益 + 第i天的卖出价格 - 手续费
		notHold[i] = Max(notHold[i-1], hold[i-1]+prices[i]-fee)
	}
	return notHold[N-1]
}

// 股票问题6最优解
func maxProfit6Max(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	N := len(prices)
	hold := -prices[0]
	notHold := 0
	for i := 1; i < N; i++ {
		// hold[i] 持有一只股票
		// 1. i位置不参与
		// 2. 第i天要持有股票，i-1天没有持有股票的最大收益 - 第i天买入价格
		tmp := hold
		hold = Max(hold, notHold-prices[i])
		// 第i天不持有股票，i-1天持有股票的最大收益 + 第i天的卖出价格 - 手续费
		notHold = Max(notHold, tmp+prices[i]-fee)
	}
	return notHold
}
