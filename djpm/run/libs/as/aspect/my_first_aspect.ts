// The entry file of your WebAssembly module.
import { Schedule, PeriodicSchedule, AdHocSchedule, Opts, ScheduleTx } from "../lib/scheduler";
import { Context } from "../lib/host"
import { AspectInput, AspectOutput } from "../lib/message"
import { IAspectBlock, IAspectTransaction } from "../lib/interfaces";

import { Storage } from "./contract_storage"

class MyFirstAspect implements IAspectTransaction, IAspectBlock {
    isOwner(sender: string): bool {
        // let value = Context.getProperty("owner");
        // let owners = value.split(",");
        // if (owners.includes(sender)) {
        //     return true;
        // }
        // return false;
        return true;
    }

    onContractBinding(contractAddr: string): bool {
        // let value = Context.getProperty("binding");
        // let owners = value.split(",");
        // if (owners.includes(contractAddr)) {
        //     return true;
        // }
        // return false;
        return true;
    }

    onTxReceive(input: AspectInput): AspectOutput {
        // call host api
        let block = Context.lastBlock();

        // write response values
        let ret = new AspectOutput();
        ret.success = true;

        // add test data
        Context.setContext("k1", "v1");
        Context.setContext("k2", "v2");

        // add hostapi return data
        if (block) {
            let header = block.header ? block.header : null;
            if (header) {
                Context.setContext("lastBlockNum", header.number.toString());
            } else {
                Context.setContext("lastBlockNum", "empty");
            }
        } else {
            Context.setContext("lastBlockNum", "not found");
        }

        // schedule a tx
        // this.scheduleTx();

        ret.success = true;
        return ret;
    }

    onBlockInitialize(input: AspectInput): AspectOutput {
        return new AspectOutput(true);
    }

    onTxVerify(input: AspectInput): AspectOutput {
        return new AspectOutput(true);;
    }

    onAccountVerify(input: AspectInput): AspectOutput {
        return new AspectOutput(true);
    }

    onGasPayment(input: AspectInput): AspectOutput {
        return new AspectOutput(true);
    }

    preTxExecute(input: AspectInput): AspectOutput {
        return new AspectOutput(true);
    }

    preContractCall(input: AspectInput): AspectOutput {
        return new AspectOutput(true);
    }

    postContractCall(input: AspectInput): AspectOutput {
        return new AspectOutput(true);
    }

    postTxExecute(input: AspectInput): AspectOutput {
        let ret = new AspectOutput();
        if (input.tx != null) {
            let num1 = new Storage.number1(input.tx!.to);
            let num1_latest = num1.latest();
            Context.setContext("number1_latest", num1_latest!.change.toString())

            // let num2 = new Storage.number2(input.tx!.to);
            // let num2_latest = num2.latest();
            // Context.setContext("number2_latest", num2_latest!.change.toString())

            // let num3 = new Storage.number3(input.tx!.to);
            // let num3_latest = num3.latest();
            // Context.setContext("number3_latest", num2_latest!.change.toString())

            // let str1 = new Storage.str1(input.tx!.to);
            // let str1_latest = str1.latest();
            // Context.setContext("str1_latest", str1_latest!.change.toString())

            // let bool1 = new Storage.bool1(input.tx!.to);
            // let bool1_latest = bool1.latest();
            // Context.setContext("bool1_latest", bool1_latest!.change.toString())

            let account = new Storage.accounts(input.tx!.to);
            let tom_balance_latest = account.person("tom").balance().latest();
            if (tom_balance_latest == null) {
                Context.setContext("account_person_tome_account_latest", "is null");
            } else {
                Context.setContext("account_person_tome_account_latest_acct", tom_balance_latest.account);
                Context.setContext("account_person_tome_balance_latest_change", tom_balance_latest.change.toString());
            }
        }
        ret.success = true;
        return ret;
    }

    onTxCommit(input: AspectInput): AspectOutput {
        return new AspectOutput(true);
    }

    onBlockFinalize(input: AspectInput): AspectOutput {
        return new AspectOutput(true);
    }

    private scheduleTx(): bool {
        let scheduleTo = Context.getProperty("ScheduleTo");
        let broker = Context.getProperty("Broker");

        // let tx = new MyContract(scheduleTo).store100(new Option(0, "200000000", "30000", broker))
        let tx = new ScheduleTx(scheduleTo).New(
            "0x6057361d00000000000000000000000000000000000000000000000000000000000003e8",
            new Opts(0, "200000000", "30000", broker))

        var periodicSch: Schedule = PeriodicSchedule
            .builder("myPeriodicSchedule")
            .startAfter(3)
            .count(1000)
            .everyNBlocks(5)
            .maxRetry(2);
        return periodicSch.submit(tx);
    }
}

export default MyFirstAspect;
