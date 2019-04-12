#include "textflag.h"

TEXT ·Rdtsc(SB),NOSPLIT,$0-8
        // N.B. We need LFENCE on Intel, AMD is more complicated.
        // Modern AMD CPUs with modern kernels make LFENCE behave like it does
        // on Intel with MSR_F10H_DECFG_LFENCE_SERIALIZE_BIT. MFENCE is
        // otherwise needed on AMD.
        LFENCE
        RDTSC
        SHLQ    $32, DX
        ADDQ    DX, AX
        MOVQ    AX, ret+0(FP)
        RET
