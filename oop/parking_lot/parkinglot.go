package parkinglot

import (
	"fmt"
	"sort"
)

// ===== 车位类型 =====

// SpotSize 表示车位的大小类别。
type SpotSize int

const (
	MotorcycleSpot SpotSize = iota // 摩托车位
	CompactSpot                    // 紧凑车位
	LargeSpot                      // 大型车位
	numSpotSizes                   // 哨兵值，用于数组大小
)

// spotSizeForIndex 根据车位在行中的索引确定其类型。
//
//	摩托车位: [0, k/4)
//	紧凑车位: [k/4, k/4*3)
//	大型车位: [k/4*3, k)
func spotSizeForIndex(index, spotsPerRow int) SpotSize {
	if index < spotsPerRow/4 {
		return MotorcycleSpot
	}
	if index < spotsPerRow/4*3 {
		return CompactSpot
	}
	return LargeSpot
}

// ===== 停车位 =====

// ParkingSpot 表示停车场中的一个车位。
type ParkingSpot struct {
	Size    SpotSize
	Vehicle Vehicle // 为 nil 表示空闲
	Level   int
	Row     int
	Index   int
}

// IsAvailable 返回该车位是否空闲。
func (ps *ParkingSpot) IsAvailable() bool {
	return ps.Vehicle == nil
}

// ===== 空闲区间 =====

// FreeInterval 表示连续空闲大型车位的范围 [Start, End)（左闭右开）。
type FreeInterval struct {
	Start, End int
}

// Len 返回该区间包含的车位数量。
func (fi FreeInterval) Len() int { return fi.End - fi.Start }

// ===== 行 =====

// Row 维护一行车位及其高效查找结构。
//
// 核心数据结构：
//   - avail[车位类型] map[int]bool: O(1) 查找任意空闲车位
//   - freeIntervals []FreeInterval: 排序的空闲大型车位区间，用于巴士停车
type Row struct {
	spots       []*ParkingSpot
	spotsPerRow int

	// 按车位类型记录的空闲车位索引集合。单车位停车 O(1) 查找。
	avail [numSpotSizes]map[int]bool

	// 排序的空闲大型车位区间列表，用于高效的连续车位（巴士）停车。
	freeIntervals []FreeInterval
}

func newRow(level, row, spotsPerRow int) *Row {
	r := &Row{
		spots:       make([]*ParkingSpot, spotsPerRow),
		spotsPerRow: spotsPerRow,
	}

	for i := 0; i < int(numSpotSizes); i++ {
		r.avail[i] = make(map[int]bool)
	}

	for i := 0; i < spotsPerRow; i++ {
		size := spotSizeForIndex(i, spotsPerRow)
		r.spots[i] = &ParkingSpot{
			Size:  size,
			Level: level,
			Row:   row,
			Index: i,
		}
		r.avail[size][i] = true
	}

	// 初始化一个覆盖所有大型车位的空闲区间。
	largeStart := spotsPerRow / 4 * 3
	if largeStart < spotsPerRow {
		r.freeIntervals = []FreeInterval{{Start: largeStart, End: spotsPerRow}}
	}

	return r
}

// parkSpotOfType 在指定类型的车位中停入一辆车。
// 返回停入的车位，如果该类型无空闲车位则返回 nil。
//
// 时间复杂度: O(1)
func (r *Row) parkSpotOfType(v Vehicle, st SpotSize) *ParkingSpot {
	if len(r.avail[st]) == 0 {
		return nil
	}
	// 从 map 中取任意一个空闲车位索引 — O(1)。
	var idx int
	for idx = range r.avail[st] {
		break
	}
	spot := r.spots[idx]
	spot.Vehicle = v
	delete(r.avail[st], idx)

	// 如果占用了大型车位，需更新区间追踪。
	if st == LargeSpot {
		r.removeFromIntervals(idx)
	}
	return spot
}

// canFitBus 返回该行是否有至少 5 个连续空闲大型车位。
func (r *Row) canFitBus() bool {
	for _, interval := range r.freeIntervals {
		if interval.Len() >= 5 {
			return true
		}
	}
	return false
}

// parkBus 将巴士停入 5 个连续大型车位。
// 返回占用的车位列表，如果无法停入则返回 nil。
//
// 时间复杂度: O(I)，I = 空闲区间数量（通常远小于 k）
func (r *Row) parkBus(v Vehicle) []*ParkingSpot {
	for i, interval := range r.freeIntervals {
		if interval.Len() < 5 {
			continue
		}
		// 从该区间头部取 5 个车位。
		spots := make([]*ParkingSpot, 5)
		for j := 0; j < 5; j++ {
			idx := interval.Start + j
			r.spots[idx].Vehicle = v
			delete(r.avail[LargeSpot], idx)
			spots[j] = r.spots[idx]
		}
		// 缩减或移除该区间。
		if interval.Len() == 5 {
			r.freeIntervals = append(r.freeIntervals[:i], r.freeIntervals[i+1:]...)
		} else {
			r.freeIntervals[i].Start += 5
		}
		return spots
	}
	return nil
}

