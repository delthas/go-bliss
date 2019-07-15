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
	song, err := Analyze("audio/song.mp3")
	if err != nil {
		t.Error(err)
	}
	// optional Close, a runtime.SetFinalizer is set anyway
	defer song.Close()

	assertFloat(t, -1.349859, song.Force, "song force")

	assertFloat(t, -0.110247, song.ForceVector.Tempo, "song tempo")
	assertFloat(t, 0.197553, song.ForceVector.Amplitude, "song amplitude")
	assertFloat(t, -1.547412, song.ForceVector.Frequency, "song frequency")
	assertFloat(t, -1.621171, song.ForceVector.Attack, "song attack")

	assertInt(t, 2, song.Channels, "song channels")
	assertInt(t, 12508554, len(song.Samples), "song samples count")
	assertInt(t, 22050, song.SampleRate, "song sample rate")
	assertInt(t, 198332, song.Bitrate, "song bitrate")
	assertInt(t, 2, song.BytesPerSample, "song bytes per sample")
	assertInt(t, 283, int(song.Duration), "song duration")

	assertString(t, "David TMX", song.Artist, "song artist")
	assertString(t, "Lost in dreams", song.Title, "song title")
	assertString(t, "Renaissance", song.Album, "song album")
	assertString(t, "14", song.TrackNumber, "song track number")
	assertString(t, "(255)", song.Genre, "song genre")
}
