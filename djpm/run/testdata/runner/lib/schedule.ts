import {AspTransaction} from "../aspect/v1/AspTransaction";
import {Context} from "./lib";
import {Schedule as ScheduleMsg} from "../scheduler/v1/Schedule";
import {ScheduleId as ScheduleMsgId} from "../scheduler/v1/ScheduleId";
import {ScheduleStatus} from "../scheduler/v1/ScheduleStatus";

export interface Schedule {
    submit(tran: AspTransaction): bool
}

export class PeriodicSchedule implements Schedule {
    public submit(tran: AspTransaction): bool {
        let sch = new ScheduleMsg();
        sch.count = this._count;
        sch.startBlock = this._startBlock;
        sch.maxRetry = this._maxRetry;
        sch.everyNBlock = this._everyNBlocks;
        sch.tx = tran;
        sch.status=ScheduleStatus.Open

        sch.id = new ScheduleMsgId(this._name, "");
        // sch.id.aspectId, createHeight will be set in the hostapi
        return Context.scheduleTx(sch);
    }

    public static builder(name: string): PeriodicSchedule {
        return new PeriodicSchedule(name);
    }

    public startAfter(num: u64): PeriodicSchedule {
        this._startBlock = num; // add current height in host api.
        return this;
    }

    public everyNBlocks(num: u32): PeriodicSchedule {
        this._everyNBlocks = num;
        return this;
    }

    public maxRetry(num: u32): PeriodicSchedule {
        this._maxRetry = num;
        return this;
    }

    public count(num: u64): PeriodicSchedule {
        this._count = num;
        return this;
    }

    private _name: string;
    private _everyNBlocks: u32;
    private _maxRetry: u32;
    private _count: u64;
    private _startBlock: u64;

    constructor(
        name: string = "",
        startAfter: u64 = 0,
        everyNBlocks: u32 = 0,
        count: u64 = 0,
        maxRetry: u32 = 0,
    ) {
        // startBlock add current height in host api.
        this._startBlock = startAfter;
        this._name = name;
        this._everyNBlocks = everyNBlocks;
        this._maxRetry = maxRetry;
        this._count = count;
    }
}

export class AdHocSchedule implements Schedule {
    public submit(tran: AspTransaction): bool {
        let sch = new ScheduleMsg();
        sch.count = 1;
        // startBlock add current height in host api.
        sch.startBlock = this._nextNBlocks;
        sch.maxRetry = this._maxRetry;

        sch.id = new ScheduleMsgId(this._name, "");
        // sch.id.aspectId, createHeight will be set in the hostapi
        return Context.scheduleTx(sch);
    }

    public static builder(name: string): AdHocSchedule {
        return new AdHocSchedule(name);
    }

    public maxRetry(num: u32): AdHocSchedule {
        this._maxRetry = num;
        return this;
    }

    public nextNBlocks(num: u64): AdHocSchedule {
        this._nextNBlocks = num;
        return this;
    }

    private _name: string;
    private _maxRetry: u32;
    private _nextNBlocks: u64;

    constructor(
        name: string = "",
        maxRetry: u32 = 0,
        nextNBlocks: u64 = 0,
    ) {
        this._name = name;
        this._nextNBlocks = nextNBlocks;
        this._maxRetry = maxRetry;
    }
}
