#include <stdio.h>
#include <string.h>
#include <assert.h>
#include <errno.h> 

#if defined(_WIN32) && !defined(__CYGWIN__)
#include <windows.h>
#else
#include <sys/select.h>
#endif

#include <xyzsphinxbase/err.h>
#include <xyzsphinxbase/ad.h>

#include "pocketsphinx.h"

POCKETSPHINX_EXPORT
//ps_decoder_t *ps_init_buffered(cmd_ln_t *config, void *buffer, size_t size);

// #include "sphinxbase/include/sphinxbase/err.h"
// #include "sphinxbase/include/sphinxbase/ad.h"

// #include "pocketsphinx/include/pocketsphinx.h"

//#include "continuous.h"

static const arg_t cont_args_def[] = {
    POCKETSPHINX_OPTIONS,
    /* Argument file. */
    {"-argfile",
     ARG_STRING,
     NULL,
     "Argument file giving extra arguments."},
    {"-adcdev",
     ARG_STRING,
     NULL,
     "Name of audio device to use for input."},
    {"-infile",
     ARG_STRING,
     NULL,
     "Audio file to transcribe."},
    {"-inmic",
     ARG_BOOLEAN,
     "no",
     "Transcribe audio from microphone."},
    {"-time",
     ARG_BOOLEAN,
     "no",
     "Print word times in file transcription."},
    CMDLN_EMPTY_OPTION
};

static ps_decoder_t *ps;
static cmd_ln_t *config;
static FILE *rawfd;

typedef struct result_s result_t;

struct result_s {
    int n;
    int* starts;
    int* ends;
    int* phonemes_idx;

};

static void
print_word_times()
{
    int frame_rate = cmd_ln_int32_r(config, "-frate");
    ps_seg_t *iter = ps_seg_iter(ps);
    while (iter != NULL) {
        int32 sf, ef, pprob;
        float conf;

        ps_seg_frames(iter, &sf, &ef);
        pprob = ps_seg_prob(iter, NULL, NULL, NULL);
        conf = logmath_exp(ps_get_logmath(ps), pprob);
        printf("%s %.3f %.3f %f\n", ps_seg_word(iter), ((float)sf / frame_rate),
               ((float) ef / frame_rate), conf);
        iter = ps_seg_next(iter);
    }
}

static int
check_wav_header(char *header, int expected_sr)
{
    int sr;

    if (header[34] != 0x10) {
        E_ERROR("Input audio file has [%d] bits per sample instead of 16\n", header[34]);
        return 0;
    }
    if (header[20] != 0x1) {
        E_ERROR("Input audio file has compression [%d] and not required PCM\n", header[20]);
        return 0;
    }
    if (header[22] != 0x1) {
        E_ERROR("Input audio file has [%d] channels, expected single channel mono\n", header[22]);
        return 0;
    }
    sr = ((header[24] & 0xFF) | ((header[25] & 0xFF) << 8) | ((header[26] & 0xFF) << 16) | ((header[27] & 0xFF) << 24));
    if (sr != expected_sr) {
        E_ERROR("Input audio file has sample rate [%d], but decoder expects [%d]\n", sr, expected_sr);
        return 0;
    }
    return 1;
}


/*
 * Continuous recognition from a buffered file
 */
static void
recognize_from_buffered_file(void* audio_buffer, size_t bsize)
{
    int16 adbuf[2048];
    const char *fname;
    const char *hyp;
    int32 k;
    uint8 utt_started, in_speech;
    int32 print_times = cmd_ln_boolean_r(config, "-time");



    fname = cmd_ln_str_r(config, "-infile");
    // if ((rawfd = fopen(fname, "rb")) == NULL) {
    //     E_FATAL_SYSTEM("Failed to open file '%s' for reading",
    //                    fname);
    // }
    FILE* file = NULL;
    file = fmemopen(audio_buffer, bsize ,"rb");
    // FILE* fresult = NULL;
    // fresult = fopen("./result.txt","w");
    // if (fresult == NULL ) {
    //     printf("Couldn't open file for results.\n");
    // }
    
    //------------------- Needs better checking for wav format -----------------------------------------
    if (strlen(fname) > 4 && strcmp(fname + strlen(fname) - 4, ".wav") == 0) {
        char waveheader[44];
        k=fread(waveheader, 1, 44, file); //warning:  ignoring return value of ‘fread’
    
	if (!check_wav_header(waveheader, (int)cmd_ln_float32_r(config, "-samprate")))
    	    E_FATAL("Failed to process file '%s' due to format mismatch.\n", fname);
    }

    if (strlen(fname) > 4 && strcmp(fname + strlen(fname) - 4, ".mp3") == 0) {
	E_FATAL("Can not decode mp3 files, convert input file to WAV 16kHz 16-bit mono before decoding.\n");
    }
    //---------------------------------------------------------------------------------------------------
    
    ps_start_utt(ps);
    utt_started = FALSE;

    while ((k = fread(adbuf, sizeof(int16), 2048, file)) > 0) {
        ps_process_raw(ps, adbuf, k, FALSE, FALSE);
        in_speech = ps_get_in_speech(ps);
        if (in_speech && !utt_started) {
            utt_started = TRUE;
        } 
        if (!in_speech && utt_started) {
            ps_end_utt(ps);
            hyp = ps_get_hyp(ps, NULL);
            if (hyp != NULL)
        	printf("%s\n", hyp);
            
            if (print_times)
        	print_word_times();
            fflush(stdout);

            ps_start_utt(ps);
            utt_started = FALSE;
        }
    }
    ps_end_utt(ps);
    if (utt_started) {
        hyp = ps_get_hyp(ps, NULL);
        if (hyp != NULL) {
    	    //printf("%s\n", hyp);
            //fprintf(fresult, "%s\n", hyp);
    	    if (print_times) {
    		//print_word_times();
	    }
	}
    }
    //fclose(fresult);
    fclose(file);
}

/*
 * Continuous recognition from a file
 */
static void
recognize_from_file()
{
    int16 adbuf[2048];
    const char *fname;
    const char *hyp;
    int32 k;
    uint8 utt_started, in_speech;
    int32 print_times = cmd_ln_boolean_r(config, "-time");

    fname = cmd_ln_str_r(config, "-infile");
    if ((rawfd = fopen(fname, "rb")) == NULL) {
        E_FATAL_SYSTEM("Failed to open file '%s' for reading",
                       fname);
    }
    
    if (strlen(fname) > 4 && strcmp(fname + strlen(fname) - 4, ".wav") == 0) {
        char waveheader[44];
    fread(waveheader, 1, 44, rawfd); //warning:  ignoring return value of ‘fread’
    
	if (!check_wav_header(waveheader, (int)cmd_ln_float32_r(config, "-samprate")))
    	    E_FATAL("Failed to process file '%s' due to format mismatch.\n", fname);
    }

    if (strlen(fname) > 4 && strcmp(fname + strlen(fname) - 4, ".mp3") == 0) {
	E_FATAL("Can not decode mp3 files, convert input file to WAV 16kHz 16-bit mono before decoding.\n");
    }
    
    ps_start_utt(ps);
    utt_started = FALSE;

    while ((k = fread(adbuf, sizeof(int16), 2048, rawfd)) > 0) {
        ps_process_raw(ps, adbuf, k, FALSE, FALSE);
        in_speech = ps_get_in_speech(ps);
        if (in_speech && !utt_started) {
            utt_started = TRUE;
        } 
        if (!in_speech && utt_started) {
            ps_end_utt(ps);
            hyp = ps_get_hyp(ps, NULL);
            if (hyp != NULL)
        	printf("%s\n", hyp);
            if (print_times)
        	print_word_times();
            fflush(stdout);

            ps_start_utt(ps);
            utt_started = FALSE;
        }
    }
    ps_end_utt(ps);
    if (utt_started) {
        hyp = ps_get_hyp(ps, NULL);
        if (hyp != NULL) {
    	    printf("%s\n", hyp);
    	    if (print_times) {
    		print_word_times();
	    }
	}
    }
    
    fclose(rawfd);
}

/* Sleep for specified msec */
static void
sleep_msec(int32 ms)
{
#if (defined(_WIN32) && !defined(GNUWINCE)) || defined(_WIN32_WCE)
    Sleep(ms);
#else
    /* ------------------- Unix ------------------ */
    struct timeval tmo;

    tmo.tv_sec = 0;
    tmo.tv_usec = ms * 1000;

    select(0, NULL, NULL, NULL, &tmo);
#endif
}

/*
 * Main utterance processing loop:
 *     for (;;) {
 *        start utterance and wait for speech to process
 *        decoding till end-of-utterance silence will be detected
 *        print utterance result;
 *     }
 */
static void
recognize_from_microphone()
{
    ad_rec_t *ad;
    int16 adbuf[2048];
    uint8 utt_started, in_speech;
    int32 k;
    char const *hyp;

    if ((ad = ad_open_dev(cmd_ln_str_r(config, "-adcdev"),
                          (int) cmd_ln_float32_r(config,
                                                 "-samprate"))) == NULL)
        E_FATAL("Failed to open audio device\n");
    if (ad_start_rec(ad) < 0)
        E_FATAL("Failed to start recording\n");

    if (ps_start_utt(ps) < 0)
        E_FATAL("Failed to start utterance\n");
    utt_started = FALSE;
    E_INFO("Ready....\n");

    for (;;) {
        if ((k = ad_read(ad, adbuf, 2048)) < 0)
            E_FATAL("Failed to read audio\n");
        ps_process_raw(ps, adbuf, k, FALSE, FALSE);
        in_speech = ps_get_in_speech(ps);
        if (in_speech && !utt_started) {
            utt_started = TRUE;
            E_INFO("Listening...\n");
        }
        if (!in_speech && utt_started) {
            /* speech -> silence transition, time to start new utterance  */
            ps_end_utt(ps);
            hyp = ps_get_hyp(ps, NULL );
            if (hyp != NULL) {
                printf("%s\n", hyp);
                fflush(stdout);
            }

            if (ps_start_utt(ps) < 0)
                E_FATAL("Failed to start utterance\n");
            utt_started = FALSE;
            E_INFO("Ready....\n");
        }
        sleep_msec(100);
    }
    ad_close(ad);
}

// int append_field(char* str, const char* conststr, const char* field_separator, int previous_size) {
//     int separatorsize =strlen(field_separator);
//     int fieldsize = strlen(conststr);
//     int size = previous_size + fieldsize + separatorsize;

//     str=(char*)realloc(str, sizeof(char)*(size+1)); //+1 for the '\0' character
//     memcpy(str+previous_size, conststr, sizeof(char)*fieldsize);
//     memcpy(str+previous_size+fieldsize, field_separator, sizeof(char)*(separatorsize)); 
//     str[size]='\0';
//     return size;
// }


