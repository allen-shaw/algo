package parkinglot

// ===== 车辆接口 =====

// Vehicle 是所有车辆类型必须实现的接口。
// 每种具体类型（Motorcycle、Car、Bus）定义各自的停车规则。
type Vehicle interface {
	// GetName 返回车辆标识符。
	GetName() string
	// SpotsNeeded 返回需要的连续车位数量。
	SpotsNeeded() int
	// CanFitIn 返回该车辆是否能停在指定大小的车位。
	CanFitIn(SpotSize) bool
	// PreferredSpots 返回优先尝试的车位类型列表，按优先级排序。
	PreferredSpots() []SpotSize
}

// ===== 基础车辆（共享字段）=====

// baseVehicle 提供所有车辆类型的公共字段。
type baseVehicle struct {
	name string
}

func (v *baseVehicle) GetName() string { return v.name }

// ===== 摩托车 =====

// Motorcycle 可以停在任何类型的车位（摩托车位、紧凑车位、大型车位）。
type Motorcycle struct {
	baseVehicle
}

func NewMotorcycle(name string) *Motorcycle {
	return &Motorcycle{baseVehicle{name: name}}
}

func (m *Motorcycle) SpotsNeeded() int         { return 1 }
func (m *Motorcycle) CanFitIn(_ SpotSize) bool { return true }
func (m *Motorcycle) PreferredSpots() []SpotSize {
	return []SpotSize{MotorcycleSpot, CompactSpot, LargeSpot}
}

// ===== 小汽车 =====

// Car 可以停在单个紧凑车位或单个大型车位。
type Car struct {
	baseVehicle
}

func NewCar(name string) *Car {
	return &Car{baseVehicle{name: name}}
}

func (c *Car) SpotsNeeded() int { return 1 }
func (c *Car) CanFitIn(s SpotSize) bool {
	return s == CompactSpot || s == LargeSpot
}
func (c *Car) PreferredSpots() []SpotSize {
	return []SpotSize{CompactSpot, LargeSpot}
}

// ===== 巴士 =====

// Bus 需要同一行内 5 个连续的大型车位。
type Bus struct {
	baseVehicle
}

func NewBus(name string) *Bus {
	return &Bus{baseVehicle{name: name}}
}

func (b *Bus) SpotsNeeded() int { return 5 }
func (b *Bus) CanFitIn(s SpotSize) bool {
	return s == LargeSpot
}
func (b *Bus) PreferredSpots() []SpotSize {
	return []SpotSize{LargeSpot}
}

// ===== 工厂方法 =====

// NewVehicle 根据名称字符串（如 "Car_1"、"Bus_2"、"Motorcycle_3"）创建对应的车辆实例。
func NewVehicle(name string) Vehicle {
	switch {
	case len(name) >= 10 && name[:10] == "Motorcycle":
		return NewMotorcycle(name)
	case len(name) >= 3 && name[:3] == "Bus":
		return NewBus(name)
	case len(name) >= 3 && name[:3] == "Car":
		return NewCar(name)
	default:
		return NewCar(name)
	}
}
