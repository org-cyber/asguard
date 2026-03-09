# Python API client for Asguard

Complete Asguard platform SDK - Fraud Detection + Face Verification.
Two independent microservices, one unified client.

## Overview

This SDK is the official Python bindings for the Asguard ecosystem. It allows simple and type-safe integration for both backend identity verification and deep algorithmic fraud checks.

- SDK package: `asguard`
- Requirements: `Python >= 3.9`

## Installation

Install the package formally using `pip`:

```bash
pip install asguard
```

## Getting Started

Because Asguard runs its domains as independent services on disparate ports, we use two separate `ApiClient` connections to route logic.

Below is an extensive quickstart to interface with and capture AI risk scores, and map image bytes to Face vectors!

### Typical Setup & Auth Implementation

```python
import base64
from pprint import pprint

# Import the core library and classes
import asguard
from asguard.api import fraud_detection_api, face_verification_api
from asguard.models import FraudCheckRequest, AnalyzeFaceRequest, CompareFacesRequest

# ==========================================
# 1. Config Object for Fraud Analysis
# ==========================================
fraud_config = asguard.Configuration(host="http://localhost:8081")
fraud_client = asguard.ApiClient(fraud_config)
# Force apply the x-api-key explicitly
fraud_client.set_default_header("X-API-Key", "devsecret")

fraud_api = fraud_detection_api.FraudDetectionApi(fraud_client)

# ==========================================
# 2. Config Object for Face Verification
# ==========================================
face_config = asguard.Configuration(host="http://localhost:8082")
face_client = asguard.ApiClient(face_config)
# Mount the Bearer token directly
face_client.set_default_header("Authorization", "Bearer dev-key-123")

face_api = face_verification_api.FaceVerificationApi(face_client)
```

### Performing a Fraud Verification

Once your SDK endpoints are mapped, creating requests takes seconds, and triggers Groq intelligence for analysis:

```python
print("Evaluating transaction...")
fraud_request = FraudCheckRequest(
    user_id="user_123",
    transaction_id="txn_456",
    amount=250000.0,
    currency="USD",
    ip_address="192.168.1.5",
    device_id="device_789",
    location="Lagos, Nigeria"
)

try:
    response = fraud_api.check_fraud(fraud_request)
    print(f"Risk Score: {response.risk_score}")
    print(f"Level: {response.risk_level}")
    if response.ai_triggered:
        print(f"AI Opinion: {response.ai_recommendation} - {response.ai_summary}")
except asguard.ApiException as e:
    print(f"Fraud API returned an error: {e}")
```

### Encoding and Checking Face Similarity

For biometrics, handle file I/O safely and pass the Base64 data to Asguard to execute high-tier validation routines.

```python
print("Encoding local image logic...")

# 1. Provide an image
with open("testface.jpg", "rb") as image_file:
    encoded_string = base64.b64encode(image_file.read()).decode('utf-8')
data_uri = f"data:image/jpeg;base64,{encoded_string}"

# 2. Check for image quality metrics and retrieve an embed
analyze_request = AnalyzeFaceRequest(
    image=data_uri,
    quality_checks=True
)

try:
    # We call analyze to extract facial node vectors
    face_resp = face_api.analyze_face(analyze_request)

    if face_resp.success:
        print(f"Extracted a length {len(face_resp.embedding)} facial matrix.")
        print(f"Image metric score: {face_resp.quality_score}")

        # 3. Supply the matrix alongside an origin test image to compare matching values
        compare_request = CompareFacesRequest(
            probe_image=data_uri,
            reference_embedding=face_resp.embedding,
            threshold=0.6
        )

        cmp_resp = face_api.compare_faces(compare_request)
        if cmp_resp.match:
            print(f"System Confirmed User Identify at {cmp_resp.confidence * 100:.1f}% Match")

except asguard.ApiException as e:
    print(f"Face API errored out: {e}")
```

## Exception Handling

All generated API wrappers throw explicit `ApiException` exceptions. Examine `e.status`, `e.reason`, and `e.body` to resolve integration errors and gracefully downgrade flows.
