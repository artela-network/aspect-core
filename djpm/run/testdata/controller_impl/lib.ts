import { AString, AUint8Array } from "./types";
import { Protobuf } from 'as-proto/assembly';
import { BlockOutput } from "./ethermint/aspect/v1/BlockOutput"
import { EthBlock } from "./ethermint/aspect/v1/EthBlock";
import { AspectInput } from "./ethermint/aspect/v1/AspectInput"
import { AspectOutput } from "./ethermint/aspect/v1/AspectOutput"

export interface Aspect {
    onTxReceive(arg: AspectInput): AspectOutput
    onBlockInitialize(arg: AspectInput): AspectOutput
    onTxVerify(arg: AspectInput): AspectOutput
    onAccountVerify(arg: AspectInput): AspectOutput
    onGasPayment(arg: AspectInput): AspectOutput
    preTxExecute(arg: AspectInput): AspectOutput
    preContractCall(arg: AspectInput): AspectOutput
    postContractCall(arg: AspectInput): AspectOutput
    postTxExecute(arg: AspectInput): AspectOutput
    onTxCommit(arg: AspectInput): AspectOutput
    onBlockFinalize(arg: AspectInput): AspectOutput
}

class DummyAspect implements Aspect {
    onTxReceive(arg: AspectInput): AspectOutput {
        return new AspectOutput();
    }
    onBlockInitialize(arg: AspectInput): AspectOutput {
        return new AspectOutput();
    }
    onTxVerify(arg: AspectInput): AspectOutput {
        return new AspectOutput();
    }
    onAccountVerify(arg: AspectInput): AspectOutput {
        return new AspectOutput();
    }
    onGasPayment(arg: AspectInput): AspectOutput {
        return new AspectOutput();
    }
    preTxExecute(arg: AspectInput): AspectOutput {
        return new AspectOutput();
    }
    preContractCall(arg: AspectInput): AspectOutput {
        return new AspectOutput();
    }
    postContractCall(arg: AspectInput): AspectOutput {
        return new AspectOutput();
    }
    postTxExecute(arg: AspectInput): AspectOutput {
        return new AspectOutput();
    }
    onTxCommit(arg: AspectInput): AspectOutput {
        return new AspectOutput();
    }
    onBlockFinalize(arg: AspectInput): AspectOutput {
        return new AspectOutput();
    }
}

class Entry {
    public buildAspect: () => Aspect;

    constructor() {
        this.buildAspect = function () {
            return new DummyAspect();
        }
    }

    public execute(methodPtr: i32, argPtr: i32): i32 {
        let methodArg = new AString();
        methodArg.load(methodPtr);
        let method = methodArg.get();

        let encodedArg = new AUint8Array();
        encodedArg.load(argPtr);

        let aspect = this.buildAspect();
        if (aspect instanceof DummyAspect) {
            throw new Error("invalid aspect code");
        }

        const input = Protobuf.decode<AspectInput>(encodedArg.get(), AspectInput.decode);
        var output: AspectOutput
        if (method == "onTxReceive") {
            output = aspect.onTxReceive(input);
        } else if (method == "onBlockInitialize") {
            output = aspect.onBlockInitialize(input);
        } else if (method == "onTxVerify") {
            output = aspect.onTxVerify(input);
        } else if (method == "onAccountVerify") {
            output = aspect.onAccountVerify(input);
        } else if (method == "onGasPayment") {
            output = aspect.onGasPayment(input);
        } else if (method == "preTxExecute") {
            output = aspect.preTxExecute(input);
        } else if (method == "preContractCall") {
            output = aspect.preContractCall(input);
        } else if (method == "postTxExecute") {
            output = aspect.postTxExecute(input);
        } else if (method == "onTxCommit") {
            output = aspect.onTxCommit(input);
        } else if (method == "onBlockFinalize") {
            output = aspect.onBlockFinalize(input);
        } else {
            throw new Error("method " + method + " not valid");
        }

        let encodedOutput = Protobuf.encode(output, AspectOutput.encode);
        let ret = new AUint8Array();
        ret.set(encodedOutput);
        let retPtr = ret.store();
        return retPtr;
    }
}

export let entry = new Entry();

export function execute(methodPtr: i32, argPtr: i32): i32 {
    return entry.execute(methodPtr, argPtr)
}

export function allocate(size: i32): i32 {
    return heap.alloc(size) as i32;
}

declare namespace __HostApi__ {
    function lastBlock(): i32
    function currentBlock(): i32
    function localCall(ptr: i32): i32
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
}

// Util part of hostapis
export class Util {
}

// Crypto part of hostapis
export class Crypto {
}

// Tx part of hostapis
export class Tx {
}
