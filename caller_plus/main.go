// https://stackoverflow.com/questions/61821424/how-to-use-channels-to-gather-response-from-various-goroutines

package main

import (
	"bufio"
	"fmt"
	_ "net/http/pprof"
	"os"
	"sort"
	"sync"
	"time"

	//"os"
	//"xyz"
	//"github.com/colinarticulate/scanScheduler"
	"github.com/davidbarbera/articulate-pocketsphinx-go/xyz_plus"
	//"github.com/pkg/profile"
)

//Checking cores
const n int = 5

// func call_to_ps(jsgf_buffer []byte, audio_buffer []byte, params []string, c chan []xyz_plus.Utt) {

// 	c <- Ps(jsgf_buffer, audio_buffer, params)

// }

func call_to_ps_wg_chan(jsgf_buffer []byte, audio_buffer []byte, params []string, wg *sync.WaitGroup, resultChan chan<- []xyz_plus.Utt) {
	defer wg.Done()

	//resultChan <- Ps(jsgf_buffer, audio_buffer, params)
	start := time.Now()
	resultChan <- xyz_plus.Ps_plus_call(jsgf_buffer, audio_buffer, params)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("elapsed =", elapsed)

}

// func call_to_ps_batch_wg_chan(audio_buffer []byte, params []string, wg *sync.WaitGroup, resultChan chan<- []string) {
// 	defer wg.Done()

// 	//resultChan <- Ps(jsgf_buffer, audio_buffer, params)
// 	start := time.Now()
// 	resultChan <- xyz_plus.Ps_batch_plus_call(audio_buffer, params)
// 	t := time.Now()
// 	elapsed := t.Sub(start)
// 	fmt.Println("elapsed =", elapsed)

// }

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

func concurrently_n(frates [n]string, parameters [n][]string, jsgf_buffers [n][]byte, audio_buffers [n][]byte) [][]xyz_plus.Utt {
	m := len(audio_buffers)
	var results [][]xyz_plus.Utt
	ch := make(chan []xyz_plus.Utt, 1)
	//var id = []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	//n := len(wavs)
	//wg.Add(n)
	fmt.Printf(">>>> multithreaded:\n")
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
	// for result := range results {
	// 	fmt.Println(result)
	// }
	//fmt.Println(results)
	// for _, result := range results {
	// 	fmt.Println(result)
	// }
	fmt.Printf(">>>> Timing multithreaded: %s\n", elapsed)
	fmt.Println()

	return results
}

func concurrently(frates [n]string, parameters [n][]string, jsgf_buffers [n][]byte, audio_buffers [n][]byte) [][]xyz_plus.Utt {
	m := len(audio_buffers)
	var results [][]xyz_plus.Utt
	ch := make(chan []xyz_plus.Utt, 1)
	//var id = []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup

	//n := len(wavs)
	//wg.Add(n)
	p := 1
	fmt.Printf(">>>> %dx%d Continuous multithreaded:\n", p, m) // p groups of 5 scans
	start := time.Now()
	for j := 0; j < p; j++ {
		for i := 0; i < m; i++ {

			wg.Add(1)

			go call_to_ps_wg_chan(jsgf_buffers[i], audio_buffers[i], parameters[i], &wg, ch)
		}
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
	// for result := range results {
	// 	fmt.Println(result)
	// }
	//fmt.Println(results)
	// for _, result := range results {
	// 	fmt.Println(result)
	// }
	fmt.Printf(">>>> Timing multithreaded: %s\n", elapsed)
	fmt.Println()

	return results
}

// func concurrently_batch(frates [n]string, parameters [n][]string, audio_buffers [n][]byte) [][]string {
// 	m := len(audio_buffers)
// 	var results [][]string
// 	ch := make(chan []string, 1)
// 	//var id = []int{1, 2, 3, 4, 5}
// 	var wg sync.WaitGroup

// 	//n := len(wavs)
// 	//wg.Add(n)
// 	p := 1
// 	fmt.Printf(">>>> %dx%d Batch multithreaded:\n", p, m) // p groups of 5 scans
// 	start := time.Now()
// 	for j := 0; j < p; j++ {
// 		for i := 0; i < m; i++ {

// 			wg.Add(1)

// 			go call_to_ps_batch_wg_chan(audio_buffers[i], parameters[i], &wg, ch)
// 		}
// 	}

// 	// go func() {
// 	// 	for v := range ch {
// 	// 		results = append(results, v)
// 	// 	}
// 	// }()
// 	// wg.Wait()
// 	// close(ch)

// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()

// 	//time.Sleep(1000 * time.Millisecond)
// 	//Gathering or displaying results:
// 	for v := range ch {
// 		results = append(results, v)
// 	}
// 	// for elem := range ch {
// 	// 	fmt.Println(elem)
// 	// }

// 	elapsed := time.Since(start)

// 	// fmt.Println("Concurrently (multithreaded-encapsulated): ")
// 	// // for result := range results {
// 	// // 	fmt.Println(result)
// 	// // }
// 	// fmt.Println(results)

// 	// for _, result := range results {
// 	// 	fmt.Println(result)
// 	// }
// 	fmt.Printf(">>>> Timing multithreaded: %s\n", elapsed)

// 	// fmt.Println()

// 	return results
// }

func sequentially(frates [n]string, parameters [n][]string, jsgfs [n][]byte, wavs [n][]byte) {
	m := len(wavs)
	fmt.Printf(">>>> Sequential:\n")
	starti := time.Now()
	for i := 0; i < m; i++ {
		test_ps(frates[i], jsgfs[i], wavs[i], parameters[i])
	}
	elapsedi := time.Since(starti)
	fmt.Printf(">>>> Timing Sequential: %s\n", elapsedi)
	fmt.Println()
}

// func sequentially_batch(frates [n]string, parameters [n][]string, wavs [n][]byte) {
// 	m := len(wavs)
// 	fmt.Printf(">>>> Sequential:\n")
// 	starti := time.Now()
// 	for i := 0; i < m; i++ {
// 		test_ps_batch(frates[i], wavs[i], parameters[i])
// 	}
// 	elapsedi := time.Since(starti)
// 	fmt.Printf(">>>> Timing Sequential: %s\n", elapsedi)
// 	fmt.Println()
// }

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

	//fmt.Println("--- frate = ", frate)
	starti := time.Now()
	//var r = Ps(jsgf_buffer, audio_buffer, parameters)
	var r = xyz_plus.Ps_plus_call(jsgf_buffer, audio_buffer, parameters)
	elapsedi := time.Since(starti)
	//fmt.Printf(">>>> Timing: %s\n", elapsedi)
	//fmt.Println(">>> ", elapsedi, "\t--- frate = ", frate, r)
	fmt.Println("--- frate = ", frate, ": ", r, "    \t>>>> Timing: ", elapsedi)
	//fmt.Println()

}

