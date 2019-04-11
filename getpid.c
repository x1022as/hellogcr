#include<stdio.h>
#include<stdlib.h>
#include <sys/ipc.h>
#include <sys/msg.h>
#include <sys/types.h>
#include <unistd.h>

#define DECLARE_ARGS(val, low, high) unsigned long low, high
#define EAX_EDX_VAL(val, low, high) ((low) | (high) << 32)
#define EAX_EDX_RET(val, low, high) "=a" (low), "=d" (high)

static __always_inline unsigned long long rdtsc(void)
{
	DECLARE_ARGS(val, low, high);

	asm volatile("rdtscp" : EAX_EDX_RET(val, low, high));

	return EAX_EDX_VAL(val, low, high);
}


#define BULK (16*1024*1024)
#define TOTAL_SIZE (16*1024*1024)
#define TOTAL_TIMES 100000

void main(void)
{
	unsigned long long start,end, single, avg;
	char *testmem;
	int cnt;

	/*
	start = rdtsc();
	end = rdtsc();
	printf("tsc cost: %llu\n", end - start);
	*/

	// to eliminate the influece of first syscall
	getpid();

	start = rdtsc();
	getpid();
	end = rdtsc();
	single = end-start;

	start = rdtsc();
	for(cnt=0;cnt<TOTAL_TIMES;cnt++) {
		getpid();
	}
	end = rdtsc();
	avg=(end-start)/TOTAL_TIMES;
	printf("getpid - single takes: %llu \n", single);
	printf("getpid - avrage takes: %llu\n", avg);
}
