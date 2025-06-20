package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

func pointerTrans(p *int) {
	*p = 25
}
func valueTrans(va int) {
	va = 3
}

func addTwo(slice *[]int) {
	for i := range *slice {
		(*slice)[i] += 2
	}
}

// 定义函数类型
type task func()

type taskResult struct {
	TaskID     int
	starttaime time.Time
	endtime    time.Time
	duration   time.Duration
}

type Scheduler struct {
	task       []task
	taskresult []taskResult
	startime   time.Time
	wg         sync.WaitGroup
}

func NewScheduler(t []task) *Scheduler {
	return &Scheduler{
		task:       t,
		taskresult: make([]taskResult, len(t)),
	}
}

func (s *Scheduler) Run() {
	s.startime = time.Now()
	s.wg.Add(len(s.task))
	for i, v := range s.task {
		go s.executeTask(i, v)
	}
	s.wg.Wait()
}

func (s *Scheduler) executeTask(id int, task task) {
	defer s.wg.Done()

	result := taskResult{
		starttaime: time.Now(),
		TaskID:     id,
	}

	defer func() {
		result.endtime = time.Now()
		result.duration = result.endtime.Sub(result.starttaime)
		s.taskresult[id] = result
	}()

	task()
}

func (s *Scheduler) printStats() {

	totalDuration := time.Since(s.startime)

	fmt.Println("\n任务执行统计:")
	fmt.Printf("总任务数: %d\n", len(s.task))
	fmt.Printf("总执行时间: %v\n", totalDuration)
	for _, result := range s.taskresult {
		fmt.Printf(
			"任务%d - 开始: %v, 结束: %v, 耗时: %v\n",
			result.TaskID,
			result.starttaime.Format("15:04:05.000"),
			result.endtime.Format("15:04:05.000"),
			result.duration,
		)
	}
}

type Shap interface {
	Area() float64
	Perimeter() float64
	GetShap() string
}

type Recttangle struct {
	Width  float64
	Height float64
	shap   string
}

func (r Recttangle) Area() float64 {
	return r.Width * r.Height
}

func (r Recttangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
func (r Recttangle) GetShap() string {
	return r.shap
}

type Circle struct {
	r    float64
	shap string
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (c Circle) GetShap() string {
	return c.shap
}

func printShap(s Shap) {
	fmt.Println("形状：", s.GetShap())
	fmt.Println("面积", s.Area())
	fmt.Println("周长", s.Perimeter())
}

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	baseInfo   Person
	EmployeeID int
}

func (e Employee) printInfo() {
	fmt.Printf("员工信息:姓名:%s 年龄:%d", e.baseInfo.Name, e.baseInfo.Age)
	fmt.Println("\nID:", e.EmployeeID)
}

func sendData(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		fmt.Println("发送数据:", i)
		ch <- i
	}
	close(ch) // 发送完成后关闭通道
}
func reciveData(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("接收数据:", v)
	}
}

func main() {

	// i := 1

	// fmt.Println("值传递前", i)
	// valueTrans(i)
	// fmt.Println("值传递后", i)

	// fmt.Println("指针传递前", i)
	// pointerTrans(&i)
	// fmt.Println("指针传递前", i)

	// slice := []int{2, 4, 5, 6, 7, 7}
	// addTwo(&slice)
	// fmt.Println(slice)

	// var wg sync.WaitGroup
	// wg.Add(2)

	// go func() {
	// 	defer wg.Done()
	// 	for i := 2; i <= 10; i += 2 {
	// 		fmt.Println("偶数", i)
	// 		time.Sleep(time.Millisecond * 100) // 添加微小延迟
	// 	}
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	for i := 1; i <= 10; i += 2 {
	// 		fmt.Println(("奇数"), i)
	// 		time.Sleep(time.Millisecond * 100) // 添加微小延迟
	// 	}
	// }()
	// wg.Wait()
	// fmt.Println("所有数字打印完成")

	// tasks := []task{
	// 	func() {
	// 		time.Sleep(1 * time.Second)
	// 		fmt.Println("任务1完成")
	// 	},
	// 	func() {
	// 		time.Sleep(2 * time.Second)
	// 		fmt.Println("任务2完成")
	// 	},
	// 	func() {
	// 		time.Sleep(500 * time.Millisecond)
	// 		fmt.Println("任务3完成")
	// 	},
	// 	func() {
	// 		time.Sleep(300 * time.Millisecond)
	// 		fmt.Println("任务4完成")
	// 	},
	// }

	// // 创建调度器并运行
	// scheduler := NewScheduler(tasks)
	// scheduler.Run()
	// scheduler.printStats()

	// rectangle := Recttangle{Width: 5, Height: 10, shap: "矩形"}
	// circle := Circle{r: 3, shap: "圆形"}

	// printShap(rectangle)
	// printShap(circle)

	// employee :=
	// 	Employee{
	// 		baseInfo: Person{
	// 			Name: "deng",
	// 			Age:  26,
	// 		},
	// 		EmployeeID: 1,
	// 	}
	// employee.printInfo()
	// var wg sync.WaitGroup
	// wg.Add(2)
	// ch := make(chan int, 10)
	// go sendData(ch, &wg)
	// go reciveData(ch, &wg)

	// wg.Wait()

	// var wg sync.WaitGroup
	// wg.Add(1)

	// ch := make(chan int, 3)

	// go func() {
	// 	defer wg.Done()
	// 	for i := 0; i < 100; i++ {
	// 		ch <- i
	// 		fmt.Println("生产：", i)
	// 		time.Sleep(200 * time.Millisecond)

	// 		select {
	// 		case num := <-ch:
	// 			fmt.Println("接收:", num)
	// 		default:
	// 			println("没有数据")
	// 		}
	// 	}
	// 	close(ch)
	// }()
	// wg.Wait()

	var (
		counter int32

		wg sync.WaitGroup
		// mu sync.Mutex
	)

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				// mu.Lock()
				// counter++
				// mu.Unlock()

				atomic.AddInt32(&counter, 1)
			}

		}()
	}
	wg.Wait()
	fmt.Println(counter)

}
