package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//var mutex sync.Mutex
	input := []string{"Anomalous", "cosmic", "rays", "(ACRs),", "or", "the", "anomalous", "component", "of", "cosmic", "rays,", "are", "energetic", "ions", "of", "interstellar", "origin", "that", "are", "observed", "inside", "the", "heliosphere.", "Interstellar", "neutral", "atoms", "drift", "inward,", "become", "ionized", "by", "solar", "ultraviolet", "photons,", "electron", "impact,", "or", "charge", "exchange,", "are", "picked", "up", "by", "the", "solar", "wind,", "and", "are", "then", "accelerated—mainly", "near", "the", "solar-wind", "termination", "shock", "and", "throughout", "the", "heliosheath—by", "diffusive", "shock", "acceleration", "and", "related", "processes.", "The", "particles", "were", "called", "anomalous", "because", "their", "presence", "and", "characteristics", "didn't", "fit", "with", "the", "existing", "understanding", "of", "cosmic", "rays", "at", "the", "time.", "The", "enhancement", "was", "particularly", "notable", "in", "that", "the", "GCR", "[galactic", "cosmic", "rays]", "intensity", "did", "not", "decrease", "with", "decreasing", "energy,", "as", "was", "expected", "based", "on", "our", "understanding", "that", "low-energy", "GCRs", "were", "unable", "to", "reach", "1", "AU", "due", "to", "their", "interaction", "with", "the", "solar", "wind", "Giacalone", "et", "al.", "(2022)", "argue", "that", "anomalous", "cosmic", "rays", "is", "a", "confusing", "and", "not", "descriptive", "name,", "and", "propose", "to", "use", "Heliospheric", "Energetic", "Particles", "instead", "of", "it."}

	output := make(map[string]int)

	for _, in := range input {
		wg.Add(1)
		go func(in string) {
			defer wg.Done()

			//runes := strings.Split(in, "")
			//runes := []rune(in)
			//			mutex.Lock()
			output[in] = len(in)
			//			mutex.Unlock()
		}(in)
		//runes := strings.Split(in, "")
	}

	wg.Wait()
	fmt.Println(output)
}