// freeSpot 将车位归还到空闲池。在取车时调用。
//
// 时间复杂度: 非大型车位 O(1)，大型车位 O(I)（区间合并）
func (r *Row) freeSpot(spot *ParkingSpot) {
	spot.Vehicle = nil
	r.avail[spot.Size][spot.Index] = true
	if spot.Size == LargeSpot {
		r.addToIntervals(spot.Index)
	}
}

// removeFromIntervals 从空闲区间列表中移除单个车位索引。
// 在任何车辆停入大型车位时调用。
func (r *Row) removeFromIntervals(idx int) {
	for i, interval := range r.freeIntervals {
		if idx < interval.Start || idx >= interval.End {
			continue
		}
		switch {
		case interval.Len() == 1:
			// 移除整个区间。
			r.freeIntervals = append(r.freeIntervals[:i], r.freeIntervals[i+1:]...)
		case idx == interval.Start:
			r.freeIntervals[i].Start++
		case idx == interval.End-1:
			r.freeIntervals[i].End--
		default:
			// 分裂为 [Start, idx) 和 [idx+1, End) 两个区间。
			right := FreeInterval{Start: idx + 1, End: interval.End}
			r.freeIntervals[i].End = idx
			r.freeIntervals = append(r.freeIntervals[:i+1],
				append([]FreeInterval{right}, r.freeIntervals[i+1:]...)...)
		}
		return
	}
}

// addToIntervals 将释放的大型车位加回空闲区间，并与相邻区间合并。
// 在车辆从大型车位离开时调用。
func (r *Row) addToIntervals(idx int) {
	// 二分查找插入位置（区间按 Start 排序）。
	pos := sort.Search(len(r.freeIntervals), func(i int) bool {
		return r.freeIntervals[i].Start > idx
	})

	mergeLeft := pos > 0 && r.freeIntervals[pos-1].End == idx
	mergeRight := pos < len(r.freeIntervals) && r.freeIntervals[pos].Start == idx+1

	switch {
	case mergeLeft && mergeRight:
		// 左区间 + 新车位 + 右区间 合并为一个。
		r.freeIntervals[pos-1].End = r.freeIntervals[pos].End
		r.freeIntervals = append(r.freeIntervals[:pos], r.freeIntervals[pos+1:]...)
	case mergeLeft:
		r.freeIntervals[pos-1].End = idx + 1
	case mergeRight:
		r.freeIntervals[pos].Start = idx
	default:
		// 插入新的单车位区间 [idx, idx+1)。
		newIntervals := make([]FreeInterval, len(r.freeIntervals)+1)
		copy(newIntervals, r.freeIntervals[:pos])
		newIntervals[pos] = FreeInterval{Start: idx, End: idx + 1}
		copy(newIntervals[pos+1:], r.freeIntervals[pos:])
		r.freeIntervals = newIntervals
	}
}

// ===== 层 =====

// Level 表示停车场的一层。
type Level struct {
	floor   int
	rows    []*Row
	numRows int

	// 按类型统计的空闲车位数 — O(1) 判断该层是否有空位。
	availCount [numSpotSizes]int
}

func newLevel(floor, numRows, spotsPerRow int) *Level {
	l := &Level{
		floor:   floor,
		numRows: numRows,
		rows:    make([]*Row, numRows),
	}
	for r := 0; r < numRows; r++ {
		l.rows[r] = newRow(floor, r, spotsPerRow)
		for st := SpotSize(0); st < numSpotSizes; st++ {
			l.availCount[st] += len(l.rows[r].avail[st])
		}
	}
	return l
}

// ===== 停车场 =====

// ParkingLot 是停车场的顶层结构。
type ParkingLot struct {
	levels []*Level

	// vehicleSpots 映射：车辆名称 → 占用的车位列表。
	// 这是核心优化：取车从 O(m·n·k) 降到 O(1)。
	vehicleSpots map[string][]*ParkingSpot
}

// NewParkingLot 创建一个拥有 n 层、每层 m 行、每行 k 个车位的停车场。
func NewParkingLot(numLevels, numRows, spotsPerRow int) *ParkingLot {
	levels := make([]*Level, numLevels)
	for i := 0; i < numLevels; i++ {
		levels[i] = newLevel(i, numRows, spotsPerRow)
	}
	return &ParkingLot{
		levels:       levels,
		vehicleSpots: make(map[string][]*ParkingSpot),
	}
}

