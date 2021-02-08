package board

//PVEnteryStruct The struct for each entry in the PV table
type PVEnteryStruct struct {
	posKey uint64
	move   int
}

//enterytBytes calculated using  unsafe.Sizeof(pos.PVTable.Entry[0])
//
//if more values are added to the struct (PVEnteryStruct) it will need to be recalculated
const enterytBytes = 16

//PVTableStruct Struct to hole the PVTable
type PVTableStruct struct {
	Entry []PVEnteryStruct
}

var pvEntrys int

//initPVTables Initilize our PV Table slice with a exact amount of memory based on hashSizeMB
func (pos *PositionStruct) initPVTables(hashSizeMB int) {
	hashBytes := hashSizeMB * 1000000
	pvEntrys = hashBytes / enterytBytes
	pos.PVTable.Entry = make([]PVEnteryStruct, pvEntrys)
}
