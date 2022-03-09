// https://stackoverflow.com/questions/61821424/how-to-use-channels-to-gather-response-from-various-goroutines

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	//"os"
	//"xyz"
	//"github.com/colinarticulate/scanScheduler"
	"github.com/davidbarbera/articulate-pocketsphinx-go/xyz_plus"
)

func call_to_ps(jsgf_buffer []byte, audio_buffer []byte, params []string, c chan []xyz_plus.Utt) {

	c <- Ps(jsgf_buffer, audio_buffer, params)

}

func call_to_ps_wg_chan(jsgf_buffer []byte, audio_buffer []byte, params []string, wg *sync.WaitGroup, resultChan chan<- []xyz_plus.Utt) {
	defer wg.Done()

	resultChan <- Ps(jsgf_buffer, audio_buffer, params)

}

func collect_ps_result(c chan []xyz_plus.Utt) {
	for {
		select {
		case msg := <-c:
			fmt.Println((msg))
		}
	}
}

func process(input int, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done()

	// rand.Seed(time.Now().UnixNano())
	// n := rand.Intn(5)
	// time.Sleep(time.Duration(n) * time.Second)

	resultChan <- input //* 10
}

func concurrently_int(n int) {
	var wg sync.WaitGroup

	resultChan := make(chan int)

	for i := range []int{1, 2, 3, 4, 5} {
		wg.Add(1)
		go process(i, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var result []int
	for r := range resultChan {
		result = append(result, r)
	}

	fmt.Println(result)
}

func concurrently(frates [5]string, parameters [5][]string, jsgfs [5]string, wavs [5]string) {

	ch := make(chan []xyz_plus.Utt)
	var wg sync.WaitGroup

	//n := len(wavs)
	//wg.Add(n)
	start := time.Now()
	for i := 4; i < 5; i++ {
		jsgf_buffer, err := os.ReadFile(jsgfs[i])
		check(err)
		audio_buffer, err := os.ReadFile(wavs[i])
		check(err)
		wg.Add(1)

		go call_to_ps_wg_chan(jsgf_buffer, audio_buffer, parameters[i], &wg, ch)
	}

	// go func() {
	// 	wg.Wait()
	// 	close(ch)
	// }()

	var results [][]xyz_plus.Utt
	// go func() {
	// 	for v := range ch {
	// 		results = append(results, v)
	// 	}
	// }()

	wg.Wait()
	close(ch)
	for v := range ch {
		results = append(results, v)
	}
	elapsed := time.Since(start)

	// for i := range messages {
	// 	fmt.Println(i)
	// }

	fmt.Println("Concurrently (multithreaded-encapsulated): ")
	for result := range results {
		fmt.Println(result)
	}
	fmt.Println(results)
	fmt.Printf(">>>> Timing: %s\n", elapsed)
	fmt.Println()
	// var results[][]int
	// for r := range resultChan {
	// 	result = append(result, r)
	// }

	// fmt.Println(result)
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func test_ps(frate string, jsgf_filename string, wav_filename string, parameters []string) {
	jsgf_buffer, err := os.ReadFile(jsgf_filename)
	check(err)
	audio_buffer, err := os.ReadFile(wav_filename)
	check(err)

	fmt.Println("--- frate = ", frate)
	starti := time.Now()
	var r = Ps(jsgf_buffer, audio_buffer, parameters)
	elapsedi := time.Since(starti)
	fmt.Printf(">>>> Timing: %s\n", elapsedi)
	fmt.Println(r)
	fmt.Println()

}

func Ps(jsgf_buffer []byte, audio_buffer []byte, params []string) []xyz_plus.Utt {
	result := []string{"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"}

	xyz_plus.Ps_plus_call(jsgf_buffer, audio_buffer, params, result)

	//Adapting result from coded string to utt struct
	if strings.Contains(result[0], "**") {
		raw := strings.Split(result[0], "**")

		if len(raw) < 2 {
			fmt.Println("xyzpocketsphinx: problems!")
		}

		//fmt.Printf("%T", raw)
		fields := strings.Split(raw[0], "*")

		//fmt.Println(fields)
		// hyp := fields[0]
		// score := fields[1]

		//fmt.Println(hyp)
		//fmt.Println(strings.Split(score, ","))
		utts := []xyz_plus.Utt{}
		//var utts = make([]Utt, len(fields)-2)

		for i := 0; i < len(fields)-2; i++ {
			parts := strings.Split(fields[2:][i], ",")
			phoneme := parts[0]
			text_start := parts[1]
			text_end := parts[2]
			start, serr := strconv.Atoi(text_start)
			end, eerr := strconv.Atoi(text_end)

			if phoneme != "(NULL)" {
				//fmt.Println(phoneme, start, end)
				//utts = append(utts, xyz_plus.Utt{phoneme, int32(start), int32(end)})
				utts = append(utts, xyz_plus.Utt{Text: phoneme, Start: int32(start), End: int32(end)})

				if serr != nil || eerr != nil {
					fmt.Println(serr, eerr)
				}
			}
		}
		return utts
	} else {
		return nil
	}
}

func readParams(args []string) map[string]string {

	param_values := make(map[string]string)

	//Create dictionary: key:value == -ps_option:value
	for i := 1; i < len(args)-1; i = i + 2 {
		param_values[args[i]] = args[i+1]
	}

	//Order the dictionary by key to make it easier to inspect visually
	keys := make([]string, 0, len(param_values))
	for k := range param_values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return param_values
}

func getValue(key string, params []string) string {

	params_dict := readParams(params)

	value := params_dict[key]

	return value

}

func getParamsFromFile(file string) {
	contents, err := os.ReadFile(file)
	check(err)
	fmt.Println(string(contents))

}

func main() {
	//Data:
	//Filenames
	//filename_jsgf := "./../data/kl_ay_m.jsgf"
	//filename_wav := "./../data/climb1_colin.wav"

	// filename_jsgf1 := "/home/dbarbera/Repositories/sphinx/data/kl_ay_m.jsgf"
	// filename_wav1 := "/home/dbarbera/Repositories/mySphinx/data/climb1_colin.wav"
	// //Parameters
	// params1 := []string{
	// 	"pocketsphinx_continuous",
	// 	"-alpha", "0.97",
	// 	"-backtrace", "yes",
	// 	"-beam", "1e-10000",
	// 	"-bestpath", "no",
	// 	"-cmn", "live",
	// 	"-cmninit", "52.55,0.14,-3.23,14.29,-7.74,9.03,-7.17,-6.31,-0.13,1.09,5.23,-2.69,1.01",
	// 	"-dict", "/home/dbarbera/Repositories/mySphinx/data/art_db.phone",
	// 	"-dither", "no",
	// 	"-doublebw", "no",
	// 	"-featparams", "/home/dbarbera/Repositories/mySphinx/data/en-us/en-us/feat.params",
	// 	//"-featparams", "/usr/local/share/pocketsphinx/model/en-us/en-us/feat.params",
	// 	"-fsgusefiller", "no",
	// 	"-frate", "125",
	// 	"-fwdflat", "no",
	// 	"-hmm", "/home/dbarbera/Repositories/mySphinx/data/en-us/en-us",
	// 	//"-hmm", "/usr/local/share/pocketsphinx/model/en-us/en-us",
	// 	"-infile", "/home/dbarbera/Repositories/mySphinx/data/climb1_colin.wav",
	// 	"-jsgf", "/home/dbarbera/Repositories/sphinx/data/kl_ay_m.jsgf",
	// 	"-logfn", "/home/dbarbera/Repositories/mySphinx/data/output.log",
	// 	"-lpbeam", "1e-10000",
	// 	"-lponlybeam", "1e-10000",
	// 	"-lw", "6",
	// 	"-maxhmmpf", "-1",
	// 	"-maxwpf", "-1",
	// 	"-nfft", "256",
	// 	"-nwpen", "1",
	// 	"-pbeam", "1e-10000",
	// 	"-pip", "1",
	// 	"-pl_window", "0",
	// 	"-remove_dc", "no",
	// 	"-remove_noise", "yes",
	// 	"-remove_silence", "no",
	// 	"-topn", "4",
	// 	"-vad_postspeech", "25",
	// 	"-vad_prespeech", "5",
	// 	"-vad_startspeech", "8",
	// 	"-vad_threshold", "1",
	// 	"-wbeam", "1e-10000",
	// 	"-wip", "0.5",
	// 	"-wlen", "0.016"}

	// params2 := []string{
	// 	"pocketsphinx_continuous",
	// 	"-alpha", "0.97",
	// 	"-backtrace", "yes",
	// 	"-beam", "1e-10000",
	// 	"-bestpath", "no",
	// 	"-cmn", "live",
	// 	"-cmninit", "43.46,-0.55,-4.37,11.73,-6.42,8.67,-8.58,-7.35,-0.16,2.92,6.63,0.05,4.06",
	// 	"-dict", "/home/dbarbera/Repositories/mySphinx/data/art_db.phone",
	// 	"-dither", "no",
	// 	"-doublebw", "no",
	// 	//"-featparams", "/home/dbarbera/Repositories/mySphinx/data/en-us/en-us/feat.params",
	// 	"-featparams", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us/feat.params",
	// 	"-fsgusefiller", "no",
	// 	"-frate", "125",
	// 	"-fwdflat", "no",
	// 	//"-hmm", "/home/dbarbera/Repositories/mySphinx/data/en-us/en-us",
	// 	"-hmm", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us",
	// 	"-infile", "/home/dbarbera/Data/climb/climb1_colin_fixed_trimmed.wav",
	// 	"-jsgf", "/home/dbarbera/Data/climb/forced_align_5a02788e-3d91-446f-828a-428b7f2d8785_climb1_colin_fixed_trimmed.wav.jsgf",
	// 	"-logfn", "/home/dbarbera/Data/climb/output/output.log",
	// 	"-lpbeam", "1e-10000",
	// 	"-lponlybeam", "1e-10000",
	// 	"-lw", "6",
	// 	"-maxhmmpf", "-1",
	// 	"-maxwpf", "-1",
	// 	"-nfft", "256",
	// 	"-nwpen", "1",
	// 	"-pbeam", "1e-10000",
	// 	"-pip", "1",
	// 	"-pl_window", "0",
	// 	"-remove_dc", "no",
	// 	"-remove_noise", "yes",
	// 	"-remove_silence", "no",
	// 	"-topn", "4",
	// 	"-vad_postspeech", "25",
	// 	"-vad_prespeech", "5",
	// 	"-vad_startspeech", "8",
	// 	"-vad_threshold", "1",
	// 	"-wbeam", "1e-10000",
	// 	"-wip", "0.5",
	// 	"-wlen", "0.016"}
	// filename_jsgf2 := "/home/dbarbera/Data/climb/forced_align_5a02788e-3d91-446f-828a-428b7f2d8785_climb1_colin_fixed_trimmed.wav.jsgf"
	// filename_wav2 := "/home/dbarbera/Data/climb/climb1_colin_fixed_trimmed.wav"

	// params3 := []string{
	// 	"pocketsphinx_continuous",
	// 	"-nwpen", "1",
	// 	"-backtrace", "yes",
	// 	"-maxwpf", "-1",
	// 	"-lw", "6",
	// 	"-featparams", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us/feat.params",
	// 	"-hmm", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us",
	// 	//"-lm", "/usr/local/share/xyzpocektsphinx/model/en-us/en-us.lm.bin",
	// 	"-dict", "/home/dbarbera/Data/art_db.phone",
	// 	"-fwdflat", "no",
	// 	"-wlen", "0.016",
	// 	"-frate", "125",
	// 	"-wbeam", "1e-10000",
	// 	"-remove_silence", "no",
	// 	"-vad_postspeech", "25",
	// 	"-doublebw", "no",
	// 	"-vad_threshold", "1",
	// 	"-fsgusefiller", "no",
	// 	//"-jsgf", "/home/dbarbera/Repositories/test_pronounce/audio_clips/Temp_7302731c-188f-4dfe-83ce-a5a004f1cab2/forced_align_450654b3-3a8f-4709-821f-9341848ccd86_climb1_colin_fixed_trimmed.wav.jsgf", "-pl_window", "0", "-beam", "1e-10000", "-lponlybeam", "1e-10000", "-pbeam", "1e-10000", "-vad_startspeech", "8", "-alpha", "0.97", "-pip", "1", "-bestpath", "no", "-lpbeam", "1e-10000", "-maxhmmpf", "-1",
	// 	"-jsgf", "/home/dbarbera/Data/climb/forced_align_000_climb1_colin_fixed_trimmed.wav.jsgf",
	// 	"-pl_window", "0", "-beam", "1e-10000", "-lponlybeam", "1e-10000", "-pbeam", "1e-10000", "-vad_startspeech", "8", "-alpha", "0.97", "-pip", "1", "-bestpath", "no", "-lpbeam", "1e-10000", "-maxhmmpf", "-1",
	// 	//"-infile", "/home/dbarbera/Repositories/test_pronounce/audio_clips/climb1_colin_fixed_trimmed.wav",
	// 	"-infile", "/home/dbarbera/Data/climb/climb1_colin_fixed_trimmed.wav",
	// 	"-cmninit", "43.46,-0.55,-4.37,11.73,-6.42,8.67,-8.58,-7.35,-0.16,2.92,6.63,0.05,4.06",
	// 	"-vad_prespeech", "5",
	// 	"-dither", "no",
	// 	"-topn", "4",
	// 	"-remove_noise", "yes",
	// 	"-remove_dc", "no",
	// 	//"-nfft", "256", //this works
	// 	"-nfft", "512", //this works too
	// 	"-logfn", "/home/dbarbera/Data/climb/output/forced_align_000_climb1_colin_fixed_trimmed.log",
	// 	"-cmn", "-live",
	// 	"-wip", "0.5",
	// } //Just
	// filename_jsgf3 := "/home/dbarbera/Data/climb/forced_align_000_climb1_colin_fixed_trimmed.wav.jsgf"
	// filename_wav3 := "/home/dbarbera/Data/climb/climb1_colin_fixed_trimmed.wav"

	params72 := []string{
		"pocketsphinx_continuous",
		"-alpha", "0.97",
		"-backtrace", "yes",
		"-beam", "1e-10000",
		"-bestpath", "no",
		"-cmn", "live",
		"-cmninit", "61.61,8.03,-6.54,4.13,-3.74,9.61,-5.77,-16.52,-13.85,-3.98,2.30,2.59,-1.94",
		"-dict", "/home/dbarbera/Data/art_db.phone",
		"-dither", "yes", //yes
		"-doublebw", "yes", //yes
		"-featparams", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us/feat.params",
		"-frate", "72",
		"-fsgusefiller", "no",
		"-fwdflat", "no",
		"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
		"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_2668db47-d3ce-4760-ab4b-60b9b8a6c46e_allowed1_philip_fixed_trimmed.wav.jsgf",
		"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/d0c65f23-d9cc-4047-8f3c-3a91db3623ff_2668db47_frate_72_from_go_.log",
		"-lpbeam", "1e-10000",
		"-lponlybeam", "1e-10000",
		"-lw", "6",
		"-maxhmmpf", "-1",
		"-maxwpf", "-1",
		"-nfft", "512",
		"-nwpen", "1",
		"-pbeam", "1e-10000",
		"-pip", "1.15", //1.15
		"-pl_window", "0",
		"-remove_dc", "no",
		"-remove_noise", "yes",
		"-remove_silence", "yes", //yes
		"-topn", "6",
		"-vad_postspeech", "20",
		"-vad_prespeech", "5",
		"-vad_startspeech", "5",
		"-vad_threshold", "1.5",
		"-wbeam", "1e-10000",
		"-wip", "0.25",
		"-wlen", "0.032",
	}

	params125 := []string{
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
		"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/e1e0d844-812b-496c-83fb-712de847f8a7_a3ecf04d_frate_125_from_go_.log",
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
		"-wlen", "0.016",
	}

	params105 := []string{
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
		"-featparams", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us/feat.params",
		"-frate", "105",
		"-fsgusefiller", "no",
		"-fwdflat", "no",
		"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
		"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_9ddb7131-fa08-4bc4-b44c-814b2ed9917e_allowed1_philip_fixed_trimmed.wav.jsgf",
		"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/c061830a-8106-4e06-ae16-18feb072ea45_9ddb7131_frate_105_from_go_.log",
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
		"-wlen", "0.020",
	}

	params80 := []string{
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
		"-featparams", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us/feat.params",
		"-frate", "80",
		"-fsgusefiller", "no",
		"-fwdflat", "no",
		"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
		"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_4311b957-22a7-446c-85d9-d154d4156d02_allowed1_philip_fixed_trimmed.wav.jsgf",
		"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/4f0e00fa-9d20-4096-8bb5-8aeedc110e52_4311b957_frate_80_from_go_.log",
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
		"-remove_silence", "yes", //yes
		"-topn", "6",
		"-vad_postspeech", "20",
		"-vad_prespeech", "5",
		"-vad_startspeech", "5",
		"-vad_threshold", "1.5",
		"-wbeam", "1e-10000",
		"-wip", "0.25",
		"-wlen", "0.028",
	}

	params91 := []string{
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
		"-featparams", "/usr/local/share/pocketsphinx/model/en-us/en-us/feat.params",
		"-frate", "91",
		"-fsgusefiller", "no",
		"-fwdflat", "no",
		"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
		"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_1eed3902-f7e5-444b-a4b8-29b5c47ea52e_allowed1_philip_fixed_trimmed.wav.jsgf",
		"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/f7a95619-d7c5-42a9-b548-561187b350da_1eed3902_frate_91_from_go_.log",
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
		"-wlen", "0.024",
	}

	// paramsMinimal := []string{
	// 	"pocketsphinx_continuous",
	// 	// "-alpha", "0.97",
	// 	// "-backtrace", "yes",
	// 	// "-beam", "1e-10000",
	// 	// "-bestpath", "no",
	// 	// "-cmn", "live",
	// 	// "-cmninit", "53.92,4.73,-7.01,5.40,-1.92,8.97,-4.24,-17.95,-17.00,-6.15,2.58,1.61,-1.69",
	// 	"-dict", "/home/dbarbera/Data/art_db.phone",
	// 	// "-dither", "no",
	// 	// "-doublebw", "no",
	// 	"-featparams", "/usr/local/share/xyzpocketsphinx/model/en-us/en-us/feat.params",
	// 	// "-frate", "105",
	// 	// "-fsgusefiller", "no",
	// 	// "-fwdflat", "no",
	// 	"-infile", "/home/dbarbera/Repositories/test_pronounce/audio_clips/allowed1_philip_fixed_trimmed.wav",
	// 	"-jsgf", "/home/dbarbera/Repositories/test_pronounce/audio_clips/Temp_f1d289d6-c152-4e26-a2f4-483175d7a12c/forced_align_a434c972-810b-40fc-b9a5-9e6697dcf5c5_allowed1_philip_fixed_trimmed.wav.jsgf",
	// 	"-logfn", "/home/dbarbera/Repositories/test_pronounce/audio_clips/Temp_f1d289d6-c152-4e26-a2f4-483175d7a12c/08d56917-1343-485a-b159-98c512623710_frate_105.log",
	// 	// "-lpbeam", "1e-10000",
	// 	// "-lponlybeam", "1e-10000",
	// 	// "-lw", "6",
	// 	// "-maxhmmpf", "-1",
	// 	// "-maxwpf", "-1",
	// 	// "-nfft", "512",
	// 	// "-nwpen", "1",
	// 	// "-pbeam", "1e-10000",
	// 	// "-pip", "1",
	// 	// "-pl_window", "0",
	// 	// "-remove_dc", "no",
	// 	// "-remove_noise", "yes",
	// 	// "-remove_silence", "no",
	// 	// "-topn", "4",
	// 	// "-vad_postspeech", "20",
	// 	// "-vad_prespeech", "5",
	// 	// "-vad_startspeech", "5",
	// 	// "-vad_threshold", "1.5",
	// 	// "-wbeam", "1e-10000",
	// 	// "-wip", "0.5",
	// 	// "-wlen", "0.020",
	// }

	// filename_jsgfMinimal := "/home/dbarbera/Repositories/test_pronounce/audio_clips/Temp_f1d289d6-c152-4e26-a2f4-483175d7a12c/forced_align_a434c972-810b-40fc-b9a5-9e6697dcf5c5_allowed1_philip_fixed_trimmed.wav.jsgf"
	// filename_wavMinimal := "/home/dbarbera/Repositories/test_pronounce/audio_clips/allowed1_philip_fixed_trimmed.wav"

	//[]string len: 77, cap: 96, ["pocketsphinx_continuous","-beam","1e-10000","-alpha","0.97","-nfft","512","-remove_silence","yes","-vad_startspeech","5","-lw","6","-maxwpf","-1","-doublebw","yes","-dither","yes","-backtrace","yes","-pl_window","0","-wip","0.25","-featparams","/usr/local/share/xyzpocketsphinx/model/en-us/en-us/feat.params","-jsgf","/home/dbarbera/Repositories/test_pronounce/audio_clips/Temp_ede78846-4ed3-444b-b1bb-ccfc4ddb9465/forced_align_3dec0235-cb32-4b6f-a4bb-24c9f123aa56_allowed1_philip_fixed_trimmed.wav.jsgf","-frate","72","-bestpath","no","-lponlybeam","1e-10000","-vad_prespeech","5","-nwpen","1","-cmn","live","-logfn","/home/dbarbera/Repositories/test_pronounce/audio_clips/Temp_ede78846-4ed3-444b-b1bb-ccfc4ddb9465/cb2ea7d3-f961-409f-aa4a-b7708a9dae36.log","-wbeam","1e-10000","-infile","/home/dbarbera/Repositories/test_pronounce/audio_clips/allowed1_philip_fixed_trimmed.wav","-pip","1.15","-fwdflat","no","-lpbeam","1e-10000","-cmninit","","-vad_threshold","1.5","-topn","6","-vad_postspeech","20","-fsgusefiller","no","-wlen",...+13 more]

	//result := []string{"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"}

	// params := params1
	// params = params2
	// params = params3
	// // params = params72
	// params = params125
	// params = params105
	// params = params80

	// filename_jsgf := filename_jsgf1
	// filename_wav := filename_wav1
	// filename_jsgf = filename_jsgf2
	// filename_wav = filename_wav2
	// filename_jsgf = filename_jsgf3
	// filename_wav = filename_wav3
	// filename_jsgf = filename_jsgf72
	// filename_wav = filename_wav72
	// filename_jsgf = filename_jsgf125
	// filename_wav = filename_wav125
	// filename_jsgf = filename_jsgf105
	// filename_wav = filename_wav105
	// filename_jsgf = filename_jsgf80
	// filename_wav = filename_wav80

	// // init params:
	// params = params
	// filename_jsgf = filename_jsgf
	// filename_wav = filename_wav

	// jsgf_buffer, err := os.ReadFile(filename_jsgf)
	// check(err)
	// audio_buffer, err := os.ReadFile(filename_wav)
	// check(err)

	// //Checks:
	// fmt.Printf("fjsgf %T, %d\n", jsgf_buffer, len(jsgf_buffer))
	// fmt.Printf("fwav %T, %d\n", audio_buffer, len(audio_buffer))
	// fmt.Printf("Params: %T, %d\n", params, len(params))

	//fmt.Printf("%s\n", filename_jsgf[:len(filename_jsgf)-5]+"__from_c.jsgf")
	//fmt.Printf("%s\n", filename_wav[:len(filename_wav)-4]+"__from_c.jsgf")

	//r := []xyz.Utt{}

	//Sorry, quick and dirty:

	var frates [5]string
	frates[0] = getValue("-frate", params72)
	frates[1] = getValue("-frate", params125)
	frates[2] = getValue("-frate", params105)
	frates[3] = getValue("-frate", params80)
	frates[4] = getValue("-frate", params91)

	var parameters [5][]string
	parameters[0] = params72
	parameters[1] = params125
	parameters[2] = params105
	parameters[3] = params80
	parameters[4] = params91

	var jsgfs [5]string
	jsgfs[0] = getValue("-jsgf", params72)
	jsgfs[1] = getValue("-jsgf", params125)
	jsgfs[2] = getValue("-jsgf", params105)
	jsgfs[3] = getValue("-jsgf", params80)
	jsgfs[4] = getValue("-jsgf", params91)

	var wavs [5]string
	wavs[0] = getValue("-infile", params72)
	wavs[1] = getValue("-infile", params125)
	wavs[2] = getValue("-infile", params105)
	wavs[3] = getValue("-infile", params80)
	wavs[4] = getValue("-infile", params91)

	// //This works, because it is serialised
	// n := len(wavs)
	// starti := time.Now()
	// for i := 0; i < n; i++ {
	// 	test_ps(frates[i], jsgfs[i], wavs[i], parameters[i])
	// }
	// elapsedi := time.Since(starti)
	// fmt.Printf(">>>> Timing: %s\n", elapsedi)
	// fmt.Println()

	concurrently(frates, parameters, jsgfs, wavs)
	//concurrently_int(5)

}
