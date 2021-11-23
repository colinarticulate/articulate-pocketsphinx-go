#include <stdio.h>
#include <string.h>
#include <assert.h>
#include <stdlib.h>
#include <ctype.h>

#include "ps.h"

//For testing in c:
char *argv_[] = {
            "/home/dbarbera/Repositories/mySphinx/ps-debug",
            "-alpha", "0.97",
            "-backtrace", "yes",
            "-beam", "1e-10000",
            "-bestpath", "no",
            "-cmn", "live",
            "-cmninit", "52.55,0.14,-3.23,14.29,-7.74,9.03,-7.17,-6.31,-0.13,1.09,5.23,-2.69,1.01",
            "-dict", "/home/dbarbera/Repositories/mySphinx/data/art_db.phone",
            "-dither", "no",
            "-doublebw", "no",
            "-featparams", "/home/dbarbera/Repositories/mySphinx/data/en-us/en-us/feat.params",
            "-fsgusefiller", "no",
            "-frate", "125",
            "-fwdflat", "no",
            "-hmm", "/home/dbarbera/Repositories/mySphinx/data/en-us/en-us",
            "-infile", "/home/dbarbera/Repositories/mySphinx/data/climb1_colin.wav",
            "-jsgf", "/home/dbarbera/Repositories/mySphinx/data/kl_ay_m.jsgf",
            "-logfn", "/home/dbarbera/Repositories/mySphinx/data/output.log",
            "-lpbeam", "1e-10000",
            "-lponlybeam", "1e-10000",
            "-lw", "6",
            "-maxhmmpf", "-1",
            "-maxwpf", "-1",
            "-nfft", "256",
            "-nwpen", "1",
            "-pbeam", "1e-10000",
            "-pip", "1",
            "-pl_window", "0",
            "-remove_dc", "no",
            "-remove_noise", "yes",
            "-remove_silence", "no",
            "-topn", "4",
            "-vad_postspeech", "25",
            "-vad_prespeech", "5",
            "-vad_startspeech", "8",
            "-vad_threshold", "1",
            "-wbeam", "1e-10000",
            "-wip", "0.5",
            "-wlen", "0.016"
        };

int argc_ = 77;

const char* jsgf_filename_path = "/home/dbarbera/Repositories/mySphinx/data/kl_ay_m.jsgf";
const char* audio_filename_path = "/home/dbarbera/Repositories/mySphinx/data/climb1_colin.wav";

//To test how to return the result
//char *result="aaaaaaaaaa\0";



//Tested:

// const char *jsgf_filename="/home/dbarbera/Repositories/mySphinx/data/_kl_ay_m__from_wrapper_from_c.jsgf";
// const char *wav_filename="/home/dbarbera/Repositories/mySphinx/data/_climb1_colin__from_wrapper_from_c.wav";
// const char *params_filename="/home/dbarbera/Repositories/mySphinx/data/_params__from_wrapper_from_c.txt";
// const char *c_filename="/home/dbarbera/Repositories/mySphinx/data/_file_from_c.txt";
// const char *c_binary_filename="/home/dbarbera/Repositories/mySphinx/data/_binary_file_from_c.wav";

const char *jsgf_filename="./../data/_kl_ay_m__from_wrapper_from_c.jsgf";
const char *wav_filename="./../data/_climb1_colin__from_wrapper_from_c.wav";
const char *params_filename="/./../data/_params__from_wrapper_from_c.txt";
const char *c_filename="./../data/_file_from_c.txt";
const char *c_binary_filename="./../data/_binary_file_from_c.wav";
char *text_results[] = {
        "sil kl ay m v b sil (-3641)",
        "word-start-end",
        "sil-3-90",
        "(NULL)-90-90",
        "kl-91-109",
        "(NULL)-109-109",
        "ay-110-147",
        "(NULL)-147-147",
        "m-148-165",
        "(NULL)-165-165",
        "v-166-172",
        "b-173-176",
        "sil-177-245"
        };
int n_len = 13;

