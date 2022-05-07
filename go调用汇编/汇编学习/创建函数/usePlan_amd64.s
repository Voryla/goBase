#include "textflag.h"

// func add(a, b int) int
TEXT ·add(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX // 参数 a
    MOVQ b+8(FP), BX // 参数 b
    ADDQ BX, AX    // AX += BX
    MOVQ AX, ret+16(FP) // 返回
    RET

TEXT ·swap(SB),NOSPLIT,$0
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    MOVQ BX, ret0+16(FP)
    MOVQ AX, ret1+24(FP)
    RET
