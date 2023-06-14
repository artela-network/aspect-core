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

    getBeforeFunc(typeTag: string, paramPrefix: string, valueFunc: string): string {
        const param1 : string = typeTag;
        const param2 : string = paramPrefix;
        const param3 : string = valueFunc;
        let message: string = 
    `public before(): State<${param1}> | null {
      let changes = Context.getStateChanges(this.addr, ${param2}, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let account = changes.all[0].account;
      let value = Utils.uint8ArrayTo${param3}(changes.all[0].value);
      return new State(account, value);
    }\n`;
        return message;
    }

    getChangesFunc(typeTag: string, paramPrefix: string, valueFunc: string): string {
        const param1 : string = typeTag;
        const param2 : string = paramPrefix;
        const param3 : string = valueFunc;
        let message: string = 
    `public changes(): Array<State<${param1}>> | null {
      let changes = Context.getStateChanges(this.addr, ${param2}, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let res = new Array<State<${param1}>>(changes.all.length);
      for (let i = 0; i < changes.all.length; i++) {
          let account = changes.all[i].account;
          let value = Utils.uint8ArrayTo${param3}(changes.all[0].value);
          res[i] = new State(account, value)
      }
      return res;
    }\n`;
        return message;
    }

    getLatestFunc(typeTag: string, paramPrefix: string, valueFunc: string): string {
        const param1 : string = typeTag;
        const param2 : string = paramPrefix;
        const param3 : string = valueFunc;
        let message: string = 
    `public latest(): State<${param1}> | null {
      let changes = Context.getStateChanges(this.addr, ${param2}, this.prefix);
      if (changes.all.length == 0) {
          return null;
      }

      let index = changes.all.length - 1;
      let account = changes.all[index].account;
      let value = Utils.uint8ArrayTo${param3}(changes.all[0].value);
      return new State(account, value);
    }\n`;
        return message;
    }

    getDiffFunc(typeTag: string, paramPrefix: string, valueFunc: string): string {
        const param1 : string = typeTag;
        const param2 : string = paramPrefix;
        const param3 : string = valueFunc;
        let message: string = 
    `public diff(): ${param1}  | null {
      let changes = Context.getStateChanges(this.addr, ${param2}, this.prefix);
      if (changes.all.length < 2) {
          return null;
      }

      let before = Utils.uint8ArrayTo${param3}(changes.all[0].value);
      let after = Utils.uint8ArrayTo${param3}(changes.all[changes.all.length - 1].value);
      return after - before;
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
}