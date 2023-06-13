// generated by abi.json

import { Protobuf } from 'as-proto/assembly';
import { Context, Abi } from "../lib/host";
import { State } from "../lib/states"
import { Utils } from "../lib/utils"
import { BigInt } from "../lib/message"

export namespace Storage {
    ///
    /// The following codes for properity: uint256 number1;
    ///
    export class number1 {
        public before(): State<BigInt> | null {
            let changes = Context.getStateChanges(this.addr, "Storage.number1", this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let account = changes.all[0].account;
            let valueHex = Utils.uint8ArrayToHex(changes.all[0].value);
            let value = BigInt.fromString(valueHex, 16);
            return new State(account, value);
        }

        public changes(): Array<State<BigInt>> | null {
            let changes = Context.getStateChanges(this.addr, "Storage.number1", this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let res = new Array<State<BigInt>>(changes.all.length);
            for (let i = 0; i < changes.all.length; i++) {
                let account = changes.all[i].account;
                let valueHex = Utils.uint8ArrayToHex(changes.all[0].value);
                let value = BigInt.fromString(valueHex, 16);
                res[i] = new State(account, value)
            }
            return res;
        }

        public latest(): State<BigInt> | null {
            let changes = Context.getStateChanges(this.addr, "Storage.number1", this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let index = changes.all.length - 1;
            let account = changes.all[index].account;
            let valueHex = Utils.uint8ArrayToHex(changes.all[index].value);
            let value = BigInt.fromString(valueHex, 16);
            return new State(account, value);
        }

        public diff(): BigInt {
            let changes = Context.getStateChanges(this.addr, "Storage.number1", this.prefix);
            if (changes.all.length < 2) {
                return BigInt.ZERO;
            }

            let beforeHex = Utils.uint8ArrayToHex(changes.all[0].value);
            let before = BigInt.fromString(beforeHex, 16);

            let afterHex = Utils.uint8ArrayToHex(changes.all[changes.all.length - 1].value);
            let after = BigInt.fromString(beforeHex, 16);

            return after.sub(before);
        }

        addr: string;
        prefix: Uint8Array;

        constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
            this.addr = addr;
            this.prefix = prefix;
        }
    }

    ///
    /// The following codes for properity: int32 number2;
    ///
    export class number2 {
        public before(): State<i32> | null {
            let changes = Context.getStateChanges(this.addr, "Storage.number2", this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let account = changes.all[0].account;
            let valueHex = Utils.uint8ArrayToHex(changes.all[0].value);
            let value = BigInt.fromString(valueHex, 16).toInt32();
            return new State(account, value);
        }

        public changes(): Array<State<i32>> | null {
            let changes = Context.getStateChanges(this.addr, "Storage.number2", this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let res = new Array<State<i32>>(changes.all.length);
            for (let i = 0; i < changes.all.length; i++) {
                let account = changes.all[i].account;
                let valueHex = Utils.uint8ArrayToHex(changes.all[0].value);
                let value = BigInt.fromString(valueHex, 16).toInt32();
                res[i] = new State(account, value)
            }
            return res;
        }

        public latest(): State<i32> | null {
            let changes = Context.getStateChanges(this.addr, "Storage.number2", this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let index = changes.all.length - 1;
            let account = changes.all[index].account;
            let valueHex = Utils.uint8ArrayToHex(changes.all[index].value);
            let value = BigInt.fromString(valueHex, 16).toInt32();
            return new State(account, value);
        }

        public diff(): i32 {
            let changes = Context.getStateChanges(this.addr, "Storage.number2", this.prefix);
            if (changes.all.length < 2) {
                return 0;
            }

            let beforeHex = Utils.uint8ArrayToHex(changes.all[0].value);
            let before = BigInt.fromString(beforeHex, 16).toInt32();

            let afterHex = Utils.uint8ArrayToHex(changes.all[changes.all.length - 1].value);
            let after = BigInt.fromString(beforeHex, 16).toInt32();

            return after - before;
        }

        addr: string;
        prefix: Uint8Array;

        constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
            this.addr = addr;
            this.prefix = prefix;
        }
    }

    ///
    /// The following codes for properity: uint64 number3;
    ///
    export class number3 {
        public before(): State<u64> | null {
            let changes = Context.getStateChanges(this.addr, "Storage.number3", this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let account = changes.all[0].account;
            let valueHex = Utils.uint8ArrayToHex(changes.all[0].value);
            let value = BigInt.fromString(valueHex, 16).toUInt64();
            return new State(account, value);
        }

        public changes(): Array<State<u64>> | null {
            let changes = Context.getStateChanges(this.addr, "Storage.number3", this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let res = new Array<State<u64>>(changes.all.length);
            for (let i = 0; i < changes.all.length; i++) {
                let account = changes.all[i].account;
                let valueHex = Utils.uint8ArrayToHex(changes.all[0].value);
                let value = BigInt.fromString(valueHex, 16).toUInt64();
                res[i] = new State(account, value)
            }
            return res;
        }

        public latest(): State<u64> | null {
            let changes = Context.getStateChanges(this.addr, "Storage.number3", this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let index = changes.all.length - 1;
            let account = changes.all[index].account;
            let valueHex = Utils.uint8ArrayToHex(changes.all[index].value);
            let value = BigInt.fromString(valueHex, 16).toUInt64();
            return new State(account, value);
        }

        public diff(): u64 {
            let changes = Context.getStateChanges(this.addr, "Storage.number3", this.prefix);
            if (changes.all.length < 2) {
                return 0;
            }

            let beforeHex = Utils.uint8ArrayToHex(changes.all[0].value);
            let before = BigInt.fromString(beforeHex, 16).toUInt64();

            let afterHex = Utils.uint8ArrayToHex(changes.all[changes.all.length - 1].value);
            let after = BigInt.fromString(beforeHex, 16).toUInt64();

            return after - before;
        }

        addr: string;
        prefix: Uint8Array;

        constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
            this.addr = addr;
            this.prefix = prefix;
        }
    }

    ///
    /// The following codes for properity: string str1;
    ///
    export class str1 {
        public before(): State<string> | null {
            let changes = Context.getStateChanges(this.addr, "Storage.str1", this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let account = changes.all[0].account;
            let value = Utils.uint8ArrayToString(changes.all[0].value);
            return new State(account, value)
        }

        public changes(): Array<State<string>> | null {
            let changes = Context.getStateChanges(this.addr, "Storage.str1", this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let res = new Array<State<string>>(changes.all.length);
            for (let i = 0; i < changes.all.length; i++) {
                let account = changes.all[i].account;
                let value = Utils.uint8ArrayToString(changes.all[i].value)
                res[i] = new State(account, value)
            }
            return res;
        }

        public latest(): State<string> | null {
            let changes = Context.getStateChanges(this.addr, "Storage.str1", this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let index = changes.all.length - 1;
            let account = changes.all[changes.all.length - 1].account;
            let value = Utils.uint8ArrayToString(changes.all[changes.all.length - 1].value)
            return new State(account, value);
        }

        addr: string;
        prefix: Uint8Array;

        constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
            this.addr = addr;
            this.prefix = prefix;
        }
    }

    ///
    /// The following codes for properity: bool bool1;
    ///
    export class bool1 {
        public before(): State<bool> | null {
            let changes = Context.getStateChanges(this.addr, "Storage.bool1", this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let account = changes.all[0].account;
            let value = Utils.uint8ArrayToBool(changes.all[0].value);
            return new State(account, value);
        }

        public changes(): Array<State<bool>> | null {
            let changes = Context.getStateChanges(this.addr, "Storage.bool1", this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let res = new Array<State<bool>>(changes.all.length);
            for (let i = 0; i < changes.all.length; i++) {
                let account = changes.all[i].account;
                let value = Utils.uint8ArrayToBool(changes.all[i].value);
                res[i] = new State(account, value)
            }
            return res;
        }

        public latest(): State<bool> | null {
            let changes = Context.getStateChanges(this.addr, "Storage.bool1", this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let index = changes.all.length - 1;
            let account = changes.all[index].account;
            let value = Utils.uint8ArrayToBool(changes.all[index].value);
            return new State(account, value);
        }

        addr: string
        prefix: Uint8Array

        constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
            this.addr = addr;
            this.prefix = prefix;
        }
    }

    ///
    /// The following codes for properity: mapping(string => Person) public accounts;
    ///
    export class accounts {
        public person(key: string): Person {
            let encoded = Abi.encodeString(key);
            return new Person(this.addr, "Storage.accounts", Utils.concatUint8Arrays(this.prefix, encoded))
        }

        addr: string;
        prefix: Uint8Array;

        constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
            this.addr = addr;
            this.prefix = prefix;
        }
    }

    export class Person {
        public id(): Person_id {
            let encoded = Abi.encodeString("id");
            return new Person_id(this.addr, this.variable, Utils.concatUint8Arrays(this.prefix, encoded));
        }

        public balance(): Person_balance {
            let encoded = Abi.encodeString("balance");
            return new Person_balance(this.addr, this.variable, Utils.concatUint8Arrays(this.prefix, encoded));
        }

        addr: string;
        variable: string;
        prefix: Uint8Array;

        constructor(addr: string, varibale: string, prefix: Uint8Array = new Uint8Array(0)) {
            this.addr = addr;
            this.variable = varibale;
            this.prefix = prefix;
        }
    }

    // uint64 id;
    export class Person_id {
        public before(): State<u64> | null {
            let changes = Context.getStateChanges(this.addr, this.variable, this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let account = changes.all[0].account;
            let valueHex = Utils.uint8ArrayToHex(changes.all[0].value);
            let value = BigInt.fromString(valueHex, 16).toUInt64();
            return new State(account, value);
        }

        public changes(): Array<State<u64>> | null {
            let changes = Context.getStateChanges(this.addr, this.variable, this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let res = new Array<State<u64>>(changes.all.length);
            for (let i = 0; i < changes.all.length; i++) {
                let account = changes.all[i].account;
                let valueHex = Utils.uint8ArrayToHex(changes.all[0].value);
                let value = BigInt.fromString(valueHex, 16).toUInt64();
                res[i] = new State(account, value)
            }
            return res;
        }

        public latest(): State<u64> | null {
            let changes = Context.getStateChanges(this.addr, this.variable, this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let index = changes.all.length - 1;
            let account = changes.all[index].account;
            let valueHex = Utils.uint8ArrayToHex(changes.all[index].value);
            let value = BigInt.fromString(valueHex, 16).toUInt64();
            return new State(account, value);
        }

        public diff(): u64 {
            let changes = Context.getStateChanges(this.addr, this.variable, this.prefix);
            if (changes.all.length < 2) {
                return 0;
            }

            let beforeHex = Utils.uint8ArrayToHex(changes.all[0].value);
            let before = BigInt.fromString(beforeHex, 16).toUInt64();

            let afterHex = Utils.uint8ArrayToHex(changes.all[changes.all.length - 1].value);
            let after = BigInt.fromString(beforeHex, 16).toUInt64();

            return after - before;
        }

        variable: string;
        addr: string;
        prefix: Uint8Array;

        constructor(addr: string, variable: string = "", prefix: Uint8Array = new Uint8Array(0)) {
            this.addr = addr;
            this.variable = variable;
            this.prefix = prefix;
        }
    }

    // uint32 balance;
    export class Person_balance {
        public before(): State<u32> | null {
            let changes = Context.getStateChanges(this.addr, this.variable, this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let account = changes.all[0].account;
            let valueHex = Utils.uint8ArrayToHex(changes.all[0].value);
            let value = BigInt.fromString(valueHex, 16).toUInt32();
            return new State(account, value);
        }

        public changes(): Array<State<u32>> | null {
            let changes = Context.getStateChanges(this.addr, this.variable, this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let res = new Array<State<u32>>(changes.all.length);
            for (let i = 0; i < changes.all.length; i++) {
                let account = changes.all[i].account;
                let valueHex = Utils.uint8ArrayToHex(changes.all[0].value);
                let value = BigInt.fromString(valueHex, 16).toUInt32();
                res[i] = new State(account, value)
            }
            return res;
        }

        public latest(): State<u32> | null {
            let changes = Context.getStateChanges(this.addr, this.variable, this.prefix);
            if (changes.all.length == 0) {
                return null;
            }

            let index = changes.all.length - 1;
            let account = changes.all[index].account;
            let valueHex = Utils.uint8ArrayToHex(changes.all[index].value);
            let value = BigInt.fromString(valueHex, 16).toUInt32();
            return new State(account, value);
        }

        public diff(): u32 {
            let changes = Context.getStateChanges(this.addr, this.variable, this.prefix);
            if (changes.all.length < 2) {
                return 0;
            }

            let beforeHex = Utils.uint8ArrayToHex(changes.all[0].value);
            let before = BigInt.fromString(beforeHex, 16).toUInt32();

            let afterHex = Utils.uint8ArrayToHex(changes.all[changes.all.length - 1].value);
            let after = BigInt.fromString(beforeHex, 16).toUInt32();

            return after - before;
        }

        addr: string;
        variable: string;
        prefix: Uint8Array;

        constructor(addr: string, variable: string = "", prefix: Uint8Array = new Uint8Array(0)) {
            this.addr = addr;
            this.variable = variable;
            this.prefix = prefix;
        }
    }
}


