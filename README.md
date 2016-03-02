# go-dpdk
[Intel DPDK](https://dpdk.org) for golang

# Tutorial
+ Download and build dpdk with `CONFIG_RTE_COMBINE_LIBS=y`
+ Run `export RTE_SDK=/path/to/dpdk`
+ Run `export RTE_TARGET=some_target` e.g. `RTE_TARGET=x86_64-native-linuxapp-gcc`
+ Run `export CGO_CFLAGS="-I{$RTE_SDK}/${RTE_TARGET}/include"`
+ Run `export CGO_LDFLAGS="-L{$RTE_SDK}/${RTE_TARGET}/lib -ldpdk -m64 -pthread -march=native -lm -ldl"`
+ Run `go get github.com/melvinw/go-dpdk`
+ Build your project

See DPDK/godocs for API
