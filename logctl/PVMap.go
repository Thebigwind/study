package main

import "sync"

type PVMap struct {
	PvData map[string]int64
	lc     *sync.Mutex
}

var GlobalStatPvData *PVMap = nil

func GetStatPVMap() *PVMap {
	return GlobalStatPvData
}

//var StatPvData PVMap
func NewStatPvMap() *PVMap {
	pvdata := make(map[string]int64)
	pvmap := &PVMap{
		PvData: pvdata,
		lc:     new(sync.Mutex),
	}

	GlobalStatPvData = pvmap

	return pvmap
}
