ifeq ($(RTE_SDK),)
	$(error "RTE_SDK undefined")
endif

RTE_TARGET ?= x86_64-native-linuxapp-gcc

CGO_CFLAGS = -I${RTE_SDK}/${RTE_TARGET}/include
CGO_LDFLAGS = -L${RTE_SDK}/${RTE_TARGET}/lib -ldpdk -m64 -pthread -march=native -lm -ldl

install:
	go build main.go
