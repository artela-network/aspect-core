import { AI32, AUint8Array } from "./types";

declare namespace __Abi__ {
    function decodeInt32(ptr: i32): i32
}

export class Abi {
    static asInt32(data: Uint8Array): i32 {
        let input = new AUint8Array(data);
        let output = new AI32();
        output.load(__Abi__.decodeInt32(input.store()));
        return output.body;
    }

    static asUint64(data: Uint8Array): u64 {
        return 0;
    }
}