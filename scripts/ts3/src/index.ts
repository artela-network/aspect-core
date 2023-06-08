import Generator from './Generator';
import {StorageItem} from './Generator';

function isStringEmpty(str: string): boolean {
    return !str.trim();
  }
  
function getStrAfterLastColon(input: string): string {
    const lastColonIndex = input.lastIndexOf(":");
    if (lastColonIndex === -1) {
      return "";
    }
    return input.slice(lastColonIndex + 1).trim();
}

function getStrBetLastCommaAndParen(input: string): string {
    if (!input.startsWith("t_mapping")) {
        return input;
    }
    const lastCommaIndex = input.lastIndexOf(',');
    const lastParenthesisIndex = input.lastIndexOf(')');
    
    if (lastCommaIndex === -1 || lastParenthesisIndex === -1 || lastCommaIndex >= lastParenthesisIndex) {
      return "";
    }
    
    return input.slice(lastCommaIndex + 1, lastParenthesisIndex).trim();
  }  

function getTypeTag(item: StorageItem): string {
    const paramType = getStrBetLastCommaAndParen(item.type);
    switch(paramType) {
        case "t_int32":
            return "i32";
        case "t_uint64":
            return "u64";
        case "t_uint256":
            return "BigInt";
        case "t_string":
            return "string";
        case "t_bool":
            return "bool";
        default:
            return "";
    }
}

function getParamPrefix(item: StorageItem): string {
    let contractName = getStrAfterLastColon(item.contract);
    if (isStringEmpty(contractName))
      return "";
    return contractName + "." + item.label;
}

function getValueFunc(item: StorageItem): string {
    const paramType = getStrBetLastCommaAndParen(item.type);
    switch(paramType) {
        case "t_uint32":
            return "Int32";
        case "t_uint256":
            return "Int256";
        case "t_string":
            return "String";
        case "t_bool":
            return "Bool";
        default:
            return "";
    }
}

const tracer: Generator = new Generator(
    "/Users/yuanyuan/go/src/github.com/artela-network/artelasdk/scripts/ts3/src/LayoutHoneyPot.json",
    "/Users/yuanyuan/go/src/github.com/artela-network/artelasdk/scripts/ts3/src/HoneyPot.ts");

const obj = tracer.getLayout();
const items = obj.storage;

// 1. append reference
tracer.append(tracer.refLib, 0);
// 2.1 append namespace start
tracer.append(tracer.getNameSpace(getStrAfterLastColon(items[0].contract)), 0);


// -----Loop to handle multi params start------
items.forEach(function (item) {
    // 3.1 append class start
    tracer.append(tracer.getClass(item.label), 1);
    // 4.1 append addr and prefix
    tracer.append(tracer.argsTemplage ,2);
    // 4.2 append constructor
    tracer.append(tracer.constructorTemplate ,2);
    // 4.3 append before func
    tracer.append(tracer.getBeforeFunc(getTypeTag(item), 
        getParamPrefix(item), getValueFunc(item)) ,2);
    // 4.4 append changes func
    tracer.append(tracer.getChangesFunc(getTypeTag(item), 
        getParamPrefix(item), getValueFunc(item)) ,2);
    // 4.5 append lastest func
    tracer.append(tracer.getLatestFunc(getTypeTag(item), 
        getParamPrefix(item), getValueFunc(item)) ,2);
    // 4.6 append diff func
    tracer.append(tracer.getDiffFunc(getTypeTag(item), 
        getParamPrefix(item), getValueFunc(item)) ,2);
    // 3.2 append class end
    tracer.append(tracer.endBracket, 1); 
});
// -----Loop to handle multi params end------



// 2.2 append namespace end
tracer.append(tracer.endBracket, 0);