import Generator from './Generator';
import {StorageItem} from './Generator';

function isStringEmpty(str: string): boolean {
    return !str.trim();
}
  
function getStrBetweenColon(str: string): string {
    const startIndex = str.indexOf('(');
  const endIndex = str.indexOf(')');

  if (startIndex === -1 || endIndex === -1 || startIndex >= endIndex) {
    return '';
  }

  return str.substring(startIndex + 1, endIndex);
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

function getTypeTag(itemType: string): string {
    const paramType = getStrBetLastCommaAndParen(itemType);
    switch(paramType) {
        case "t_int32":
            return "i32";
        case "t_int64":
            return "i64";
        case "t_int256":
            return "BigInt";
        case "t_uint32":
            return "u32";
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

function getStructName(typeStr: string): string {
    if (typeStr.startsWith("t_mapping")) {
        return typeStr;
    }
    const paramType = getStrBetLastCommaAndParen(typeStr);
    if (paramType.startsWith("t_struct")) {
        const structName = getStrBetweenColon(paramType);
        return structName;
    }
    return "";
}

function getValueFunc(itemType: string): string {
    const paramType = getStrBetLastCommaAndParen(itemType);
    switch(paramType) {
        case "t_int32":
            return "Int32";
        case "t_int64":
            return "Int64";
        case "t_int256":
            return "Int256";
        case "t_uint32":
            return "UInt32";
        case "t_uint64":
            return "UInt64";
        case "t_uint256":
            return "UInt256";
        case "t_string_storage":
            return "String";
        case "t_bool":
            return "Bool";
        default:
            return "";
    }
}

function isNumber(itemType: string): boolean {
    const paramType = getStrBetLastCommaAndParen(itemType);
    switch(paramType) {
        case "t_int32":
            return true;
        case "t_int64":
            return true;
        case "t_int256":
            return true;
        case "t_uint32":
            return true;
        case "t_uint64":
            return true;
        case "t_uint256":
            return true;
        case "t_string_storage":
            return false;
        case "t_bool":
            return false;
        default:
            return false;
    }
}

function handleBasic(className: string, item: StorageItem, 
    tracer: Generator, isStruct: boolean) {
    // 1 append class start
    tracer.append(tracer.getClass(className), 1);
    // 2 append addr and prefix
    if (isStruct) {
        tracer.append(tracer.argsTemplageStruct ,2);
    } else {
        tracer.append(tracer.argsTemplage ,2);
    }
    // 3 append constructor
    if (isStruct) {
        tracer.append(tracer.constructorTemplateStruct ,2);
    } else {
        tracer.append(tracer.constructorTemplate ,2);
    }
    // 4 append before func
    tracer.append(tracer.getBeforeFunc(getTypeTag(item.type), 
        getParamPrefix(item), getValueFunc(item.type), isStruct, isNumber(item.type)) ,2);
    // 5 append changes func
    tracer.append(tracer.getChangesFunc(getTypeTag(item.type), 
        getParamPrefix(item), getValueFunc(item.type), isStruct, isNumber(item.type)) ,2);
    // 6 append lastest func
    tracer.append(tracer.getLatestFunc(getTypeTag(item.type), 
        getParamPrefix(item), getValueFunc(item.type), isStruct, isNumber(item.type)) ,2);
    // 7 append diff func (only for number type)
    if (isNumber(item.type)) {
        tracer.append(tracer.getDiffFunc(getTypeTag(item.type), 
        getParamPrefix(item), getValueFunc(item.type), isStruct, isNumber(item.type)) ,2);
    }
    
    // 1' append class end
    tracer.append(tracer.endBracket, 1);
}

function handleStruct(item: StorageItem, tracer: Generator, 
    structName: string, members: StorageItem[]) {
    // 1 append class start
    tracer.append(tracer.getClass(structName), 1);
    // 2 append addr and variable and prefix
    tracer.append(tracer.argsTemplageStruct ,2);
    // 3 append constructor
    tracer.append(tracer.constructorTemplateStruct ,2);
    // 4 handle params
    members.forEach(function (item) {
        tracer.append(tracer.getStructParam(item.label, structName+"_"+item.label) ,2)
    });
    // 1' append class end
    tracer.append(tracer.endBracket, 1);

    // 5 handle struct params to class
    members.forEach(function (item) {
        handleBasic(structName+"_"+item.label, item, tracer, true);
    });
}

function handleMapping(item: StorageItem, tracer: Generator, structNameSet: Set<string>) {
    // 1 append class start
    tracer.append(tracer.getClass(item.label), 1);
    // 2 append addr and prefix
    tracer.append(tracer.argsTemplage ,2);
    // 3 append constructor
    tracer.append(tracer.constructorTemplate ,2);
    // 4 handle map second param
    let secondParamIsBasic = true;
    let secondParamType = getStrBetLastCommaAndParen(item.type);
    if (secondParamType.startsWith("t_struct")) {
        secondParamIsBasic = false;
    }
    if (!secondParamIsBasic) {
        let structName = getStructName(secondParamType);
        let prefix = getStrAfterLastColon(item.contract) + "." + item.label;
        tracer.append(tracer.getMappintSecondParam(structName.toLowerCase(), structName, prefix), 2);
        // if struct has not been hadle
        if (!structNameSet.has(structName)) {
            let members = obj.types[getStrBetLastCommaAndParen(item.type)].members as StorageItem[];
            structNameSet.add(structName);
            handleStruct(item, tracer, structName, members);
        }
    } else {
        // 4.1 append before func
        tracer.append(tracer.getBeforeFuncMap(getTypeTag(item.type), 
            getParamPrefix(item), getValueFunc(item.type), isNumber(item.type)) ,2);
        // 4.2 append changes func
        tracer.append(tracer.getChangesFuncMap(getTypeTag(item.type), 
            getParamPrefix(item), getValueFunc(item.type), isNumber(item.type)) ,2);
        // 4.3 append lastest func
        tracer.append(tracer.getLatestFuncMap(getTypeTag(item.type), 
            getParamPrefix(item), getValueFunc(item.type), isNumber(item.type)) ,2);
        // 4.4 append diff func (only for number type)
        if (isNumber(item.type)) {
            tracer.append(tracer.getDiffFuncMap(getTypeTag(item.type), 
            getParamPrefix(item), getValueFunc(item.type), isNumber(item.type)) ,2);
        }
    }
    
    // 1' append class end
    tracer.append(tracer.endBracket, 1);        
}

const tracer: Generator = new Generator(
    "/Users/yuanyuan/go/src/github.com/artela-network/artelasdk/scripts/ts3/src/LayoutHoneyPot.json",
    "/Users/yuanyuan/go/src/github.com/artela-network/artelasdk/scripts/ts3/src/HoneyPot.ts");
    
const structNameSet: Set<string> = new Set();
const jsonStr = tracer.getLayoutJson();
const obj = tracer.getStorage(jsonStr);
const items = obj.storage;

// 1. append reference
tracer.append(tracer.refLib, 0);
// 2.1 append namespace start
tracer.append(tracer.getNameSpace(getStrAfterLastColon(items[0].contract)), 0);


// ----- 3.1 Loop to handle multi params start ------
items.forEach(function (item) {
    let structName = getStructName(item.type);
    if (structName.startsWith("t_mapping")) {
        handleMapping(item, tracer, structNameSet);
    } else if(!isStringEmpty(structName)) {
        if (!structNameSet.has(structName)) {
            let members = obj.types[item.type].members as StorageItem[];
            structNameSet.add(structName);
            handleStruct(item, tracer, structName, members);
        }
    } else {
        handleBasic(item.label ,item, tracer, false);
    }
});
// ----- 3.2 Loop to handle multi params end ------



// 2.2 append namespace end
tracer.append(tracer.endBracket, 0);