// ParkVehicle 按名称停入车辆。成功返回 true。
//
// 策略：类型优先遍历，最小化大型车位碎片化。
//   - 先在所有行尝试优先车位类型，再降级到下一类型。
//   - 当非巴士车辆需要使用大型车位时，优先选择已无法停巴士的行，集中碎片。
//
// 时间复杂度：
//   - 单车位车辆: 最坏 O(m·n)
//   - 巴士: 最坏 O(m·n·I)，I = 每行空闲区间数
func (pl *ParkingLot) ParkVehicle(name string) bool {
	if _, exists := pl.vehicleSpots[name]; exists {
		return false
	}
	v := NewVehicle(name)
	if v.SpotsNeeded() > 1 {
		return pl.parkBus(v)
	}
	return pl.parkSingle(v)
}

// parkSingle 使用类型优先遍历：
//  1. 先在所有层和行中尝试摩托车位/紧凑车位。
//  2. 全部用完后才降级到大型车位。
//  3. 使用大型车位时，优先选择已无法停巴士的行。
func (pl *ParkingLot) parkSingle(v Vehicle) bool {
	pref := v.PreferredSpots()

	// 第一阶段：在所有层/行中尝试非大型车位。
	for _, st := range pref {
		if st == LargeSpot {
			continue // 大型车位留到第二阶段处理
		}
		for _, level := range pl.levels {
			if level.availCount[st] == 0 {
				continue
			}
			for _, row := range level.rows {
				if spot := row.parkSpotOfType(v, st); spot != nil {
					level.availCount[st]--
					pl.vehicleSpots[v.GetName()] = []*ParkingSpot{spot}
					return true
				}
			}
		}
	}

	// 第二阶段：优先车位已用完，必须使用大型车位。
	// 优先选择已无法停巴士的行，集中碎片。
	if !v.CanFitIn(LargeSpot) {
		return false
	}
	return pl.parkInLargeSpotSmart(v)
}

// parkInLargeSpotSmart 使用巴士友好策略分配大型车位：
//
//	第 1 轮：优先选无法停巴士的行（已被破坏，损失最小）
//	第 2 轮：所有行都能停巴士时，使用任意行
func (pl *ParkingLot) parkInLargeSpotSmart(v Vehicle) bool {
	// 第 1 轮：选已无法停巴士的行（canFitBus == false）
	for _, level := range pl.levels {
		if level.availCount[LargeSpot] == 0 {
			continue
		}
		for _, row := range level.rows {
			if row.canFitBus() {
				continue // 跳过 — 保留该行给巴士
			}
			if spot := row.parkSpotOfType(v, LargeSpot); spot != nil {
				level.availCount[LargeSpot]--
				pl.vehicleSpots[v.GetName()] = []*ParkingSpot{spot}
				return true
			}
		}
	}
	// 第 2 轮：没有已破坏的行，使用任意有大型车位的行
	for _, level := range pl.levels {
		if level.availCount[LargeSpot] == 0 {
			continue
		}
		for _, row := range level.rows {
			if spot := row.parkSpotOfType(v, LargeSpot); spot != nil {
				level.availCount[LargeSpot]--
				pl.vehicleSpots[v.GetName()] = []*ParkingSpot{spot}
				return true
			}
		}
	}
	return false
}

func (pl *ParkingLot) parkBus(v Vehicle) bool {
	for _, level := range pl.levels {
		if level.availCount[LargeSpot] < 5 {
			continue // 该层大型车位不足 5 个。
		}
		for _, row := range level.rows {
			if spots := row.parkBus(v); spots != nil {
				level.availCount[LargeSpot] -= 5
				pl.vehicleSpots[v.GetName()] = spots
				return true
			}
		}
	}
	return false
}

// UnParkVehicle 按名称取出车辆。
//
// 时间复杂度: O(spotsNeeded) ≈ O(1) — 通过 vehicleSpots 哈希表直接查找。
func (pl *ParkingLot) UnParkVehicle(name string) {
	spots, exists := pl.vehicleSpots[name]
	if !exists {
		return
	}
	for _, spot := range spots {
		level := pl.levels[spot.Level]
		row := level.rows[spot.Row]
		row.freeSpot(spot)
		level.availCount[spot.Size]++
	}
	delete(pl.vehicleSpots, name)
}

// Example 演示题目中的示例用法。
func Example() {
	lot := NewParkingLot(1, 1, 11)

	fmt.Println(lot.ParkVehicle("Motorcycle_1")) // true
	fmt.Println(lot.ParkVehicle("Car_1"))        // true
	fmt.Println(lot.ParkVehicle("Car_2"))        // true
	fmt.Println(lot.ParkVehicle("Car_3"))        // true
	fmt.Println(lot.ParkVehicle("Car_4"))        // true
	fmt.Println(lot.ParkVehicle("Car_5"))        // true
	fmt.Println(lot.ParkVehicle("Bus_1"))        // false
	lot.UnParkVehicle("Car_5")
	fmt.Println(lot.ParkVehicle("Bus_1")) // true
}
