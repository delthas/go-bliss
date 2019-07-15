package bliss

/*
#cgo LDFLAGS: -lbliss
#include <bliss.h>

#define xstr(a) str(a)
#define str(a) #a

static inline const char * bl_version_str() {
	return xstr(BL_VERSION);
}
*/
import "C"
import (
	"errors"
	"reflect"
	"runtime"
	"unsafe"
)

var version string

func init() {
	version = C.GoString(C.bl_version_str())
}

func newBool(i C.int) bool {
	if i == 1 {
		return true
	}
	return false
}

func newString(stringC **C.char) string {
	if *stringC == nil {
		return ""
	}
	s := C.GoString(*stringC)
	C.free(unsafe.Pointer(*stringC))
	*stringC = nil
	return s
}

/*
ForceVector is a vector representing ratings of a song (tempo, attack, amplitude, frequency).
*/
type ForceVector struct {
	/*
		The tempo rating is the beats per minute of a song.
	*/
	Tempo float32
	/*
		The attack rating is a sum of the intensity of all the attacks divided by
		the song's length.
	*/
	Attack float32
	/*
		The amplitude rating reprents the physical "force" of the song, that is, how
		much the speaker's membrane will move in order to create the sound.
	*/
	Amplitude float32
	/*
		The frequency rating is a ratio between high and low frequencies: a song with a
		lot of high-pitched sounds tends to wake humans up far more easily.
	*/
	Frequency float32
}

func newForceVector(forceVectorC C.struct_force_vector_s) ForceVector {
	return ForceVector{
		Tempo:     float32(forceVectorC.tempo),
		Amplitude: float32(forceVectorC.amplitude),
		Frequency: float32(forceVectorC.frequency),
		Attack:    float32(forceVectorC.attack),
	}
}

func newForceVectorC(forceVector ForceVector) C.struct_force_vector_s {
	return C.struct_force_vector_s{
		tempo:     C.float(forceVector.Tempo),
		amplitude: C.float(forceVector.Amplitude),
		frequency: C.float(forceVector.Frequency),
		attack:    C.float(forceVector.Attack),
	}
}

/*
Envelope stores envelope-related characteristics of a Song.
*/
type Envelope struct {
	/*
		The tempo rating is the beats per minute of a song.
	*/
	Tempo float32
	/*
		The attack rating is a sum of the intensity of all the attacks divided by
		the song's length.
	*/
	Attack float32
}

/*
ForceRating represents the overall force category of a Song (from its Force).
*/
type ForceRating int

const (
	/*
		Loud means the song has a positive force (exact meaning may change).
	*/
	Loud ForceRating = 0
	/*
		Calm means the song has a negative force (exact meaning may change).
	*/
	Calm ForceRating = 1
	/*
		Unknown means the song has a force of zero (exact meaning may change).
	*/
	Unknown ForceRating = 2
)

const unexpected = -2

