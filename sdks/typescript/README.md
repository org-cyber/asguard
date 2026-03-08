## asguard@1.0.0

This generator creates TypeScript/JavaScript client that utilizes [axios](https://github.com/axios/axios). The generated Node module can be used in the following environments:

Environment
* Node.js
* Webpack
* Browserify

Language level
* ES5 - you must have a Promises/A+ library installed
* ES6

Module system
* CommonJS
* ES6 module system

It can be used in both TypeScript and JavaScript. In TypeScript, the definition will be automatically resolved via `package.json`. ([Reference](https://www.typescriptlang.org/docs/handbook/declaration-files/consumption.html))

### Building

To build and compile the typescript sources to javascript use:
```
npm install
npm run build
```

### Publishing

First build the package then run `npm publish`

### Consuming

navigate to the folder of your consuming project and run one of the following commands.

_published:_

```
npm install asguard@1.0.0 --save
```

_unPublished (not recommended):_

```
npm install PATH_TO_GENERATED_PACKAGE --save
```

### Documentation for API Endpoints

All URIs are relative to *http://localhost:8081*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FaceVerificationApi* | [**analyzeFace**](docs/FaceVerificationApi.md#analyzeface) | **POST** /v1/analyze | Analyze face image quality and extract embedding
*FaceVerificationApi* | [**compareFaces**](docs/FaceVerificationApi.md#comparefaces) | **POST** /v1/compare | Compare a probe face image with a reference embedding
*FraudDetectionApi* | [**checkFraud**](docs/FraudDetectionApi.md#checkfraud) | **POST** /analyze | Check transaction for fraud risk
*SystemApi* | [**healthCheckFace**](docs/SystemApi.md#healthcheckface) | **GET** /face/health | Face service health check
*SystemApi* | [**healthCheckFraud**](docs/SystemApi.md#healthcheckfraud) | **GET** /fraud/health | Fraud service health check


### Documentation For Models

 - [AnalyzeFaceRequest](docs/AnalyzeFaceRequest.md)
 - [AnalyzeResponse](docs/AnalyzeResponse.md)
 - [CompareFacesRequest](docs/CompareFacesRequest.md)
 - [CompareResponse](docs/CompareResponse.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [FraudCheckRequest](docs/FraudCheckRequest.md)
 - [FraudCheckResponse](docs/FraudCheckResponse.md)
 - [HealthResponse](docs/HealthResponse.md)


<a id="documentation-for-authorization"></a>
## Documentation For Authorization


Authentication schemes defined for the API:
<a id="ApiKeyAuth"></a>
### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: X-API-Key
- **Location**: HTTP header

<a id="BearerAuth"></a>
### BearerAuth

- **Type**: Bearer authentication (API Key or JWT)

