/*
go-bliss are Go bindings to bliss, a small library that can analyze songs and compute the distance between two songs. It can be useful for creating "intelligent" playlists of songs that are similar.

The main algorithm works by outputting a vector V = (a,b,c,d) for each song. The euclidean distance between these vectors corresponds to the actual distance felt by listening to them: a playlist can then be built by queuing close songs, or do other stuff using mathematical tools available on euclidean 4-dimensionnal spaces.

Setup

go-bliss depends on bliss (compile-time and runtime), which itself depends on some libav* libraries (compile-time and runtime). Check the project README for detailed setup instructions.

Technical details

Check the project README for links to technical details about the analysis process.

go-bliss can compute:

- for each song, a force (float) for the overall song intensity, corresponding to a force rating (ForceRating), either calm, loud, or unknown

- for each song, a vector containg 4 floats (ForceVector), each rating an aspect of the song: tempo is the BPM of the song, amplitude is the physical force of the song, frequency is the ratio between high and low frequencies, attack is a sum of intensity of all the attacks divided by the song length

- for two songs, the distance between them (float) (the euclidian distance between their force vectors), or their cosine similarity

API

There is no global library setup or teardown method.

go-bliss can decode an audio file with Decode into a struct, Song, that contains metadata about the file and its samples. Song also contains fields corresponding to the song analysis result, that is not filled by this call.

A Song owns internal memory that can be freed explicitly by using Close when done using it, or automatically by the GC after some time (with runtime.SetFinalizer).

go-bliss can analyze an audio file with Analyze into a Song, like Decode, except it also analyzes the song and fills its analysis-related fields.

go-bliss can compute the distance between two songs (either audio files with DistanceFile, or files already decoded to songs with Distance), or their cosine similarity (with CosineSimilarity and CosineSimilarityFile).

go-bliss can also compute a specific value of a song rather than all of them with EnvelopeSort, AmplitudeSort, and FrequencySort.

Misc

go-bliss also has Mean, Variance, and RectangularFilter helpers, as well as a Version function.

Multi-threaded use

As far as I know bliss does not store global state, so processing two different songs concurrently should be fine.
*/
package bliss
