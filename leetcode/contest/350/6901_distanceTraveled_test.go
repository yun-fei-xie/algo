package _50

import (
	"fmt"
	"testing"
)

/*


模拟问题：有点像经典的换啤酒问题(啤酒盖可以换啤酒)(华为笔试题)

*/

func distanceTraveled(mainTank int, additionalTank int) int {
	if mainTank < 5 {
		return mainTank * 10
	}
	var ans int
	for {
		//这一次无法用到5升
		if mainTank < 5 {
			ans += mainTank * 10
			return ans
		}
		mainTank = mainTank - 5
		ans += 5 * 10
		if additionalTank >= 1 {
			additionalTank--
			mainTank++
		}
	}
}

func TestDistanceTravel(t *testing.T) {
	fmt.Println(distanceTraveled(5, 10))
	fmt.Println(distanceTraveled(1, 2))
}
