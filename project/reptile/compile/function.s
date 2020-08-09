	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 10, 15	sdk_version 10, 15
	.globl	_my_function            ## -- Begin function my_function
	.p2align	4, 0x90
_my_function:                           ## @my_function
	.cfi_startproc
## %bb.0:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	movl	%edi, -4(%rbp)
	movl	%esi, -8(%rbp)
	movl	-4(%rbp), %esi
	addl	-8(%rbp), %esi
	movl	%esi, %eax
	popq	%rbp
	retq
	.cfi_endproc
                                        ## -- End function
	.globl	_main                   ## -- Begin function main
	.p2align	4, 0x90
_main:                                  ## @main
	.cfi_startproc
## %bb.0:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	subq	$16, %rsp
	movl	$1, %edi
	movl	$2, %esi
	callq	_my_function
	xorl	%esi, %esi
	movl	%eax, -4(%rbp)
	movl	%esi, %eax
	addq	$16, %rsp
	popq	%rbp
	retq
	.cfi_endproc
                                        ## -- End function

.subsections_via_symbols

在 my_function 调用之前，调用方 main 函数将 my_function 的两个参数分别存到 edi 和 esi 寄存器中；
在 my_function 执行时，它会将寄存器 edi 和 esi 中的数据存储到 eax 和 edx 两个寄存器中，随后通过汇编指令 addl 计算两个入参之和；
在 my_function 调用之后，使用寄存器 eax 传递返回值，main 函数将 my_function 的返回值存储到栈上的 i 变量中；


