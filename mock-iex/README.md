# Mock IEX

## DESCRIPTION
This package generates a test server for IEX which returns recorded responses.

Useful for testing `github.com/jonwho/go-iex` library.

Another package in `recorder/` is also helpful to record responses into YAML files to assert on.

## USE
### RECORDER
Export ENV VAR

```sh
export IEX_TEST_SECRET_TOKEN=Tsk_ahsvyao12u4u0ausvn1o3rhw988120yf_FAKE
export IEX_TEST_PUBLISHABLE_TOKEN=Tpk_la091720ihakbso128uihotbfao_FAKE
export IEX_SECRET_TOKEN=Tsk_ahsvyao12u4u0ausvn1o3rhw988120yf_REAL
export IEX_PUBLISHABLE_TOKEN=Tpk_la091720ihakbso128uihotbfao_REAL
```
`go run recorder.go`
