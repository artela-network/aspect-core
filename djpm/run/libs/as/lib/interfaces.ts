import { AspectOutput } from "./message";
import {
    stateCtx,
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
    isOwner(ctx: stateCtx, sender: string): bool
    onBlockInitialize(ctx: onBlockInitializeCtx): AspectOutput
    onBlockFinalize(ctx: onBlockFinalizeCtx): AspectOutput
}

export interface IAspectTransaction {
    isOwner(ctx: stateCtx, sender: string): bool
    onContractBinding(ctx: stateCtx, contractAddr: string): bool
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