/*
Song represents a decoded audio file.

It can be analyzed or not, depending on how the Song was obtained. If it is,
Force, ForceRating, ForceVector, will have meaningful values. Otherwise, their
value is undefined.

A Song owns internal memory that can be freed explicitly by using Close when done
using it, or automatically by the GC after some time (with runtime.SetFinalizer).
*/
type Song struct {
	/*
		Force is the overall force / strength of the Song.
		Its value is only defined if the Song was analyzed.
		Lower values means the song is calm, higher values means it is loud.
	*/
	Force float32
	/*
		ForceRating is the overall force / strength category of the Song.
		Its value is only defined if the Song was analyzed.
		It can either be Calm, Loud, or Unknown.
	*/
	ForceRating ForceRating
	/*
		ForceVector stores the analyzed ratings of the Song.
		Its value is only defined if the Song was analyzed.
	*/
	ForceVector ForceVector
	/*
		Samples stores the decoded samples of the Song, in linear PCM format,
		interleaved per channel.

		Example, with BytesPerSample=2: left_sample_0_byte0,left_sample_0_byte1,
		right_sample_0_byte0,...
	*/
	Samples []int8
	/*
		Channels stores the number of channels of the Song. Mono is 1, stereo is 2.
	*/
	Channels int
	/*
		SampleRate stores the sampling rate of the Song in samples per second.
	*/
	SampleRate int
	/*
		Bitrate stores the average bitrate of the Song in bits per second.
	*/
	Bitrate int
	/*
		BytesPerSample stores the count of bytes per sample of the Song.

		For PCM-S16LE this is 2.
	*/
	BytesPerSample int
	/*
		Resampled is true if the Song has been resampled (either by changing its
		sampling rate, or by changing its sample format, eg float32 to uint16).
	*/
	Resampled bool
	/*
		Duration if the duration of the Song in seconds, rounded down.
	*/
	Duration uint64
	/*
		Filename is the path of the file to the Song.
	*/
	Filename string
	/*
		Artist is the value of the artist tag in the audio file metadata,
		or the empty string if not found.
	*/
	Artist string
	/*
		Title is the value of the title tag in the audio file metadata,
		or the empty string if not found.
	*/
	Title string
	/*
		Album is the value of the album tag in the audio file metadata,
		or the empty string if not found.
	*/
	Album string
	/*
		TrackNumber is the value of the track number tag in the audio file metadata,
		or the empty string if not found.
	*/
	TrackNumber string
	/*
		Genre is the value of the genre tag in the audio file metadata,
		or the empty string if not found.
	*/
	Genre string
	songC *C.struct_bl_song
}

/*
Close frees any native resources owned by this Song.

Calling Close more than once is a no-op.

The Song must not be used after it is closed.
*/
func (song *Song) Close() {
	if song.songC == nil {
		return
	}
	C.bl_free_song(song.songC)
	C.free(unsafe.Pointer(song.songC))
}

func closeSong(song *Song) {
	song.Close()
}

func closeSongC(songC *C.struct_bl_song) {
	C.bl_free_song(songC)
	C.free(unsafe.Pointer(songC))
}

func newSong(songC *C.struct_bl_song) *Song {
	song := &Song{
		Force:       float32(songC.force),
		ForceVector: newForceVector(songC.force_vector),
		Samples: *(*[]int8)(unsafe.Pointer(&reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(songC.sample_array)),
			Len:  int(songC.nSamples),
			Cap:  int(songC.nSamples),
		})),
		Channels:       int(songC.channels),
		SampleRate:     int(songC.sample_rate),
		Bitrate:        int(songC.bitrate),
		BytesPerSample: int(songC.nb_bytes_per_sample),
		ForceRating:    ForceRating(songC.calm_or_loud),
		Resampled:      newBool(songC.resampled),
		Duration:       uint64(songC.duration),
		Filename:       newString(&songC.filename),
		Artist:         newString(&songC.artist),
		Title:          newString(&songC.title),
		Album:          newString(&songC.album),
		TrackNumber:    newString(&songC.tracknumber),
		Genre:          newString(&songC.genre),
		songC:          songC,
	}
	runtime.SetFinalizer(song, closeSong)
	return song
}

/*
Decode decodes an audio file, without analyzing it and returns it as a Song.

filename is the path of the song to decode.

If there is an error reading or decoding the file, Decode returns a non-nil
error and the *Song will be nil. Otherwise, error is nil and *Song is non-nil.
*/
func Decode(filename string) (*Song, error) {
	var filenameC *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(filenameC))
	songC := (*C.struct_bl_song)(C.calloc(1, C.sizeof_struct_bl_song))
	r := int(C.bl_audio_decode(filenameC, songC))
	if r == unexpected {
		closeSongC(songC)
		return nil, errors.New("bliss: couldn't decode song")
	}
	return newSong(songC), nil
}

