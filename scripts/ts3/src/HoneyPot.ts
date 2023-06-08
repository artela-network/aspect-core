import { Protobuf } from 'as-proto/assembly';
import { Context, Pair, Abi } from './lib/index';
export namespace HoneyPot {
  export class dummy1 {
    addr: string;
    constructor(addr: string) {
        this.addr = addr;
    }
    public before(): i256 {
      let changes = Context.getStateChanges(this.addr, "HoneyPot.dummy1", "");
      if (changes.all.length == 0) {
        return 0;
      }
      
      let value = changes.all[0].value;
      return Abi.asInt256(value);
    }
    public changes(): Array<Pair<i256>> {
      let changes = Context.getStateChanges(this.addr, "HoneyPot.dummy1", "");
      if (changes.all.length < 2) {
        return new Array<Pair<i256>>(0);
      }
      
      let res = new Array<Pair<i256>>(changes.all.length - 1);
      for (let i: i32 = 1; i < changes.all.length; i++) {
        let parsedValue = Abi.asInt256(changes.all[i].value);
        res[i - 1] = new Pair(changes.all[i].account, parsedValue)
      }
      return res;
    }
    public lastest(): i256 {
      let changes = Context.getStateChanges(this.addr, "HoneyPot.dummy1", "");
      if (changes.all.length == 0) {
        return 0;
      }
      
      let value = changes.all[changes.all.length - 1].value;
      return Abi.asInt256(value);
    }
    public diff(): i256 {
      let changes = Context.getStateChanges(this.addr, "HoneyPot.dummy1", "");
      if (changes.all.length < 2) {
        return 0;
      }
      
      let before = Abi.asInt256(changes.all.values[0]);
      let end = Abi.asInt256(changes.all.values[changes.all.length - 1]);
      return end - before;
    }
  }
  export class balances {
    addr: string;
    constructor(addr: string) {
        this.addr = addr;
    }
    public before(key: string): i256 {
      let changes = Context.getStateChanges(this.addr, "HoneyPot.balances", key);
      if (changes.all.length == 0) {
        return 0;
      }
      
      let value = changes.all[0].value;
      return Abi.asInt256(value);
    }
    public changes(key: string): Array<Pair<i256>> {
      let changes = Context.getStateChanges(this.addr, "HoneyPot.balances", key);
      if (changes.all.length < 2) {
        return new Array<Pair<i256>>(0);
      }
      
      let res = new Array<Pair<i256>>(changes.all.length - 1);
      for (let i: i32 = 1; i < changes.all.length; i++) {
        let parsedValue = Abi.asInt256(changes.all[i].value);
        res[i - 1] = new Pair(changes.all[i].account, parsedValue)
      }
      return res;
    }
    public lastest(key: string): i256 {
      let changes = Context.getStateChanges(this.addr, "HoneyPot.balances", key);
      if (changes.all.length == 0) {
        return 0;
      }
      
      let value = changes.all[changes.all.length - 1].value;
      return Abi.asInt256(value);
    }
    public diff(key: string): i256 {
      let changes = Context.getStateChanges(this.addr, "HoneyPot.balances", key);
      if (changes.all.length < 2) {
        return 0;
      }
      
      let before = Abi.asInt256(changes.all.values[0]);
      let end = Abi.asInt256(changes.all.values[changes.all.length - 1]);
      return end - before;
    }
  }
}
