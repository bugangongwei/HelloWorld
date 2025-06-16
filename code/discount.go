package code

import (
	"fmt"
	"sort"
)

/*
// 原始代码（订单价格计算）
public double calculate(Order order) {
    double discount = 0;
    if (order.user.isVIP()) discount += 0.1;
    if (order.items.size() > 10) discount += 0.05;
    if (order.total > 1000) discount += 0.2;
    return order.total * (1 - discount);
}
用策略模式重构折扣逻辑
支持动态添加折扣规则
避免折扣叠加冲突

1. 互斥组机制
每个折扣策略属于一个互斥组（GroupID）
GroupID=0：无互斥组，可与任何策略叠加
GroupID>0：互斥组，同一组内策略互斥
互斥组管理器确保组ID唯一性

2. 优先级系统
每个策略有优先级（Priority）
当同一互斥组内多个策略适用时：
按优先级排序（高优先级优先）
只应用优先级最高的策略

3. 折扣计算流程
收集适用策略：
过滤出当前订单适用的策略
分为互斥组策略和非互斥组策略
处理互斥组：
对每个互斥组，选择优先级最高的策略
同一组内其他策略被忽略
叠加折扣：
所有被选中的策略折扣相加
确保总折扣不超过100%
*/

// 用户类型
type User struct {
	IsVIP     bool
	IsNewUser bool
}

// 订单项
type Item struct {
	Name  string
	Price float64
}

// 订单
type Order struct {
	name  string // 订单名称
	User  User
	Items []Item
	Total float64
}

// 折扣策略接口
type DiscountStrategy interface {
	CalculateDiscount(order Order) float64 // 计算折扣
	Description() string                   // 描述折扣规则
	GroupID() int                          // 分组ID，用于避免冲突，0表示没有冲突
	Priority() float64                     // 优先级，数字越小优先级越高
}

// 互斥组管理器
type MutexGroupManager struct {
	groups map[int]struct{} // 已使用的互斥组ID集合
}

func NewMutexGroupManager() *MutexGroupManager {
	return &MutexGroupManager{
		groups: make(map[int]struct{}),
	}
}

// 检查互斥组ID是否已存在
func (m *MutexGroupManager) IsGroupExists(groupID int) bool {
	_, exists := m.groups[groupID]
	return exists
}

// 添加互斥组
func (m *MutexGroupManager) AddGroup(groupID int) {
	if groupID > 0 {
		m.groups[groupID] = struct{}{}
	}
}

// VIP 折扣策略
type VIPDiscount struct {
	groupID int
}

func NewVIPDiscount(groupID int) *VIPDiscount {
	return &VIPDiscount{groupID: groupID}
}

func (s *VIPDiscount) CalculateDiscount(order Order) float64 {
	if order.User.IsVIP {
		return 0.1 // 10% 折扣
	}
	return 0.0
}

func (s *VIPDiscount) Description() string {
	return "VIP 折扣 (10%)"
}

func (s *VIPDiscount) GroupID() int {
	return s.groupID
}

func (s *VIPDiscount) Priority() float64 {
	return 0.1
}

// 全场满减折扣
type FullDiscount struct {
	groupID int
}

func NewFullDiscount(groupID int) *FullDiscount {
	return &FullDiscount{groupID: groupID}
}
func (s *FullDiscount) CalculateDiscount(order Order) float64 {
	if order.Total > 500 {
		return 0.1 // 满500元减10%
	}
	return 0.0
}
func (s *FullDiscount) Description() string {
	return "全场满减折扣 (满500元减10%)"
}
func (s *FullDiscount) GroupID() int {
	return s.groupID
}
func (s *FullDiscount) Priority() float64 {
	return 0.1 // 优先级最低
}

// 大宗订单折扣策略
type BulkOrderDiscount struct {
	groupID int
}

func NewBulkOrderDiscount(groupID int) *BulkOrderDiscount {
	return &BulkOrderDiscount{groupID: groupID}
}

func (s *BulkOrderDiscount) CalculateDiscount(order Order) float64 {
	if len(order.Items) > 5 {
		return 0.05 // 5% 折扣
	}
	return 0.0
}

func (s *BulkOrderDiscount) Description() string {
	return "大宗订单折扣 (5%)"
}

func (s *BulkOrderDiscount) GroupID() int {
	return s.groupID
}

func (s *BulkOrderDiscount) Priority() float64 {
	return 0.05
}

// 高额订单折扣策略
type HighValueDiscount struct {
	groupID int
}

func NewHighValueDiscount(groupID int) *HighValueDiscount {
	return &HighValueDiscount{groupID: groupID}
}

func (s *HighValueDiscount) CalculateDiscount(order Order) float64 {
	if order.Total > 1000 {
		return 0.2 // 20% 折扣
	}
	return 0.0
}

func (s *HighValueDiscount) Description() string {
	return "高额订单折扣 (20%)"
}

