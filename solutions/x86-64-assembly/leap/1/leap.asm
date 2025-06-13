section .text
global leap_year
leap_year:
        xor     rdx, rdx                ;
        mov     rax, rdi                ; result (rax) initially holds x
        mov     rcx, 4
        div     rcx
        cmp     rdx, 0
        jne     branchA
        cmp     rdx, 0
        je      branchB
        ret                             ; the max will be in rax
branchB:
       mov     rax, rdi                 ; result (rax) initially holds x
       mov     rcx, 100
       div rcx
       cmp rdx, 0
       jne branchC
       cmp rdx, 0
       je branchD

branchA:
       mov rax, 0
       ret

branchC:
       mov rax, 1
       ret

branchD:
       mov     rax, rdi                 ; result (rax) initially holds x
       mov     rcx, 400
       div rcx
       cmp rdx, 0
       jne branchA
       cmp rdx, 0
       je branchC


       

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif