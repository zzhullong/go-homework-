package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 接收int类型的指针 *int
func add10(p *int) {
	// *p 解引用，修改指针指向内存上的值
	*p = *p + 10
}



// 整数切片的指针 *[]int
func mulTwo(slicePtr *[]int) {
	// 解引用，拿到原始切片
	s := *slicePtr
	for i := range s {
		s[i] *= 2
	}
}


//题三 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
 func goroutine1(x int) {
	fmt.Println("题目四")
	go func() {
		fmt.Printf("奇数")
		for i := 0; i < x+1; i++ {
			if i%2 != 0 {
				fmt.Println("题目三奇数",i)
			}
		}
	}()
	
	go func() {
		fmt.Printf("偶数数")
		for i := 0; i < x+1; i++ {
			if i%2 == 0 {
				fmt.Println("题目3偶数",i)
			}
		}
	}()

	time.Sleep(time.Second)

 }


 //题目四 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间



type Task struct {
	id int
	cost time.Duration
}


func goroutine2(tasks []Task)  {
	
	var wg sync.WaitGroup

	for _, v := range tasks {
		wg.Add(1)
		go func(task Task) {

			defer wg.Done()
			start := time.Now()
			time.Sleep(time.Second*2)
			task.cost =time.Since(start)
			
			fmt.Println("执行任务ID:",task.id ,",执行耗时:",task.cost)
		}(v)


	}

	wg.Wait()
	fmt.Println("All tasks completed")
	
}




// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。

type Shape interface {

	Area() int

	Perimeter() int 

}


type Rectangle struct{
	
	 area int
		
	perimeter int
}

func (rectangle Rectangle) Area() int {
	
	
	return rectangle.area
}

func (rectangle Rectangle) Perimeter() int {
	
	return rectangle.perimeter
}

type Circle struct{
	
	 area int
		
	perimeter int
}

func (circle Circle) Area() int {
	
	return circle.area
}

func (circle Circle) Perimeter() int {
	
	return circle.perimeter
}



//题目六 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。


type Person struct{
	name string
	age int
}


type Employee struct{
	person Person
	employeeID int
}


func (e Employee)PrintInfo() (name string,age int) {

	return e.person.name,e.person.age
}


//题目七  题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
func channel1() {
	
	ch1 := make(chan int ,4)
	
	var wg  sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 11; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		defer wg.Done()
		for v := range ch1 {
			fmt.Println("ch2管道的值是：",v)
		}
		
	}()
		wg.Wait()
		
}



//题目八 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。

func channel2() {
	
	ch1 := make(chan int ,100)
	fmt.Println("第八题开始")
	var wg  sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		defer wg.Done()
		for v := range ch1 {
			fmt.Println("ch2管道的值是：",v)
		}
		
	}()
		wg.Wait()
		
}


//题目九 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

type Count struct {
	mu sync.Mutex
	count int
}


func (count *Count) add()  {
	count.mu.Lock()
	count.count++
	defer count.mu.Unlock()
}

func (count *Count) getCount() int {
	count.mu.Lock()
	
	defer count.mu.Unlock()

	return count.count
}


func mutex()  {
	
	
	var wg sync.WaitGroup
	count := Count{}
fmt.Println("第九题开始")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {

			for i := 0; i < 1000; i++ {
				
		count.add()
		
		}
		defer wg.Done()
		
	}()
	}
	
	
	wg.Wait()

	fmt.Println("第九题计数的总数为：",count.getCount())
	
	

}


//题目十 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

func syncc()  {
	var count atomic.Int32
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				count.Add(1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("总数:",count.Load())
}



func main() {
	
	
	//题目一
	num := 5
	// &num 取num的内存地址，传入函数
	add10(&num)
	fmt.Println("题目1:",num) // 输出：15


//题目二
	slicePtr := []int{1,2,3,4,5,}

	for _, v := range slicePtr {
		fmt.Println("题目2指针前数字:",v)
	}
	mulTwo(&slicePtr)

	for _, v := range slicePtr {
		fmt.Println("题目2指针后的数字:",v)
	}

//题目三
	goroutine1(10)


//题目四
task := make([]Task,3)
	for i := 0; i < len(task); i++ {
		task[i].id = i
	}
	fmt.Println("题目4task数目",len(task))

	goroutine2(task)


	//题目五

	fmt.Println("题目五:")
	var shape Shape
	shape = Rectangle{25,20}
	fmt.Println("Rectangle的area为",shape.Area(),"Rectangle的per为:",shape.Perimeter())

	shape = Circle{20,18}
	fmt.Println("Circle的area为",shape.Area(),"Circle的per为:",shape.Perimeter())



	//题目六
	person := Person{name: "zhangsan",age: 15}
	employee := Employee{employeeID: 1,person: person}

	name,age := employee.PrintInfo()
	fmt.Println("题目六员工名字为",name,",年龄为：",age)

	//题目七
	channel1()

	//题目八
channel2()

//题目九

mutex()

//题目十
syncc()

}

