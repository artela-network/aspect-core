import * as fs from 'fs';

export type StorageItem = {
    astId: number;
    contract: string;
    label: string;
    offset: number;
    slot: string;
    type: string;
}

export type StorageType = {
  encoding: string;
  label: string;
  numberOfBytes: string;
  members?: StorageItem[];
  key?: string;
  value?: string;
}

export type StorageLayout = {
  storage: StorageItem[];
  types: Record<string, StorageType>;
}
export default class Generator {
    private layoutPath: string;
    private tsPath: string;

    public refLib = `import { Protobuf } from 'as-proto/assembly';
import { Context, State, Abi, Utils, TypeValue } from "./lib/index";\n`;

    public endBracket  = "}\n";
    public argsTemplage = `addr: string;
    prefix: Uint8Array;\n`;
    public argsTemplageStruct = `addr: string;
    variable: string;
    prefix: Uint8Array;\n`;
    public constructorTemplate = 
    `constructor(addr: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.prefix = prefix;
    }\n`;    
    public constructorTemplateStruct = 
    `constructor(addr: string, varibale: string, prefix: Uint8Array = new Uint8Array(0)) {
      this.addr = addr;
      this.variable = varibale;
      this.prefix = prefix;
    }\n`;    

    constructor(layoutPath: string, tsPath: string) {
        this.layoutPath = layoutPath;
        this.tsPath = tsPath;
    }

    getStorage(loadJson: string): StorageLayout {
        const storageLayout = JSON.parse(loadJson) as StorageLayout;
        return storageLayout;
    }
    getLayoutJson(): string {
      if(fs.existsSync(this.layoutPath))
      {
          const loadJson = fs.readFileSync(this.layoutPath, "utf-8");
          return loadJson;
      }
      return "";
    }

    append(str: string, space: number): boolean {
        if (space > 0) {
            fs.writeFileSync(this.tsPath, "  ".repeat(space), {flag:'a'});
        }
        fs.writeFileSync(this.tsPath, str, {flag:'a'});
        return true;
    }

    getNameSpace(contract: string): string {
      const contractName: string = contract;
      return `export namespace ${contractName} {\n`;
    }

    getClass(arg: string): string {
      const argName: string = arg;
      return `export class ${argName} {\n`;
    }

    getBeforeFunc(typeTag: string, paramPrefix: string, valueFunc: string, 
      isStruct: boolean, isNumber: boolean): string {
        const param1 : string = typeTag;
        let param2 : string = "\""+paramPrefix+"\"";
        if (isStruct) {
          param2 = "this.variable";
        }
        const param3 : string = valueFunc;
        let param4 : string = `let value = Utils.uint8ArrayTo${param3}(changes.all[0].value);`;
        if (isNumber) {
          let param5 : string = "";
          if ("BigInt" != typeTag) {
            param5 = ".to"+valueFunc+"()";
          }
          param4 = `let valueHex = Utils.uint8ArrayToHex(changes.all[0].value);
          let value = BigInt.fromString(valueHex, 16)${param5};`;
        }
        let message: string = 
    `public before(): State<${param1}> | null {
      let changes = Context.getStateChanges(this.addr, ${param2}, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      ${param4}
      return new State(account, value);
    }\n`;
        return message;
    }

    getChangesFunc(typeTag: string, paramPrefix: string, valueFunc: string, 
      isStruct: boolean, isNumber: boolean): string {
        const param1 : string = typeTag;
        let param2 : string = "\""+paramPrefix+"\"";
        if (isStruct) {
          param2 = "this.variable";
        }
        const param3 : string = valueFunc;
        let param4 : string = `let value = Utils.uint8ArrayTo${param3}(changes.all[i].value);`;
        if (isNumber) {
          let param5 : string = "";
          if ("BigInt" != typeTag) {
            param5 = ".to"+valueFunc+"()";
          }
          param4 = `let valueHex = Utils.uint8ArrayToHex(changes.all[0].value);
          let value = BigInt.fromString(valueHex, 16)${param5};`;
        }
        let message: string = 
    `public changes(): Array<State<${param1}>> | null {
      let changes = Context.getStateChanges(this.addr, ${param2}, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<${param1}>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          ${param4}
          res[i] = new State(account, value)
      }
      return res;
    }\n`;
        return message;
    }

