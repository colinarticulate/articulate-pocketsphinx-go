#include <stdio.h>
#include <string.h>
#include <assert.h>

#include "ps.h"

const char *jsgf_filename="./../data/_kl_ay_m__from_wrapper_from_c.jsgf";
const char *wav_filename="./../data/_climb1_colin__from_wrapper_from_c.wav";
const char *params_filename="./../data/_params__from_wrapper_from_c.txt";
const char *c_filename="./../data/_file_from_c.txt";
const char *c_binary_filename="./../data/_binary_file_from_c.wav";

void create_file(char *buffer, int len, const char *filename) {
    //printf("Just called a function\n");
    FILE *file;// = NULL;
    int k = 0;
    //printf("About to open a file for writing.\n");
    file =fopen(filename, "wb");
    if (file == NULL) {
        printf("Failed to open %s for writing", filename);

    }
    //printf("About to write the file.");
    k = fwrite(buffer, sizeof(char), len, file);
    //printf("Just wrote the file.");
    fclose(file);
}

int passing_bytes(char *buffer, int len) {

  create_file(buffer, len, c_binary_filename);

//   for(int i = 0; i< len; i++)
//     printf("%c",buffer[i]);


  return len;
}

//void create_file_params(char *argv[], int argc, const char *filename){
int create_file_params(int argc, char *argv[], char *filename){
        //printf("Just called a function\n");

    FILE *file;// = NULL;
    int k = 0;
    //printf("About to open a file for writing.\n");
    //file =fopen(filename, "wb");
    file =fopen(filename, "wb");
    if (file == NULL) {
        printf("Failed to open %s for writing", filename);

    }

    for(int i=0; i<argc; i++) {
        fprintf(file, "%d\t%s\n", i, argv[i]);
    }

    fclose(file);

    return 0;
}

//void create_file_params(char *argv[], int argc, const char *filename){
int create_file_params_nofilename(int argc, char *argv[]){

    FILE *file;// = NULL;
    int k = 0;

    file =fopen(c_filename, "wb");
    if (file == NULL) {
        printf("Failed to open %s for writing", c_filename);

    }

    for(int i=0; i<argc; i++) {
        fprintf(file, "%d\t%s\n", i, argv[i]);
    }

    fclose(file);

    return 0;
}

int check_string(char *c_string) {

     FILE *file;// = NULL;
    int k = 0;
    //printf("About to open a file for writing.\n");
    //file =fopen(filename, "wb");
    file =fopen("./../data/test_go_c_string.txt", "w");
    if (file == NULL) {
        printf("Failed to open test file for writing");
        return 1;
    }

    
    fprintf(file, "filename from go: \n %s\n\n", c_string);
    fprintf(file, "writing more things just for the sake of it...\n");
    fflush(file);

    fclose(file);
    

    return 0;
}

//result_t ps_call(char* jsgf_buffer, int16* audio_buffer, int argc, char *argv[]);
int ps_call(void* jsgf_buffer, int jsgf_buffer_size, void* audio_buffer, int audio_buffer_size, int argc, char *argv[]){

    create_file(jsgf_buffer, jsgf_buffer_size, jsgf_filename);
    create_file(audio_buffer, audio_buffer_size, wav_filename);
    create_file_params(argc, argv, (char *)params_filename);
    check_string((char*)params_filename);

    ps_call_from_go(jsgf_buffer, (size_t)jsgf_buffer_size, audio_buffer, (size_t)audio_buffer_size, argc, argv);

    return 0;
} 

