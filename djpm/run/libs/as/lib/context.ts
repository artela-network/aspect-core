import { AspTransaction, EthBlock } from "./message";
import { Context } from "./host";
import { ScheduleMsg } from "./scheduler";
import { StateChanges } from "./message";

export interface ScheduleCtx {
    scheduleTx(sch: ScheduleMsg): bool;
}

export interface TraceCtx {
    getStateChanges(addr: string, variable: string, key: Uint8Array): StateChanges;
}

export class onTxReceiveCtx {
    public lastBlock(): EthBlock | null {
        return Context.lastBlock();
    }

    public localCall(input: string): string {
        return Context.localCall(input);
    }

    public getProperty(key: string): string {
        return Context.getProperty(key);
    }

    public setContext(key: string, value: string): bool {
        return Context.setContext(key, value);
    }

    public getContext(key: string): string {
        return Context.getContext(key);
    }

    public setAspectState(key: string, value: string): bool {
        return Context.setAspectState(key, value);
    }

    public getAspectState(key: string): string {
        return Context.getAspectState(key);
    }

    blockHeight: i64;
    tx: AspTransaction | null;

    constructor(blockHeight: i64, tx: AspTransaction | null) {
        this.blockHeight = blockHeight;
        this.tx = tx;
    };
}

export class onBlockInitializeCtx implements ScheduleCtx {
    public lastBlock(): EthBlock | null {
        return Context.lastBlock();
    }

    public currentBlock(): EthBlock | null {
        return Context.currentBlock();
    }

    public localCall(input: string): string {
        return Context.localCall(input);
    }

    public getProperty(key: string): string {
        return Context.getProperty(key);
    }

    public setContext(key: string, value: string): bool {
        return Context.setContext(key, value);
    }

    public getContext(key: string): string {
        return Context.getContext(key);
    }

    public setAspectState(key: string, value: string): bool {
        return Context.setAspectState(key, value);
    }

    public getAspectState(key: string): string {
        return Context.getAspectState(key);
    }

    public scheduleTx(sch: ScheduleMsg): bool {
        return Context.scheduleTx(sch);
    }

    blockHeight: i64;
    tx: AspTransaction | null;

    constructor(blockHeight: i64, tx: AspTransaction | null) {
        this.blockHeight = blockHeight;
        this.tx = tx;
    };
}

export class onTxVerifyCtx {
    public lastBlock(): EthBlock | null {
        return Context.lastBlock();
    }

    public currentBlock(): EthBlock | null {
        return Context.currentBlock();
    }

    public localCall(input: string): string {
        return Context.localCall(input);
    }

    public getProperty(key: string): string {
        return Context.getProperty(key);
    }

    public setContext(key: string, value: string): bool {
        return Context.setContext(key, value);
    }

    public getContext(key: string): string {
        return Context.getContext(key);
    }

    public setAspectState(key: string, value: string): bool {
        return Context.setAspectState(key, value);
    }

    public getAspectState(key: string): string {
        return Context.getAspectState(key);
    }

    blockHeight: i64;
    tx: AspTransaction | null;

    constructor(blockHeight: i64, tx: AspTransaction | null) {
        this.blockHeight = blockHeight;
        this.tx = tx;
    };
}

export class onAccountVerifyCtx {
    public lastBlock(): EthBlock | null {
        return Context.lastBlock();
    }

    public currentBlock(): EthBlock | null {
        return Context.currentBlock();
    }

    public localCall(input: string): string {
        return Context.localCall(input);
    }

    public getProperty(key: string): string {
        return Context.getProperty(key);
    }

    public setContext(key: string, value: string): bool {
        return Context.setContext(key, value);
    }

    public getContext(key: string): string {
        return Context.getContext(key);
    }

    public setAspectState(key: string, value: string): bool {
        return Context.setAspectState(key, value);
    }

    public getAspectState(key: string): string {
        return Context.getAspectState(key);
    }

    blockHeight: i64;
    tx: AspTransaction | null;

    constructor(blockHeight: i64, tx: AspTransaction | null) {
        this.blockHeight = blockHeight;
        this.tx = tx;
    };
}

export class onGasPaymentCtx {
    public lastBlock(): EthBlock | null {
        return Context.lastBlock();
    }

    public currentBlock(): EthBlock | null {
        return Context.currentBlock();
    }

    public localCall(input: string): string {
        return Context.localCall(input);
    }

    public getProperty(key: string): string {
        return Context.getProperty(key);
    }

    public setContext(key: string, value: string): bool {
        return Context.setContext(key, value);
    }

    public getContext(key: string): string {
        return Context.getContext(key);
    }

    public setAspectState(key: string, value: string): bool {
        return Context.setAspectState(key, value);
    }

    public getAspectState(key: string): string {
        return Context.getAspectState(key);
    }

    blockHeight: i64;
    tx: AspTransaction | null;

    constructor(blockHeight: i64, tx: AspTransaction | null) {
        this.blockHeight = blockHeight;
        this.tx = tx;
    };
}

export class preTxExecuteCtx {
    public lastBlock(): EthBlock | null {
        return Context.lastBlock();
    }

    public currentBlock(): EthBlock | null {
        return Context.currentBlock();
    }

    public localCall(input: string): string {
        return Context.localCall(input);
    }

    public getProperty(key: string): string {
        return Context.getProperty(key);
    }

    public setContext(key: string, value: string): bool {
        return Context.setContext(key, value);
    }

    public getContext(key: string): string {
        return Context.getContext(key);
    }

    public setAspectState(key: string, value: string): bool {
        return Context.setAspectState(key, value);
    }

