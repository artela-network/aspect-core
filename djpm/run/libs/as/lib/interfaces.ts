import { AspectInput, AspectOutput } from "./message";
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

export interface IAspectBlock {
    isOwner(sender: string): bool
    onBlockInitialize(ctx: onBlockInitializeCtx): AspectOutput
    onBlockFinalize(ctx: onBlockFinalizeCtx): AspectOutput
}

export interface IAspectTransaction {
    isOwner(sender: string): bool
    onContractBinding(contractAddr: string): bool
    onTxReceive(ctx: onTxReceiveCtx): AspectOutput
    onTxVerify(ctx: onTxVerifyCtx): AspectOutput
    onAccountVerify(ctx: onAccountVerifyCtx): AspectOutput
    onGasPayment(ctx: onGasPaymentCtx): AspectOutput
    preTxExecute(ctx: preTxExecuteCtx): AspectOutput
    preContractCall(ctx: preContractCallCtx): AspectOutput
    postContractCall(ctx: postContractCallCtx): AspectOutput
    postTxExecute(ctx: postTxExecuteCtx): AspectOutput
    onTxCommit(ctx: onTxCommitCtx): AspectOutput
}

