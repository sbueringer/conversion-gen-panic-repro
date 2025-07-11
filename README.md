# Targets

Tests:
```
make test
```

Expected output:
```test
go test ./...
--- FAIL: TestFuzzyConversion (0.00s)
    --- FAIL: TestFuzzyConversion/for_ClusterClass (0.00s)
        --- FAIL: TestFuzzyConversion/for_ClusterClass/spoke-hub-spoke (0.00s)
panic: runtime error: invalid memory address or nil pointer dereference [recovered]
        panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x2 addr=0x1 pc=0x105cd6628]

goroutine 37 [running]:
testing.tRunner.func1.2({0x106017f80, 0x106c11e90})
        ./code/pkg/mod/golang.org/toolchain@v0.0.1-go1.24.0.darwin-arm64/src/testing/testing.go:1734 +0x1ac
testing.tRunner.func1()
        ./code/pkg/mod/golang.org/toolchain@v0.0.1-go1.24.0.darwin-arm64/src/testing/testing.go:1737 +0x334
panic({0x106017f80?, 0x106c11e90?})
        ./code/pkg/mod/golang.org/toolchain@v0.0.1-go1.24.0.darwin-arm64/src/runtime/panic.go:787 +0x124
conversion-gen-panic-repro/api/v1beta1.(*ClusterClass).ConvertTo(0x106c13b70?, {0x10619fa78?, 0x140003f57a0})
        ./code/src/sigs.k8s.io/conversion-gen-panic-repro/api/v1beta1/conversion.go:23 +0x68
conversion-gen-panic-repro/api/v1beta1.TestFuzzyConversion.FuzzTestFunc.func1.1(0x140003aea80?)
        ./code/src/sigs.k8s.io/conversion-gen-panic-repro/pkg/conversion/conversion.go:90 +0x160
testing.tRunner(0x140003aea80, 0x1400018e780)
        ./code/pkg/mod/golang.org/toolchain@v0.0.1-go1.24.0.darwin-arm64/src/testing/testing.go:1792 +0xe4
created by testing.(*T).Run in goroutine 36
        .//code/pkg/mod/golang.org/toolchain@v0.0.1-go1.24.0.darwin-arm64/src/testing/testing.go:1851 +0x374
FAIL    conversion-gen-panic-repro/api/v1beta1  0.352s
?       conversion-gen-panic-repro/api/v1beta2  [no test files]
?       conversion-gen-panic-repro/pkg/conversion       [no test files]
FAIL
```

Regenerate deep-copy & conversion code:
```
make generate
```
