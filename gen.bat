go run gen.go

pushd dsp32\dsp
go test
popd

pushd dsp64\dsp
go test
popd

pause