    getLatestFunc(typeTag: string, paramPrefix: string, valueFunc: string, 
      isStruct: boolean, isNumber: boolean): string {
        const param1 : string = typeTag;
        let param2 : string = "\""+paramPrefix+"\"";
        if (isStruct) {
          param2 = "this.variable";
        }
        const param3 : string = valueFunc;
        let param4 : string = `let value = Utils.uint8ArrayTo${param3}(changes.all[index].value);`;
        if (isNumber) {
          let param5 : string = "";
          if ("BigInt" != typeTag) {
            param5 = ".to"+valueFunc+"()";
          }
          param4 = `let valueHex = Utils.uint8ArrayToHex(changes.all[index].value);
          let value = BigInt.fromString(valueHex, 16)${param5};`;
        }
        let message: string = 
    `public latest(): State<${param1}> | null {
      let changes = Context.getStateChanges(this.addr, ${param2}, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      ${param4}
      return new State(account, value);
    }\n`;
        return message;
    }

    getDiffFunc(typeTag: string, paramPrefix: string, valueFunc: string, 
      isStruct: boolean, isNumber: boolean): string {
        const param1 : string = typeTag;
        let param2 : string = "\""+paramPrefix+"\"";
        let forNumber : string = "after - before";
        if ("BigInt" == typeTag) {
          forNumber = "after.sub(before)";
        }
        if (isStruct) {
          param2 = "this.variable";
        }
        const param3 : string = valueFunc;
        let param4 : string = `let before = Utils.uint8ArrayTo${param3}(changes.all[0].value);`;
        let param5 : string = `let after = Utils.uint8ArrayTo${param3}(changes.all[changes.all.length - 1].value);`;
        if (isNumber) {
          let param6 : string = "";
          if ("BigInt" != typeTag) {
            param6 = ".to"+valueFunc+"()";
          }
          param4 = `let beforeHex = Utils.uint8ArrayToHex(changes.all[0].value);
          let before = BigInt.fromString(beforeHex, 16)${param6};`;
          param5 = `let afterHex = Utils.uint8ArrayToHex(changes.all[changes.all.length - 1].value);
          let after = BigInt.fromString(afterHex, 16)${param6};`;
        }
        let message: string = 
    `public diff(): ${param1}  | null {
      let changes = Context.getStateChanges(this.addr, ${param2}, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      ${param4}
      ${param5}
      
      return ${forNumber};
    }\n`;
        return message;
    }

    getBeforeFuncMap(typeTag: string, paramPrefix: string, valueFunc: string, 
      isNumber: boolean): string {
      const param1 : string = typeTag;
      const param2 : string = paramPrefix;
      const param3 : string = valueFunc;
      let param4 : string = `let value = Utils.uint8ArrayTo${param3}(changes.all[0].value);`;
      if (isNumber) {
        let param5 : string = "";
        if ("BigInt" != typeTag) {
          param5 = ".to"+valueFunc+"()";
        }
        param4 = `let valueHex = Utils.uint8ArrayToHex(changes.all[0].value);
        let value = BigInt.fromString(valueHex, 16)${param5};`;
      }
      let message: string = 
  `public before(key: string): State<${param1}> | null {
    let changes = Context.getStateChanges(this.addr, "${param2}", this.prefix+key);
    if (changes.all.length == 0) {
        return null;
    }

    let account = changes.all[0].account;
    ${param4}
    return new State(account, value);
  }\n`;
      return message;
  }

