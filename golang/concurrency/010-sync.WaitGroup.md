sync.WaitGroup的功能是等待协程执行完毕。    
```
type WaitGroup
    func (wg *WaitGroup) Add(delta int)
    func (wg *WaitGroup) Done()
    func (wg *WaitGroup) Wait()
```
可以用于实现获取协程执行后的返回值

```golang
func main() {
	var wg sync.WaitGroup
	var sum int 
	for i := 1; i<6; i++ {
		wg.Add(1)
		go accelate(&sum, i, &wg)
	}

	wg.Wait()
	fmt.Println("sum:", sum)
}

func accelate(sum *int, item int, wg *sync.WaitGroup) {
	fmt.Println("sum accelate ", item)
	
	defer wg.Done()

	*sum = *sum + item

	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
}
```
