#!/usr/bin/env python3
"""
Asguard SDK Test - Python
Tests Fraud Detection and Face Verification services
"""

import sys
import base64
import requests
import urllib.request
import os
from asguard import Configuration, ApiClient
from asguard.api import FraudDetectionApi, FaceVerificationApi
from asguard.models import FraudCheckRequest, AnalyzeFaceRequest, CompareFacesRequest

# Configuration
FRAUD_BASE_URL = "http://localhost:8081"
FACE_BASE_URL = "http://localhost:8082"
API_KEY = "devsecret"
FACE_TOKEN = "dev-key-123"


def ensure_test_image():
    """Download a sample image if testface.jpg doesn't exist"""
    if not os.path.exists("testface.jpg"):
        print("Downloading sample test image...")
        try:
            url = "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a2/Arme-schwarz-wei%C3%9F.jpg/440px-Arme-schwarz-wei%C3%9F.jpg"
            urllib.request.urlretrieve(url, "testface.jpg")
            print("Downloaded testface.jpg")
        except Exception as e:
            print(f"Could not download test image: {e}")
            return False
    return True


def test_fraud_health():
    """TEST 1: Fraud Service Health (Manual HTTP)"""
    print("=== TEST 1: Fraud Service Health ===")
    
    try:
        response = requests.get(f"{FRAUD_BASE_URL}/health", timeout=5)
        print(f"HTTP {response.status_code}: {response.text}")
        response.raise_for_status()
        return True
    except Exception as e:
        print(f"ERROR: {e}")
        return False


def test_sdk_fraud_check():
    """TEST 2: SDK Fraud Check"""
    print("\n=== TEST 2: SDK Fraud Check ===")
    
    # Configure client - use set_default_header on ApiClient
    config = Configuration(host=FRAUD_BASE_URL)
    client = ApiClient(config)
    client.set_default_header("X-API-Key", API_KEY)
    
    fraud_api = FraudDetectionApi(client)
    
    # Build request
    request = FraudCheckRequest(
        transaction_id="sdk-test-py-001",
        amount=250.00,
        currency="USD",
        user_id="user-123",
        device_id="device-abc",
        ip_address="192.168.1.100"
    )
    
    try:
        result = fraud_api.check_fraud(fraud_check_request=request)
        
        print(f"Transaction: {result.transaction_id}")
        print(f"Risk Score: {result.risk_score}")
        print(f"Risk Level: {result.risk_level}")
        print(f"AI Triggered: {result.ai_triggered}")
        print(f"AI Confidence: {result.ai_confidence}")
        print(f"Reasons: {result.reasons}")
        print(f"Message: {result.message}")
        return True
        
    except Exception as e:
        print(f"SDK ERROR: {e}")
        return False


def test_face_service():
    """TEST 3: Face Service SDK (Analyze & Compare)"""
    print("\n=== TEST 3: Face Service SDK ===")
    
    # Ensure we have a test image
    if not ensure_test_image():
        print("ERROR: No testface.jpg available")
        return False
    
    # Read test image
    try:
        with open("testface.jpg", "rb") as f:
            image_bytes = f.read()
    except Exception as e:
        print(f"ERROR reading testface.jpg: {e}")
        return False
    
    # Encode to base64 data URI
    base64_image = base64.b64encode(image_bytes).decode('utf-8')
    data_uri = f"data:image/jpeg;base64,{base64_image}"
    
    # Configure face client - use set_default_header on ApiClient
    config = Configuration(host=FACE_BASE_URL)
    client = ApiClient(config)
    client.set_default_header("Authorization", f"Bearer {FACE_TOKEN}")
    
    face_api = FaceVerificationApi(client)
    
    # 3a. Analyze face
    print("  -> Calling /v1/analyze ...")
    analyze_request = AnalyzeFaceRequest(
        image=data_uri,
        quality_checks=True
    )
    
    try:
        analyze_resp = face_api.analyze_face(analyze_face_request=analyze_request)
        
        if not analyze_resp.success:
            print(f"Analyze failed: {analyze_resp.error}")
            return False
        
        print(f"  Success: {analyze_resp.success}")
        print(f"  Face detected: {analyze_resp.face_detected}")
        print(f"  Embedding length: {len(analyze_resp.embedding)}")
        
        if analyze_resp.quality_score is not None:
            print(f"  Quality score: {analyze_resp.quality_score}")
        if analyze_resp.warnings:
            print(f"  Warnings: {analyze_resp.warnings}")
            
    except Exception as e:
        print(f"Analyze ERROR: {e}")
        return False
    
    # 3b. Compare faces (use same image as probe, use embedding from analyze)
    print("  -> Calling /v1/compare ...")
    compare_request = CompareFacesRequest(
        probe_image=data_uri,
        reference_embedding=analyze_resp.embedding,
        # threshold=None  # use default 0.6
    )
    
    try:
        compare_resp = face_api.compare_faces(compare_faces_request=compare_request)
        
        if not compare_resp.success:
            print(f"Compare failed: {compare_resp.error}")
            return False
        
        print(f"  Match: {compare_resp.match}")
        print(f"  Confidence: {compare_resp.confidence}")
        print(f"  Distance: {compare_resp.distance}")
        
        if compare_resp.probe_quality is not None:
            print(f"  Probe quality: {compare_resp.probe_quality}")
            
    except Exception as e:
        print(f"Compare ERROR: {e}")
        return False
    
    return True


def main():
    """Run all tests"""
    results = []
    
    # Run tests
    results.append(("Fraud Health", test_fraud_health()))
    results.append(("SDK Fraud Check", test_sdk_fraud_check()))
    results.append(("Face Service SDK", test_face_service()))
    
    # Summary
    print("\n" + "="*50)
    print("TEST SUMMARY")
    print("="*50)
    
    all_passed = True
    for name, passed in results:
        status = "✓ PASSED" if passed else "✗ FAILED"
        print(f"{name}: {status}")
        if not passed:
            all_passed = False
    
    if all_passed:
        print("\nAll SDK tests PASSED ✓")
        sys.exit(0)
    else:
        print("\nSome tests FAILED ✗")
        sys.exit(1)


if __name__ == "__main__":
    main()