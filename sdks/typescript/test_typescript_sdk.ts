/**
 * Asguard SDK Test - TypeScript/Node.js
 * Tests Fraud Detection and Face Verification services
 */

import * as fs from 'fs';
import * as http from 'http';
import { Configuration } from './configuration';
import { FraudDetectionApi, FaceVerificationApi } from './api';
import { FraudCheckRequest, AnalyzeFaceRequest, CompareFacesRequest } from './api';

// Configuration
const FRAUD_BASE_URL = 'http://localhost:8081';
const FACE_BASE_URL = 'http://localhost:8082';
const API_KEY = 'devsecret';
const FACE_TOKEN = 'dev-key-123';

// Helper: HTTP GET promise
function httpGet(url: string): Promise<{ statusCode: number; body: string }> {
    return new Promise((resolve, reject) => {
        http.get(url, (res) => {
            let data = '';
            res.on('data', chunk => data += chunk);
            res.on('end', () => resolve({ statusCode: res.statusCode || 0, body: data }));
        }).on('error', reject);
    });
}

// Helper: Download file if not exists
async function ensureTestImage(): Promise<boolean> {
    if (fs.existsSync('testface.jpg')) {
        return true;
    }
    
    console.log('Downloading sample test image...');
    try {
        const url = 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a2/Arme-schwarz-wei%C3%9F.jpg/440px-Arme-schwarz-wei%C3%9F.jpg';
        const response = await new Promise<http.IncomingMessage>((resolve, reject) => {
            http.get(url, resolve).on('error', reject);
        });
        
        const fileStream = fs.createWriteStream('testface.jpg');
        response.pipe(fileStream);
        
        await new Promise((resolve, reject) => {
            fileStream.on('finish', resolve);
            fileStream.on('error', reject);
        });
        
        console.log('Downloaded testface.jpg');
        return true;
    } catch (e) {
        console.error('Could not download test image:', e);
        return false;
    }
}

// TEST 1: Fraud Service Health (Manual HTTP)
async function testFraudHealth(): Promise<boolean> {
    console.log('=== TEST 1: Fraud Service Health ===');
    
    try {
        const response = await httpGet(`${FRAUD_BASE_URL}/health`);
        console.log(`HTTP ${response.statusCode}: ${response.body}`);
        return response.statusCode === 200;
    } catch (error) {
        console.error('ERROR:', error);
        return false;
    }
}

// TEST 2: SDK Fraud Check
async function testSdkFraudCheck(): Promise<boolean> {
    console.log('\n=== TEST 2: SDK Fraud Check ===');
    
    // Configure client - use baseOptions for headers
    const config = new Configuration({
        basePath: FRAUD_BASE_URL,
        baseOptions: {
            headers: {
                'X-API-Key': API_KEY
            }
        }
    });
    
    const fraudApi = new FraudDetectionApi(config);
    
    const request: FraudCheckRequest = {
        transaction_id: 'sdk-test-ts-001',
        amount: 250.00,
        currency: 'USD',
        user_id: 'user-123',
        device_id: 'device-abc',
        ip_address: '192.168.1.100'
    };
    
    try {
        // API returns AxiosResponse, access .data for the actual response
        const response = await fraudApi.checkFraud(request);
        const result = response.data;
        
        console.log(`Transaction: ${result.transaction_id}`);
        console.log(`Risk Score: ${result.risk_score}`);
        console.log(`Risk Level: ${result.risk_level}`);
        console.log(`AI Triggered: ${result.ai_triggered}`);
        console.log(`AI Confidence: ${result.ai_confidence}`);
        console.log(`Reasons: ${JSON.stringify(result.reasons)}`);
        console.log(`Message: ${result.message}`);
        
        return true;
    } catch (error) {
        console.error('SDK ERROR:', error);
        return false;
    }
}

