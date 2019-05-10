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

			//URL kan være annerledes hvis brukeren av programmet har valgt å benytte seg
			//av en annen server. API-nøkkel er forskjellig for hver unike bruker av IBM.
			//Her har en URL og API-nøkkel blitt fylt inn, slik at programmet kan brukes.
			URL:       "https://gateway-lon.watsonplatform.net/speech-to-text/api",
			IAMApiKey: "kpZNfLM0WZnb7qR8OpyizZDC4zQYAzWT1l3rwk-QZvE7",
		})
	if speechToTextErr != nil {
		panic(speechToTextErr)
	}

	//For å velge hvilken lydfil som skal konverteres til tekst, må man skrive inn riktig
	//path. Det er også mulig å konvertere flere lydfiler, dersom man lager en array som
	//består av flere verdier enn kun [1].
	files := [1]string{"C:/Users/Sindre/go/src/github.com/Brofo/Audio/audio-file.flac"}
	for _, file := range files {
		var audioFile io.ReadCloser
		var audioFileErr error
		audioFile, audioFileErr = os.Open(file)
		if audioFileErr != nil {
			panic(audioFileErr)
		}
		//Hvis man ønsker, kan man under her implementere funksjoner som gir mer informasjon
		//om lydfilen, for eksempel timestamps og keywords. Dette er ikke relevant for vår
		//oppgave, ettersom vi kun er interessert i konversjonen av tale til tekst.
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
