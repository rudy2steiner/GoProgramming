package tempconv

import "fmt"

// 类型声明
type Celsius float64
type Fahrenheit float64

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC Celsius = 0
    BoilingC  Celsius = 100
)

func (c Celsius) String() string {return fmt.Sprintf("%g ^0C",c)}
func (f Fahrenheit) String() string {return fmt.Sprintf("%g ^0F",f)}
