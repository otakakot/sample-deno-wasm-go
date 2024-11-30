/// <reference path="./wasm_exec.d.ts" />

import "./wasm_exec.js";

import mainwasm from "./mainwasm.ts";

import { decodeBase64 } from "https://deno.land/std@0.224.0/encoding/base64.ts";

const module = decodeBase64(mainwasm);

const go = new Go();

// ref: https://developer.mozilla.org/ja/docs/WebAssembly/JavaScript_interface/instantiate_static
const { instance } = await WebAssembly.instantiate(module, go.importObject);

go.run(instance);

// deno-lint-ignore no-explicit-any
const hello = (globalThis as any).hello;

hello("Hello", "World");
