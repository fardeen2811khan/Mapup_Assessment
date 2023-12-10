package models

type SortResponse struct {
	SortedArrays      [][]int `json:"sorted_arrays"`
	TimeNanoseconds   int64   `json:"time_ns"`
}