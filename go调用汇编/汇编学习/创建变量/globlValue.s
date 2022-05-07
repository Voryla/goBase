#include "textflag.h" // Go汇编语言还在"textflag.h"文件定义了一些标志。其中用于变量的标志有DUPOK、RODATA和NOPTR几个
// GLOBL ·count(SB),RODATA,$8 // 全局变量的width 必须为2的指数大小  GLOBL symbol(SB), width
// 定义全局变量语法：GLOBL ·name(SB),[NOPTR|RODATA|DUPOK],$16
// GLOBL：表示该变量为全局变量对应.go文件中的var
// ·name:· 固定符号，中文模式下的按~可以输出,标识符可以由绝对的包路径加标识符本身定位，因此不同包中的标识符即使同名也不会有问题
    // GLOBL ·pkg_name1(SB),$1
    // GLOBL main·pkg_name2(SB),$1
    // GLOBL my/pkg·pkg_name(SB),$1
    // 此外，当定义仅当前文件可以访问的私有标识符时，可以使用<>后缀
    // GLOBL file_private<>(SB),$1
// [NOPTR|RODATA|DUPOK] 类型信息
    // NOPTR：标识此变量内部不包含指针数据，让垃圾回收器忽略对该变量的扫描。如果变量已经在.go文件中声明，Go编译器会自动分析出该变量是否包含指针，这种时候可以不用手写NOPTR标志。
    // RODATA：标识此变量将被定义在只读内存段，因此后续任何对比变量的修改操作将导致异常(recover也无法捕获)
    // DUPOK: 标识此变量对应的标识符可能会有多个，在链接时只选择其中一个即可(一般用于合并相同的常量字符串，减少重复数据占用的空间)
    // 必须包含三个中的其中一个或多个(多个使用|分隔)，否则会报：runtime.gcdata: missing Go type information for global symbol main.int32Value: size 4
// $16：该变量的内存大小，该值必须是2的指数大小
// 全局变量赋值语法：DATA ·name+0(SB)/width,$0
// Go汇编语言定义变量无法指定类型信息，因此需要先通过Go语言声明变量的类型

// int32 在汇编定义中，并没有正负数之说，当汇编值为16进制
// 正负数是在Go中定义后，根据数据类型是否有符号，如int32为有符号，那么会将汇编值作为补码进行解析，
// 而当uint32为无符号，直接按照汇编值进行解析
// 汇编值可以用十进制表示，汇编值为10进制时，直接采用十进制值作为变量的值,但是要解析为无符号类型时，假设uint32的汇编值为-1，那么该变量解析后的最终值为MaxUInt32-1
GLOBL ·int32Value(SB),RODATA,$4
DATA ·int32Value+0(SB)/1,$0x9c  // 第0字节
DATA ·int32Value+1(SB)/1,$0xff  // 第1字节
DATA ·int32Value+2(SB)/2,$0xffff  // 第3-4字节
GLOBL ·uint32Value(SB),RODATA,$4
DATA ·uint32Value(SB)/4,$0x00000002 // 第1-4字节
GLOBL ·nint32Value(SB),RODATA,$4
DATA ·nint32Value(SB)/4,$0xfffffff2

// bool
GLOBL ·boolValue(SB),RODATA,$1      // 定义未初始
GLOBL ·trueValue(SB),RODATA,$1      // 定义初始 true
DATA ·trueValue+0(SB)/1,$1
GLOBL ·falseValue(SB),RODATA,$1     // 定义初始 false
DATA ·falseValue+0(SB)/1,$0

// float
GLOBL ·float32Value(SB),RODATA,$4
DATA ·float32Value+0(SB)/4,$1.5        // var
GLOBL ·float64Value(SB),RODATA,$8
DATA ·float64Value+0(SB)/8,$0x01020304 // bit

// string
GLOBL ·helloworld(SB),RODATA,$16     // StringHeader 结构占用16字节
DATA ·helloworld+0(SB)/8,$text<>(SB) // StringHeader.Data
DATA ·helloworld+8(SB)/8,$12         // StringHeader.Len
GLOBL text<>(SB),NOPTR,$16           // 创建数据1.<>表示该变量为当前文件的私有变量。2.字符串的内容只有12个字节，这里设置16的原因是，内存大小必须是2的指数倍数
DATA text<>+0(SB)/8,$"Hello Wo"
DATA text<>+8(SB)/8,$"rld!"

// array
GLOBL ·num(SB), RODATA, $16         // 数组没有特殊的结构，直接解析一段内存
DATA ·num+0(SB)/8, $0
DATA ·num+8(SB)/8, $1

// slice
GLOBL ·mySlice(SB),RODATA,$24      // SliceHeader 结构体占24字节
DATA ·mySlice+0(SB)/8,$text<>(SB)  // SliceHeader.Data
DATA ·mySlice+8(SB)/8,$12          // SliceHeader.Len
DATA ·mySlice+16(SB)/8,$16         // SliceHeader.Cap

