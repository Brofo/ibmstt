package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/watson-developer-cloud/go-sdk/speechtotextv1"
)

func main() {
	speechToText, speechToTextErr := speechtotextv1.
		NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			URL:       "https://gateway-lon.watsonplatform.net/speech-to-text/api",
			IAMApiKey: "kpZNfLM0WZnb7qR8OpyizZDC4zQYAzWT1l3rwk-QZvE7",
		})
	if speechToTextErr != nil {
		panic(speechToTextErr)
	}

	files := [1]string{"C:/Users/Sindre/go/src/github.com/Brofo/Audio/audio-file.flac"}
	for _, file := range files {
		var audioFile io.ReadCloser
		var audioFileErr error
		audioFile, audioFileErr = os.Open(file)
		if audioFileErr != nil {
			panic(audioFileErr)
		}
		response, responseErr := speechToText.Recognize(
			&speechtotextv1.RecognizeOptions{
				Audio: &audioFile,
			},
		)
		if responseErr != nil {
			panic(responseErr)
		}
		result := speechToText.GetRecognizeResult(response)
		b, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(b))
	}
}