// TEST 3: Face Service SDK (Analyze & Compare)
async function testFaceService(): Promise<boolean> {
    console.log('\n=== TEST 3: Face Service SDK ===');
    
    // Ensure we have a test image
    if (!(await ensureTestImage())) {
        console.error('ERROR: No testface.jpg available');
        return false;
    }
    
    // Read test image
    let imageBytes: Buffer;
    try {
        imageBytes = fs.readFileSync('testface.jpg');
    } catch (error) {
        console.error('ERROR reading testface.jpg:', error);
        return false;
    }
    
    // Encode to base64 data URI
    const base64Image = imageBytes.toString('base64');
    const dataURI = `data:image/jpeg;base64,${base64Image}`;
    
    // Configure face client
    const config = new Configuration({
        basePath: FACE_BASE_URL,
        baseOptions: {
            headers: {
                'Authorization': `Bearer ${FACE_TOKEN}`
            }
        }
    });
    
    const faceApi = new FaceVerificationApi(config);
    
    // 3a. Analyze face
    console.log('  -> Calling /v1/analyze ...');
    const analyzeRequest: AnalyzeFaceRequest = {
        image: dataURI,
        quality_checks: true
    };
    
    let embedding: number[];
    try {
        const response = await faceApi.analyzeFace(analyzeRequest);
        const analyzeResp = response.data;
        
        if (!analyzeResp.success) {
            console.error(`Analyze failed: ${analyzeResp.error}`);
            return false;
        }
        
        console.log(`  Success: ${analyzeResp.success}`);
        console.log(`  Face detected: ${analyzeResp.face_detected}`);
        console.log(`  Embedding length: ${analyzeResp.embedding.length}`);
        
        if (analyzeResp.quality_score !== undefined) {
            console.log(`  Quality score: ${analyzeResp.quality_score}`);
        }
        if (analyzeResp.warnings) {
            console.log(`  Warnings: ${JSON.stringify(analyzeResp.warnings)}`);
        }
        
        embedding = analyzeResp.embedding;
    } catch (error) {
        console.error('Analyze ERROR:', error);
        return false;
    }
    
    // 3b. Compare faces
    console.log('  -> Calling /v1/compare ...');
    const compareRequest: CompareFacesRequest = {
        probe_image: dataURI,
        reference_embedding: embedding
    };
    
    try {
        const response = await faceApi.compareFaces(compareRequest);
        const compareResp = response.data;
        
        if (!compareResp.success) {
            console.error(`Compare failed: ${compareResp.error}`);
            return false;
        }
        
        console.log(`  Match: ${compareResp.match}`);
        console.log(`  Confidence: ${compareResp.confidence}`);
        console.log(`  Distance: ${compareResp.distance}`);
        
        if (compareResp.probe_quality !== undefined) {
            console.log(`  Probe quality: ${compareResp.probe_quality}`);
        }
        
    } catch (error) {
        console.error('Compare ERROR:', error);
        return false;
    }
    
    return true;
}

// Main execution
async function main() {
    const results: Array<{ name: string; passed: boolean }> = [];
    
    // Run tests
    results.push({ name: 'Fraud Health', passed: await testFraudHealth() });
    results.push({ name: 'SDK Fraud Check', passed: await testSdkFraudCheck() });
    results.push({ name: 'Face Service SDK', passed: await testFaceService() });
    
    // Summary
    console.log('\n' + '='.repeat(50));
    console.log('TEST SUMMARY');
    console.log('='.repeat(50));
    
    let allPassed = true;
    for (const { name, passed } of results) {
        const status = passed ? '✓ PASSED' : '✗ FAILED';
        console.log(`${name}: ${status}`);
        if (!passed) allPassed = false;
    }
    
    if (allPassed) {
        console.log('\nAll SDK tests PASSED ✓');
        process.exit(0);
    } else {
        console.log('\nSome tests FAILED ✗');
        process.exit(1);
    }
}

main().catch(error => {
    console.error('Unexpected error:', error);
    process.exit(1);
});