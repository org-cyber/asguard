# TypeScript (Axios) API client for Asguard

Complete Asguard platform SDK - Fraud Detection + Face Verification.
Two independent microservices, one unified client.

## Overview

A Promise-based, TypeScript compatible API client wrapper extending Axios syntax to interact seamlessly with Asguard capabilities.

- SDK package: `@org-cyber/asguard`
- Requirements: `Node.js >= 14`, `Axios >= 1.0`

## Installation

Using `npm` to pull down the compiled SDK interfaces:

```bash
npm install @org-cyber/asguard
```

## Getting Started

The TypeScript SDK abstracts all endpoints gracefully behind Promises. To configure these successfully, be sure to declare endpoints appropriately in initial `Configuration` dictionaries. Let's cover establishing connections to **both** Face API and Fraud API.

### 1. Connecting Clients

```typescript
import {
  Configuration,
  FraudDetectionApi,
  FaceVerificationApi,
} from "@org-cyber/asguard";

// Set up Fraud - Note the X-API-Key injected into Axios headers
const fraudConfig = new Configuration({
  basePath: "http://localhost:8081",
  baseOptions: {
    headers: {
      "X-API-Key": "devsecret",
    },
  },
});
const fraudApi = new FraudDetectionApi(fraudConfig);

// Set up Face - Note Bearer Authentication injected via headers
const faceConfig = new Configuration({
  basePath: "http://localhost:8082",
  baseOptions: {
    headers: {
      Authorization: "Bearer dev-key-123",
    },
  },
});
const faceApi = new FaceVerificationApi(faceConfig);
```

### 2. Issuing Fraud Checks

The SDK manages type generation perfectly, utilizing defined Interfaces like `FraudCheckRequest` to resolve type anomalies before compilation:

```typescript
import { FraudCheckRequest } from "@org-cyber/asguard";

async function verifySale() {
  const payload: FraudCheckRequest = {
    user_id: "user_123",
    transaction_id: "txn_456",
    amount: 250000.0,
    currency: "USD",
    ip_address: "192.168.1.5",
    device_id: "device_789",
  };

  try {
    // Under Axios, the parsed payload resolves within the '.data' attribute
    const response = await fraudApi.checkFraud(payload);
    const result = response.data;

    console.log(`Risk Score: ${result.risk_score}`);
    console.log(`Risk Level: ${result.risk_level}`);

    if (result.ai_triggered) {
      console.log(`Alert AI Analysis: ${result.ai_recommendation}`);
      console.log(`Review details: ${result.ai_summary}`);
    }
  } catch (e: any) {
    console.error("HTTP Failure:", e.response?.data || e.message);
  }
}
```

### 3. Evaluating Biometrics

With TypeScript/NodeJS scripts, handling Base64 encodes is essential to parsing image attributes correctly to our Golang Face API parser.

```typescript
import * as fs from "fs";
import { AnalyzeFaceRequest, CompareFacesRequest } from "@org-cyber/asguard";

async function verifyFaceId() {
  // 1. Buffer the image into Base64 formats
  const imageBytes = fs.readFileSync("testface.jpg");
  const base64Image = imageBytes.toString("base64");
  const dataURI = `data:image/jpeg;base64,${base64Image}`;

  try {
    // 2. Transmit to Extract Neural Embeddings
    const extractOpts: AnalyzeFaceRequest = {
      image: dataURI,
      quality_checks: true,
    };

    const analysis = await faceApi.analyzeFace(extractOpts);
    const aResult = analysis.data;

    if (!aResult.success) {
      console.error(aResult.error);
      return;
    }

    console.log(
      `System extracted embeddings vector of size ${aResult.embedding.length}`,
    );

    // 3. Forward the Embedding and DataURI backward into the comparison algorithm!
    const compareOpts: CompareFacesRequest = {
      probe_image: dataURI,
      reference_embedding: aResult.embedding,
      threshold: 0.6,
    };

    const compare = await faceApi.compareFaces(compareOpts);
    const cResult = compare.data;

    if (cResult.match) {
      console.log(
        `Similarity Identified! Final Confidence Rate Modeled at ${cResult.confidence}`,
      );
    }
  } catch (e: any) {
    console.error(
      "Biometrics Engine Rejection:",
      e.response?.data || e.message,
    );
  }
}
```

## Types and Documentation

As part of the install base, complete TypeScript index typings (`index.d.ts`) natively populate hints across VS Code and JetBrains IDE deployments outlining field scopes implicitly!

Refer to `/docs` nested within the TypeScript installation locally for raw markdown evaluations of Models and Definitions.
