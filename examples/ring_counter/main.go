package main

import (
	"C"
	"fmt"
	"os"
	"unsafe"

	"github.com/melvinw/go-dpdk"
	"github.com/melvinw/go-dpdk/util"
)

type ctx struct {
	mp *dpdk.RteMemPool
	r  *dpdk.RteRing
}

func runS(arg unsafe.Pointer) int {
	c := (*ctx)(arg)
	mp := c.mp
	r := c.r
	burst := uint(32)
	tbl := util.GetCArray(burst)
	tblSlice := util.SliceFromCArray(tbl, burst)

	n := uint(0)
	for {
		ret := mp.GetBulk(tbl, burst)
		if ret == 0 {
			for i := uint(0); i < burst; i++ {
				*((*uint)(tblSlice[i])) = n
				n++
			}
			nb_tx := uint(0)
			for nb_tx < burst {
				addr := uintptr(unsafe.Pointer(tbl)) + unsafe.Sizeof(tblSlice[0])*uintptr(nb_tx)
				p := unsafe.Pointer(addr)
				nb_tx += r.EnqueueBurst((*unsafe.Pointer)(p), burst-nb_tx)
			}
		}
		if n >= 100100 {
			fmt.Printf("Done sending!\n")
			break
		}
	}
	return 0
}

func main() {
	dpdk.RteEalInit(os.Args)

	mp := dpdk.RteMemPoolCreate("test_mp", 1024, 64, 512, 0, 0, 0)
	r := dpdk.RteRingCreate("test_ring", 1024, 0, 0)

	c := ctx{mp, r}
	dpdk.RteEalRemoteLaunch(runS, unsafe.Pointer(&c), 1)

	n := uint(32)
	tbl := util.GetCArray(n)
	tblSlice := util.SliceFromCArray(tbl, n)

	total := uint(0)
	sum := uint(0)
	for {
		nb_rx := r.DequeueBurst(tbl, 32)
		for i := uint(0); i < nb_rx; i++ {
			obj := tblSlice[i]
			x := *((*uint)(obj))
			sum += x
		}
		total += nb_rx
		mp.PutBulk(tbl, uint(nb_rx))
		if total > 100000 {
			break
		}
	}
	dpdk.RteEalMpWaitLCore()
	fmt.Printf("Done! n=%d avg=%.2f\n", total, float32(sum)/float32(total))
}
