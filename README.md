# javy + wasmtime-go

## js を wasm にコンパイルする

```bash
$ pnpx javy-cli compile index.js --wit index.wit -n index-world -o index.wasm
```

## wasmtime-go で動かす

```bash
$ go run main.go
```

## wasmtime で動かす

```bash
$ wasmtime run --invoke foo index.wasm
```
