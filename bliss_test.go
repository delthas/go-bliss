package bliss

import (
	"fmt"
	"testing"
)

func assertFloat(t *testing.T, expected float32, actual float32, message string) {
	r := expected - actual
	if r < 0 {
		r = -r
	}
	if r > 0.000001 {
		t.Error(fmt.Sprintf("%s: float mismatch: expected: [%f], got: [%f]", message, expected, actual))
	}
}

func assertInt(t *testing.T, expected int, actual int, message string) {
	if expected != actual {
		t.Error(fmt.Sprintf("%s: float mismatch: expected: [%d], got: [%d]", message, expected, actual))
	}
}

func assertString(t *testing.T, expected string, actual string, message string) {
	if expected != actual {
		t.Error(fmt.Sprintf("%s: float mismatch: expected: [%s], got: [%s]", message, expected, actual))
	}
}

func TestAnalyze(t *testing.T) {
	song, err := Analyze("audio/song.flac")
	if err != nil {
		t.Error(err)
	}
	// optional Close, a runtime.SetFinalizer is set anyway
	defer song.Close()

	assertFloat(t, -25.165920, song.Force, "song force")

	assertFloat(t, -8.945454, song.ForceVector.Tempo, "song tempo")
	assertFloat(t, -15.029835, song.ForceVector.Amplitude, "song amplitude")
	assertFloat(t, -10.136086, song.ForceVector.Frequency, "song frequency")
	assertFloat(t, -15.560563, song.ForceVector.Attack, "song attack")

	assertInt(t, 2, song.Channels, "song channels")
	assertInt(t, 488138, len(song.Samples), "song samples count")
	assertInt(t, 22050, song.SampleRate, "song sample rate")
	assertInt(t, 233864, song.Bitrate, "song bitrate")
	assertInt(t, 2, song.BytesPerSample, "song bytes per sample")
	assertInt(t, 11, int(song.Duration), "song duration")

	assertString(t, "David TMX", song.Artist, "song artist")
	assertString(t, "Renaissance", song.Title, "song title")
	assertString(t, "Renaissance", song.Album, "song album")
	assertString(t, "02", song.TrackNumber, "song track number")
	assertString(t, "Pop", song.Genre, "song genre")
}