// func test_ps_batch(frate string, audio_buffer []byte, parameters []string) {
// 	// jsgf_buffer, err := os.ReadFile(jsgf_filename)
// 	// check(err)
// 	// audio_buffer, err := os.ReadFile(wav_filename)
// 	// check(err)

// 	//fmt.Println("--- frate = ", frate)
// 	starti := time.Now()
// 	//var r = Ps(jsgf_buffer, audio_buffer, parameters)
// 	var r = xyz_plus.Ps_batch_plus_call(audio_buffer, parameters)
// 	elapsedi := time.Since(starti)
// 	//fmt.Printf(">>>> Timing: %s\n", elapsedi)
// 	fmt.Println("--- frate = ", frate, ": ", r, "    \t>>>> Timing: ", elapsedi)
// 	//fmt.Println()

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

func testing_ps_continuous() {
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
	sequentially(frates, parameters, jsgf_buffers, wav_buffers)

	//concurrently(frates, parameters, jsgf_buffers, wav_buffers)
	// fmt.Println(results)
	// concurrently_int(5)

	// //Testing how many threads in parallel can we do:
	// var pjsgf_buffers [n][]byte
	// var pwav_buffers [n][]byte
	// var pwavs [n]string
	// var pparameters [n][]string
	// var pjsgfs [n]string
	// var pfrates [n]string

	// var f = 0
	// for i := 0; i < n; i++ {
	// 	pjsgfs[i] = jsgfs[f]
	// 	pwavs[i] = wavs[f]
	// 	pfrates[i] = frates[f]
	// 	pparameters[i] = parameters[f]
	// 	pjsgf_buffers[i], err = os.ReadFile(pjsgfs[i])
	// 	check(err)
	// 	pwav_buffers[i], err = os.ReadFile(pwavs[i])
	// 	check(err)

	// }
	// //sequentially(pfrates, pparameters, pjsgf_buffers, pwav_buffers)
	// fmt.Println(n, " scans:")
	// concurrently(pfrates, pparameters, pjsgf_buffers, pwav_buffers)

	// results := concurrently(pfrates, pparameters, pjsgf_buffers, pwav_buffers)
	// fmt.Println(results)

}

