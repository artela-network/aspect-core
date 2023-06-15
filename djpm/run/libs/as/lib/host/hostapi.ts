import { BlockOutput, EthBlock, StateChanges, ABool, AString, AUint8Array } from "../message";
import { Protobuf } from "as-proto/assembly";
import { ScheduleMsg } from "../scheduler";


declare namespace __HostApi__ {
    function lastBlock(): i32

    function currentBlock(): i32

    function localCall(ptr: i32): i32

    function getProperty(ptr: i32): i32

    function scheduleTx(ptr: i32): i32

    function setContext(keyPtr: i32, valuePtr: i32): bool

    function getContext(ptr: i32): i32

    function setAspectState(keyPtr: i32, valuePtr: i32): bool

    function getAspectState(ptr: i32): i32

    function getStateChanges(addr: i32, variable: i32, key: i32): i32
}

// Context part of hostapis
export class Context {
    static lastBlock(): EthBlock | null {
        let ret = __HostApi__.lastBlock();
        // read bytes from the output, and then unmarshal via proto
        let bytes = new AUint8Array();
        bytes.load(ret);
        const output = Protobuf.decode<BlockOutput>(bytes.get(), BlockOutput.decode);
        // here we can read more error message from output.res.error
        return output.block
    }

    static currentBlock(): EthBlock | null {
        let ret = __HostApi__.currentBlock();
        let bytes = new AUint8Array();
        bytes.load(ret);
        const output = Protobuf.decode<BlockOutput>(bytes.get(), BlockOutput.decode);
        return output.block;
    }

    static localCall(input: string): string {
        // TODO support local call input/output
        return "localCall params is not support for now"
    }

    static getProperty(key: string): string {
        let input = new AString();
        input.set(key);
        let inPtr = input.store();
        let outPtr = __HostApi__.getProperty(inPtr);
        let output = new AString();
        output.load(outPtr);
        return output.get();
    }

    static setContext(key: string, value: string): bool {
        let inputKey = new AString();
        inputKey.set(key);
        let inPtr = inputKey.store();
        let inputValue = new AString();
        inputValue.set(value);
        let ptrValue = inputValue.store();
        return __HostApi__.setContext(inPtr, ptrValue);
    }

    static getContext(key: string): string {
        let input = new AString();
        input.set(key);
        let inPtr = input.store();
        let outPtr = __HostApi__.getContext(inPtr);
        let output = new AString();
        output.load(outPtr);
        return output.get();
    }

    static setAspectState(key: string, value: string): bool {
        let inputKey = new AString();
        inputKey.set(key);
        let inPtr = inputKey.store();
        let inputValue = new AString();
        inputValue.set(value);
        let ptrValue = inputValue.store();

        return __HostApi__.setAspectState(inPtr, ptrValue);

    }

    static getAspectState(key: string): string {
        let input = new AString();
        input.set(key);
        let inPtr = input.store();
        let outPtr = __HostApi__.getAspectState(inPtr);
        let output = new AString();
        output.load(outPtr);
        return output.get();
    }

    static scheduleTx(sch: ScheduleMsg): bool {
        const encoded = Protobuf.encode(sch, ScheduleMsg.encode);
        let input = new AUint8Array();
        input.set(encoded);
        let inputPtr = input.store();
        let ret = __HostApi__.scheduleTx(inputPtr);
        let output = new ABool();
        output.load(ret);
        return output.get();
    }

    static getStateChanges(addr: string, variable: string, key: Uint8Array): StateChanges {
        let addrPtr = (new AString(addr)).store();
        let variablePtr = (new AString(variable)).store();
        let keyPtr = (new AUint8Array(key)).store();

        let dataPtr = __HostApi__.getStateChanges(addrPtr, variablePtr, keyPtr);
        let data = new AUint8Array();
        data.load(dataPtr);

        const changes = Protobuf.decode<StateChanges>(data.get(), StateChanges.decode);
        return changes;
    }
}