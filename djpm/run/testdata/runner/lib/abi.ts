import { AI32, AString, AUint8Array, typeIndex } from "./types";
import { Utils } from "./utils"
import { ValueKind } from "../aspect/v1/ValueKind";
import { Values } from "../aspect/v1/Values";
import { Protobuf } from 'as-proto/assembly';
import { Value } from "../aspect/v1/Value";


declare namespace __Abi__ {
    function encode(types: i32, val: i32): i32;
    function decode(types: i32, data: i32): i32;
}

export class Abi {
    // encode receives the types and values, return the hex of abi codes
    // the types is a array of type with a separator of ','
    // the values should be perfect match to the types.
    static encode(types: string, val: Values): string {
        let bytes = Protobuf.encode(val, Values.encode);
        let typesPtr = new AString(types).store();
        let valPtr = new AUint8Array(bytes).store();

        let ret = new AString();
        ret.load(__Abi__.encode(typesPtr, valPtr));
        return ret.get();
    }

    // decode receives types and hex of abi codes, return the values
    static decode(types: string, data: Uint8Array): Values {
        let typePtr = new AString(types).store();
        let dataPtr = new AUint8Array(data).store();
        let ret = new AUint8Array();

        ret.load(__Abi__.decode(typePtr, dataPtr));
        const output = Protobuf.decode<Values>(ret.get(), Values.decode);
        return output;
    }
}

export class TypeValue {
    fromString(s: string): void {
        this.value = new Value(ValueKind.STRING)
        this.value.data = Utils.stringToUint8Arrary(s);
    }

    toString(): string {
        if (this.value.kind != ValueKind.STRING) {
            return "";
        }
        return Utils.uint8ArrayToString(this.value.data);
    }

    // little endian
    toInt32(): i32 {
        if (this.value.kind != ValueKind.INT || this.value.data.length < 4) {
            return 0;
        }
        let ret: i32 = 0;
        for (let i = 0; i < 4; i++) {
            let add = i32(this.value.data[i]) << (8 * i);
            ret += add;
        }
        return ret;
    }

    // little endian
    toInt64(): i64 {
        if (this.value.kind != ValueKind.INT || this.value.data.length < 8) {
            return 0;
        }
        let ret: i64 = 0;
        for (let i = 0; i < 8; i++) {
            let add = i64(this.value.data[i]) << (8 * i);
            ret += add;
        }
        return ret;
    }

    constructor(value: Value = new Value()) {
        this.value = value;
    }

    value: Value;
}