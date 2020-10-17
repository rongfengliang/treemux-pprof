# pprof for treemux 


## How to using

* register handler

```code

go get github.com/rongfengliang/treemux-pprof


router := treemux.New()
pprof.RouterpprofRegister(router)

view result:

/debug/pprof

```