  getChangesFuncMap(typeTag: string, paramPrefix: string, valueFunc: string, 
    isNumber: boolean): string {
      const param1 : string = typeTag;
      const param2 : string = paramPrefix;
      const param3 : string = valueFunc;
      let param4 : string = `let value = Utils.uint8ArrayTo${param3}(changes.all[i].value);`;
      if (isNumber) {
        let param5 : string = "";
        if ("BigInt" != typeTag) {
          param5 = ".to"+valueFunc+"()";
        }
        param4 = `let valueHex = Utils.uint8ArrayToHex(changes.all[0].value);
        let value = BigInt.fromString(valueHex, 16)${param5};`;
      }
      let message: string = 
  `public changes(key: string): Array<State<${param1}>> | null {
    let changes = Context.getStateChanges(this.addr, "${param2}", this.prefix+key);
    if (changes.all.length == 0) {
        return null;
    }

    let res = new Array<State<${param1}>>(changes.all.length);
    for (let i = 0; i < changes.all.length; i++) {
        let account = changes.all[i].account;
        ${param4}
        res[i] = new State(account, value)
    }
    return res;
  }\n`;
      return message;
  }

  getLatestFuncMap(typeTag: string, paramPrefix: string, valueFunc: string, 
    isNumber: boolean): string {
      const param1 : string = typeTag;
      const param2 : string = paramPrefix;
      const param3 : string = valueFunc;
      let param4 : string = `let value = Utils.uint8ArrayTo${param3}(changes.all[index].value);`;
      if (isNumber) {
        let param5 : string = "";
        if ("BigInt" != typeTag) {
          param5 = ".to"+valueFunc+"()";
        }
        param4 = `let valueHex = Utils.uint8ArrayToHex(changes.all[index].value);
        let value = BigInt.fromString(valueHex, 16)${param5};`;
      }
      let message: string = 
  `public latest(key: string): State<${param1}> | null {
    let changes = Context.getStateChanges(this.addr, "${param2}", this.prefix+key);
    if (changes.all.length == 0) {
        return null;
    }

    let index = changes.all.length - 1;
    let account = changes.all[index].account;
    ${param4}
    return new State(account, value);
  }\n`;
      return message;
  }

  getDiffFuncMap(typeTag: string, paramPrefix: string, valueFunc: string, 
    isNumber: boolean): string {
      const param1 : string = typeTag;
      const param2 : string = paramPrefix;
      const param3 : string = valueFunc;
      let param4 : string = `let before = Utils.uint8ArrayTo${param3}(changes.all[0].value);`;
      let param5 : string = `let after = Utils.uint8ArrayTo${param3}(changes.all[changes.all.length - 1].value);`;
      let forNumber : string = "after - before";
      if ("BigInt" == typeTag) {
        forNumber = "after.sub(before)";
      }
      if (isNumber) {
        let param6 : string = "";
        if ("BigInt" != typeTag) {
          param6 = ".to"+valueFunc+"()";
        }
        param4 = `let beforeHex = Utils.uint8ArrayToHex(changes.all[0].value);
        let before = BigInt.fromString(beforeHex, 16)${param6};`;
        param5 = `let afterHex = Utils.uint8ArrayToHex(changes.all[changes.all.length - 1].value);
        let after = BigInt.fromString(afterHex, 16)${param6};`;
      }
      let message: string = 
  `public diff(key: string): ${param1}  | null {
    let changes = Context.getStateChanges(this.addr, "${param2}", this.prefix+key);
    if (changes.all.length < 2) {
        return null;
    }

    ${param4}
    ${param5}

    return ${forNumber};
  }\n`;
      return message;
  }

    getStructParam(name: string, wrapName: string): string {
      const param1 : string = name;
      const param2 : string = wrapName;
      let message: string =
      `public ${param1}(): ${param2} {
        let encoded = Abi.encodeString("${param1}");
        return new ${param2}(this.addr, this.variable, Utils.concatUint8Arrays(this.prefix, encoded));
    }\n`;
      return message;
    }

    getMappintSecondParam(name: string, type: string, prefix: string): string {
      const param1 : string = name;
      const param2 : string = type;
      const param3 : string = prefix; //ContractName.ParamNameInContract
      let message: string =
      `public ${param1}(key: string): ${param2} {
        let encoded = Abi.encodeString(key);
        return new ${param2}(this.addr, "${param3}", Utils.concatUint8Arrays(this.prefix, encoded))
    }\n`;
      return message;
    }
}