/*
Analyze decodes an audio file, then analyzes it and returns it as an analyzed Song.

filename is the path of the song to analyze.

If there is an error reading or decoding the file, Analyze returns a non-nil
error and the *Song will be nil. Otherwise, error is nil and *Song is non-nil.
*/
func Analyze(filename string) (*Song, error) {
	var filenameC *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(filenameC))
	songC := (*C.struct_bl_song)(C.calloc(1, C.sizeof_struct_bl_song))
	r := int(C.bl_analyze(filenameC, songC))
	if r == unexpected {
		closeSongC(songC)
		return nil, errors.New("bliss: couldn't decode song")
	}
	return newSong(songC), nil
}

/*
DistanceFile computes the distance between two songs stored in audio files, and additionally
returns them as analyzed Songs.

The distance is computed using a standard euclidian distance between the force vectors of the songs.

filename1 and filename2 are the paths of the songs to compare.

If there is an error reading or decoding the files, DistanceFile returns a non-nil
error and the *Song values will be nil. Otherwise, error is nil and *Song values are non-nil.

The return values correspond to, in that order, the first song, the second song, the distance
between the two songs, and any error.
*/
func DistanceFile(filename1 string, filename2 string) (*Song, *Song, float32, error) {
	var filename1C *C.char = C.CString(filename1)
	defer C.free(unsafe.Pointer(filename1C))
	var filename2C *C.char = C.CString(filename2)
	defer C.free(unsafe.Pointer(filename2C))
	song1C := (*C.struct_bl_song)(C.calloc(1, C.sizeof_struct_bl_song))
	song2C := (*C.struct_bl_song)(C.calloc(1, C.sizeof_struct_bl_song))
	r := float32(C.bl_distance_file(filename1C, filename2C, song1C, song2C))
	if r == unexpected {
		closeSongC(song1C)
		closeSongC(song2C)
		return nil, nil, 0, errors.New("bliss: couldn't decode songs")
	}
	return newSong(song1C), newSong(song2C), r, nil
}

/*
Distance computes the distance between two force vectors.

The distance is computed using a standard euclidian distance between the force vectors.

song1 and song2 are the force vectors to compare. They can typically be obtained with
the ForceVector field of a Song after it has been returned by Analyze.
*/
func Distance(song1 ForceVector, song2 ForceVector) float32 {
	r := float32(C.bl_distance(newForceVectorC(song1), newForceVectorC(song2)))
	return r
}

/*
CosineSimilarityFile computes the cosine similairty between two songs stored in audio files, and additionally
returns them as analyzed Songs.

The cosine similarity is a value between -1 and 1; -1 means songs are total opposites,
1 means that they are completely similar.

filename1 and filename2 are the paths of the songs to compare.

If there is an error reading or decoding the files, CosineSimilarityFile returns a non-nil
error and the *Song values will be nil. Otherwise, error is nil and *Song values are non-nil.

The return values correspond to, in that order, the first song, the second song, the cosine similiarity
between the two songs, and any error.
*/
func CosineSimilarityFile(filename1 string, filename2 string) (*Song, *Song, float32, error) {
	var filename1C *C.char = C.CString(filename1)
	defer C.free(unsafe.Pointer(filename1C))
	var filename2C *C.char = C.CString(filename2)
	defer C.free(unsafe.Pointer(filename2C))
	song1C := (*C.struct_bl_song)(C.calloc(1, C.sizeof_struct_bl_song))
	song2C := (*C.struct_bl_song)(C.calloc(1, C.sizeof_struct_bl_song))
	r := float32(C.bl_cosine_similarity_file(filename1C, filename2C, song1C, song2C))
	if r == unexpected {
		closeSongC(song1C)
		closeSongC(song2C)
		return nil, nil, 0, errors.New("bliss: couldn't decode songs")
	}
	return newSong(song1C), newSong(song2C), r, nil
}

