package main

import (
	"fmt"
	tokenizers "github.com/daulet/tokenizers"
	ort "github.com/yalue/onnxruntime_go"
	"log"
	"os"
)

func main() {

	// 1. Setup Paths (Update these to your actual file locations)
	// UPDATED PATH to match your 'ls' output
	libPath := "./models/onnx_runtime/onnxruntime-linux-x64-1.24.2/lib/libonnxruntime.so"
	modelPath := "./models/minilm/model.onnx"
	tokenizerPath := "./models/minilm/tokenizer.json"
	// 2. Test ONNX Shared Library Loading
	fmt.Println("Step 1: Loading ONNX Shared Library...")
	ort.SetSharedLibraryPath(libPath)
	err := ort.InitializeEnvironment()
	if err != nil {
		log.Fatalf("❌ Failed to initialize ONNX Runtime: %v\nCheck if %s exists and is compatible with your OS.", err, libPath)
	}
	defer ort.DestroyEnvironment()
	fmt.Println("✅ ONNX Runtime Initialized successfully!")

	// 3. Test Tokenizer Loading
	fmt.Println("\nStep 2: Loading Tokenizer...")
	if _, err := os.Stat(tokenizerPath); os.IsNotExist(err) {
		log.Fatalf("❌ Tokenizer file not found at %s", tokenizerPath)
	}
	tk, err := tokenizers.FromFile(tokenizerPath)
	if err != nil {
		log.Fatalf("❌ Failed to load tokenizer: %v", err)
	}
	defer tk.Close()
	fmt.Println("✅ Tokenizer loaded successfully!")

	// 4. Test ONNX Model Loading (Creating a Session)
	fmt.Println("\nStep 3: Loading ONNX Model into Session...")
	if _, err := os.Stat(modelPath); os.IsNotExist(err) {
		log.Fatalf("❌ Model file not found at %s", modelPath)
	}

	// We define the input/output names expected by all-MiniLM-L6-v2
	inputNames := []string{"input_ids", "attention_mask", "token_type_ids"}
	outputNames := []string{"last_hidden_state"}

	session, err := ort.NewDynamicAdvancedSession(modelPath, inputNames, outputNames, nil)
	if err != nil {
		log.Fatalf("❌ Failed to create ONNX session: %v\nThis usually means the model file is corrupt or incompatible.", err)
	}
	defer session.Destroy()
	fmt.Println("✅ ONNX Model loaded into session successfully!")

	fmt.Println("\n🎉 ALL SYSTEMS GO! Your embedding engine is ready.")

	// 5. Example: Tokenizing a string
	userInput := "Hello, how are you today?"

	// In your version, Encode returns: ([]uint32, []string)
	// It does NOT return an error, so we don't use 'err' here.
	ids, tokens := tk.Encode(userInput, true)

	// 6. Manual Attention Mask & Type IDs
	// For MiniLM, the Attention Mask is 1 for every token you have.
	// Corrected Attention Mask Logic
	mask := make([]int64, len(ids))
	typeIDs := make([]int64, len(ids))
	finalIDs := make([]int64, len(ids))

	for i, id := range ids {
		finalIDs[i] = int64(id) // ONNX wants int64, not uint32
		if id != 0 {
			mask[i] = 1
		} else {
			mask[i] = 0
		}
		typeIDs[i] = 0
	}
	// 7. Print the results to verify
	fmt.Printf("Tokens:         %v\n", tokens)
	fmt.Printf("Input IDs:      %v\n", ids)
	fmt.Printf("Attention Mask: %v\n", mask)
	fmt.Printf("Type IDs:       %v\n", typeIDs)

	// 1. Create the input tensors
	// Shape is [1][Length] (1 sentence, 128 tokens)
	shape := ort.NewShape(1, int64(len(finalIDs)))

	inputTensor, _ := ort.NewTensor(shape, finalIDs)
	maskTensor, _ := ort.NewTensor(shape, mask)
	typeTensor, _ := ort.NewTensor(shape, typeIDs)

	defer inputTensor.Destroy()
	defer maskTensor.Destroy()
	defer typeTensor.Destroy()

	// 2. Run the model!
	// The order MUST match the names defined in Step 3: input_ids, attention_mask, token_type_ids
	outputTensor, err := ort.NewEmptyTensor[float32](ort.NewShape(1, int64(len(finalIDs)), 384)) // MiniLM-L6-v2 outputs 384 dimensions
	defer outputTensor.Destroy()

	err = session.Run([]ort.ArbitraryTensor{inputTensor, maskTensor, typeTensor}, []ort.ArbitraryTensor{outputTensor})
	if err != nil {
		log.Fatalf("❌ Inference failed: %v", err)
	}

	// 3. Get the results
	embeddings := outputTensor.GetData()
	fmt.Printf("\n🚀 SUCCESS! Generated embedding with %d dimensions.\n", len(embeddings))
	fmt.Printf("First 5 values: %v\n", embeddings[:5])

	// The raw output contains 128 vectors of 384 dims each.
	// We only want the FIRST vector (index 0), which corresponds to the [CLS] token.
	sentenceEmbedding := embeddings[:384]

	fmt.Printf("\n--- Final Validation ---\n")
	fmt.Printf("Sentence Vector Size: %d\n", len(sentenceEmbedding))
	fmt.Printf("Sentence Vector (first 3): %v\n", sentenceEmbedding[:3])

	// Check for valid numbers (Not NaN or Infinity)
	if len(sentenceEmbedding) == 384 && !containsInvalid(sentenceEmbedding) {
		fmt.Println("✅ VALIDATION PASSED: This vector is ready for ChromaDB.")
	}
}

// Helper to check for broken math
func containsInvalid(slice []float32) bool {
	for _, v := range slice {
		if v != v {
			return true
		} // NaN check
	}
	return false
}
