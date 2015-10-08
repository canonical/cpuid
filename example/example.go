package main

import (
	"github.com/intel-go/cpuid"
	"fmt"
)

func main() {
	fmt.Printf("VendorString:   %s\n", cpuid.VendorIdentificatorString)
	fmt.Printf("SteppingId:     %d\n", cpuid.SteppingId)
	fmt.Printf("ProcessorType:  %d\n", cpuid.ProcessorType)
	fmt.Printf("DisplayFamily:  %d\n", cpuid.DisplayFamily)
	fmt.Printf("DisplayModel:   %d\n", cpuid.DisplayModel)
	fmt.Printf("CacheLineSize:  %d\n", cpuid.CacheLineSize)
	fmt.Printf("MaxLogocalCPUId:%d\n", cpuid.MaxLogocalCPUId)
	fmt.Printf("InitialAPICId:  %d\n", cpuid.InitialAPICId)
	fmt.Printf("Smallest monitor-line size in bytes:  %d\n", cpuid.MonLineSizeMin)
	fmt.Printf("Largest monitor-line size in bytes:  %d\n", cpuid.MonLineSizeMax)
	fmt.Printf("Monitor Interrupt break-event is supported:  %v\n", cpuid.MonitorIBE)
	fmt.Printf("MONITOR/MWAIT extensions are supported:  %v\n", cpuid.MonitorEMX)
	fmt.Printf("AVX state %v\n", cpuid.EnabledAVX)
	fmt.Printf("AVX-512 state %v\n", cpuid.EnabledAVX512)

	fmt.Printf("Features: ")
	for i := uint64(0); i < 64; i++ {
		if cpuid.HasFeature(1 << i) {
			fmt.Printf("%s ", cpuid.FeatureNames[1<<i])
		}
	}
	fmt.Printf("\n")

	fmt.Printf("ExtendedFeatures: ")
	for i := uint64(0); i < 64; i++ {
		if cpuid.HasExtendedFeature(1 << i) {
			fmt.Printf("%s ", cpuid.ExtendedFeatureNames[1<<i])
		}
	}
	fmt.Printf("\n")

	fmt.Printf("ExtraFeatures: ")
	for i := uint64(0); i < 64; i++ {
		if cpuid.HasExtraFeature(1 << i) {
			fmt.Printf("%s ", cpuid.ExtraFeatureNames[1<<i])
		}
	}
	fmt.Printf("\n")

	for _, cacheDescription := range cpuid.CacheDescriptors {
		fmt.Printf("CacheDescriptor: %v\n", cacheDescription)
	}

}