//redefinition:
//char *ps_call_from_go(void* jsgf_buffer, size_t jsgf_buffer_size, void* audio_buffer, size_t audio_buffer_size, int argc, char *argv[], int *result_size);

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
        printf("Failed to open test file for writing.\n");
        return 1;
    }
    printf("Stream opened!!!!!!!!!!!\n");
    
    fprintf(file, "filename from go: \n %s\n\n", c_string);
    fprintf(file, "writing more things just for the sake of it...\n");
    fprintf(file, "Just another bit as this is a differetn version where we flush the file results.\n");
    fprintf(file, "now woring locally to avoid endless commits to take effect into calling the caller.\n");

    fclose(file);
    

    return 0;
}

void* create_buffer(int* bsize, const char* filename, const char* mode){
    FILE* file = NULL;
    file = fopen(filename, mode);
    if (file == NULL) {
        E_ERROR_SYSTEM("Failed to open %s for parsing", filename);
        return 0;
    }
    fseek(file, 0, SEEK_END);
    long fsize = ftell(file);
    *bsize=fsize;
    fseek(file, 0, SEEK_SET);  /* same as rewind(f); */

    
    void* contents = (void*)malloc(fsize + 1);
    fread(contents, 1, fsize, file);
    fclose(file);

    return contents;
}

void delete_buffer(void* buffer){
    free(buffer);
}

int ps_call(void* jsgf_buffer, int jsgf_buffer_size, void* audio_buffer, int audio_buffer_size, int argc, char *argv[], char* result, int rsize){

    char *sresult=NULL;
    int resultsize=0;
 
    // create_file(jsgf_buffer, jsgf_buffer_size, jsgf_filename);
    // create_file(audio_buffer, audio_buffer_size, wav_filename);
    // create_file_params(argc, argv, (char *)params_filename);
    // check_string((char*)params_filename);

    sresult = (char*)malloc(sizeof(char)*rsize);
    sresult[0]='\0';
    resultsize=ps_call_from_go(jsgf_buffer, (size_t)jsgf_buffer_size, audio_buffer, (size_t)audio_buffer_size, argc, argv, sresult);

    if ( (resultsize < rsize) && (resultsize>1)){
//        printf("%s\n", result);
        for(int i=0;i<resultsize; i++){
            result[i]=(char)sresult[i];
            //printf("-> %c, %c\n", result[i], sresult[i]);
        }
        // printf("%s\n", sresult);
        // printf("%s\n", result);

        //Alternatively:
        //memcpy(result, sresult, sizeof(char)*resultsize);
    } 

    free(sresult);  

    return resultsize;
 } 

void modify_go_string(char *str, int len) {
//We only change the contents of the string. This seems safe.
  int i;
  
  for (i = 0; i < len; i++) {
    str[i] = (char)toupper(str[i]);
  }
}

void mock_ps_call(){

    void* jsgf_buffer = NULL;
    void* wav_data = NULL;

    int jsgf_buffer_size = 0; 
    int wav_data_size = 0;
    
    jsgf_buffer = create_buffer( &jsgf_buffer_size, jsgf_filename_path, "rb");
    wav_data = create_buffer(&wav_data_size, audio_filename_path, "rb");

    char result[512];
    memset(result,'a',sizeof(char)*512);
    int rsize = 512;
        //result_t result = ps_call(jsgf_buffer, audio_buffer, argc, argv);
    ps_call(jsgf_buffer, jsgf_buffer_size, wav_data, wav_data_size, argc_, argv_, result, rsize);

    //print_result(result);

    delete_buffer(jsgf_buffer);
    delete_buffer(wav_data);

}


//This doesnt allocate well in Go, I suppose go is not memory continuous
// int modify_go_strings(_goslicestring_ strings) {
//     int c = strings.len;
//     int n = n_len;
    
//     if (n > c) {
//         return 0;
//     } 

//     for(int i=0; i<n; i++){
//         intgo m = n_len;
//         strings.array[i].p =  (char*)realloc(strings.array[i].p, sizeof(char)*m);
//         memcpy(strings.array[i].p, text_results[i], sizeof(char)*m);
//         strings.array[i].n = m;
//     }

//     return n;
// }

// int
// main()
// {


//     mock_ps_call();

//     return 0;

    
// }