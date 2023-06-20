import { ABool, AspectInput, AspectOutput, AString, AUint8Array } from "./message";

import { IAspectBlock, IAspectTransaction } from "./interfaces"
import { Protobuf } from 'as-proto/assembly';
import {
    onTxReceiveCtx,
    onBlockInitializeCtx,
    onTxVerifyCtx,
    onAccountVerifyCtx,
    onGasPaymentCtx,
    preTxExecuteCtx,
    preContractCallCtx,
    postContractCallCtx,
    postTxExecuteCtx,
    onTxCommitCtx,
    onBlockFinalizeCtx
} from "../lib/context";

export class Entry {
    private readonly blockAspect: IAspectBlock;
    private readonly transactionAspect: IAspectTransaction;

    constructor(blockAspect: IAspectBlock, transactionAspect: IAspectTransaction) {
        this.blockAspect = blockAspect;
        this.transactionAspect = transactionAspect;
    }

    public isBlockLevel(): i32 {
        return this.storeOutputBool(this.blockAspect != null);
    }

    public isTransactionLevel(): i32 {
        return this.storeOutputBool(this.transactionAspect != null);
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

        if (this.blockAspect == null && this.transactionAspect == null) {
            throw new Error("invalid aspect code");
        }

        switch (true) {
            case method === "onContractBinding" && this.transactionAspect != null:
                let arg = this.loadInputString(argPtr);
                let out = this.transactionAspect.onContractBinding(arg);
                return this.storeOutputBool(out);

            case method === "isOwner":
                let arg = this.loadInputString(argPtr);
                if (this.transactionAspect != null) {
                    let out = this.transactionAspect.isOwner(arg);
                    return this.storeOutputBool(out);
                }

                let out = this.blockAspect.isOwner(arg);
                return this.storeOutputBool(out);
        }

        let arg = this.loadAspectInput(argPtr);
        var out: AspectOutput;
        switch (true) {
            case (method == "onTxReceive" && this.transactionAspect != null):
                let ctx = new onTxReceiveCtx(arg.blockHeight, arg.tx);
                out = this.transactionAspect.onTxReceive(ctx);
                break;

            case method == "onBlockInitialize" && this.blockAspect != null:
                let ctx = new onBlockInitializeCtx(arg.blockHeight, arg.tx);
                out = this.blockAspect.onBlockInitialize(ctx);
                break;

            case method == "onTxVerify" && this.transactionAspect != null:
                let ctx = new onTxVerifyCtx(arg.blockHeight, arg.tx);
                out = this.transactionAspect.onTxVerify(ctx);
                break

            case method == "onAccountVerify" && this.transactionAspect != null:
                let ctx = new onAccountVerifyCtx(arg.blockHeight, arg.tx);
                out = this.transactionAspect.onAccountVerify(ctx);
                break;

            case method == "onGasPayment" && this.transactionAspect != null:
                let ctx = new onGasPaymentCtx(arg.blockHeight, arg.tx);
                out = this.transactionAspect.onGasPayment(ctx);
                break;

            case method == "preTxExecute" && this.transactionAspect != null:
                let ctx = new preTxExecuteCtx(arg.blockHeight, arg.tx);
                out = this.transactionAspect.preTxExecute(ctx);
                break;

            case method == "preContractCall" && this.transactionAspect != null:
                let ctx = new preContractCallCtx(arg.blockHeight, arg.tx);
                out = this.transactionAspect.preContractCall(ctx);
                break;

            case method == "postContractCall" && this.transactionAspect != null:
                let ctx = new postContractCallCtx(arg.blockHeight, arg.tx);
                out = this.transactionAspect.postContractCall(ctx);
                break;

            case method == "postTxExecute" && this.transactionAspect != null:
                let ctx = new postTxExecuteCtx(arg.blockHeight, arg.tx);
                out = this.transactionAspect.postTxExecute(ctx);
                break;

            case method == "onTxCommit" && this.transactionAspect != null:
                let ctx = new onTxCommitCtx(arg.blockHeight, arg.tx);
                out = this.transactionAspect.onTxCommit(ctx);
                break;

            case method == "onBlockFinalize" && this.blockAspect != null:
                let ctx = new onBlockFinalizeCtx(arg.blockHeight, arg.tx);
                out = this.blockAspect.onBlockFinalize(ctx);
                break;

            default:
                throw new Error("method " + method + " not found");
        }
        return this.storeAspectOutput(out);
    }
}