    public getAspectState(key: string): string {
        return Context.getAspectState(key);
    }

    blockHeight: i64;
    tx: AspTransaction | null;

    constructor(blockHeight: i64, tx: AspTransaction | null) {
        this.blockHeight = blockHeight;
        this.tx = tx;
    };
}

export class preContractCallCtx implements TraceCtx {
    public lastBlock(): EthBlock | null {
        return Context.lastBlock();
    }

    public currentBlock(): EthBlock | null {
        return Context.currentBlock();
    }

    public getProperty(key: string): string {
        return Context.getProperty(key);
    }

    public setContext(key: string, value: string): bool {
        return Context.setContext(key, value);
    }

    public getContext(key: string): string {
        return Context.getContext(key);
    }

    public setAspectState(key: string, value: string): bool {
        return Context.setAspectState(key, value);
    }

    public getAspectState(key: string): string {
        return Context.getAspectState(key);
    }

    public getStateChanges(addr: string, variable: string, key: Uint8Array): StateChanges {
        return Context.getStateChanges(addr, variable, key);
    }

    blockHeight: i64;
    tx: AspTransaction | null;

    constructor(blockHeight: i64, tx: AspTransaction | null) {
        this.blockHeight = blockHeight;
        this.tx = tx;
    };
}

export class postContractCallCtx implements TraceCtx {
    public lastBlock(): EthBlock | null {
        return Context.lastBlock();
    }

    public currentBlock(): EthBlock | null {
        return Context.currentBlock();
    }

    public getProperty(key: string): string {
        return Context.getProperty(key);
    }

    public setContext(key: string, value: string): bool {
        return Context.setContext(key, value);
    }

    public getContext(key: string): string {
        return Context.getContext(key);
    }

    public setAspectState(key: string, value: string): bool {
        return Context.setAspectState(key, value);
    }

    public getAspectState(key: string): string {
        return Context.getAspectState(key);
    }

    public getStateChanges(addr: string, variable: string, key: Uint8Array): StateChanges {
        return Context.getStateChanges(addr, variable, key);
    }

    blockHeight: i64;
    tx: AspTransaction | null;

    constructor(blockHeight: i64, tx: AspTransaction | null) {
        this.blockHeight = blockHeight;
        this.tx = tx;
    };
}

export class postTxExecuteCtx implements TraceCtx {
    public lastBlock(): EthBlock | null {
        return Context.lastBlock();
    }

    public currentBlock(): EthBlock | null {
        return Context.currentBlock();
    }

    public localCall(input: string): string {
        return Context.localCall(input);
    }

    public getProperty(key: string): string {
        return Context.getProperty(key);
    }

    public setContext(key: string, value: string): bool {
        return Context.setContext(key, value);
    }

    public getContext(key: string): string {
        return Context.getContext(key);
    }

    public setAspectState(key: string, value: string): bool {
        return Context.setAspectState(key, value);
    }

    public getAspectState(key: string): string {
        return Context.getAspectState(key);
    }

    public getStateChanges(addr: string, variable: string, key: Uint8Array): StateChanges {
        return Context.getStateChanges(addr, variable, key);
    }

    blockHeight: i64;
    tx: AspTransaction | null;

    constructor(blockHeight: i64, tx: AspTransaction | null) {
        this.blockHeight = blockHeight;
        this.tx = tx;
    };
}

export class onTxCommitCtx implements ScheduleCtx, TraceCtx {
    public lastBlock(): EthBlock | null {
        return Context.lastBlock();
    }

    public currentBlock(): EthBlock | null {
        return Context.currentBlock();
    }

    public localCall(input: string): string {
        return Context.localCall(input);
    }

    public getProperty(key: string): string {
        return Context.getProperty(key);
    }

    public setContext(key: string, value: string): bool {
        return Context.setContext(key, value);
    }

    public getContext(key: string): string {
        return Context.getContext(key);
    }

    public setAspectState(key: string, value: string): bool {
        return Context.setAspectState(key, value);
    }

    public getAspectState(key: string): string {
        return Context.getAspectState(key);
    }

    public scheduleTx(sch: ScheduleMsg): bool {
        return Context.scheduleTx(sch);
    }

    public getStateChanges(addr: string, variable: string, key: Uint8Array): StateChanges {
        return Context.getStateChanges(addr, variable, key);
    }

    blockHeight: i64;
    tx: AspTransaction | null;

    constructor(blockHeight: i64, tx: AspTransaction | null) {
        this.blockHeight = blockHeight;
        this.tx = tx;
    };
}

export class onBlockFinalizeCtx implements ScheduleCtx {
    public lastBlock(): EthBlock | null {
        return Context.lastBlock();
    }

    public currentBlock(): EthBlock | null {
        return Context.currentBlock();
    }

    public localCall(input: string): string {
        return Context.localCall(input);
    }

    public getProperty(key: string): string {
        return Context.getProperty(key);
    }

    public setContext(key: string, value: string): bool {
        return Context.setContext(key, value);
    }

    public getContext(key: string): string {
        return Context.getContext(key);
    }

    public setAspectState(key: string, value: string): bool {
        return Context.setAspectState(key, value);
    }

    public getAspectState(key: string): string {
        return Context.getAspectState(key);
    }

    public scheduleTx(sch: ScheduleMsg): bool {
        return Context.scheduleTx(sch);
    }

    blockHeight: i64;
    tx: AspTransaction | null;

    constructor(blockHeight: i64, tx: AspTransaction | null) {
        this.blockHeight = blockHeight;
        this.tx = tx;
    };
}