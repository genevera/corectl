#!/usr/sbin/dtrace -C -s
/*
 * eptfault.d - report all EPT faults for particular VM
 *
 * USAGE: sudo eptfault.d -p <pid of corectld.runner>
 *        sudo eptfault.d -D TOTAL -p <pid of corectld.runner>
 *        sudo eptfault.d -D INTERVAL=1 -p <pid of corectld.runner>
 *
 * If '-D INTERVAL=<seconds>' is specified periodically print a
 * summary of EPT Faults per vCPU.
 *
 * It seems that tracing starts before the end of dtrace:::BEGIN, i.e.
 * before the 'lapic_map' array is fully initialised. This results in
 * some empty/corrupted Local APIC register names being
 * printed. According to the documentation, this should not happen, so
 * it might be a bug in High Sierra.
 *
 * As a workaround, specifying '-D·NUMERIC' stores and prints the Local
 * APIC register offsets instead of names.
 */

#pragma D option quiet

string lapic_map[uint32_t];

dtrace:::BEGIN
{
        start = timestamp;

        // from src/include/xhyve/vmm/io/vlapic_priv.h
        lapic_map[0x20] = "ID";
        lapic_map[0x30] = "VER";
        lapic_map[0x80] = "TPR";
        lapic_map[0x90] = "APR";
        lapic_map[0xA0] = "PPR";
        lapic_map[0xB0] = "EOI";
        lapic_map[0xC0] = "RRR";
        lapic_map[0xD0] = "LDR";
        lapic_map[0xE0] = "DFR";
        lapic_map[0xF0] = "SVR";
        lapic_map[0x100] = "ISR0";
        lapic_map[0x110] = "ISR1";
        lapic_map[0x130] = "ISR3";
        lapic_map[0x150] = "ISR5";
        lapic_map[0x170] = "ISR7";
        lapic_map[0x190] = "TMR1";
        lapic_map[0x1B0] = "TMR3";
        lapic_map[0x1D0] = "TMR5";
        lapic_map[0x1F0] = "TMR7";
        lapic_map[0x210] = "IRR1";
        lapic_map[0x230] = "IRR3";
        lapic_map[0x250] = "IRR5";
        lapic_map[0x270] = "IRR7";
        lapic_map[0x2F0] = "CMCI_LVT";
        lapic_map[0x300] = "ICR_LOW";
        lapic_map[0x310] = "ICR_HI";
        lapic_map[0x330] = "THERM_LVT";
        lapic_map[0x340] = "PERF_LVT";
        lapic_map[0x350] = "LINT0_LVT";
        lapic_map[0x360] = "LINT1_LVT";
        lapic_map[0x370] = "ERROR_LVT";
        lapic_map[0x380] = "TIMER_ICR";
        lapic_map[0x390] = "TIMER_CCR";
        lapic_map[0x3E0] = "TIMER_DCR";
        lapic_map[0x3F0] = "SELF_IPI";

        printf("Tracing... Hit Ctrl-C to end.\n");
}

hyperkit$target:::vmx-ept-fault
/(arg1 & 0xfff00000) == 0xfee00000/
{
        // LAPIC FAULTS

        @lapic_faults[lapic_map[arg1 & 0x000fffff], arg0] = count();
}

hyperkit$target:::vmx-ept-fault
{
        @all_faults[arg1, arg0] = count();
}

corectld.runner$target:::vmx-ept-fault
{
        #ifdef INTERVAL
        /* Per vCPU count for periodic reporting */
        @total[arg0] = count();
        #endif

        @all_faults[arg1, arg0] = count();
}

#ifdef INTERVAL
/* timer */
profile:::tick-1sec
{
       secs--;
}

/* Periodically print per vCPU VM Exits */
profile:::tick-1sec
/secs == 0/
{
        printa(@total);
        clear(@total);
        secs = INTERVAL;
}
#endif

dtrace:::END
{
        #ifdef TOTAL
        printf("%18s %-4s %10s\n", "ADDRESS", "vCPU", "COUNT");
        #else
        printf("%18s %-4s %10s\n", "ADDRESS", "vCPU", "RATE (1/s)");
        normalize(@all_faults, (timestamp - start) / 1000000000);
        #endif
        printa("%18x %-4d %@10d\n", @all_faults);

        #ifdef TOTAL
        printf("%18s %-4s %10s\n", "LAPIC REGISTER", "vCPU", "COUNT");
        #else
        printf("%18s %-4s %10s\n", "LAPIC REGISTER", "vCPU", "RATE (1/s)");
        normalize(@lapic_faults, (timestamp - start) / 1000000000);
        #endif
        printa("%18s %-4d %@10d\n", @lapic_faults);
}
