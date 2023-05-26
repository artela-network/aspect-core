// The entry file of your WebAssembly module.
import { Aspect, Context } from "./lib/lib";
import { Schedule, PeriodicSchedule, AdHocSchedule } from "./lib/schedule";
import { Msg } from "./lib/msg";

import { AspectInput } from "./aspect/v1/AspectInput"
import { AspectOutput } from "./aspect/v1/AspectOutput"

import { MyContract } from "./generated/my_contract";

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
        return new AspectOutput();
    }

    onTxCommit(input: AspectInput): AspectOutput {
        return new AspectOutput();
    }

    onBlockFinalize(input: AspectInput): AspectOutput {
        return new AspectOutput();
    }

    private scheduleTx(): bool {
        let tx = new MyContract("0x11f4d0a3c12e86b4b5f39b213f7e19d048276dae").getBalance(new Msg(2, "20000", "30000", "0x00c5496aee77c1ba1f0854206a26dda82a81d6d8"))
        var periodicSch: Schedule = PeriodicSchedule.builder("myPeriodicSchedule").startAfter(10).count(1000).everyNBlocks(5).maxRetry(2);
        return periodicSch.submit(tx);
    }
}

export default MyFirstAspect;
