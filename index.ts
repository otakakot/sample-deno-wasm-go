/// <reference path="./wasm_exec.d.ts" />

import "./wasm_exec.js";

const module = await Deno.readFile("./main.wasm");

const go = new Go();

const { instance } = await WebAssembly.instantiate(module, go.importObject);

go.run(instance);

// deno-lint-ignore no-explicit-any
const golog = (globalThis as any).golog;

golog("Hello, World!");
