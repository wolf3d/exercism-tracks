section .text
global leap_year
leap_year:
        xor     rdx, rdx                ;
        mov     rax, rdi                ;
        mov     rcx, 4
        div     rcx
        cmp     rdx, 0                  ; compare division remainder with 0
        jne     _not_leap                 ; year x is not a leap year
        cmp     rdx, 0
        je      _div100                 ; year x probably is a leap year, needs additional checks

_div100:
       mov     rax, rdi                 ;
       mov     rcx, 100                 ;
       div rcx                          ; divide by 100
       cmp rdx, 0                       ; compare division remainder with 0
       jne _leap                      ; division by 100 produced a remainder greater than 0
       cmp rdx, 0
       je _div400                       ; division by 100 produced a remainder equal to 0, need to check one more thing

_not_leap:
       mov rax, 0
       ret

_leap:
       mov rax, 1                       ; year x is a leap year
       ret

_div400:
       mov     rax, rdi                 ; result (rax) initially holds x
       mov     rcx, 400
       div rcx
       cmp rdx, 0
       jne _not_leap                    ; year x is not divisible by 400
       cmp rdx, 0
       je _leap                       ; year x is divisible by 400


       

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif