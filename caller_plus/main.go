// https://stackoverflow.com/questions/61821424/how-to-use-channels-to-gather-response-from-various-goroutines

package main

import (
	"fmt"
	"os"
	"sort"
	"sync"
	"time"

	//"os"
	//"xyz"
	//"github.com/colinarticulate/scanScheduler"
	"github.com/davidbarbera/articulate-pocketsphinx-go/xyz_plus"
)

//Checking cores
const n int = 64

// func call_to_ps(jsgf_buffer []byte, audio_buffer []byte, params []string, c chan []xyz_plus.Utt) {

// 	c <- Ps(jsgf_buffer, audio_buffer, params)

// }

func call_to_ps_wg_chan(jsgf_buffer []byte, audio_buffer []byte, params []string, wg *sync.WaitGroup, resultChan chan<- []xyz_plus.Utt) {
	defer wg.Done()

	//resultChan <- Ps(jsgf_buffer, audio_buffer, params)
	resultChan <- xyz_plus.Ps_plus_call(jsgf_buffer, audio_buffer, params)

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

func concurrently(frates [n]string, parameters [n][]string, jsgf_buffers [n][]byte, audio_buffers [n][]byte) [][]xyz_plus.Utt {
	m := len(audio_buffers)
	var results [][]xyz_plus.Utt
	ch := make(chan []xyz_plus.Utt, 1)
	//var id = []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	//n := len(wavs)
	//wg.Add(n)
	start := time.Now()
	for i := 0; i < m; i++ {

		wg.Add(1)

		go call_to_ps_wg_chan(jsgf_buffers[i], audio_buffers[i], parameters[i], &wg, ch)
	}

	// go func() {
	// 	for v := range ch {
	// 		results = append(results, v)
	// 	}
	// }()
	// wg.Wait()
	// close(ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	//time.Sleep(1000 * time.Millisecond)
	//Gathering or displaying results:
	for v := range ch {
		results = append(results, v)
	}
	// for elem := range ch {
	// 	fmt.Println(elem)
	// }

	elapsed := time.Since(start)

	// fmt.Println("Concurrently (multithreaded-encapsulated): ")
	// // for result := range results {
	// // 	fmt.Println(result)
	// // }
	// fmt.Println(results)
	fmt.Printf(">>>> Timing: %s\n", elapsed)
	// fmt.Println()

	return results
}

func sequentially(frates [n]string, parameters [n][]string, jsgfs [n][]byte, wavs [n][]byte) {
	m := len(wavs)
	starti := time.Now()
	for i := 0; i < m; i++ {
		test_ps(frates[i], jsgfs[i], wavs[i], parameters[i])
	}
	elapsedi := time.Since(starti)
	fmt.Printf(">>>> Timing: %s\n", elapsedi)
	fmt.Println()
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func test_ps(frate string, jsgf_buffer []byte, audio_buffer []byte, parameters []string) {
	// jsgf_buffer, err := os.ReadFile(jsgf_filename)
	// check(err)
	// audio_buffer, err := os.ReadFile(wav_filename)
	// check(err)

	fmt.Println("--- frate = ", frate)
	starti := time.Now()
	//var r = Ps(jsgf_buffer, audio_buffer, parameters)
	var r = xyz_plus.Ps_plus_call(jsgf_buffer, audio_buffer, parameters)
	elapsedi := time.Since(starti)
	fmt.Printf(">>>> Timing: %s\n", elapsedi)
	fmt.Println(r)
	fmt.Println()

}

// func Ps(jsgf_buffer []byte, audio_buffer []byte, params []string) []xyz_plus.Utt {
// 	result := []string{"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"}

// 	xyz_plus.Ps_plus_call(jsgf_buffer, audio_buffer, params, result)

// 	//Adapting result from coded string to utt struct
// 	if strings.Contains(result[0], "**") {
// 		raw := strings.Split(result[0], "**")

// 		if len(raw) < 2 {
// 			fmt.Println("xyzpocketsphinx: problems!")
// 		}

// 		//fmt.Printf("%T", raw)
// 		fields := strings.Split(raw[0], "*")

// 		//fmt.Println(fields)
// 		// hyp := fields[0]
// 		// score := fields[1]

// 		//fmt.Println(hyp)
// 		//fmt.Println(strings.Split(score, ","))
// 		utts := []xyz_plus.Utt{}
// 		//var utts = make([]Utt, len(fields)-2)

// 		for i := 0; i < len(fields)-2; i++ {
// 			parts := strings.Split(fields[2:][i], ",")
// 			phoneme := parts[0]
// 			text_start := parts[1]
// 			text_end := parts[2]
// 			start, serr := strconv.Atoi(text_start)
// 			end, eerr := strconv.Atoi(text_end)

// 			if phoneme != "(NULL)" {
// 				//fmt.Println(phoneme, start, end)
// 				//utts = append(utts, xyz_plus.Utt{phoneme, int32(start), int32(end)})
// 				utts = append(utts, xyz_plus.Utt{Text: phoneme, Start: int32(start), End: int32(end)})

// 				if serr != nil || eerr != nil {
// 					fmt.Println(serr, eerr)
// 				}
// 			}
// 		}
// 		return utts
// 	} else {
// 		return nil
// 	}
// }

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

//func main() {

// var params72 = []string{
// 	"pocketsphinx_continuous",
// 	"-alpha", "0.97",
// 	"-backtrace", "yes",
// 	"-beam", "1e-10000",
// 	"-bestpath", "no",
// 	"-cmn", "live",
// 	"-cmninit", "61.61,8.03,-6.54,4.13,-3.74,9.61,-5.77,-16.52,-13.85,-3.98,2.30,2.59,-1.94",
// 	"-dict", "/home/dbarbera/Data/art_db.phone",
// 	"-dither", "yes", //yes
// 	"-doublebw", "yes", //yes
// 	"-featparams", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us/feat.params",
// 	"-frate", "72",
// 	"-fsgusefiller", "no",
// 	"-fwdflat", "no",
// 	"-hmm", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us",
// 	"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
// 	"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_2668db47-d3ce-4760-ab4b-60b9b8a6c46e_allowed1_philip_fixed_trimmed.wav.jsgf",
// 	"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/d0c65f23-d9cc-4047-8f3c-3a91db3623ff_2668db47_frate_72_from_go_.log",
// 	"-lpbeam", "1e-10000",
// 	"-lponlybeam", "1e-10000",
// 	"-lw", "6",
// 	"-maxhmmpf", "-1",
// 	"-maxwpf", "-1",
// 	"-nfft", "512",
// 	"-nwpen", "1",
// 	"-pbeam", "1e-10000",
// 	"-pip", "1.15", //1.15
// 	"-pl_window", "0",
// 	"-remove_dc", "no",
// 	"-remove_noise", "yes",
// 	"-remove_silence", "yes", //yes
// 	"-topn", "6",
// 	"-vad_postspeech", "20",
// 	"-vad_prespeech", "5",
// 	"-vad_startspeech", "5",
// 	"-vad_threshold", "1.5",
// 	"-wbeam", "1e-10000",
// 	"-wip", "0.25",
// 	"-wlen", "0.032",
// }

// var params125 = []string{
// 	"pocketsphinx_continuous",
// 	"-alpha", "0.97",
// 	"-backtrace", "yes",
// 	"-beam", "1e-10000",
// 	"-bestpath", "no",
// 	"-cmn", "live",
// 	"-cmninit", "48.66,4.31,-7.12,5.61,-1.63,9.01,-4.65,-17.99,-16.52,-5.18,3.45,2.53,-1.34",
// 	"-dict", "/home/dbarbera/Data/art_db.phone",
// 	"-dither", "no",
// 	"-doublebw", "no",
// 	"-featparams", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us/feat.params",
// 	"-frate", "125",
// 	"-fsgusefiller", "no",
// 	"-fwdflat", "no",
// 	"-hmm", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us",
// 	"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
// 	"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_a3ecf04d-a77a-4269-9eb5-395f8dfbdd8a_allowed1_philip_fixed_trimmed.wav.jsgf",
// 	"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/e1e0d844-812b-496c-83fb-712de847f8a7_a3ecf04d_frate_125_from_go_.log",
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
// 	"-wlen", "0.016",
// }

// var params105 = []string{
// 	"pocketsphinx_continuous",
// 	"-alpha", "0.97",
// 	"-backtrace", "yes",
// 	"-beam", "1e-10000",
// 	"-bestpath", "no",
// 	"-cmn", "live",
// 	"-cmninit", "53.92,4.73,-7.01,5.40,-1.92,8.97,-4.24,-17.95,-17.00,-6.15,2.58,1.61,-1.69",
// 	"-dict", "/home/dbarbera/Data/art_db.phone",
// 	"-dither", "no",
// 	"-doublebw", "no",
// 	"-featparams", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us/feat.params",
// 	"-frate", "105",
// 	"-fsgusefiller", "no",
// 	"-fwdflat", "no",
// 	"-hmm", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us",
// 	"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
// 	"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_9ddb7131-fa08-4bc4-b44c-814b2ed9917e_allowed1_philip_fixed_trimmed.wav.jsgf",
// 	"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/c061830a-8106-4e06-ae16-18feb072ea45_9ddb7131_frate_105_from_go_.log",
// 	"-lpbeam", "1e-10000",
// 	"-lponlybeam", "1e-10000",
// 	"-lw", "6",
// 	"-maxhmmpf", "-1",
// 	"-maxwpf", "-1",
// 	"-nfft", "512",
// 	"-nwpen", "1",
// 	"-pbeam", "1e-10000",
// 	"-pip", "1",
// 	"-pl_window", "0",
// 	"-remove_dc", "no",
// 	"-remove_noise", "yes",
// 	"-remove_silence", "no",
// 	"-topn", "4",
// 	"-vad_postspeech", "20",
// 	"-vad_prespeech", "5",
// 	"-vad_startspeech", "5",
// 	"-vad_threshold", "1.5",
// 	"-wbeam", "1e-10000",
// 	"-wip", "0.5",
// 	"-wlen", "0.020",
// }

// var params80 = []string{
// 	"pocketsphinx_continuous",
// 	"-alpha", "0.97",
// 	"-backtrace", "yes",
// 	"-beam", "1e-10000",
// 	"-bestpath", "no",
// 	"-cmn", "live",
// 	"-cmninit", "60.83,7.43,-6.07,4.54,-3.88,9.48,-5.85,-16.81,-13.89,-3.82,2.39,2.53,-1.88",
// 	"-dict", "/home/dbarbera/Data/art_db.phone",
// 	"-dither", "yes",
// 	"-doublebw", "yes",
// 	"-featparams", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us/feat.params",
// 	"-frate", "80",
// 	"-fsgusefiller", "no",
// 	"-fwdflat", "no",
// 	"-hmm", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us",
// 	"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
// 	"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_4311b957-22a7-446c-85d9-d154d4156d02_allowed1_philip_fixed_trimmed.wav.jsgf",
// 	"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/4f0e00fa-9d20-4096-8bb5-8aeedc110e52_4311b957_frate_80_from_go_.log",
// 	"-lpbeam", "1e-10000",
// 	"-lponlybeam", "1e-10000",
// 	"-lw", "6",
// 	"-maxhmmpf", "-1",
// 	"-maxwpf", "-1",
// 	"-nfft", "512",
// 	"-nwpen", "1",
// 	"-pbeam", "1e-10000",
// 	"-pip", "1.15",
// 	"-pl_window", "0",
// 	"-remove_dc", "no",
// 	"-remove_noise", "yes",
// 	"-remove_silence", "yes", //yes
// 	"-topn", "6",
// 	"-vad_postspeech", "20",
// 	"-vad_prespeech", "5",
// 	"-vad_startspeech", "5",
// 	"-vad_threshold", "1.5",
// 	"-wbeam", "1e-10000",
// 	"-wip", "0.25",
// 	"-wlen", "0.028",
// }

// var params91 = []string{
// 	"pocketsphinx_continuous",
// 	"-alpha", "0.97",
// 	"-backtrace", "yes",
// 	"-beam", "1e-10000",
// 	"-bestpath", "no",
// 	"-cmn", "live",
// 	"-cmninit", "55.22,5.33,-6.78,5.07,-2.13,9.10,-4.03,-17.60,-16.77,-6.25,2.43,1.62,-1.56",
// 	"-dict", "/home/dbarbera/Data/art_db.phone",
// 	"-dither", "no",
// 	"-doublebw", "no",
// 	"-featparams", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us/feat.params",
// 	"-frate", "91",
// 	"-fsgusefiller", "no",
// 	"-fwdflat", "no",
// 	"-hmm", "/usr/local/share/xyzpocketsphinx/model/art-en-us/en-us",
// 	"-infile", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/allowed1_philip_fixed_trimmed.wav",
// 	"-jsgf", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/forced_align_1eed3902-f7e5-444b-a4b8-29b5c47ea52e_allowed1_philip_fixed_trimmed.wav.jsgf",
// 	"-logfn", "/home/dbarbera/Data/test_cases/allowed1_philip/Temp_990ba583-5249-41f3-8d42-0617f9eea6cd/log/f7a95619-d7c5-42a9-b548-561187b350da_1eed3902_frate_91_from_go_.log",
// 	"-lpbeam", "1e-10000",
// 	"-lponlybeam", "1e-10000",
// 	"-lw", "6",
// 	"-maxhmmpf", "-1",
// 	"-maxwpf", "-1",
// 	"-nfft", "512",
// 	"-nwpen", "1",
// 	"-pbeam", "1e-10000",
// 	"-pip", "1",
// 	"-pl_window", "0",
// 	"-remove_dc", "no",
// 	"-remove_noise", "yes",
// 	"-remove_silence", "no",
// 	"-topn", "4",
// 	"-vad_postspeech", "20",
// 	"-vad_prespeech", "5",
// 	"-vad_startspeech", "5",
// 	"-vad_threshold", "0.5",
// 	"-wbeam", "1e-10000",
// 	"-wip", "0.5",
// 	"-wlen", "0.024",
// }

//Sorry, quick and dirty:
func main() {
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

	var err error
	var jsgf_buffers [5][]byte
	for i := 0; i < 5; i++ {
		jsgf_buffers[i], err = os.ReadFile(jsgfs[i])
		check(err)
	}

	var wav_buffers [5][]byte
	for i := 0; i < 5; i++ {
		wav_buffers[i], err = os.ReadFile(wavs[i])
		check(err)
	}

	//This works, because it is serialised
	// sequentially(frates, parameters, jsgf_buffers, wav_buffers)

	// results := concurrently(frates, parameters, jsgf_buffers, wav_buffers)
	// fmt.Println(results)
	//concurrently_int(5)

	//Testing how many threads in parallel can we do:
	var pjsgf_buffers [n][]byte
	var pwav_buffers [n][]byte
	var pwavs [n]string
	var pparameters [n][]string
	var pjsgfs [n]string
	var pfrates [n]string

	var f = 1
	for i := 0; i < n; i++ {
		pjsgfs[i] = jsgfs[f]
		pwavs[i] = wavs[f]
		pfrates[i] = frates[f]
		pparameters[i] = parameters[f]
		pjsgf_buffers[i], err = os.ReadFile(pjsgfs[i])
		check(err)
		pwav_buffers[i], err = os.ReadFile(pwavs[i])
		check(err)

	}
	//sequentially(pfrates, pparameters, pjsgf_buffers, pwav_buffers)

	results := concurrently(pfrates, pparameters, pjsgf_buffers, pwav_buffers)
	fmt.Println(results)

}
