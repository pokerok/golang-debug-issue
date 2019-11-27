# Golang debug issue

A simple program to run and display issue

```bash
go get github.com/pokeroklife/golang-debug-issue
```

When we debug the program (Shift+F9) using the debugger,
 the download hangs, the compilation is hanging even when building from command line:
 
 ```bash
 /home/sergey/go/go1.13.4/bin/go build -o /tmp/___go_build_golang_debug_issue -gcflags "all=-N -l" /home/sergey/go/src/github.com/pokeroklife/golang-debug-issue/main.go #gosetup
 ```

When send SIGQUIT to the compile program to generate a stack trace we have this result:

 ```bash
# github.com/zonedb/zonedb
SIGQUIT: quit
PC=0x45e461 m=0 sigcode=0

goroutine 0 [idle]:
runtime.futex(0x1596c28, 0x80, 0x0, 0x0, 0x0, 0x7fb409818ac0, 0x7fb409818a00, 0x7fb409818ac0, 0x7ffc39170db8, 0x40a76f, ...)
        /usr/local/go/src/runtime/sys_linux_amd64.s:535 +0x21
runtime.futexsleep(0x1596c28, 0x7fb400000000, 0xffffffffffffffff)
        /usr/local/go/src/runtime/os_linux.go:44 +0x46
runtime.notesleep(0x1596c28)
        /usr/local/go/src/runtime/lock_futex.go:151 +0x9f
runtime.stopm()
        /usr/local/go/src/runtime/proc.go:1928 +0xc0
runtime.findrunnable(0xc000026000, 0x0)
        /usr/local/go/src/runtime/proc.go:2391 +0x53f
runtime.schedule()
        /usr/local/go/src/runtime/proc.go:2524 +0x2be
runtime.park_m(0xc000000c00)
        /usr/local/go/src/runtime/proc.go:2610 +0x9d
runtime.mcall(0x0)
        /usr/local/go/src/runtime/asm_amd64.s:318 +0x5b

goroutine 1 [semacquire, 1 minutes]:
sync.runtime_Semacquire(0xc004b60f28)
        /usr/local/go/src/runtime/sema.go:56 +0x42
sync.(*WaitGroup).Wait(0xc004b60f20)
        /usr/local/go/src/sync/waitgroup.go:130 +0x64
cmd/compile/internal/gc.compileFunctions()
        /usr/local/go/src/cmd/compile/internal/gc/pgen.go:373 +0x1ce
cmd/compile/internal/gc.Main(0xe5bfe8)
        /usr/local/go/src/cmd/compile/internal/gc/main.go:695 +0x3241
main.main()
        /usr/local/go/src/cmd/compile/main.go:51 +0xac

goroutine 11 [runnable]:
cmd/compile/internal/ssa.(*desiredState).addList(0xc008e91720, 0xffffff070002ac1d)
        /usr/local/go/src/cmd/compile/internal/ssa/regalloc.go:2578 +0x6b
cmd/compile/internal/ssa.(*desiredState).merge(0xc008e91720, 0xc11edc1f20)
        /usr/local/go/src/cmd/compile/internal/ssa/regalloc.go:2638 +0x7d
cmd/compile/internal/ssa.(*regAllocState).computeLive(0xc004c4a000)
        /usr/local/go/src/cmd/compile/internal/ssa/regalloc.go:2436 +0x158d
cmd/compile/internal/ssa.(*regAllocState).init(0xc004c4a000, 0xc0003dfe40)
        /usr/local/go/src/cmd/compile/internal/ssa/regalloc.go:674 +0x6c1
cmd/compile/internal/ssa.regalloc(0xc0003dfe40)
        /usr/local/go/src/cmd/compile/internal/ssa/regalloc.go:145 +0x4a
cmd/compile/internal/ssa.Compile(0xc0003dfe40)
        /usr/local/go/src/cmd/compile/internal/ssa/compile.go:92 +0x994
cmd/compile/internal/gc.buildssa(0xc0003df8c0, 0x3, 0x0)
        /usr/local/go/src/cmd/compile/internal/gc/ssa.go:289 +0xc0e
cmd/compile/internal/gc.compileSSA(0xc0003df8c0, 0x3)
        /usr/local/go/src/cmd/compile/internal/gc/pgen.go:298 +0x4d
cmd/compile/internal/gc.compileFunctions.func2(0xc004b8b3e0, 0xc004b60f20, 0x3)
        /usr/local/go/src/cmd/compile/internal/gc/pgen.go:363 +0x49
created by cmd/compile/internal/gc.compileFunctions
        /usr/local/go/src/cmd/compile/internal/gc/pgen.go:361 +0x128

rax    0xca
rbx    0x1596ae0
rcx    0x45e463
rdx    0x0
rdi    0x1596c28
rsi    0x80
rbp    0x7ffc39170d80
rsp    0x7ffc39170d38
r8     0x0
r9     0x0
r10    0x0
r11    0x286
r12    0x0
r13    0x7fb415d15008
r14    0x1
r15    0x1
rip    0x45e461
rflags 0x286
cs     0x33
fs     0x0
gs     0x0 
```

Any ideas about what's the problem?
