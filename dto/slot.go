package dto

import "fmt"

// SlotResp 回复
type SlotResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Slot `json:"data"`
}

// Slot 车位锁
type Slot struct {
	Name     string  `json:"name"`
	Addr     uint32  `json:"addr"`
	IP       string  `json:"ip"`
	Position string  `json:"position"`
	State    string  `json:"state"`
	Battery  float64 `json:"battery"`
}

func (s Slot) String() string {
	return fmt.Sprintf("%s: %d", s.Name, s.Addr)
}

// ByAddr 排序
type ByAddr []Slot

func (a ByAddr) Len() int {
	return len(a)
}

func (a ByAddr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAddr) Less(i, j int) bool {
	return a[i].Addr < a[j].Addr
}
