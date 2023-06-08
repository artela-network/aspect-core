import * as fs from 'fs';

export type StorageItem = {
    astId: number;
    contract: string;
    label: string;
    offset: number;
    slot: string;
    type: string;
}

type Layout = {
    storage:  StorageItem[];
}

export default class Generator {
    private layoutPath: string;
    private tsPath: string;

    public refPro = "import { Protobuf } from 'as-proto/assembly';\n";
    public refLib = "import { Context, Pair, Abi } from './lib/index';\n";

    public endBracket  = "}\n";
    public addrTemplate = "addr: string;\n";
    public constructorTemplate = 
    "constructor(addr: string) {\n\
        this.addr = addr;\n\
    }\n";    

    constructor(layoutPath: string, tsPath: string) {
        this.layoutPath = layoutPath;
        this.tsPath = tsPath;
    }

    getLayout(): Layout {
        if(fs.existsSync(this.layoutPath))
        {
            const loadJson = fs.readFileSync(this.layoutPath, "utf-8");
            const storageLayout = JSON.parse(loadJson) as Layout;
            return storageLayout;
        }
        return <Layout>{};
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

    getBeforeFunc(isMap: boolean, typeTag: string, paramPrefix: string, abiTag: string): string {
        const param1 : string = typeTag;
        const param2 : string = paramPrefix;
        const param3 : string = abiTag;
        let message: string = 
    `public before(): ${param1} {\n\
      let changes = Context.getStateChanges(this.addr, "${param2}", "");\n\
      if (changes.all.length == 0) {\n\
        return 0;\n\
      }\n\
      \n\
      let value = changes.all[0].value;\n\
      return Abi.as${param3}(value);\n\
    }\n`;
        let msgMap: string = 
    `public before(key: string): ${param1} {\n\
      let changes = Context.getStateChanges(this.addr, "${param2}", key);\n\
      if (changes.all.length == 0) {\n\
        return 0;\n\
      }\n\
      \n\
      let value = changes.all[0].value;\n\
      return Abi.as${param3}(value);\n\
    }\n`;
        if (isMap) {
          return msgMap;
        }
        return message;
    }

    getChangesFunc(isMap: boolean, typeTag: string, paramPrefix: string, abiTag: string): string {
        const param1 : string = typeTag;
        const param2 : string = paramPrefix;
        const param3 : string = abiTag;
        let message: string = 
    `public changes(): Array<Pair<${param1}>> {\n\
      let changes = Context.getStateChanges(this.addr, "${param2}", "");\n\
      if (changes.all.length < 2) {\n\
        return new Array<Pair<${param1}>>(0);\n\
      }\n\
      \n\
      let res = new Array<Pair<${param1}>>(changes.all.length - 1);\n\
      for (let i: i32 = 1; i < changes.all.length; i++) {\n\
        let parsedValue = Abi.as${param3}(changes.all[i].value);\n\
        res[i - 1] = new Pair(changes.all[i].account, parsedValue)\n\
      }\n\
      return res;\n\
    }\n`;
      let msgMap: string = 
    `public changes(key: string): Array<Pair<${param1}>> {\n\
      let changes = Context.getStateChanges(this.addr, "${param2}", key);\n\
      if (changes.all.length < 2) {\n\
        return new Array<Pair<${param1}>>(0);\n\
      }\n\
      \n\
      let res = new Array<Pair<${param1}>>(changes.all.length - 1);\n\
      for (let i: i32 = 1; i < changes.all.length; i++) {\n\
        let parsedValue = Abi.as${param3}(changes.all[i].value);\n\
        res[i - 1] = new Pair(changes.all[i].account, parsedValue)\n\
      }\n\
      return res;\n\
    }\n`;
        if (isMap) {
          return msgMap;
        }
        return message;
    }

    getLatestFunc(isMap: boolean, typeTag: string, paramPrefix: string, abiTag: string): string {
        const param1 : string = typeTag;
        const param2 : string = paramPrefix;
        const param3 : string = abiTag;
        let message: string = 
    `public lastest(): ${param1} {\n\
      let changes = Context.getStateChanges(this.addr, "${param2}", "");\n\
      if (changes.all.length == 0) {\n\
        return 0;\n\
      }\n\
      \n\
      let value = changes.all[changes.all.length - 1].value;\n\
      return Abi.as${param3}(value);\n\
    }\n`;
        let msgMap: string = 
    `public lastest(key: string): ${param1} {\n\
      let changes = Context.getStateChanges(this.addr, "${param2}", key);\n\
      if (changes.all.length == 0) {\n\
        return 0;\n\
      }\n\
      \n\
      let value = changes.all[changes.all.length - 1].value;\n\
      return Abi.as${param3}(value);\n\
    }\n`;
        if (isMap) {
          return msgMap;
        }
        return message;
    }

    getDiffFunc(isMap: boolean, typeTag: string, paramPrefix: string, abiTag: string): string {
        const param1 : string = typeTag;
        const param2 : string = paramPrefix;
        const param3 : string = abiTag;
        let message: string = 
    `public diff(): ${param1} {\n\
      let changes = Context.getStateChanges(this.addr, "${param2}", "");\n\
      if (changes.all.length < 2) {\n\
        return 0;\n\
      }\n\
      \n\
      let before = Abi.as${param3}(changes.all.values[0]);\n\
      let end = Abi.as${param3}(changes.all.values[changes.all.length - 1]);\n\
      return end - before;\n\
    }\n`;
      let msgMap: string = 
    `public diff(key: string): ${param1} {\n\
      let changes = Context.getStateChanges(this.addr, "${param2}", key);\n\
      if (changes.all.length < 2) {\n\
        return 0;\n\
      }\n\
      \n\
      let before = Abi.as${param3}(changes.all.values[0]);\n\
      let end = Abi.as${param3}(changes.all.values[changes.all.length - 1]);\n\
      return end - before;\n\
    }\n`;
        if (isMap) {
          return msgMap;
        }
        return message;
    }
}