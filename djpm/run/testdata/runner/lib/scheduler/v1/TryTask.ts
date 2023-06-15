// Code generated by protoc-gen-as. DO NOT EDIT.
// Versions:
//   protoc-gen-as v1.3.0
//   protoc        v3.21.12

import { Writer, Reader } from "as-proto/assembly";
import { TaskTx } from "./TaskTx";

export class TryTask {
  static encode(message: TryTask, writer: Writer): void {
    writer.uint32(8);
    writer.bool(message.needRetry);

    const taskTxs = message.taskTxs;
    for (let i: i32 = 0; i < taskTxs.length; ++i) {
      writer.uint32(18);
      writer.fork();
      TaskTx.encode(taskTxs[i], writer);
      writer.ldelim();
    }
  }

  static decode(reader: Reader, length: i32): TryTask {
    const end: usize = length < 0 ? reader.end : reader.ptr + length;
    const message = new TryTask();

    while (reader.ptr < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.needRetry = reader.bool();
          break;

        case 2:
          message.taskTxs.push(TaskTx.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  }

  needRetry: bool;
  taskTxs: Array<TaskTx>;

  constructor(needRetry: bool = false, taskTxs: Array<TaskTx> = []) {
    this.needRetry = needRetry;
    this.taskTxs = taskTxs;
  }
}
