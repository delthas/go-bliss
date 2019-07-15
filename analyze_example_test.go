package bliss

import (
	"fmt"
)

func ExampleAnalyze() {
	song, err := Analyze("audio/song.mp3")
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
	// Analysis for music: audio/song.mp3
	// Force: -1.349859
	// Force vector: (-0.110247, 0.197553, -1.547412, -1.621171)
	// Channels: 2
	// Number of samples: 12508554
	// Sample rate: 22050
	// Bitrate: 198332
	// Number of bytes per sample: 2
	// Calm or loud: Calm
	// Duration: 283
	// Artist: David TMX
	// Title: Lost in dreams
	// Album: Renaissance
	// Track number: 14
	// Genre: (255)
}

func ExampleAnalyze_ML() {
	song, err := Analyze("audio/song.mp3")
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
	// Lost in dreams;-0.110247;0.197553;-1.547412;-1.621171
}
