#include <stdint.h> //for uintptr_t
#include <stdlib.h>
#include <stdio.h>

//A Go function
extern void MyGoPrint(uintptr_t handle);
//extern char* GoBytes(uintptr_t handle);
//A C function
void myprint_with_handle(uintptr_t handle){
    MyGoPrint(handle);
    // char *cstring = GoBytes(handle);
    // printf("From c: %s", cstring);
}