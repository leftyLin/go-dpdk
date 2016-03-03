# go-dpdk
[Intel DPDK](http://dpdk.org) wrapper for Go using [cgo](https://github.com/golang/go/wiki/cgo)

This is not meant to be fast, just to allow using some of Go's features to play
around with packet processing

# Tutorial
+ [Download and build](http://dpdk.org/doc/guides/linux_gsg/build_dpdk.html) 
DPDK with `CONFIG_RTE_COMBINE_LIBS=y`
+ Run `export RTE_SDK=/path/to/dpdk`
+ Run `export RTE_TARGET=some_target` e.g. `RTE_TARGET=x86_64-native-linuxapp-gcc`
+ Then just run
```
export CGO_LDFLAGS="-L${RTE_SDK}/${RTE_TARGET}/lib -Wl,--whole-archive -ldpdk \
-lz -Wl,--start-group -lrt -lm -ldl -Wl,--end-group -Wl,--no-whole-archive"

export CGO_CFLAGS="-m64 -pthread -O3 -march=native \
-I${RTE_SDK}/${RTE_TARGET}/include"

go get github.com/melvinw/go-dpdk
```
+ Add `import github.com/melvinw/go-dpdk` to your code and run `go build`, etc...

See [the examples](https://github.com/melvinw/go-dpdk-examples)

See DPDK/godocs for API
