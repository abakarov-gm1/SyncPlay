package main

//func lengthOfLongestSubstring(s string) int {
//	if len(s) == 1 {
//		return 1
//	}
//	m := map[byte]int{}
//	count := 0
//	windowStart := 0
//
//	for windowEnd := 0; windowEnd < len(s); windowEnd++ {
//
//		m[s[windowEnd]]++
//
//		for m[s[windowEnd]] > 1 {
//			m[s[windowStart]]--
//			windowStart++
//			fmt.Println(windowEnd)
//		}
//		if windowEnd-windowStart+1 > count {
//			count = windowEnd - windowStart + 1
//		}
//
//	}
//
//	return count
//}

//type User struct {
//	ID   uint `gorm:"primaryKey"`
//	Name string
//	Age  int
//}

//dsn := "host=localhost user=hajik password=secret dbname=english port=5432 sslmode=disable"
//
//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//if err != nil {
//	panic("Не удалось подключиться к базе данных")
//}
//
//er := db.AutoMigrate(&User{Name: "Hajik", Age: 24})
//
//if er != nil {
//	panic("какая то ошибка")
//}
//
//fmt.Println("✅ Успешное подключение к базе!")

//s := lengthOfLongestSubstring("pawwwwwwwwkew")

//package main
//
//import (
//"fmt"
//"sync"
//)
//
//func test() <-chan int {
//
//	ch := make(chan int)
//	wg := sync.WaitGroup{}
//	wg.Add(2)
//
//	go func() {
//		for i := 0; i < 2; i++ {
//			ch <- i + 1
//		}
//		//defer wg.Done()
//		close(ch)
//
//	}()
//
//	go func() {
//		for i := 10; i < 15; i++ {
//			ch <- i + 1
//		}
//		//defer wg.Done()
//		close(ch)
//	}()
//
//	//go func() {
//	//	wg.Wait()
//	//	close(ch)
//	//}()
//
//	return ch
//}
//
//func main() {
//	ch := test()
//	for range 13 {
//		b, ok := <-ch
//		if !ok {
//			break
//		}
//		fmt.Println(b, ok)
//	}
//}
