// package main

// import (
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

//	func main() {
//		dsn := "host=localhost user=postgres password=secret dbname=app port=5432 sslmode=disable"
//		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//		if err != nil {
//			panic(err)
//		}
//		db.AutoMigrate(&model.Product{})
//	}
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func main() {
	const (
		totalRequests = 10000                                                                                           // 総リクエスト数
		concurrency   = 500                                                                                             // 並列数（ここを増やす）
		dataSize      = 10 * 1024 * 1024                                                                                // 10MBのデータ
		url           = "https://sites.google.com/stu.yamato-u.ac.jp/2025test23/%E3%83%9B%E3%83%BC%E3%83%A0?authuser=0" // ← テスト用API
	)

	// 10MBのダミーデータを作成
	data := bytes.Repeat([]byte("A"), dataSize)

	sem := make(chan struct{}, concurrency)
	done := make(chan struct{})
	start := time.Now()

	for i := 0; i < totalRequests; i++ {
		sem <- struct{}{}
		go func(i int) {
			defer func() {
				<-sem
				done <- struct{}{}
			}()

			resp, err := http.Post(url, "application/octet-stream", bytes.NewReader(data))
			if err != nil {
				fmt.Printf("❌ Request %d failed: %v\n", i, err)
				return
			}
			resp.Body.Close()
			fmt.Printf("✅ Sent request %d\n", i)
		}(i)
	}

	// 全てのリクエストの完了を待つ
	for i := 0; i < totalRequests; i++ {
		<-done
	}

	elapsed := time.Since(start)
	fmt.Printf("🎉 Completed %d POST requests in %s\n", totalRequests, elapsed)
}
