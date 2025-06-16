package code

import "testing"

func TestDiscountCalculator(t *testing.T) {
	oldVipUser := User{IsVIP: true, IsNewUser: false}
	newVipUser := User{IsVIP: true, IsNewUser: true}
	// newUser := User{IsVIP: false, IsNewUser: true}
	// oldUser := User{IsVIP: false, IsNewUser: false}

	regular := createTestRegular()

	testCases := []struct {
		name             string
		order            Order
		regular          *DiscountCalculator
		expectedDiscount float64
	}{
		{
			name: "VIP用户，小额订单",
			order: Order{
				name: "VIP小额订单",
				User: oldVipUser, // 10%
				Items: []Item{
					{Name: "Item1", Price: 500},
					{Name: "Item2", Price: 300},
				},
				Total: 800,
			},
			regular:          regular,
			expectedDiscount: 0.8 * 800, // 10% VIP + 10% 全场满减
		},
		{
			name: "VIP用户，大额订单",
			order: Order{
				name: "VIP大额订单",
				User: oldVipUser, // 10%
				Items: []Item{
					{Name: "Item1", Price: 1000},
					{Name: "Item2", Price: 800},
					{Name: "Item3", Price: 400},
					{Name: "Item4", Price: 200},
					{Name: "Item5", Price: 1400},
					{Name: "Item6", Price: 600},
				},
				Total: 4400,
			},
			regular:          regular,
			expectedDiscount: 0.6 * 4400, // 10% VIP + 10% 全场满减 + 20% 高额订单
		},
		{
			name: "VIP新用户，大额订单",
			order: Order{
				name: "VIP新用户大额订单",
				User: newVipUser, // 15%
				Items: []Item{
					{Name: "Item1", Price: 1000},
					{Name: "Item2", Price: 800},
					{Name: "Item3", Price: 400},
					{Name: "Item4", Price: 200},
					{Name: "Item5", Price: 1400},
					{Name: "Item6", Price: 600},
				},
				Total: 4400,
			},
			regular:          regular,
			expectedDiscount: 0.55 * 4400, // 15% 新用户 + 10% 全场满减 + 20% 高额订单
		},
		{
			name: "VIP新用户，微订单",
			order: Order{
				name: "VIP新用户微订单",
				User: newVipUser, // 15%
				Items: []Item{
					{Name: "Item1", Price: 10},
					{Name: "Item2", Price: 10},
					{Name: "Item3", Price: 10},
					{Name: "Item4", Price: 10},
					{Name: "Item5", Price: 20},
					{Name: "Item6", Price: 20},
				},
				Total: 80,
			},
			regular:          regular,
			expectedDiscount: 0.8 * 80, // 15% 新用户 + 5% 大宗订单
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			discount := c.regular.Calculate(c.order)
			if discount != c.expectedDiscount {
				t.Errorf("[%s] expected %f, got %f", c.name, c.expectedDiscount, discount)
			}
		})
	}
}

func createTestRegular() *DiscountCalculator {
	regular := NewDiscountCalculator()
	// 用户相关折扣
	regular.AddStrategy(NewNewUserDiscount(1)) // 新用户折扣 (15%)
	regular.AddStrategy(NewVIPDiscount(1))     // VIP 折扣 (10%)
	// 订单特征折扣
	regular.AddStrategy(NewBulkOrderDiscount(2)) // 大宗订单折扣(超过5单5%)
	regular.AddStrategy(NewHighValueDiscount(2)) // 高额订单折扣 (20%)1000
	// 全场满减折扣
	regular.AddStrategy(NewFullDiscount(0)) // 全场满减折扣,满500元减10%

	return regular
}
