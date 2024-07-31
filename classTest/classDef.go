package main

import "fmt"

// 首字母大写
type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

type SuperHero struct {
	Hero //表示当前类继承Hero，包含属性和方法
	age  int
}

func (this *SuperHero) GetName() string {
	return "111" + this.Hero.Name
}

func main() {
	hero := Hero{
		Name:  "zhangsan",
		Ad:    19,
		Level: 44,
	}
	fmt.Printf("%v\n", hero)
	hero.SetName("lisi")
	fmt.Printf("%v\n", hero)

	superHero := SuperHero{
		Hero{
			Name:  "2222",
			Ad:    0,
			Level: 0,
		},
		32,
	}

	println(superHero.GetName())
}
