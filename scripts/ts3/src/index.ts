import Generator from './Generator';
import {StorageItem} from './Generator';

function isStringEmpty(str: string): boolean {
    return !str.trim();
}
  
function getStrBetweenColon(str: string): string {
    if (str.length < 2) {
        return '';
    }

    const firstChar = str.charAt(0);
    const lastChar = str.charAt(str.length - 1);
    const startIndex = str.indexOf(firstChar) + 1;
    const endIndex = str.lastIndexOf(lastChar);

    if (startIndex >= endIndex) {
        return '';
    }

    return str.substring(startIndex, endIndex);
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
        case "t_string_storage":
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

function getStructName(item: StorageItem): string {
    const paramType = getStrBetLastCommaAndParen(item.type);
    if (paramType.startsWith("t_struct")) {
        const structName = getStrBetweenColon(paramType);
        return structName;
    }
    return "";
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

function handleBasic(item: StorageItem, tracer: Generator) {
    // 1 append class start
    tracer.append(tracer.getClass(item.label), 1);
    // 2 append addr and prefix
    tracer.append(tracer.argsTemplage ,2);
    // 3 append constructor
    tracer.append(tracer.constructorTemplate ,2);
    // 4 append before func
    tracer.append(tracer.getBeforeFunc(getTypeTag(item), 
        getParamPrefix(item), getValueFunc(item)) ,2);
    // 5 append changes func
    tracer.append(tracer.getChangesFunc(getTypeTag(item), 
        getParamPrefix(item), getValueFunc(item)) ,2);
    // 6 append lastest func
    tracer.append(tracer.getLatestFunc(getTypeTag(item), 
        getParamPrefix(item), getValueFunc(item)) ,2);
    // 7 append diff func
    tracer.append(tracer.getDiffFunc(getTypeTag(item), 
        getParamPrefix(item), getValueFunc(item)) ,2);
    // 1' append class end
    tracer.append(tracer.endBracket, 1);
}

function handleStruct(item: StorageItem, tracer: Generator, structName: string) {
    // 1 append class start
    tracer.append(tracer.getClass(structName), 1);
    // 2 append addr and variable and prefix
    tracer.append(tracer.argsTemplageStruct ,2);
    // 3 append constructor
    tracer.append(tracer.constructorTemplateStruct ,2);

    // 4 handle params

    
    // 1' append class end
    tracer.append(tracer.endBracket, 1);




    // 5 append before func
    tracer.append(tracer.getBeforeFunc(getTypeTag(item), 
    getParamPrefix(item), getValueFunc(item)) ,2);
    // 6 append changes func
    tracer.append(tracer.getChangesFunc(getTypeTag(item), 
        getParamPrefix(item), getValueFunc(item)) ,2);
    // 7 append lastest func
    tracer.append(tracer.getLatestFunc(getTypeTag(item), 
        getParamPrefix(item), getValueFunc(item)) ,2);
    // 8 append diff func
    tracer.append(tracer.getDiffFunc(getTypeTag(item),
        getParamPrefix(item), getValueFunc(item)) ,2);
}

const tracer: Generator = new Generator(
    "/Users/yuanyuan/go/src/github.com/artela-network/artelasdk/scripts/ts3/src/LayoutHoneyPot.json",
    "/Users/yuanyuan/go/src/github.com/artela-network/artelasdk/scripts/ts3/src/HoneyPot.ts");

const jsonStr = tracer.getLayoutJson();
const obj = tracer.getStorage(jsonStr);
const items = obj.storage;

// 1. append reference
tracer.append(tracer.refLib, 0);
// 2.1 append namespace start
tracer.append(tracer.getNameSpace(getStrAfterLastColon(items[0].contract)), 0);


// ----- 3.1 Loop to handle multi params start ------
items.forEach(function (item) {
    let structName = getStructName(item);
    if (isStringEmpty(structName)) {
        handleBasic(item, tracer);
    }
    handleStruct(item, tracer, structName);
});
// ----- 3.2 Loop to handle multi params end ------



// 2.2 append namespace end
tracer.append(tracer.endBracket, 0);