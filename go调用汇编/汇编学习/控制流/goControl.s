#include "textflag.h"

TEXT ·If(SB),NOSPLIT,$0-32
    MOVQ ok+8*0(FP),CX
    MOVQ a+8*1(FP), AX
    MOVQ b+8*2(FP), BX
    CMPQ CX, $0
    JZ   L
    MOVQ AX, ret+24(FP)
    RET
L:
    MOVQ BX,ret+24(FP)
    RET


//func LoopAdd(cnt, v0, step int) int {
//  result := v0
//    for i := 0; i < cnt; i++ {
//        result += step
//    }
//   return result
//}

// func LoopAdd(cnt, v0, step int) int
TEXT ·LoopAdd(SB), NOSPLIT,  $0-32
    MOVQ cnt+0(FP), AX   // cnt
    MOVQ v0+8(FP), BX    // v0/result
    MOVQ step+16(FP), CX // step

LOOP_BEGIN:
    MOVQ $0, DX          // i

LOOP_IF:
    CMPQ DX, AX          // compare i, cnt
    JL   LOOP_BODY       // if i < cnt: goto LOOP_BODY
    JMP LOOP_END

LOOP_BODY:
    ADDQ $1, DX          // i++
    ADDQ CX, BX          // result += step
    JMP LOOP_IF

LOOP_END:

    MOVQ BX, ret+24(FP)  // return result
    RET