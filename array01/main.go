package main 

import "fmt"

// temperatureAnalyze scans a fixed-size telemetry array
// and returns both the minimum and maximum recorded values.
func temperatureAnalyze(arr[6]int) (int,int) { 
	// Initialize min and max with the first element of the batch
	max := arr[0]
	min := arr[0]
	for _ , val := range arr {
		if val > max{
			max = val
		}else if val < min {
			min = val
		}
	}	
		return min,max
}

func main(){
	// Simulated telemetry batch from ground station
	temperatures := [6]int{40,15,-24,0,4,6}
	minVal, maxVal := temperatureAnalyze(temperatures)
	fmt.Printf("min: %d , max: %d\n" , minVal, maxVal)
}