func get_audiofilename(params []string) string {
	audio_dir := getValue("-cepdir", params)
	extension := getValue("-cepext", params)
	ctl_filename := getValue("-ctl", params)

	ctl_file, err := os.Open(ctl_filename)
	check(err)

	fileScanner := bufio.NewScanner(ctl_file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	audio := audio_dir + "/" + fileLines[0] + extension

	return audio

}

func testing_ps_batch() {
	var frates [5]string
	frates[0] = getValue("-frate", batch_params72)
	frates[1] = getValue("-frate", batch_params125)
	frates[2] = getValue("-frate", batch_params105)
	frates[3] = getValue("-frate", batch_params80)
	frates[4] = getValue("-frate", batch_params91)

	var parameters [5][]string
	parameters[0] = batch_params72
	parameters[1] = batch_params125
	parameters[2] = batch_params105
	parameters[3] = batch_params80
	parameters[4] = batch_params91

	var wavs [5]string
	wavs[0] = get_audiofilename(batch_params72)
	wavs[1] = get_audiofilename(batch_params125)
	wavs[2] = get_audiofilename(batch_params105)
	wavs[3] = get_audiofilename(batch_params80)
	wavs[4] = get_audiofilename(batch_params91)

	var err error
	var wav_buffers [5][]byte
	for i := 0; i < 5; i++ {
		wav_buffers[i], err = os.ReadFile(wavs[i])
		check(err)
	}

	//One case:
	//test_ps_batch(frates[0], wav_buffers[0], parameters[0])

	//nees n=5
	//sequentially_batch(frates, parameters, wav_buffers)

	//needs n=5
	//concurrently_batch(frates, parameters, wav_buffers)
	// results := concurrently_batch(frates, parameters, wav_buffers)
	// for _, result := range results {
	// 	fmt.Println(result)
	// }

}

func testing_continuous_n() {
	var frates [n]string
	var parameters [n][]string
	var jsgfs [n]string
	var wavs [n]string

	var err error
	var jsgf_buffers [n][]byte
	var wav_buffers [n][]byte

	for i := 0; i < n; i++ {
		frates[i] = getValue("-frate", params72)
		parameters[i] = params72
		jsgfs[i] = getValue("-jsgf", params72)
		wavs[i] = getValue("-infile", params72)
		jsgf_buffers[i], err = os.ReadFile(jsgfs[i])
		check(err)
		wav_buffers[i], err = os.ReadFile(wavs[i])
		check(err)
	}

	//This works, because it is serialised
	//sequentially(frates, parameters, jsgf_buffers, wav_buffers)

	concurrently_n(frates, parameters, jsgf_buffers, wav_buffers)
	//fmt.Println(results)
	//concurrently_int(5)

	//Testing how many threads in parallel can we do:
	// var pjsgf_buffers [n][]byte
	// var pwav_buffers [n][]byte
	// var pwavs [n]string
	// var pparameters [n][]string
	// var pjsgfs [n]string
	// var pfrates [n]string

	// var f = 0
	// for i := 0; i < n; i++ {
	// 	pjsgfs[i] = jsgfs[f]
	// 	pwavs[i] = wavs[f]
	// 	pfrates[i] = frates[f]
	// 	pparameters[i] = parameters[f]
	// 	pjsgf_buffers[i], err = os.ReadFile(pjsgfs[i])
	// 	check(err)
	// 	pwav_buffers[i], err = os.ReadFile(pwavs[i])
	// 	check(err)

	// }

}

// func profiling() {
// 	// CPUProfile enables cpu profiling. Note: Default is CPU
// 	//defer profile.Start(profile.CPUProfile).Stop()

// 	// GoroutineProfile enables goroutine profiling.
// 	// It returns all Goroutines alive when defer occurs.
// 	//defer profile.Start(profile.GoroutineProfile, profile.ProfilePath(".")).Stop()

// 	// BlockProfile enables block (contention) profiling.
// 	//defer profile.Start(profile.BlockProfile).Stop()

// 	// ThreadcreationProfile enables thread creation profiling.
// 	//defer profile.Start(profile.ThreadcreationProfile).Stop()

// 	// MemProfileHeap changes which type of memory profiling to
// 	// profile the heap.
// 	//defer profile.Start(profile.MemProfileHeap).Stop()

// 	// MemProfileAllocs changes which type of memory to profile
// 	// allocations.
// 	//defer profile.Start(profile.MemProfileAllocs).Stop()

// 	// MutexProfile enables mutex profiling.
// 	//defer profile.Start(profile.MutexProfile).Stop()

// 	testing_ps_continuous()
// }

// var cpuprofile = flag.String("cpuprofile", "", "Writes cpu profile to a file.")
// var traceprofile = flag.String("traceprofile", "", "writest trace profile to a file.")

//Sorry, quick and dirty:
func main() {
	// flag.Parse()
	// if *cpuprofile != "" {
	// 	f, err := os.Create(*cpuprofile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	pprof.StartCPUProfile(f)
	// 	defer pprof.StopCPUProfile()
	// }

	// flag.Parse()
	// if *traceprofile != "" {
	// 	f, err := os.Create(*traceprofile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	trace.Start(f)
	// 	defer trace.Stop()
	// }

	//pprof.StartCPUProfile()
	//profiling()
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()
	testing_ps_continuous()
	//pprof.StopCPUProfile()
	//testing_ps_batch()
	//testing_continuous_n()
	//fmt.Println("working on it.")

}
