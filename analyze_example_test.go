package bliss

import (
	"fmt"
)

func ExampleAnalyze() {
	song, err := Analyze("audio/song.flac")
	if err != nil {
		fmt.Println("couldn't analyze song")
		return
	}
	defer song.Close()

	var rating string
	switch song.ForceRating {
	case Calm:
		rating = "Calm"
	case Loud:
		rating = "Loud"
	default:
		rating = "Unknown"
	}

	fmt.Printf("Analysis for music: %s\n", song.Filename)
	fmt.Printf("Force: %f\n", song.Force)
	fmt.Printf("Force vector: (%f, %f, %f, %f)\n",
		song.ForceVector.Tempo,
		song.ForceVector.Amplitude,
		song.ForceVector.Frequency,
		song.ForceVector.Attack)
	fmt.Printf("Channels: %d\n", song.Channels)
	fmt.Printf("Number of samples: %d\n", len(song.Samples))
	fmt.Printf("Sample rate: %d\n", song.SampleRate)
	fmt.Printf("Bitrate: %d\n", song.Bitrate)
	fmt.Printf("Number of bytes per sample: %d\n", song.BytesPerSample)
	fmt.Printf("Calm or loud: %s\n", rating)
	fmt.Printf("Duration: %d\n", song.Duration)
	fmt.Printf("Artist: %s\n", song.Artist)
	fmt.Printf("Title: %s\n", song.Title)
	fmt.Printf("Album: %s\n", song.Album)
	fmt.Printf("Track number: %s\n", song.TrackNumber)
	fmt.Printf("Genre: %s\n", song.Genre)

	// Output:
	// Analysis for music: audio/song.flac
	// Force: -25.165920
	// Force vector: (-8.945454, -15.029835, -10.136086, -15.560563)
	// Channels: 2
	// Number of samples: 488138
	// Sample rate: 22050
	// Bitrate: 233864
	// Number of bytes per sample: 2
	// Calm or loud: Calm
	// Duration: 11
	// Artist: David TMX
	// Title: Renaissance
	// Album: Renaissance
	// Track number: 02
	// Genre: Pop
}

func ExampleAnalyze_ML() {
	song, err := Analyze("audio/song.flac")
	if err != nil {
		fmt.Println("couldn't analyze song")
		return
	}
	defer song.Close()

	fmt.Printf("%s;%f;%f;%f;%f\n",
		song.Title,
		song.ForceVector.Tempo,
		song.ForceVector.Amplitude,
		song.ForceVector.Frequency,
		song.ForceVector.Attack)

	// Output:
	// Renaissance;-8.945454;-15.029835;-10.136086;-15.560563
}
