package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"

	asguard "asguardsdk"
)

func main() {
	ctx := context.Background()

	// ============================================
	// TEST 1: Fraud Service Health (Manual HTTP)
	// ============================================
	fmt.Println("=== TEST 1: Fraud Service Health ===")

	resp, err := http.Get("http://localhost:8081/health")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("HTTP %d: %s\n", resp.StatusCode, string(body))

	// ============================================
	// TEST 2: SDK Fraud Check
	// ============================================
	fmt.Println("\n=== TEST 2: SDK Fraud Check ===")

	config := asguard.NewConfiguration()
	config.Host = "localhost:8081"
	config.Scheme = "http"
	config.AddDefaultHeader("X-API-Key", "devsecret")

	client := asguard.NewAPIClient(config)

	deviceId := "device-abc"
	ipAddress := "192.168.1.100"

	req := asguard.FraudCheckRequest{
		TransactionId: "sdk-test-001",
		Amount:        250.00,
		Currency:      "USD",
		UserId:        "user-123",
		DeviceId:      &deviceId,
		IpAddress:     &ipAddress,
	}

	result, httpResp, err := client.FraudDetectionAPI.CheckFraud(ctx).FraudCheckRequest(req).Execute()
	if err != nil {
		fmt.Printf("SDK ERROR: %v\n", err)
		if httpResp != nil {
			fmt.Printf("HTTP Status: %d\n", httpResp.StatusCode)
		}
		os.Exit(1)
	}

	fmt.Printf("HTTP Status: %d\n", httpResp.StatusCode)
	fmt.Printf("Transaction: %s\n", result.TransactionId)
	fmt.Printf("Risk Score: %.2f\n", result.RiskScore)
	fmt.Printf("Risk Level: %s\n", result.RiskLevel)
	fmt.Printf("AI Triggered: %v\n", result.AiTriggered)
	fmt.Printf("AI Confidence: %.2f\n", result.AiConfidence)
	fmt.Printf("Reasons: %v\n", result.Reasons)
	fmt.Printf("Message: %s\n", result.Message)

	// ============================================
	// TEST 3: Face Service SDK (Analyze & Compare)
	// ============================================
	fmt.Println("\n=== TEST 3: Face Service SDK ===")

	// Read and encode a sample image (place "face.jpg" in the current directory)
	imageBytes, err := os.ReadFile("testface.jpg")
	if err != nil {
		fmt.Printf("ERROR reading testface.jpg: %v\n", err)
		fmt.Println("Please place a JPEG image named 'testface.jpg' in the current directory.")
		os.Exit(1)
	}
	base64Image := base64.StdEncoding.EncodeToString(imageBytes)
	dataURI := "data:image/jpeg;base64," + base64Image

	// Create face service client (different base URL and auth)
	faceConfig := asguard.NewConfiguration()
	faceConfig.Host = "localhost:8082"
	faceConfig.Scheme = "http"
	faceConfig.AddDefaultHeader("Authorization", "Bearer dev-key-123")

	faceClient := asguard.NewAPIClient(faceConfig)

	// 3a. Analyze face
	fmt.Println("  -> Calling /v1/analyze ...")
	analyzeReq := asguard.AnalyzeFaceRequest{
		Image:         dataURI,
		QualityChecks: asguard.PtrBool(true), // enable quality checks
	}
	analyzeResp, httpResp, err := faceClient.FaceVerificationAPI.AnalyzeFace(ctx).AnalyzeFaceRequest(analyzeReq).Execute()
	if err != nil {
		fmt.Printf("Analyze ERROR: %v\n", err)
		if httpResp != nil {
			fmt.Printf("HTTP Status: %d\n", httpResp.StatusCode)
		}
		os.Exit(1)
	}
	if !analyzeResp.Success {
		fmt.Printf("Analyze failed: %s\n", *analyzeResp.Error)
		os.Exit(1)
	}
	fmt.Printf("  Success: %v, Face detected: %v, Embedding length: %d\n",
		analyzeResp.Success, analyzeResp.FaceDetected, len(analyzeResp.Embedding))
	if analyzeResp.QualityScore != nil {
		fmt.Printf("  Quality score: %.2f\n", *analyzeResp.QualityScore)
	}
	if analyzeResp.Warnings != nil {
		fmt.Printf("  Warnings: %v\n", analyzeResp.Warnings)
	}

	// 3b. Compare faces (use same image as probe and the embedding we just got)
	fmt.Println("  -> Calling /v1/compare ...")
	compareReq := asguard.CompareFacesRequest{
		ProbeImage:         dataURI,
		ReferenceEmbedding: analyzeResp.Embedding,
		Threshold:          nil, // use default 0.6
	}
	compareResp, httpResp, err := faceClient.FaceVerificationAPI.CompareFaces(ctx).CompareFacesRequest(compareReq).Execute()
	if err != nil {
		fmt.Printf("Compare ERROR: %v\n", err)
		if httpResp != nil {
			fmt.Printf("HTTP Status: %d\n", httpResp.StatusCode)
		}
		os.Exit(1)
	}
	if !compareResp.Success {
		fmt.Printf("Compare failed: %s\n", *compareResp.Error)
		os.Exit(1)
	}
	fmt.Printf("  Match: %v, Confidence: %.2f, Distance: %.3f\n",
		compareResp.Match, compareResp.Confidence, compareResp.Distance)
	if compareResp.ProbeQuality != nil {
		fmt.Printf("  Probe quality: %.2f\n", *compareResp.ProbeQuality)
	}

	fmt.Println("\nAll SDK tests PASSED ✓")
}
