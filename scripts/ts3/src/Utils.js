"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.handleMapping = exports.handleStruct = exports.handleBasic = exports.isNumber = exports.getValueFunc = exports.getStructName = exports.getParamPrefix = exports.getTypeTag = exports.getMapFirstParam = exports.getMapSecondParam = exports.getStrBetLastCommaAndParen = exports.getStrAfterLastColon = exports.getStrBetweenColon = exports.isStringEmpty = void 0;
function isStringEmpty(str) {
    return !str.trim();
}
exports.isStringEmpty = isStringEmpty;
function getStrBetweenColon(str) {
    var startIndex = str.indexOf('(');
    var endIndex = str.indexOf(')');
    if (startIndex === -1 || endIndex === -1 || startIndex >= endIndex) {
        return '';
    }
    return str.substring(startIndex + 1, endIndex);
}
exports.getStrBetweenColon = getStrBetweenColon;
function getStrAfterLastColon(input) {
    var lastColonIndex = input.lastIndexOf(":");
    if (lastColonIndex === -1) {
        return "";
    }
    return input.slice(lastColonIndex + 1).trim();
}
exports.getStrAfterLastColon = getStrAfterLastColon;
function getStrBetLastCommaAndParen(input) {
    if (!input.startsWith("t_mapping")) {
        return input;
    }
    var lastCommaIndex = input.lastIndexOf(',');
    var lastParenthesisIndex = input.lastIndexOf(')');
    if (lastCommaIndex === -1 || lastParenthesisIndex === -1 || lastCommaIndex >= lastParenthesisIndex) {
        return "";
    }
    return input.slice(lastCommaIndex + 1, lastParenthesisIndex).trim();
}
exports.getStrBetLastCommaAndParen = getStrBetLastCommaAndParen;
function getMapSecondParam(input) {
    if (!input.startsWith("t_mapping")) {
        return input;
    }
    var i = input.indexOf(',');
    var j = input.lastIndexOf(')');
    if (i === -1 || j === -1 || i >= j) {
        return "";
    }
    return input.slice(i + 1, j).trim();
}
exports.getMapSecondParam = getMapSecondParam;
function getMapFirstParam(input) {
    if (!input.startsWith("t_mapping")) {
        return input;
    }
    var i = input.lastIndexOf('(');
    var j = input.lastIndexOf(',');
    if (i === -1 || j === -1 || i >= j) {
        return "";
    }
    return input.slice(i + 1, j).trim();
}
exports.getMapFirstParam = getMapFirstParam;
function getTypeTag(itemType) {
    var paramType = getStrBetLastCommaAndParen(itemType);
    switch (paramType) {
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
        case "t_address":
            return "ethereum.Address";
        default:
            return "";
    }
}
exports.getTypeTag = getTypeTag;
function getParamPrefix(item) {
    var contractName = getStrAfterLastColon(item.contract);
    if (isStringEmpty(contractName))
        return "";
    return contractName + "." + item.label;
}
exports.getParamPrefix = getParamPrefix;
function getStructName(typeStr) {
    if (typeStr.startsWith("t_mapping")) {
        return typeStr;
    }
    var paramType = getStrBetLastCommaAndParen(typeStr);
    if (paramType.startsWith("t_struct")) {
        var structName = getStrBetweenColon(paramType);
        return structName;
    }
    return "";
}
exports.getStructName = getStructName;
function getValueFunc(itemType) {
    var paramType = getStrBetLastCommaAndParen(itemType);
    switch (paramType) {
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
        case "t_address":
            return "Address";
        default:
            return "";
    }
}
exports.getValueFunc = getValueFunc;
function isNumber(itemType) {
    var paramType = getStrBetLastCommaAndParen(itemType);
    switch (paramType) {
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
        default:
            return false;
    }
}
exports.isNumber = isNumber;
function handleBasic(className, item, tracer, isStruct, n) {
    // 1 append class start
    tracer.append(tracer.getClass(className), 1 + n);
    // 2 append addr and prefix
    if (isStruct) {
        tracer.append(tracer.argsTemplageStruct, 2 + n);
    }
    else {
        tracer.append(tracer.argsTemplage, 2 + n);
    }
    // 3 append constructor
    if (isStruct) {
        tracer.append(tracer.constructorTemplateStruct, 2 + n);
    }
    else {
        tracer.append(tracer.constructorTemplate, 2 + n);
    }
    // 4 append before func
    tracer.append(tracer.getBeforeFunc(getTypeTag(item.type), getParamPrefix(item), getValueFunc(item.type), isStruct, isNumber(item.type)), 2 + n);
    // 5 append changes func
    tracer.append(tracer.getChangesFunc(getTypeTag(item.type), getParamPrefix(item), getValueFunc(item.type), isStruct, isNumber(item.type)), 2 + n);
    // 6 append lastest func
    tracer.append(tracer.getLatestFunc(getTypeTag(item.type), getParamPrefix(item), getValueFunc(item.type), isStruct, isNumber(item.type)), 2 + n);
    // 7 append diff func (only for number type)
    if (isNumber(item.type)) {
        tracer.append(tracer.getDiffFunc(getTypeTag(item.type), getParamPrefix(item), getValueFunc(item.type), isStruct, isNumber(item.type)), 2 + n);
    }
    // 1' append class end
    tracer.append(tracer.endBracket, 1 + n);
}
exports.handleBasic = handleBasic;
function handleStruct(item, tracer, structName, members) {
    // 1 append class start
    tracer.append(tracer.getClass(structName), 1);
    // 2 append addr and variable and prefix
    tracer.append(tracer.argsTemplageStruct, 2);
    // 3 append constructor
    tracer.append(tracer.constructorTemplateStruct, 2);
    // 4 handle params
    members.forEach(function (item) {
        tracer.append(tracer.getStructParam(item.label, structName + "." + item.label), 2);
    });
    // 1' append class end
    tracer.append(tracer.endBracket, 1);
    // 5 handle struct params to class
    tracer.append("export namespace ".concat(structName, " {"), 1);
    members.forEach(function (item) {
        handleBasic(item.label, item, tracer, true, 1);
    });
    tracer.append(tracer.endBracket, 1);
}
exports.handleStruct = handleStruct;
function handleMapping(item, tracer, structNameSet, obj) {
    var firstParamType = getMapFirstParam(item.type);
    var secondParamType = getMapSecondParam(item.type);
    var ft = getTypeTag(firstParamType);
    var ff = getValueFunc(firstParamType);
    if (secondParamType.startsWith("t_mapping")) {
        firstParamType = getMapFirstParam(secondParamType);
        secondParamType = getMapSecondParam(secondParamType);
        ft = getTypeTag(firstParamType);
        ff = getValueFunc(firstParamType);
        var prefix = getStrAfterLastColon(item.contract) + "." + item.label;
        tracer.append(tracer.getClass(item.label), 1);
        tracer.append(tracer.argsTemplage, 2);
        tracer.append(tracer.constructorTemplate, 2);
        tracer.append(tracer.getNestedMappingValue(ft, ff, item.label, prefix), 2);
        tracer.append(tracer.endBracket, 1);
        tracer.append("export namespace ".concat(item.label, " {\n"), 1);
        tracer.append("export class Value {\n", 2);
        tracer.append(tracer.argsTemplageStruct, 2);
        tracer.append(tracer.constructorTemplateStruct, 2);
        tracer.append(tracer.getBeforeFuncMap(ft, ff, getTypeTag(secondParamType), "this.variable", getValueFunc(secondParamType), isNumber(secondParamType)), 2);
        tracer.append(tracer.getChangesFuncMap(ft, ff, getTypeTag(secondParamType), "this.variable", getValueFunc(secondParamType), isNumber(secondParamType)), 2);
        tracer.append(tracer.getLatestFuncMap(ft, ff, getTypeTag(secondParamType), "this.variable", getValueFunc(secondParamType), isNumber(secondParamType)), 2);
        if (isNumber(secondParamType)) {
            tracer.append(tracer.getDiffFuncMap(ft, ff, getTypeTag(secondParamType), "this.variable", getValueFunc(secondParamType), isNumber(secondParamType)), 2);
        }
        tracer.append(tracer.endBracket, 2);
        tracer.append(tracer.endBracket, 1);
        return;
    }
    // 1 append class start
    tracer.append(tracer.getClass(item.label), 1);
    // 2 append addr and prefix
    tracer.append(tracer.argsTemplage, 2);
    // 3 append constructor
    tracer.append(tracer.constructorTemplate, 2);
    // 4 handle map second param
    var secondParamIsStruct = false;
    if (secondParamType.startsWith("t_struct")) {
        secondParamIsStruct = true;
    }
    if (secondParamIsStruct) {
        var structName = getStructName(secondParamType);
        var prefix = getStrAfterLastColon(item.contract) + "." + item.label;
        tracer.append(tracer.getMappintSecondParam(structName.toLowerCase(), structName, prefix), 2);
        // if struct has not been hadle
        if (!structNameSet.has(structName)) {
            var members = obj.types[getStrBetLastCommaAndParen(item.type)].members;
            structNameSet.add(structName);
            handleStruct(item, tracer, structName, members);
        }
    }
    else {
        // 4.1 append before func
        tracer.append(tracer.getBeforeFuncMap(ft, ff, getTypeTag(secondParamType), "\"" + getParamPrefix(item) + "\"", getValueFunc(secondParamType), isNumber(secondParamType)), 2);
        // 4.2 append changes func
        tracer.append(tracer.getChangesFuncMap(ft, ff, getTypeTag(secondParamType), "\"" + getParamPrefix(item) + "\"", getValueFunc(secondParamType), isNumber(secondParamType)), 2);
        // 4.3 append lastest func
        tracer.append(tracer.getLatestFuncMap(ft, ff, getTypeTag(secondParamType), "\"" + getParamPrefix(item) + "\"", getValueFunc(secondParamType), isNumber(secondParamType)), 2);
        // 4.4 append diff func (only for number type)
        if (isNumber(secondParamType)) {
            tracer.append(tracer.getDiffFuncMap(ft, ff, getTypeTag(secondParamType), "\"" + getParamPrefix(item) + "\"", getValueFunc(secondParamType), isNumber(secondParamType)), 2);
        }
    }
    // 1' append class end
    tracer.append(tracer.endBracket, 1);
}
exports.handleMapping = handleMapping;
