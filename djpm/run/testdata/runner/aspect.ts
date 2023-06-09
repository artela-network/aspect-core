// The entry file of your WebAssembly module.
import { Aspect, Context, Schedule, PeriodicSchedule, AdHocSchedule, Option, ScheduleTx } from "./lib/index";

import { AspectInput } from "./aspect/v1/AspectInput"
import { AspectOutput } from "./aspect/v1/AspectOutput"

import { MyContract } from "./generated/my_contract";
import { Storage } from "./contract_storage"

class MyFirstAspect implements Aspect {
    isOwner(sender: string): bool {
        let value = Context.getProperty("owner");
        let owners = value.split(",");
        if (owners.includes(sender)) {
            return true;
        }
        return false;
    }

    onContractBinding(contractAddr: string): bool {
        let value = Context.getProperty("binding");
        let owners = value.split(",");
        if (owners.includes(contractAddr)) {
            return true;
        }
        return false;
    }

    onTxReceive(input: AspectInput): AspectOutput {
        // call host api
        let block = Context.lastBlock();

        // write response values
        let ret = new AspectOutput();
        ret.success = true;

        // add test data
        ret.context.set("k1", "v1");
        ret.context.set("k2", "v2");

        // add hostapi return data
        if (block) {
            let header = block.header ? block.header : null;
            if (header) {
                ret.context.set("lastBlockNum", header.number.toString());
            } else {
                ret.context.set("lastBlockNum", "empty");
            }
        } else {
            ret.context.set("lastBlockNum", "not found");
        }

        // add input data to output
        const keys = input.context.keys();
        for (let i = 0, len = keys.length; i < len; i++) {
            ret.context.set(keys[i], input.context.get(keys[i]))
        }

        // schedule a tx
        this.scheduleTx();

        return ret;
    }

    onBlockInitialize(input: AspectInput): AspectOutput {
        return new AspectOutput();
    }

    onTxVerify(input: AspectInput): AspectOutput {
        return new AspectOutput();
    }

    onAccountVerify(input: AspectInput): AspectOutput {
        return new AspectOutput();
    }

    onGasPayment(input: AspectInput): AspectOutput {
        return new AspectOutput();
    }

    preTxExecute(input: AspectInput): AspectOutput {
        return new AspectOutput();
    }

    preContractCall(input: AspectInput): AspectOutput {
        return new AspectOutput();
    }

    postContractCall(input: AspectInput): AspectOutput {
        return new AspectOutput();
    }

    postTxExecute(input: AspectInput): AspectOutput {
        let ret = new AspectOutput();
        if (input.tx != null) {
            let num1 = new Storage.number1(input.tx!.to);
            let num1_latest = num1.latest();
            ret.context.set("number1_latest", num1_latest!.change.toString())

            let num2 = new Storage.number2(input.tx!.to);
            let num2_latest = num2.latest();
            ret.context.set("number2_latest", num2_latest!.change.toString())

            let num3 = new Storage.number3(input.tx!.to);
            let num3_latest = num3.latest();
            ret.context.set("number3_latest", num2_latest!.change.toString())

            let str1 = new Storage.str1(input.tx!.to);
            let str1_latest = str1.latest();
            ret.context.set("str1_latest", str1_latest!.change.toString())

            let bool1 = new Storage.bool1(input.tx!.to);
            let bool1_latest = bool1.latest();
            ret.context.set("bool1_latest", bool1_latest!.change.toString())
        }
        return ret;
    }

    onTxCommit(input: AspectInput): AspectOutput {
        return new AspectOutput();
    }

    onBlockFinalize(input: AspectInput): AspectOutput {
        return new AspectOutput();
    }

    private scheduleTx(): bool {
        let scheduleTo = Context.getProperty("ScheduleTo");
        let broker = Context.getProperty("Broker");

        // let tx = new MyContract(scheduleTo).store100(new Option(0, "200000000", "30000", broker))
        let tx = new ScheduleTx(scheduleTo).New(
            "0x6057361d00000000000000000000000000000000000000000000000000000000000003e8",
            new Option(0, "200000000", "30000", broker))

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
