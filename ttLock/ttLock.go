package ttLock

import (
	"fmt"
	"sync"
	"time"
)

type TtLock struct {
	mLockI     sync.Mutex
	iLockCount int
	iLockCount2 int
	timer1     *time.Timer
	mLock      sync.Mutex
}

func NewTtLock() *TtLock {
	aTtLock :=new(TtLock)
	aTtLock.timer1 = time.NewTimer(1)
	aTtLock.timer1.Stop()
	//go aTtLock.run()
	return aTtLock
}
func (this *TtLock)run()  {
	for {
		<-this.timer1.C
		{
			this.mLockI.Lock()
			fmt.Println("AAAAAAAAA", this.iLockCount)
			for i:=0;i<this.iLockCount;i++{
				fmt.Println("AAAAAAAAA ii",i)
				if this.iLockCount2>0 {
					this.mLock.Unlock()
					this.iLockCount2--
					fmt.Println("bbbb ii",i)
				}
			}
			this.mLockI.Unlock()
		}
	}
}


func (this *TtLock)Lock(c int)  {
	this.mLockI.Lock()
	this.iLockCount++
	this.mLockI.Unlock()

	this.timer1.Reset(time.Duration(c)*time.Second)
	this.mLock.Lock()
	this.iLockCount2++
}


func (this *TtLock)UnLock()  {
	this.mLockI.Lock()
	if this.iLockCount >0 {
		this.iLockCount--
		this.mLock.Unlock()
		this.iLockCount2--
	}

	this.mLockI.Unlock()

}


