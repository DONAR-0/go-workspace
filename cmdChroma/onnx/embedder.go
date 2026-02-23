package onnx

import (
	"fmt"

	"github.com/daulet/tokenizers"
	ort "github.com/yalue/onnxruntime_go"
)

type Embedder struct {
	session   *ort.DynamicAdvancedSession
	tokenizer *tokenizers.Tokenizer
}

// Embedder initialize the dictionary and the brain
func NewEmbedder(modelPath, tokenizersPath, libpath string) (*Embedder, error) {

	//1. Setup the ONNX Library
	ort.SetSharedLibraryPath(libpath)

	if err := ort.InitializeEnvironment(); err != nil {
		return nil, fmt.Errorf("error received when initialize the ONNX Library:  %w", err)
	}

	//2. Load dictionary
	tk, err := tokenizers.FromFile(tokenizersPath)
	if err != nil {
		return nil, fmt.Errorf("error received when initialize tokenizers from file: %w", err)
	}

	//3. Load brain
	inputNames := []string{"input_ids", "attention_mask", "token_type_ids"}
	outputNames := []string{"last_hidden_state"}
	sess, err := ort.NewDynamicAdvancedSession(modelPath, inputNames, outputNames, nil)
	if err != nil {
		return nil, fmt.Errorf("error received when starting a session")
	}

	return &Embedder{tokenizer: tk, session: sess}, nil
}

// Embed converts text into a 384-dimmension vector
func (e *Embedder) Embed(text string) ([]float32, error) {
	// Step A: Tokenize (Text -> IDs)
	ids, _ := e.tokenizer.Encode(text, true)
	// Step B: Prepare tensors (Numbers -> Math Format)
	lenght := int64(len(ids))
	shape := ort.NewShape(1, lenght)

	finalIDs := make([]int64, lenght)
	mask := make([]int64, lenght)
	types := make([]int64, lenght)

	for i, id := range ids {
		finalIDs[i] = int64(id)
		if id != 0 {
			mask[i] = 1
		}
		types[i] = 0
	}

	inT, _ := ort.NewTensor(shape, finalIDs)
	maT, _ := ort.NewTensor(shape, mask)
	tyT, _ := ort.NewTensor(shape, types)

	defer inT.Destroy()
	defer maT.Destroy()
	defer tyT.Destroy()

	// Step C: Run Brain (Math -> Raw Output)
	outT, _ := ort.NewEmptyTensor[float32](ort.NewShape(1, lenght, 384))
	defer outT.Destroy()

	err := e.session.Run([]ort.ArbitraryTensor{inT, maT, tyT}, []ort.ArbitraryTensor{outT})
	if err != nil {
		return nil, err
	}

	// Step D: Pooling (Raw Output -> 384 Sentence Vector)
	return outT.GetData()[:384], nil
}

func (e *Embedder) Close() {
	e.tokenizer.Close()
	e.session.Destroy()
	ort.DestroyEnvironment()
}
