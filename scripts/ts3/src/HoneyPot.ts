import { Protobuf } from 'as-proto/assembly';
import { Context, State, Abi, Utils, TypeValue } from "./lib/index"
export namespace Storage {
  export class number1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<BigInt> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayToInt256(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<BigInt>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<BigInt>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayToInt256(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<BigInt> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayToInt256(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): BigInt  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayToInt256(changes.all[0].value);
      let after = Utils.uint8ArrayToInt256(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class  {
    addr: string;
    variable: string;
    prefix: Uint8Array;
    constructor(addr: string, varibale: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.variable = varibale;
      this.prefix = prefix;
    }
  }
    public before(): State<BigInt> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayToInt256(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<BigInt>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<BigInt>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayToInt256(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<BigInt> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayToInt256(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): BigInt  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayToInt256(changes.all[0].value);
      let after = Utils.uint8ArrayToInt256(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  export class number2 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<i32> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<i32>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<i32>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<i32> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): i32  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class  {
    addr: string;
    variable: string;
    prefix: Uint8Array;
    constructor(addr: string, varibale: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.variable = varibale;
      this.prefix = prefix;
    }
  }
    public before(): State<i32> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<i32>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<i32>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<i32> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): i32  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  export class number3 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<u64> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<u64>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<u64>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<u64> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): u64  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class  {
    addr: string;
    variable: string;
    prefix: Uint8Array;
    constructor(addr: string, varibale: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.variable = varibale;
      this.prefix = prefix;
    }
  }
    public before(): State<u64> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<u64>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<u64>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<u64> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): u64  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  export class str1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<string> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<string>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<string>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<string> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): string  | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class  {
    addr: string;
    variable: string;
    prefix: Uint8Array;
    constructor(addr: string, varibale: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.variable = varibale;
      this.prefix = prefix;
    }
  }
    public before(): State<string> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<string>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<string>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<string> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): string  | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  export class bool1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<bool> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayToBool(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<bool>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<bool>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayToBool(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<bool> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayToBool(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): bool  | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayToBool(changes.all[0].value);
      let after = Utils.uint8ArrayToBool(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class  {
    addr: string;
    variable: string;
    prefix: Uint8Array;
    constructor(addr: string, varibale: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.variable = varibale;
      this.prefix = prefix;
    }
  }
    public before(): State<bool> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayToBool(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<bool>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<bool>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayToBool(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<bool> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayToBool(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): bool  | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayToBool(changes.all[0].value);
      let after = Utils.uint8ArrayToBool(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  export class _struct(Person)17_storag {
    addr: string;
    variable: string;
    prefix: Uint8Array;
    constructor(addr: string, varibale: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.variable = varibale;
      this.prefix = prefix;
    }
  }
    public before(): State<> | null {
      let changes = Context.getStateChanges(this.addr, Storage.p1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.p1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<> | null {
      let changes = Context.getStateChanges(this.addr, Storage.p1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff():   | null {
      let changes = Context.getStateChanges(this.addr, Storage.p1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  export class _struct(Person)17_storag {
    addr: string;
    variable: string;
    prefix: Uint8Array;
    constructor(addr: string, varibale: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.variable = varibale;
      this.prefix = prefix;
    }
  }
    public before(): State<> | null {
      let changes = Context.getStateChanges(this.addr, Storage.accounts, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.accounts, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<> | null {
      let changes = Context.getStateChanges(this.addr, Storage.accounts, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff():   | null {
      let changes = Context.getStateChanges(this.addr, Storage.accounts, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
}
import { Protobuf } from 'as-proto/assembly';
    import { Context, State, Abi, Utils, TypeValue } from "./lib/index";
export namespace Storage {
  export class number1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<BigInt> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayToInt256(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<BigInt>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<BigInt>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayToInt256(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<BigInt> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayToInt256(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): BigInt  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayToInt256(changes.all[0].value);
      let after = Utils.uint8ArrayToInt256(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class number2 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<i32> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<i32>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<i32>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<i32> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): i32  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class number3 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<u64> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<u64>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<u64>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<u64> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): u64  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class str1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<string> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<string>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<string>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<string> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): string  | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class bool1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<bool> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayToBool(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<bool>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<bool>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayToBool(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<bool> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayToBool(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): bool  | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayToBool(changes.all[0].value);
      let after = Utils.uint8ArrayToBool(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
import { Protobuf } from 'as-proto/assembly';
    import { Context, State, Abi, Utils, TypeValue } from "./lib/index";
export namespace Storage {
  export class number1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<BigInt> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayToInt256(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<BigInt>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<BigInt>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayToInt256(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<BigInt> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayToInt256(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): BigInt  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayToInt256(changes.all[0].value);
      let after = Utils.uint8ArrayToInt256(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class number2 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<i32> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<i32>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<i32>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<i32> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): i32  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class number3 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<u64> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<u64>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<u64>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<u64> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): u64  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class str1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<string> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<string>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<string>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<string> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): string  | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class bool1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<bool> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayToBool(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<bool>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<bool>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayToBool(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<bool> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayToBool(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): bool  | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayToBool(changes.all[0].value);
      let after = Utils.uint8ArrayToBool(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
import { Protobuf } from 'as-proto/assembly';
    import { Context, State, Abi, Utils, TypeValue } from "./lib/index";
export namespace Storage {
  export class number1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<BigInt> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayToInt256(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<BigInt>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<BigInt>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayToInt256(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<BigInt> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayToInt256(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): BigInt  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayToInt256(changes.all[0].value);
      let after = Utils.uint8ArrayToInt256(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class number2 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<i32> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<i32>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<i32>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<i32> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): i32  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class number3 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<u64> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<u64>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<u64>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<u64> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): u64  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class str1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<string> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<string>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<string>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<string> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): string  | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class bool1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<bool> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayToBool(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<bool>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<bool>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayToBool(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<bool> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayToBool(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): bool  | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayToBool(changes.all[0].value);
      let after = Utils.uint8ArrayToBool(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
import { Protobuf } from 'as-proto/assembly';
    import { Context, State, Abi, Utils, TypeValue } from "./lib/index";
export namespace Storage {
  export class number1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<BigInt> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayToInt256(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<BigInt>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<BigInt>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayToInt256(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<BigInt> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayToInt256(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): BigInt  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayToInt256(changes.all[0].value);
      let after = Utils.uint8ArrayToInt256(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class number2 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<i32> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<i32>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<i32>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<i32> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): i32  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class number3 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<u64> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<u64>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<u64>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<u64> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): u64  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class str1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<string> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<string>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<string>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<string> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): string  | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class bool1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<bool> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayToBool(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<bool>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<bool>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayToBool(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<bool> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayToBool(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): bool  | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayToBool(changes.all[0].value);
      let after = Utils.uint8ArrayToBool(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
import { Protobuf } from 'as-proto/assembly';
    import { Context, State, Abi, Utils, TypeValue } from "./lib/index";
export namespace Storage {
  export class number1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<BigInt> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayToInt256(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<BigInt>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<BigInt>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayToInt256(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<BigInt> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayToInt256(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): BigInt  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayToInt256(changes.all[0].value);
      let after = Utils.uint8ArrayToInt256(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class number2 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<i32> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<i32>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<i32>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<i32> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): i32  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number2, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class number3 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<u64> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<u64>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<u64>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<u64> | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): u64  | null {
      let changes = Context.getStateChanges(this.addr, Storage.number3, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class str1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<string> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<string>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<string>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<string> | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): string  | null {
      let changes = Context.getStateChanges(this.addr, Storage.str1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo(changes.all[0].value);
      let after = Utils.uint8ArrayTo(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
  export class bool1 {
    addr: string;
    prefix: Uint8Array;
    constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }
    public before(): State<bool> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayToBool(changes.all[0].value);
      return new State(account, value);
    }
    public changes(): Array<State<bool>> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<bool>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayToBool(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }
    public latest(): State<bool> | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayToBool(changes.all[0].value);
      return new State(account, value);
    }
    public diff(): bool  | null {
      let changes = Context.getStateChanges(this.addr, Storage.bool1, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayToBool(changes.all[0].value);
      let after = Utils.uint8ArrayToBool(changes.all[changes.all.length - 1].value);
      return after - before;
    }
  }
