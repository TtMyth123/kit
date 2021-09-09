package ttLock

import "sync"

var (
	serverLock sync.Mutex
	//LockManage LockManager
)

func init() {

}

type LockManager struct {
	mpTtLock sync.Map
	mLock sync.Mutex
}


func (this *LockManager) Lock(key interface{},c int )  {
	serverLock.Lock()
	defer serverLock.Unlock()
	if l,ok:=this.mpTtLock.Load(key) ;ok{
		lock := l.(*TtLock)
		lock.Lock(c)
	} else {
		aTtLock := NewTtLock()
		aTtLock.Lock(c)
		this.mpTtLock.Store(key,aTtLock)
	}
}
func (this *LockManager) UnLock(key interface{})  {
	serverLock.Lock()
	defer serverLock.Unlock()
	if l,ok:=this.mpTtLock.Load(key) ;ok{
		lock := l.(*TtLock)
		lock.UnLock()
	} else {
		aTtLock := NewTtLock()
		aTtLock.UnLock()
		this.mpTtLock.Store(key,aTtLock)
	}
}
