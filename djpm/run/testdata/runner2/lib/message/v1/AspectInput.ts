// Code generated by protoc-gen-as. DO NOT EDIT.
// Versions:
//   protoc-gen-as v1.3.0
//   protoc        v3.21.12

import { Writer, Reader } from "as-proto/assembly";
import { AspTransaction } from "./AspTransaction";

export class AspectInput {
  static encode(message: AspectInput, writer: Writer): void {
    writer.uint32(8);
    writer.int64(message.blockHeight);

    const tx = message.tx;
    if (tx !== null) {
      writer.uint32(18);
      writer.fork();
      AspTransaction.encode(tx, writer);
      writer.ldelim();
    }

    const context = message.context;
    if (context !== null) {
      const contextKeys = context.keys();
      for (let i: i32 = 0; i < contextKeys.length; ++i) {
        const contextKey = contextKeys[i];
        writer.uint32(26);
        writer.fork();
        writer.uint32(10);
        writer.string(contextKey);
        writer.uint32(18);
        writer.string(context.get(contextKey));
        writer.ldelim();
      }
    }
  }

  static decode(reader: Reader, length: i32): AspectInput {
    const end: usize = length < 0 ? reader.end : reader.ptr + length;
    const message = new AspectInput();

    while (reader.ptr < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.blockHeight = reader.int64();
          break;

        case 2:
          message.tx = AspTransaction.decode(reader, reader.uint32());
          break;

        case 3:
          let contextKey: string = "";
          let contextValue: string = "";
          let contextHasKey: bool = false;
          let contextHasValue: bool = false;
          for (
            const end: usize = reader.ptr + reader.uint32();
            reader.ptr < end;

          ) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
              case 1:
                contextKey = reader.string();
                contextHasKey = true;
                break;

              case 2:
                contextValue = reader.string();
                contextHasValue = true;
                break;

              default:
                reader.skipType(tag & 7);
                break;
            }
            if (message.context === null) {
              message.context = new Map<string, string>();
            }
            const context = message.context;
            if (context !== null && contextHasKey && contextHasValue) {
              context.set(contextKey, contextValue);
            }
          }
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  }

  blockHeight: i64;
  tx: AspTransaction | null;
  context: Map<string, string>;

  constructor(
    blockHeight: i64 = 0,
    tx: AspTransaction | null = null,
    context: Map<string, string> = new Map()
  ) {
    this.blockHeight = blockHeight;
    this.tx = tx;
    this.context = context;
  }
}
