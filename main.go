package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bytecodealliance/wasmtime-go/v15"
)

func main() {
	// Wasm ファイルのパス
	wasmFilePath := "index.wasm"

	wasiConfig := wasmtime.NewWasiConfig()
	wasiConfig.InheritStdout()
	wasiConfig.InheritStderr()
	wasiConfig.InheritStdin()

	// Wasmtime 環境の作成
	engine := wasmtime.NewEngine()
	store := wasmtime.NewStore(engine)
	store.SetWasi(wasiConfig)

	// Wasm ファイルの読み込み
	module, err := wasmtime.NewModuleFromFile(store.Engine, wasmFilePath)
	if err != nil {
		fmt.Println("Error loading Wasm module:", err)
		os.Exit(1)
	}

	// モジュールが要求するインポートの一覧を出力
	// for _, import_ := range module.Imports() {
	// 	fmt.Printf("Module: %s, Name: %s\n", import_.Module(), *import_.Name())
	// }

	// WASI モジュールのインポートを取得
	linker := wasmtime.NewLinker(engine)
	err = linker.DefineWasi()
	if err != nil {
		log.Fatalf("failed to define WASI: %v", err)
	}

	// インスタンスの作成
	instance, err := linker.Instantiate(store, module)
	if err != nil {
		log.Fatalf("failed to instantiate module: %v", err)
	}

	// 'foo' 関数の取得
	fooFunc := instance.GetFunc(store, "foo")
	if fooFunc == nil {
		fmt.Println("Error: 'main' function not found in Wasm module")
		os.Exit(1)
	}

	// 'foo' 関数の実行
	_, err = fooFunc.Call(store)
	if err != nil {
		fmt.Println("Error calling 'main' function:", err)
		os.Exit(1)
	}

	// プログラムの終了
	fmt.Println("Wasm execution completed.")
}
