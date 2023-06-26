// The entry file of your WebAssembly module.
import { Opts, PeriodicSchedule, Schedule, ScheduleTx } from "../lib/scheduler";
import { AspectOutput } from "../lib/message"
import { IAspectBlock, IAspectTransaction } from "../lib/interfaces";

import { Storage } from "./contract_storage"
import { ethereum } from "../lib/abi/ethereum/coders";
import { debug } from "../lib/host/debug";
import { ScheduleCtx } from "../lib/context";
import { BigInt } from "../lib/message";

import {
    StateCtx,
    OnTxReceiveCtx,
    OnBlockInitializeCtx,
    OnTxVerifyCtx,
    OnAccountVerifyCtx,
    OnGasPaymentCtx,
    PreTxExecuteCtx,
    PreContractCallCtx,
    PostContractCallCtx,
    PostTxExecuteCtx,
    OnTxCommitCtx,
    OnBlockFinalizeCtx
} from "../lib/context";

class MyFirstAspect implements IAspectTransaction, IAspectBlock {
    isOwner(ctx: StateCtx, sender: string): bool {
        // let value = ctx.getProperty("owner");
        // let owners = value.split(",");
        // if (owners.includes(sender)) {
        //     return true;
        // }
        // return false;
        return true;
    }

    onContractBinding(ctx: StateCtx, contractAddr: string): bool {
        // let value = ctx.getProperty("binding");
        // let owners = value.split(",");
        // if (owners.includes(contractAddr)) {
        //     return true;
        // }
        // return false;
        return true;
    }

    onTxReceive(ctx: OnTxReceiveCtx): AspectOutput {
        // call host api
        let block = ctx.lastBlock();

        // write response values
        let ret = new AspectOutput();
        ret.success = true;

        // add test data
        ctx.setContext("k1", "v1");
        ctx.setContext("k2", "v2");

        // add hostapi return data
        if (block) {
            let header = block.header ? block.header : null;
            if (header) {
                ctx.setContext("lastBlockNum", header.number.toString());
            } else {
                ctx.setContext("lastBlockNum", "empty");
            }
        } else {
            ctx.setContext("lastBlockNum", "not found");
        }

        ret.success = true;
        return ret;
    }

    onBlockInitialize(ctx: OnBlockInitializeCtx): AspectOutput {
        // this.scheduleTx(ctx, ctx.getProperty("ScheduleTo"), ctx.getProperty("Broker"));
        return new AspectOutput(true);
    }

    onTxVerify(ctx: OnTxVerifyCtx): AspectOutput {
        return new AspectOutput(true);;
    }

    onAccountVerify(ctx: OnAccountVerifyCtx): AspectOutput {
        return new AspectOutput(true);
    }

    onGasPayment(ctx: OnGasPaymentCtx): AspectOutput {
        return new AspectOutput(true);
    }

    preTxExecute(ctx: PreTxExecuteCtx): AspectOutput {
        ///
        /// example to get sender's balance
        ///
        if (ctx.tx != null) {
            let balance = ctx.currentBalance(ctx.tx!.from);
            if (balance) {
                ctx.setContext("current-balance-pre", balance.toString(16));
            } else {
                ctx.setContext("current-balance-pre", "is null");
            }
        } else {
            ctx.setContext("current-balance-pre", "ctx.tx is null");
        }

        return new AspectOutput(true);
    }

    preContractCall(ctx: PreContractCallCtx): AspectOutput {
        return new AspectOutput(true);
    }

    postContractCall(ctx: PostContractCallCtx): AspectOutput {
        return new AspectOutput(true);
    }

    postTxExecute(ctx: PostTxExecuteCtx): AspectOutput {
        ///
        /// example to get sender's balance
        ///
        if (ctx.tx != null) {
            let balance = ctx.currentBalance(ctx.tx!.from);
            if (balance) {
                let post_balance = balance;
                let pre_balance_str = ctx.getContext("current-balance-pre");
                let pre_balance = BigInt.fromString(pre_balance_str, 16);
                let diff = post_balance.sub(pre_balance);
            }
        }


        ///
        /// example to trace states
        ///
        let ret = new AspectOutput();
        if (ctx.tx != null) {
            let num1 = new Storage.number1(ctx, ctx.tx!.to);
            let num1_latest = num1.latest();
            ctx.setContext("number1_latest", num1_latest!.change.toString())

            // let num2 = new Storage.number2(ctx.tx!.to);
            // let num2_latest = num2.latest();
            // ctx.setContext("number2_latest", num2_latest!.change.toString())

            // let num3 = new Storage.number3(ctx.tx!.to);
            // let num3_latest = num3.latest();
            // ctx.setContext("number3_latest", num2_latest!.change.toString())

            // let str1 = new Storage.str1(ctx.tx!.to);
            // let str1_latest = str1.latest();
            // ctx.setContext("str1_latest", str1_latest!.change.toString())

            // let bool1 = new Storage.bool1(ctx.tx!.to);
            // let bool1_latest = bool1.latest();
            // ctx.setContext("bool1_latest", bool1_latest!.change.toString())

            let account = new Storage.accounts(ctx, ctx.tx!.to);
            let tom_balance_latest = account.person("tom").balance().latest();
            if (tom_balance_latest == null) {
                ctx.setContext("account_person_tome_account_latest", "is null");
            } else {
                ctx.setContext("account_person_tome_account_latest_acct", tom_balance_latest.account);
                ctx.setContext("account_person_tome_balance_latest_change", tom_balance_latest.change.toString());
            }

            let allowance = new Storage.allowances(ctx, ctx.tx!.to);
            let addr1 = ethereum.Address.fromHexString('0xA32576Bd17b4e9dDfFaD917068F80A4ff7b7409a');
            let addr2 = ethereum.Address.fromHexString('0xBC0E48f5d2A48350B74c81ECed3A42b35b532ef8');
            let value = allowance.value(addr1).latest(addr2);
            if (value) {
                ctx.setContext("allowance_value_latest_change", value.change.toString());
                ctx.setContext("allowance_value_latest_acct", value.account);
            } else {
                ctx.setContext("allowance_value_latest_acct", "is null");
            }
        }
        ret.success = true;
        return ret;
    }

    onTxCommit(ctx: OnTxCommitCtx): AspectOutput {
        return new AspectOutput(true);
    }

    onBlockFinalize(ctx: OnBlockFinalizeCtx): AspectOutput {
        return new AspectOutput(true);
    }

    private scheduleTx(ctx: ScheduleCtx, scheduleTo: string, broker: string): bool {
        // let tx = new MyContract(scheduleTo).store100(new Option(0, "200000000", "30000", broker))
        let num = ethereum.Number.fromU64(1);
        let addr = ethereum.Address.fromHexString('0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c');
        let str = ethereum.String.fromString('haha');

        let tuple = ethereum.Tuple.fromCoders([num, addr, str]);
        let array = ethereum.ArraySlice.fromCoders([tuple]);
        let payload = ethereum.abiEncode('myMethod', [array, tuple, num, addr, str]);

        debug.log(payload);

        let tx = new ScheduleTx(scheduleTo).New(
            payload,
            new Opts(0, "200000000", "30000", broker))

        var periodicSch: Schedule = PeriodicSchedule
            .new(ctx, "myPeriodicSchedule")
            .startAfter(3)
            .count(1000)
            .everyNBlocks(5)
            .maxRetry(2);
        return periodicSch.submit(tx);
    }
}

export default MyFirstAspect;