// // //void retrieve_results(ps_decoder_t *ps){
int retrieve_results(char *sresult){

    char buffer[256];
    buffer[0]='\0';
    /* Log a backtrace if requested. */
    if (cmd_ln_boolean_r(config, "-backtrace")) {
        // FILE *fresult=NULL;
        // fresult=fopen("result.txt","w");
        // if (fresult==NULL){
        //     printf("Couldn't open file for results.");
        // }

        ps_seg_t *seg;
        int32 score;

        const char *hyp = ps_get_hyp(ps, &score);

        if (hyp != NULL) {
    	    //E_INFO("%s (%d)\n", hyp, score);
            sprintf(buffer, "%s*%d*", hyp, score);
            strcat(sresult, buffer);
    	    //E_INFO_NOFN("%-20s %-5s %-5s\n", "word", "start", "end");

    	    // fprintf(fresult, "%s (%d)\n", hyp, score);
            // fflush(fresult);
    	    // fprintf(fresult, "%-20s %-5s %-5s\n", "word", "start", "end");
            // fflush(fresult);
 

    	    for ( seg = ps_seg_iter(ps); seg; seg = ps_seg_next(seg) ) {
                int sf, ef;
                char const *word = ps_seg_word(seg);
                ps_seg_frames(seg, &sf, &ef);
                //E_INFO_NOFN("%-20s %-5d %-5d\n", word, sf, ef);
                //printf("%-20s %-5d %-5d\n", word, sf, ef);
                //strcpy(buffer,word);
                if (sf!=ef) { //for some obscure reason this if (meant to discard (NULL) entries) breaks the hash table when ps gets free.
                
                    //fprintf(fresult, "%-20s %-5d %-5d\n", word, sf, ef);
                    sprintf(buffer, "%s,%d,%-d*", word, sf, ef);
                    strcat(sresult, buffer);
                    //printf("%s\n", sresult);
                    
                    //fflush(fresult);
                }
    	    }
            strcat(sresult,"*");
        }
        
        //err=fclose(fresult);

        // if(fclose(fresult) != 0)
        // {
        //     fprintf(stderr, "Error closing file: %s", strerror(errno));
        // }



    } 

    return strlen(sresult);
}


int ps_call_from_go(void* jsgf_buffer, size_t jsgf_buffer_size, void* audio_buffer, size_t audio_buffer_size, int argc, char *argv[], char* sresult)
{ 
    char const *cfg;
    int stringsize=0;
    

    config = cmd_ln_parse_r(NULL, cont_args_def, argc, argv, TRUE);

    /* Handle argument file as -argfile. */
    if (config && (cfg = cmd_ln_str_r(config, "-argfile")) != NULL) {
        config = cmd_ln_parse_file_r(config, cont_args_def, cfg, FALSE);
    }

    if (config == NULL || (cmd_ln_str_r(config, "-infile") == NULL && cmd_ln_boolean_r(config, "-inmic") == FALSE)) {
	    E_INFO("Specify '-infile <file.wav>' to recognize from file or '-inmic yes' to recognize from microphone.\n");
        cmd_ln_free_r(config);
	    return 1;
    }

    ps_default_search_args(config);
    ps = ps_init_buffered(config, jsgf_buffer, jsgf_buffer_size);
    if (ps == NULL) {
        cmd_ln_free_r(config);
        return 1;
    }

    E_INFO("%s COMPILED ON: %s, AT: %s\n\n", argv[0], __DATE__, __TIME__);

    if (cmd_ln_str_r(config, "-infile") != NULL) {
        //recognize_from_file();
        recognize_from_buffered_file(audio_buffer, audio_buffer_size);
    } else if (cmd_ln_boolean_r(config, "-inmic")) {
        recognize_from_microphone();
    } else {
        recognize_from_file();
    }

    stringsize = retrieve_results(sresult);
    /* Log a backtrace if requested. */
    // if (cmd_ln_boolean_r(config, "-backtrace")) {
    //     ps_seg_t *seg;
    //     int32 score;

    //     const char *hyp = ps_get_hyp(ps, &score);
        
    //     if (hyp != NULL) {
    // 	    E_INFO("%s (%d)\n", hyp, score);
    // 	    E_INFO_NOFN("%-20s %-5s %-5s\n", "word", "start", "end");

    // 	    for ( seg = ps_seg_iter(ps); seg; seg = ps_seg_next(seg) ) {
    //             int sf, ef;
    //             char const *word = ps_seg_word(seg);
    //             ps_seg_frames(seg, &sf, &ef);
    //             E_INFO_NOFN("%-20s %-5d %-5d\n", word, sf, ef);
    // 	    }
    //     }
    // }


    ps_free(ps);
    cmd_ln_free_r(config);

    //return stringsize;
    return stringsize;

}