/*
CosineSimilarityFile computes the cosine similairty between two force vectors.

The cosine similarity is a value between -1 and 1; -1 means songs are total opposites,
1 means that they are completely similar.

song1 and song2 are the force vectors to compare. They can typically be obtained with
the ForceVector field of a Song after it has been returned by Analyze.
*/
func CosineSimilarity(song1 ForceVector, song2 ForceVector) float32 {
	r := float32(C.bl_cosine_similarity(newForceVectorC(song1), newForceVectorC(song2)))
	return r
}

/*
EnvelopeSort computes and returns envelope-related characteristics of a Song.

The return value will never be nil.

The tempo rating draws the envelope of the whole song, and then computes its
DFT, obtaining peaks at the frequency of each dominant beat. The period of
each dominant beat can then be deduced from the frequencies, hinting at the
song's tempo.

Warning: the tempo is not equal to the force of the song. As an example , a
heavy metal track can have no steady beat at all, giving a very low tempo score
while being very loud.

The attack rating computes the difference between each value in the envelope
and the next (its derivative).
The final value is obtained by dividing the sum of the positive derivates by
the number of samples, in order to avoid different results just because of
the songs' length.
*/
func EnvelopeSort(song *Song) *Envelope {
	var envelopeC C.struct_envelope_result_s
	C.bl_envelope_sort(song.songC, &envelopeC)
	return &Envelope{
		Tempo:  float32(envelopeC.tempo),
		Attack: float32(envelopeC.attack),
	}
}

/*
AmplitudeSort computes the amplitude rating of a Song.

The returned value is the same as the Amplitude field of ForceVector.

The amplitude rating reprents the physical "force" of the song, that is,
how much the speaker's membrane will move in order to create the sound.

It is obtained by applying a magic formula with magic coefficients to a
histogram of the values of all the song's samples.
*/
func AmplitudeSort(song *Song) float32 {
	r := float32(C.bl_amplitude_sort(song.songC))
	return r
}

/*
FrequencySort computes the amplitude rating of a Song.

The returned value is the same as the Frequency field of ForceVector.

The frequency rating is a ratio between high and low frequencies: a song
with a lot of high-pitched sounds tends to wake humans up far more easily.

This rating is obtained by performing a DFT over the sample array, and
splitting the resulting array in 4 frequency bands: low, mid-low, mid,
mid-high, and high. Using the value in dB for each band, the final formula
corresponds to freq_result = high + mid-high + mid - (low + mid-low)
*/
func FrequencySort(song *Song) float32 {
	r := float32(C.bl_frequency_sort(song.songC))
	return r
}

/*
Version returns the runtime version of the C bliss library as a string, e.g: "1.1".
*/
func Version() string {
	return version
}

/*
Mean is a helper that compute the mean of an array of signed short samples.
*/
func Mean(samples []int16) int {
	r := int(C.bl_mean((*C.short)(unsafe.Pointer(&samples[0])), C.int(len(samples))))
	return r
}

/*
Variance is a helper that compute the variance of an array of signed short samples.

Variance needs the mean of the array, which can be get by using Mean.
*/
func Variance(samples []int16, mean int) int {
	r := int(C.bl_variance((*C.short)(unsafe.Pointer(&samples[0])), C.int(len(samples)), C.int(mean)))
	return r
}

/*
RectangularFilter is a helper that smoothes an array of samples (it is a one-dimension
moving average over the samples).

samplesIn is the array of input samples.

smoothWidth is the size of the filter, i.e. how many adjacent points to average.

samplesOut is the array in which to store output samples. Its length must be no smaller than
the length of samplesIn, or RectangularFilter will throw a panic.
*/
func RectangularFilter(samplesOut []float64, samplesIn []float64, smoothWidth int) {
	if len(samplesOut) < len(samplesIn) {
		panic("bliss: samplesOut has smaller length than samplesIn")
	}
	C.bl_rectangular_filter((*C.double)(unsafe.Pointer(&samplesOut[0])), (*C.double)(unsafe.Pointer(&samplesIn[0])), C.int(len(samplesIn)), C.int(smoothWidth))
}