func (s *HighValueDiscount) GroupID() int {
	return s.groupID
}

func (s *HighValueDiscount) Priority() float64 {
	return 0.2
}

// 新用户折扣策略
type NewUserDiscount struct {
	groupID int
}

func NewNewUserDiscount(groupID int) *NewUserDiscount {
	return &NewUserDiscount{groupID: groupID}
}

func (s *NewUserDiscount) CalculateDiscount(order Order) float64 {
	if order.User.IsNewUser {
		return 0.15 // 15% 折扣
	}
	return 0.0
}

func (s *NewUserDiscount) Description() string {
	return "新用户折扣 (15%)"
}

func (s *NewUserDiscount) GroupID() int {
	return s.groupID
}

func (s *NewUserDiscount) Priority() float64 {
	return 0.15
}

// 折扣计算器
type DiscountCalculator struct {
	// Strategies   []DiscountStrategy
	// groupManager *MutexGroupManager
	nonMutexStrategies []DiscountStrategy         // 非互斥组策略
	mutexStrategies    map[int][]DiscountStrategy // 互斥组策略
}

func NewDiscountCalculator() *DiscountCalculator {
	return &DiscountCalculator{
		// groupManager: NewMutexGroupManager(),
		nonMutexStrategies: make([]DiscountStrategy, 0),
		mutexStrategies:    make(map[int][]DiscountStrategy),
	}
}

// 添加折扣策略（带互斥组检查）
func (dc *DiscountCalculator) AddStrategy(strategy DiscountStrategy) error {
	groupID := strategy.GroupID()

	// 检查互斥组ID是否已存在
	if groupID > 0 {
		dc.mutexStrategies[groupID] = append(dc.mutexStrategies[groupID], strategy)
	} else {
		dc.nonMutexStrategies = append(dc.nonMutexStrategies, strategy)
	}

	return nil
}

// 计算最终价格
func (dc *DiscountCalculator) Calculate(order Order) float64 {

	// 找到符合条件的折扣
	nonMutexStrategies4Order := make([]DiscountStrategy, 0)
	for _, strategy := range dc.nonMutexStrategies {
		if strategy.CalculateDiscount(order) > 0 {
			nonMutexStrategies4Order = append(nonMutexStrategies4Order, strategy)
		}
	}

	mutexStrategies4Order := make(map[int][]DiscountStrategy)
	for groupID, strategies := range dc.mutexStrategies {
		selected := make([]DiscountStrategy, 0)
		for _, strategy := range strategies {
			if strategy.CalculateDiscount(order) > 0 {
				selected = append(selected, strategy)
			}
		}

		mutexStrategies4Order[groupID] = selected
	}

	// 处理互斥组：每个组选择优先级最高的策略
	appliedStrategies := []DiscountStrategy{}
	totalDiscount := 0.0
	// 处理互斥组折扣
	for _, strategies := range mutexStrategies4Order {
		if len(strategies) == 0 {
			continue
		}

		// 按折扣力度（折扣力度高的在前）
		sort.Slice(strategies, func(i, j int) bool {
			return strategies[i].Priority() > strategies[j].Priority()
		})

		// 选择组内优先级最高的策略
		selected := strategies[0]
		appliedStrategies = append(appliedStrategies, selected)
		totalDiscount += selected.CalculateDiscount(order)
	}

	// 添加非互斥组折扣
	for _, strategy := range nonMutexStrategies4Order {
		appliedStrategies = append(appliedStrategies, strategy)
		totalDiscount += strategy.CalculateDiscount(order)
	}

	// 确保总折扣不超过100%
	if totalDiscount > 1.0 {
		totalDiscount = 1.0
	}

	// 打印折扣详情
	dc.printDiscountDetails(order, totalDiscount)

	return order.Total * (1 - totalDiscount)
}

// 打印折扣详情
func (dc *DiscountCalculator) printDiscountDetails(order Order, totalDiscount float64) {
	fmt.Printf("\n[%s]订单总金额: $%.2f\n", order.name, order.Total)

	strategies := make([]DiscountStrategy, 0)
	copy(strategies, dc.nonMutexStrategies)
	for _, s := range dc.mutexStrategies {
		strategies = append(strategies, s[0])
	}

	if len(strategies) > 0 {
		fmt.Println("应用的折扣策略:")

		// 按描述排序以便更好展示
		sort.Slice(strategies, func(i, j int) bool {
			return strategies[i].Description() < strategies[j].Description()
		})

		for _, s := range strategies {
			groupInfo := ""
			if s.GroupID() > 0 {
				groupInfo = fmt.Sprintf(" [互斥组:%d]", s.GroupID())
			}
			fmt.Printf("  - %s (优先级:%f)%s\n", s.Description(), s.Priority(), groupInfo)
		}
	} else {
		fmt.Println("无适用折扣")
	}

	fmt.Printf("总折扣: %.0f%%\n", totalDiscount*100)
}
