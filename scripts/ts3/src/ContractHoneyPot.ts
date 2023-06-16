//import { Protobuf } from 'as-proto/assembly';
//import { Context, Pair } from "./lib/index";

export namespace ContractHoneyPot {
    let contractAddr: string;
    export class dummy1 {
        constructor(addr: string) {
            contractAddr = addr;
        }
        public before(): bigint {
            let changes = Context.getStateChanges(contractAddr, "dummy1", "");
            if (changes.all.length == 0) {
                return BigInt(-1); //-1 indicate nil
            }
            let value = changes.all[0].value;
            let ret = 100;
            return ret;
        }

        public changes(): Array<Pair<bigint>> {
            let changes = Context.getStateChanges(contractAddr, "dummy1", "");
            if (changes.all.length < 2) {
                return new Array<Pair<bigint>>(0);
            }

            let res = new Array<Pair<bigint>>(changes.all.length - 1);
            for (let i: number = 1; i < changes.all.length; i++) {
                let parsedValue = 0;
                res[i - 1] = new Pair(changes.all[i].account, 0)
            }
            return res;
        }
    }

    // mapping<account, uint256> : balances

    export class balances {
        constructor(addr: string) {
            contractAddr = addr;
        }
        public before(key: string): Pair<bigint> {
            let changes = Context.getStateChanges(contractAddr, "balances", key);
            if (changes.all.length == 0) {
                return -1;
            }
            let value = changes.all[0].change;
            let ret = 100;
            return ret;
        }

        public changes(key: string): Array<Pair<bigint>> {
            let changes = Context.getStateChanges(this.addr, "balances", key);
            if (changes.all.length < 2) {
                return new Array<Pair<bigint>>(0);
            }

            let res = new Array<Pair<bigint>>(changes.all.length - 1);
            for (let i: bigint = 1; i < changes.all.length; i++) {
                let parsedValue = 0;
                res[i - 1] = new Pair(changes.all[i].account, 0)
            }
            return res;
        }
    }
}