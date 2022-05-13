
#include <iostream>
#include <stdio.h>


#include "ps_continuous.h"
#include "ps_batch.h"

#ifdef __cplusplus
extern "C" {
#endif


int ps_plus_call2(void* jsgf_buffer, size_t jsgf_buffer_size, void* audio_buffer, size_t audio_buffer_size, size_t argc, char *argv[], char* result, size_t rsize){


    //C version:
    //resultsize=ps_call_from_go(jsgf_buffer, (size_t)jsgf_buffer_size, audio_buffer, (size_t)audio_buffer_size, argc, argv, sresult);

    //Encapsulated version:
    XYZ_PocketSphinx ps;
    ps.init(jsgf_buffer, jsgf_buffer_size, audio_buffer, audio_buffer_size, argc, argv);
    ps.init_recognition();
    ps.recognize_from_buffered_file();
    ps.terminate();

    //printf("Before the try-catch statement.");

    try{
        //exit(1);
        // const char* msg = "Testing....";
        // throw("Something went wrong!!!");
        //segfault:
        memset((char *)0x0, 1, 100);

    } catch( std::exception& e ) {
        printf("%s\n",e.what());
    }



    if (ps._result_size < rsize && strlen(ps._result)>0){

        for(int i=0;i<ps._result_size; i++){
            result[i]=(char)ps._result[i];
        }
    } 

    return ps._result_size;
 } 



// int ps_plus_call(void* jsgf_buffer, int jsgf_buffer_size, void* audio_buffer, int audio_buffer_size, int argc, char *argv[], char* result, int rsize){


//     //C version:
//     //resultsize=ps_call_from_go(jsgf_buffer, (size_t)jsgf_buffer_size, audio_buffer, (size_t)audio_buffer_size, argc, argv, sresult);

//     //Encapsulated version:
//     XYZ_PocketSphinx ps;
//     ps.init(jsgf_buffer, jsgf_buffer_size, audio_buffer, audio_buffer_size, argc, argv);
//     ps.init_recognition();
//     ps.recognize_from_buffered_file();
//     ps.terminate();

//     //printf("Before the try-catch statement.");

//     try{
//         //exit(1);
//         // const char* msg = "Testing....";
//         // throw("Something went wrong!!!");
//         //segfault:
//         memset((char *)0x0, 1, 100);

//     } catch( std::exception& e ) {
//         printf("%s\n",e.what());
//     }



//     if (ps._result_size < rsize && strlen(ps._result)>0){

//         for(int i=0;i<ps._result_size; i++){
//             result[i]=(char)ps._result[i];
//         }
//     } 

//     return ps._result_size;
//  } 


//  int ps_batch_plus_call(void* audio_buffer, int audio_buffer_size, int argc, char *argv[], char* result, int rsize){


//     //C version:
//     //resultsize=ps_call_from_go(jsgf_buffer, (size_t)jsgf_buffer_size, audio_buffer, (size_t)audio_buffer_size, argc, argv, sresult);

//     //Encapsulated version:
//     XYZ_Batch ps;
//     ps.init(audio_buffer, audio_buffer_size, argc, argv);
//     ps.init_recognition();
//     ps.process();
//     ps.terminate();



//     if (ps._result_size < rsize && strlen(ps._result)>0){

//         for(int i=0;i<ps._result_size; i++){
//             result[i]=(char)ps._result[i];
//         }
//     } 

//     return ps._result_size;
//  } 

 #ifdef __cplusplus
}
#endif