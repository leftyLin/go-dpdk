package dpdk

/*
#include <rte_config.h>
#include <rte_ip.h>
*/
import "C"

import (
	"encoding/binary"
	"net"
)

type IPv4Hdr C.struct_ipv4_hdr

func (h *IPv4Hdr) SourceAddr() string {
	bytes := make([]byte, 4)
	u32 := uint32((*(*C.struct_ipv4_hdr)(h)).src_addr)
	binary.BigEndian.PutUint32(bytes, u32)
	return net.IP(bytes).String()
}

func (h *IPv4Hdr) DestAddr() string {
	bytes := make([]byte, 4)
	u32 := uint32((*(*C.struct_ipv4_hdr)(h)).dst_addr)
	binary.BigEndian.PutUint32(bytes, u32)
	return net.IP(bytes).String()
}
