package main

import (
	"context"
	"log"
	"time"
)

/* context สามารถส่งค่าต่าง ๆ ระหว่าง Process ใน request scope
สามารถกำหนด ค่าต่าง ๆ ใน context ได้ดังนี้
	withCancel()
	withDeadline()
	withTimeout()
	withValue()
*/

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)
	ch := time.After(600 * time.Millisecond)

	select {
	// กรณี timeout เกินที่กำหนด ctx.Done จะส่งข้อความกลับมา
	case <-ctx.Done():
		log.Println("ctx.Done()")
		log.Printf("ctx.Err() Done : %v\n", ctx.Err())
	// กรณี timeout อยู่ในระยะเวลาที่กำหนด	
	case <-ch:
		log.Println("ch done")
		log.Printf("ctx.Err() ch : %v\n", ctx.Err())
	}

	// กรณี timeout อยู่ในระยะเวลาที่กำหนด

	/* 							Timeout

					500 milisecond ( เวลาในการทำงาน )

	context.WithTimeOut     							<-ctx.Done() (ไม่ส่งข้อความ)

	*/


	// กรณี timeout เกินที่กำหนด ctx.Done 
	
	/* 							Timeout
																			 
						500 milisecond (เวลาในการทำงาน)		                    || ========= เวลาที่เกิน ======
	context.WithTimeOut     							<-ctx.Done() (ส่งข้อความ)

	*/

}
