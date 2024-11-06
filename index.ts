/// <reference path="./wasm_exec.d.ts" />

import "./wasm_exec.js";

import mainwasm from "./mainwasm.ts";

import { decode } from "https://deno.land/std@0.139.0/encoding/base64.ts";

const go = new Go();

const module = decode(mainwasm);

const { instance } = await WebAssembly.instantiate(module, go.importObject);

go.run(instance);

// deno-lint-ignore no-explicit-any
const handle = (globalThis as any).handle;

handle("hello!", "hoge");
