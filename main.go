package main

import (
	"fmt"
	"sync"
)

type WeaponStrategy interface {
	UseWeapon()
}

type Ak47 struct{}

func (a *Ak47) UseWeapon() {
	fmt.Println("使用ak47战斗")
}

type Knife struct{}

func (a *Knife) UseWeapon() {
	fmt.Println("使用小刀战斗")
}

type Hero struct {
	Strategy WeaponStrategy
}

func (h *Hero) SetWeapon(strategy WeaponStrategy) {
	h.Strategy = strategy
}

func (h *Hero) fight() {
	h.Strategy.UseWeapon()
}

func init() {
	///fmt.Println("123321")

}

func main() {
	//name := flag.String("name", "world", "just a joke")
	//flag.Parse()
	//fmt.Println("os arg is:", os.Args)
	//fmt.Println("input parameter is ：", *name)
	//fullString := fmt.Sprintf("hello %s from go", *name)
	//fmt.Println(fullString)
	//hero := Hero{}
	//hero.SetWeapon(new(Ak47))
	//hero.fight()

	//a := 1
	//b := &a
	//c := *&a
	//*b = 2
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//
	//var x int
	//x = 1
	//
	//y := newDdx.Get(x)
	//fmt.Println(y)

	//ch := make(chan int, 10)
	//over := make(chan bool)
	//
	//go func() {
	//	ticker := time.NewTicker(1 * time.Second)
	//	for range ticker.C {
	//		select {
	//		case <-over:
	//			fmt.Println("时间到，结束啦！！！！")
	//			return
	//		default:
	//			fmt.Printf("继续啊 %d", <-ch)
	//		}
	//	}
	//
	//}()
	//
	//for i := 0; i <= 10; i++ {
	//	ch <- i
	//}
	//
	//time.Sleep(5 * time.Second)
	//close(over)
	//fmt.Println("over")

	//for i := 0; i <= 100; i++ {
	//	go func(i int) {
	//		ch <- i
	//	}(i)
	//}
	//array := make([]int, 0)
	//for {
	//	array = append(array, <-ch)
	//	if len(array) == 99 {
	//		break
	//	}
	//}
	//
	//fmt.Println(array)
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i <= 100; i++ {
		lock.Lock()
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
		lock.Unlock()
	}
	wg.Wait()

	//a := make(map[int]int, 0)
	//NewMap := newMap{NewMap: a, s: sync.Mutex{}}
	//for i := 0; i <= 100; i++ {
	//	go func(i int) {
	//		NewMap.write(i, i, a)
	//	}(i)
	//}
	//
	//time.Sleep(time.Second)
	//
	//fmt.Println(a)

	//a := make(map[int]int, 0)
	//wg := sync.WaitGroup{}
	//wg.Add(100)
	//for i := 0; i < 100; i++ {
	//	go func(i int) {
	//		a[i] = i
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()
	//
	//fmt.Println(a)

	//sync.
	//q := newPlay{
	//	NewMap: a,
	//	L:      sync.NewCond(&sync.Mutex{}),
	//}
	//q.L.L.Lock()
	//
	//defer cond.L.Unlock()
	//for i := 0; i <= 10000; i++ {
	//	if i < 9000 {
	//		cond.Wait()
	//	}
	//}

}

type newMap struct {
	NewMap map[int]int
	s      sync.Mutex
}

func (s *newMap) write(k int, v int, m map[int]int) map[int]int {
	s.s.Lock()
	defer s.s.Unlock()
	m[k] = v

	return m
}

type newPlay struct {
	NewMap map[int]int
	L      *sync.Cond
}

func (n *newPlay) waitWrite(m map[int]int) {

}

type dd interface {
	Get(x int) (y int)
}

type xx interface {
	Make(x int) (y int)
}

var newDdx = new(Ddx)

type Ddx struct {
}

func (d *Ddx) Get(x int) (y int) {
	return x
}

func (d *Ddx) Make(x int) (y int) {
	return x
}
