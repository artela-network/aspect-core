import { ABool, AString, AUint8Array } from "./types";
import { Protobuf } from 'as-proto/assembly';
import { BlockOutput } from "../aspect/v1/BlockOutput"
import { EthBlock } from "../aspect/v1/EthBlock";
import { AspectInput } from "../aspect/v1/AspectInput"
import { AspectOutput } from "../aspect/v1/AspectOutput"
import { Schedule } from "../scheduler/v1/Schedule"
import { StateChanges } from "../aspect/v1/StateChanges"

export interface Aspect {
    isOwner(sender: string): bool
    onContractBinding(contractAddr: string): bool

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
    isOwner(sender: string): bool {
        return false;
    }
    onContractBinding(contractAddr: string): bool {
        return false;
    }
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

        let aspect = this.buildAspect();
        if (aspect instanceof DummyAspect) {
            throw new Error("invalid aspect code");
        }

        if (method == "isOwner" || method == "onContractBinding") {
            let arg = new AString();
            arg.load(argPtr);
            var out: bool;
            if (method == "isOwner") {
                out = aspect.isOwner(arg.get());
            } else {
                out = aspect.onContractBinding(arg.get());
            }
            let b = new ABool();
            b.set(out);
            return b.store();
        }

        let encodedArg = new AUint8Array();
        encodedArg.load(argPtr);

        const input = Protobuf.decode<AspectInput>(encodedArg.get(), AspectInput.decode);
        var output: AspectOutput
        switch (true) {
            case method == "onTxReceive":
                output = aspect.onTxReceive(input);
                break;
            case method == "onBlockInitialize":
                output = aspect.onBlockInitialize(input);
                break;
            case method == "onTxVerify":
                output = aspect.onTxVerify(input);
                break;
            case method == "onAccountVerify":
                output = aspect.onAccountVerify(input);
                break;
            case method == "onGasPayment":
                output = aspect.onGasPayment(input);
                break;
            case method == "preTxExecute":
                output = aspect.preTxExecute(input);
                break;
            case method == "preContractCall":
                output = aspect.preContractCall(input);
                break;
            case method == "postTxExecute":
                output = aspect.postTxExecute(input);
                break;
            case method == "onTxCommit":
                output = aspect.onTxCommit(input);
                break;
            case method == "onBlockFinalize":
                output = aspect.onBlockFinalize(input);
                break;
            default:
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

