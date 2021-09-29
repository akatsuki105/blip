package blip

// #include "./blip_buf.h"
import "C"
import "unsafe"

type Blip C.struct_blip_t

// blip_new
func New(size int) *Blip {
	b := C.blip_new(C.int(size))
	return (*Blip)(b)
}

// blip_delete
func (b *Blip) Delete() {
	if b != nil {
		C.blip_delete((*C.struct_blip_t)(b))
	}
}

// blip_set_rates
func (b *Blip) SetRates(clockRate, sampleRate float64) {
	C.blip_set_rates((*C.struct_blip_t)(b), C.double(clockRate), C.double(sampleRate))
}

// blip_clear
func (b *Blip) Clear() {
	C.blip_clear((*C.struct_blip_t)(b))
}

// blip_clocks_needed
func (b *Blip) ClocksNeeded(samples int) int {
	result := C.blip_clocks_needed((*C.struct_blip_t)(b), C.int(samples))
	return int(result)
}

// blip_end_frame
func (b *Blip) EndFrame(t uint) {
	C.blip_end_frame((*C.struct_blip_t)(b), C.uint(t))
}

// blip_samples_avail
func (b *Blip) SamplesAvail() int {
	return int(C.blip_samples_avail((*C.struct_blip_t)(b)))
}

// blip_read_samples
func (b *Blip) ReadSamples(out []uint16, count, stereo int) int {
	return int(C.blip_read_samples((*C.struct_blip_t)(b), (*C.short)(unsafe.Pointer(&out[0])), C.int(count), C.int(stereo)))
}

// blip_add_delta
func (b *Blip) AddDelta(time uint, delta int) {
	C.blip_add_delta((*C.struct_blip_t)(b), C.uint(time), C.int(delta))
}

// blip_add_delta_fast
func (b *Blip) AddDeltaFast(time uint, delta int) {
	C.blip_add_delta_fast((*C.struct_blip_t)(b), C.uint(time), C.int(delta))
}
