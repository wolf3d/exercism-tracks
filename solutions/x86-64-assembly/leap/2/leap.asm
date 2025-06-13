section .text
global leap_year
leap_year:
        xor     rdx, rdx                ;
        mov     rax, rdi                ;
        mov     rcx, 4
        div     rcx
        cmp     rdx, 0                  ; compare division remainder with 0
        jne     branchA                 ; year x is not a leap year
        cmp     rdx, 0
        je      branchB                 ; year x probably is a leap year, needs additional checks
        ret                             
branchB:
       mov     rax, rdi                 ;
       mov     rcx, 100                 ;
       div rcx                          ; divide by 100
       cmp rdx, 0                       ; compare division remainder with 0
       jne branchC                      ; division by 100 produced a remainder greater than 0
       cmp rdx, 0
       je branchD                       ; division by 100 produced a remainder equal to 0, need to check one more thing

branchA:
       mov rax, 0
       ret

branchC:
       mov rax, 1                       ; year x is a leap year
       ret

branchD:
       mov     rax, rdi                 ; result (rax) initially holds x
       mov     rcx, 400
       div rcx
       cmp rdx, 0
       jne branchA                      ; year x is not divisible by 400
       cmp rdx, 0
       je branchC                       ; year x is divisible by 400


       

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif