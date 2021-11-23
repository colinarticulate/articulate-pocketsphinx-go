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

char *argv1[]= {
		"pocketsphinx_continuous",
		"-nwpen", "1",
		"-backtrace", "yes",
		"-maxwpf", "-1",
		"-lw", "6",
		"-featparams", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us/feat.params",
        "-hmm", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us",
		"-dict", "/home/dbarbera/Data/art_db.phone",
		"-fwdflat", "no",
		"-wlen", "0.016",
		"-frate", "125",
		"-wbeam", "1e-10000",
		"-remove_silence", "no",
		"-vad_postspeech", "25",
		"-doublebw", "no",
		"-vad_threshold", "1",
		"-fsgusefiller", "no",
		//"-jsgf", "/home/dbarbera/Repositories/test_pronounce/audio_clips/Temp_7302731c-188f-4dfe-83ce-a5a004f1cab2/forced_align_450654b3-3a8f-4709-821f-9341848ccd86_climb1_colin_fixed_trimmed.wav.jsgf", "-pl_window", "0", "-beam", "1e-10000", "-lponlybeam", "1e-10000", "-pbeam", "1e-10000", "-vad_startspeech", "8", "-alpha", "0.97", "-pip", "1", "-bestpath", "no", "-lpbeam", "1e-10000", "-maxhmmpf", "-1",
		"-jsgf", "/home/dbarbera/Data/climb/forced_align_000_climb1_colin_fixed_trimmed.wav.jsgf",
		"-pl_window", "0", "-beam", "1e-10000", "-lponlybeam", "1e-10000", "-pbeam", "1e-10000", "-vad_startspeech", "8", "-alpha", "0.97", "-pip", "1", "-bestpath", "no", "-lpbeam", "1e-10000", "-maxhmmpf", "-1",
		//"-infile", "/home/dbarbera/Repositories/test_pronounce/audio_clips/climb1_colin_fixed_trimmed.wav",
		"-infile", "/home/dbarbera/Data/climb/climb1_colin_fixed_trimmed.wav",
		"-cmninit", "43.46,-0.55,-4.37,11.73,-6.42,8.67,-8.58,-7.35,-0.16,2.92,6.63,0.05,4.06",
		"-vad_prespeech", "5",
		"-dither", "no",
		"-topn", "4",
		"-remove_noise", "yes",
		"-remove_dc", "no",
		"-nfft", "256",
		"-logfn", "/home/dbarbera/Data/climb/output/forced_align_000_climb1_colin_fixed_trimmed.log",
		"-cmn", "-live",
		"-wip", "0.5"
	}; //Just

int argc1 = 77;

char * argv2[] = {
		"pocketsphinx_continuous",
		"-nwpen", "1",
		"-backtrace", "yes",
		"-maxwpf", "-1",
		"-lw", "6",
		"-featparams", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us/feat.params",
		"-hmm", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us",
		"-lm", "/usr/local/share/xyzpocektsphinx/model/en-us/en-us.lm.bin",
		"-dict", "/home/dbarbera/Data/art_db.phone",
		"-fwdflat", "no",
		"-wlen", "0.016",
		"-frate", "125",
		"-wbeam", "1e-10000",
		"-remove_silence", "no",
		"-vad_postspeech", "25",
		"-doublebw", "no",
		"-vad_threshold", "1",
		"-fsgusefiller", "no",
		//"-jsgf", "/home/dbarbera/Repositories/test_pronounce/audio_clips/Temp_7302731c-188f-4dfe-83ce-a5a004f1cab2/forced_align_450654b3-3a8f-4709-821f-9341848ccd86_climb1_colin_fixed_trimmed.wav.jsgf", "-pl_window", "0", "-beam", "1e-10000", "-lponlybeam", "1e-10000", "-pbeam", "1e-10000", "-vad_startspeech", "8", "-alpha", "0.97", "-pip", "1", "-bestpath", "no", "-lpbeam", "1e-10000", "-maxhmmpf", "-1",
		"-jsgf", "/home/dbarbera/Data/climb/forced_align_000_climb1_colin_fixed_trimmed.wav.jsgf",
		"-pl_window", "0", "-beam", "1e-10000", "-lponlybeam", "1e-10000", "-pbeam", "1e-10000", "-vad_startspeech", "8", "-alpha", "0.97", "-pip", "1", "-bestpath", "no", "-lpbeam", "1e-10000", "-maxhmmpf", "-1",
		//"-infile", "/home/dbarbera/Repositories/test_pronounce/audio_clips/climb1_colin_fixed_trimmed.wav",
		"-infile", "/home/dbarbera/Data/climb/climb1_colin_fixed_trimmed.wav",
		"-cmninit", "43.46,-0.55,-4.37,11.73,-6.42,8.67,-8.58,-7.35,-0.16,2.92,6.63,0.05,4.06",
		"-vad_prespeech", "5",
		"-dither", "no",
		"-topn", "4",
		"-remove_noise", "yes",
		"-remove_dc", "no",
		"-nfft", "256",
		"-logfn", "/home/dbarbera/Data/climb/output/forced_align_000_climb1_colin_fixed_trimmed.log",
		"-cmn", "-live",
		"-wip", "0.5"
	}; //Just

int argc2 = 79;

int number_parameters(char *params[]) {

    int i=0;
    int count = 0;
    while(params[i]!=NULL) {
        count++;
        i++;
    }

    return count;
}

char* get_value(char *params[], char const *key) {
    char *value=NULL;
    int i=0;
    int count = 0;
    while((params[i]!=NULL) && (strcmp(params[i], key) != 0)) {
        count++;
        i++;
    }
    if(params[i+1] != NULL){
        value = (char*)malloc(strlen(params[i+1])*sizeof(char)+1);
        //memcpy(value, params[i], strlen(params[i+1])*sizeof(char)+1);
        strcpy(value, params[i+1]);
      //  printf(" key: %s value: %s\n", key, value);
        return value;
    } else {
        return NULL;
    }
    //return params[i+1];
}

char *params125[] = {
"pocketsphinx_continuous",
"-alpha", "0.97",
"-backtrace", "yes",
"-beam", "1e-10000",
"-bestpath", "no",
"-cmn", "live",
"-cmninit", "48.66,4.31,-7.12,5.61,-1.63,9.01,-4.65,-17.99,-16.52,-5.18,3.45,2.53,-1.34",
"-dict", "/home/dbarbera/Data/art_db.phone",
"-dither", "no",
"-doublebw", "no",
"-featparams", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us/feat.params",
"-frate", "125",
"-fsgusefiller", "no",
"-fwdflat", "no",
"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_a3ecf04d-a77a-4269-9eb5-395f8dfbdd8a_allowed1_philip_fixed_trimmed.wav.jsgf",
"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/e1e0d844-812b-496c-83fb-712de847f8a7_a3ecf04d_frate_125_debug_from_c_.log",
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

char *params72[] = {
"pocketsphinx_continuous",
"-alpha", "0.97",
"-backtrace", "yes",
"-beam", "1e-10000",
"-bestpath", "no",
"-cmn", "live",
"-cmninit", "61.61,8.03,-6.54,4.13,-3.74,9.61,-5.77,-16.52,-13.85,-3.98,2.30,2.59,-1.94",
"-dict", "/home/dbarbera/Data/art_db.phone",
"-dither", "yes",
"-doublebw", "yes",
"-featparams", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us/feat.params",
"-frate", "72",
"-fsgusefiller", "no",
"-fwdflat", "no",
"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_2668db47-d3ce-4760-ab4b-60b9b8a6c46e_allowed1_philip_fixed_trimmed.wav.jsgf",
"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/d0c65f23-d9cc-4047-8f3c-3a91db3623ff_2668db47_frate_72_debug_from_c_.log",
"-lpbeam", "1e-10000",
"-lponlybeam", "1e-10000",
"-lw", "6",
"-maxhmmpf", "-1",
"-maxwpf", "-1",
"-nfft", "512",
"-nwpen", "1",
"-pbeam", "1e-10000",
"-pip", "1.15",
"-pl_window", "0",
"-remove_dc", "no",
"-remove_noise", "yes",
"-remove_silence", "yes",
"-topn", "6",
"-vad_postspeech", "20",
"-vad_prespeech", "5",
"-vad_startspeech", "5",
"-vad_threshold", "1.5",
"-wbeam", "1e-10000",
"-wip", "0.25",
"-wlen", "0.032"
};

char *params80[] = {
"pocketsphinx_continuous",
"-alpha", "0.97",
"-backtrace", "yes",
"-beam", "1e-10000",
"-bestpath", "no",
"-cmn", "live",
"-cmninit", "60.83,7.43,-6.07,4.54,-3.88,9.48,-5.85,-16.81,-13.89,-3.82,2.39,2.53,-1.88",
"-dict", "/home/dbarbera/Data/art_db.phone",
"-dither", "yes",
"-doublebw", "yes",
"-featparams", "/usr/local/share/pocketsphinx/model/en-us/en-us/feat.params",
"-frate", "80",
"-fsgusefiller", "no",
"-fwdflat", "no",
"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_4311b957-22a7-446c-85d9-d154d4156d02_allowed1_philip_fixed_trimmed.wav.jsgf",
"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/4f0e00fa-9d20-4096-8bb5-8aeedc110e52_4311b957_frate_80_debug_from_c_.log",
"-lpbeam", "1e-10000",
"-lponlybeam", "1e-10000",
"-lw", "6",
"-maxhmmpf", "-1",
"-maxwpf", "-1",
"-nfft", "512",
"-nwpen", "1",
"-pbeam", "1e-10000",
"-pip", "1.15",
"-pl_window", "0",
"-remove_dc", "no",
"-remove_noise", "yes",
"-remove_silence", "yes",
"-topn", "6",
"-vad_postspeech", "20",
"-vad_prespeech", "5",
"-vad_startspeech", "5",
"-vad_threshold", "1.5",
"-wbeam", "1e-10000",
"-wip", "0.25",
"-wlen", "0.028"
};

char *params91[] = {
"pocketsphinx_continuous",
"-alpha", "0.97",
"-backtrace", "yes",
"-beam", "1e-10000",
"-bestpath", "no",
"-cmn", "live",
"-cmninit", "55.22,5.33,-6.78,5.07,-2.13,9.10,-4.03,-17.60,-16.77,-6.25,2.43,1.62,-1.56",
"-dict", "/home/dbarbera/Data/art_db.phone",
"-dither", "no",
"-doublebw", "no",
"-featparams", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us/feat.params",
"-frate", "91",
"-fsgusefiller", "no",
"-fwdflat", "no",
"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_1eed3902-f7e5-444b-a4b8-29b5c47ea52e_allowed1_philip_fixed_trimmed.wav.jsgf",
"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/f7a95619-d7c5-42a9-b548-561187b350da_1eed3902_frate_91_debug_from_c_.log",
"-lpbeam", "1e-10000",
"-lponlybeam", "1e-10000",
"-lw", "6",
"-maxhmmpf", "-1",
"-maxwpf", "-1",
"-nfft", "512",
"-nwpen", "1",
"-pbeam", "1e-10000",
"-pip", "1",
"-pl_window", "0",
"-remove_dc", "no",
"-remove_noise", "yes",
"-remove_silence", "no",
"-topn", "4",
"-vad_postspeech", "20",
"-vad_prespeech", "5",
"-vad_startspeech", "5",
"-vad_threshold", "0.5",
"-wbeam", "1e-10000",
"-wip", "0.5",
"-wlen", "0.024"
};



char *params105[] = {
"pocketsphinx_continuous",
"-alpha", "0.97",
"-backtrace", "yes",
"-beam", "1e-10000",
"-bestpath", "no",
"-cmn", "live",
"-cmninit", "53.92,4.73,-7.01,5.40,-1.92,8.97,-4.24,-17.95,-17.00,-6.15,2.58,1.61,-1.69",
"-dict", "/home/dbarbera/Data/art_db.phone",
"-dither", "no",
"-doublebw", "no",
"-featparams", "/usr/local/share/pocketsphinx/model/en-us/en-us/feat.params",
"-frate", "105",
"-fsgusefiller", "no",
"-fwdflat", "no",
"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_9ddb7131-fa08-4bc4-b44c-814b2ed9917e_allowed1_philip_fixed_trimmed.wav.jsgf",
"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/c061830a-8106-4e06-ae16-18feb072ea45_9ddb7131_frate_105_debug_from_c_.log",
"-lpbeam", "1e-10000",
"-lponlybeam", "1e-10000",
"-lw", "6",
"-maxhmmpf", "-1",
"-maxwpf", "-1",
"-nfft", "512",
"-nwpen", "1",
"-pbeam", "1e-10000",
"-pip", "1",
"-pl_window", "0",
"-remove_dc", "no",
"-remove_noise", "yes",
"-remove_silence", "no",
"-topn", "4",
"-vad_postspeech", "20",
"-vad_prespeech", "5",
"-vad_startspeech", "5",
"-vad_threshold", "1.5",
"-wbeam", "1e-10000",
"-wip", "0.5",
"-wlen", "0.020"
};
// const char* jsgf_filename_path = "/home/dbarbera/Repositories/mySphinx/data/kl_ay_m.jsgf";
// const char* audio_filename_path = "/home/dbarbera/Repositories/mySphinx/data/climb1_colin.wav";
const char* jsgf_filename_path = "/home/dbarbera/Repositories/test_pronounce/audio_clips/Temp_8f99a435-2a40-4ef6-a615-51d3d6d662e0/forced_align_f69956fa-3333-4cc7-898b-c69f1e637193_allowed1_philip_fixed_trimmed.wav.jsgf";
const char* audio_filename_path = "/home/dbarbera/Repositories/test_pronounce/audio_clips/allowed1_philip_fixed_trimmed.wav";

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

void print_params(char *params[], int n) {
    int i = 0;
    printf("\nParams: \n");
    printf("%s\n", params[i]);
    i++;
    int count_cmd_ln = 1;
    int count_flags =0;
    while((i+1) < n) {
        printf("%s %s\n", params[i], params[i+1]);
        i=i+2;
        count_cmd_ln=count_cmd_ln+2;
        count_flags++;
    }
    printf("\nCmd_ln parameters: %d\n", count_cmd_ln);
    printf("Flags passed to pocketsphinx: %d\n", count_flags);
}

int ps_call(void* jsgf_buffer, int jsgf_buffer_size, void* audio_buffer, int audio_buffer_size, int argc, char *argv[], char* result, int rsize){

    // print_params(argv, argc);
    // printf("\n");
    char *sresult=NULL;
    int resultsize=0;
 
    // create_file(jsgf_buffer, jsgf_buffer_size, jsgf_filename);
    // create_file(audio_buffer, audio_buffer_size, wav_filename);
    // create_file_params(argc, argv, (char *)params_filename);
    // check_string((char*)params_filename);

    sresult = (char*)malloc(sizeof(char)*rsize);
    sresult[0]='\0';
    resultsize=ps_call_from_go(jsgf_buffer, (size_t)jsgf_buffer_size, audio_buffer, (size_t)audio_buffer_size, argc, argv, sresult);

    if (resultsize < rsize && strlen(sresult)>0){
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

void mock_ps_call(char *params[]){

   // char *_argv[] = params125;
    int _argc = number_parameters(params);
    char *jsgf_filename_path = get_value(params, "-jsgf");
    char *audio_filename_path = get_value(params, "-infile");

    printf("num parameters: %d\n", _argc);
    printf("jsgf file: %s\n", jsgf_filename_path);
    printf("audio file: %s\n", audio_filename_path);

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
    ps_call(jsgf_buffer, jsgf_buffer_size, wav_data, wav_data_size, _argc, params, result, rsize);

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

//params72
//params125
//params91
//params80
//params105

// int
// main()
// {

//     printf("\nfrate %s\n", get_value(params72, "-frate"));
//     mock_ps_call(params72);
//     printf("\nfrate %s\n", get_value(params125, "-frate"));
//     mock_ps_call(params125);
//     printf("\nfrate %s\n", get_value(params80, "-frate"));
//     mock_ps_call(params80);
//     printf("\nfrate %s\n", get_value(params105, "-frate"));
//     mock_ps_call(params105);
//     printf("\nfrate %s\n", get_value(params91, "-frate"));
//     mock_ps_call(params91);

//     return 0;

    
// }