#include "textflag.h"

//type MyInt int
//
//func (v MyInt) Twice() int {
//    return int(v)*2
//}
//
//func MyInt_Twice(v MyInt) int {
//    return int(v)*2
//}
// 无法实现：会panic makeABIWrapper support for wrapping methods not implemented
// func (v MyInt) Twice() int
TEXT ·MyInt·Twice(SB), NOSPLIT, $0-16
    MOVQ a+0(FP), AX   // v
    ADDQ AX, AX        // AX *= 2
    MOVQ AX, ret+8(FP) // return v
    RET
