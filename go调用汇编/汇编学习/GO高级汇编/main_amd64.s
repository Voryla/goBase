#include "textflag.h"
// 原汇编指令
TEXT ·printnl_nosplit(SB), NOSPLIT, $8
    CALL runtime·printnl(SB)
    RET
// 执行 go tool asm -S main_amd64.s 编译成GO汇编
// 1. 实际会编译成
// TEXT "".printnl(SB), NOSPLIT, $16    // 我们定义汇编时，分配了8字节，但Go汇编会额外插入8字节
//      SUBQ $16, SP                    // 为当前栈帧分配16字节空间
//          MOVQ BP, 8(SP)              // 额外插入的8字节是为了保存调用函数时BP寄存器的值，即caller的BP值
//          LEAQ 8(SP), BP              // 将8(SP)地址重新保持到BP寄存器中，即当前calle的BP值
//              CALL runtime.printnl(SB)
//          MOVQ 8(SP), BP              // 取出第9行时保存在栈空间8(SP)的caller的BP值，重新放回到BP寄存器中
//       ADDQ $16, SP                   // 回收当前栈空间
//   RET

// 2. 尝试去掉原汇编指令中的NOSPLIT标志
// 执行 go tool asm -S main_amd64.s 编译成GO汇编
// 编译结果
// TEXT "".printnl_nosplit(SB), $16
// L_BEGIN:
//     MOVQ (TLS), CX                   // 用来加载g结构体指针
//     CMPQ SP, 16(CX)                  // SP 栈指针和g结构体中stackguard0成员比较
//     JLS  L_MORE_STK                  // 若结果小于0则跳转至L_MORE_STK获取更多的栈空间
//
//         SUBQ $16, SP
//             MOVQ BP, 8(SP)
//             LEAQ 8(SP), BP
//                 CALL runtime.printnl(SB)
//             MOVQ 8(SP), BP
//         ADDQ $16, SP
//
// L_MORE_STK:
//     CALL runtime.morestack_noctxt(SB) // 获取更多的栈空间
//     JMP  L_BEGIN                      // 栈空间获取完毕后，重新跳转至L_BEGIN开始重新执行函数
// RET
