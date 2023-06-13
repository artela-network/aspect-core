import { ABool, AString, AUint8Array, AspectInput, AspectOutput } from "./message";

import { IAspectBlock, IAspectTransaction } from "./interfaces"
import { Protobuf } from 'as-proto/assembly';

export class Entry {
    private readonly blockAspect: IAspectBlock;
    private readonly transactionAspect: IAspectTransaction;

    constructor(blockAspect: IAspectBlock, transactionAspect: IAspectTransaction) {
        this.blockAspect = blockAspect;
        this.transactionAspect = transactionAspect;
    }

    public isBlockLevel(): bool {
        return this.blockAspect != null;
    }

    public isTransactionLevel(): bool {
        return this.transactionAspect != null;
    }

    loadAspectInput(argPtr: i32): AspectInput {
        let encodedArg = new AUint8Array();
        encodedArg.load(argPtr);
        return Protobuf.decode<AspectInput>(encodedArg.get(), AspectInput.decode);
    }

    loadInputString(argPtr: i32): string {
        let arg = new AString();
        arg.load(argPtr);
        return arg.get();
    }

    storeOutputBool(out: bool): i32 {
        let b = new ABool();
        b.set(out);
        return b.store();
    }

    storeAspectOutput(output: AspectOutput): i32 {
        let encodedOutput = Protobuf.encode(output, AspectOutput.encode);
        let ret = new AUint8Array();
        ret.set(encodedOutput);
        return ret.store();
    }

    public execute(methodPtr: i32, argPtr: i32): i32 {
        let methodArg = new AString();
        methodArg.load(methodPtr);
        let method = methodArg.get();

        if (!this.isBlockLevel() && !this.isTransactionLevel()) {
            throw new Error("invalid aspect code");
        }

        switch (true) {
            case method === "onContractBinding" && this.isTransactionLevel():
                let arg = this.loadInputString(argPtr);
                let out = this.transactionAspect.onContractBinding(arg);
                return this.storeOutputBool(out);

            case method === "isOwner":
                let arg = this.loadInputString(argPtr);
                if (this.isTransactionLevel()) {
                    let out = this.transactionAspect.isOwner(arg);
                    return this.storeOutputBool(out);
                }

                let out = this.blockAspect.isOwner(arg);
                return this.storeOutputBool(out);
        }

        let arg = this.loadAspectInput(argPtr);
        var out: AspectOutput;
        switch (true) {
            case (method == "onTxReceive" && this.isTransactionLevel()):
                out = this.transactionAspect.onTxReceive(arg);
                break;

            case method == "onBlockInitialize" && this.isBlockLevel():
                out = this.blockAspect.onBlockInitialize(arg);
                break;

            case method == "onTxVerify" && this.isTransactionLevel():
                out = this.transactionAspect.onTxVerify(arg);
                break

            case method == "onAccountVerify" && this.isTransactionLevel():
                out = this.transactionAspect.onAccountVerify(arg);
                break;

            case method == "onGasPayment" && this.isTransactionLevel():
                out = this.transactionAspect.onGasPayment(arg);
                break;

            case method == "preTxExecute" && this.isTransactionLevel():
                out = this.transactionAspect.preTxExecute(arg);
                break;

            case method == "preContractCall" && this.isTransactionLevel():

                out = this.transactionAspect.preContractCall(arg);
                break;

            case method == "postContractCall" && this.isTransactionLevel():
                out = this.transactionAspect.postContractCall(arg);
                break;

            case method == "postTxExecute" && this.isTransactionLevel():
                out = this.transactionAspect.postTxExecute(arg);
                break;

            case method == "onTxCommit" && this.isTransactionLevel():
                out = this.transactionAspect.onTxCommit(arg);
                break;

            case method == "onBlockFinalize" && this.isBlockLevel():
                out = this.blockAspect.onBlockFinalize(arg);
                break;

            default:
                throw new Error("method " + method + " not found");
        }
        return this.storeAspectOutput(out);
    }
}
