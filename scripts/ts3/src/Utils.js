"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.handleMapping = exports.handleStruct = exports.handleBasic = exports.isNumber = exports.getValueFunc = exports.getStructName = exports.getParamPrefix = exports.getTypeTag = exports.getStrBetLastCommaAndParen = exports.getStrAfterLastColon = exports.getStrBetweenColon = exports.isStringEmpty = void 0;
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
        case "t_string_storage":
            return false;
        case "t_bool":
            return false;
        default:
            return false;
    }
}
exports.isNumber = isNumber;
function handleBasic(className, item, tracer, isStruct) {
    // 1 append class start
    tracer.append(tracer.getClass(className), 1);
    // 2 append addr and prefix
    if (isStruct) {
        tracer.append(tracer.argsTemplageStruct, 2);
    }
    else {
        tracer.append(tracer.argsTemplage, 2);
    }
    // 3 append constructor
    if (isStruct) {
        tracer.append(tracer.constructorTemplateStruct, 2);
    }
    else {
        tracer.append(tracer.constructorTemplate, 2);
    }
    // 4 append before func
    tracer.append(tracer.getBeforeFunc(getTypeTag(item.type), getParamPrefix(item), getValueFunc(item.type), isStruct, isNumber(item.type)), 2);
    // 5 append changes func
    tracer.append(tracer.getChangesFunc(getTypeTag(item.type), getParamPrefix(item), getValueFunc(item.type), isStruct, isNumber(item.type)), 2);
    // 6 append lastest func
    tracer.append(tracer.getLatestFunc(getTypeTag(item.type), getParamPrefix(item), getValueFunc(item.type), isStruct, isNumber(item.type)), 2);
    // 7 append diff func (only for number type)
    if (isNumber(item.type)) {
        tracer.append(tracer.getDiffFunc(getTypeTag(item.type), getParamPrefix(item), getValueFunc(item.type), isStruct, isNumber(item.type)), 2);
    }
    // 1' append class end
    tracer.append(tracer.endBracket, 1);
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
        tracer.append(tracer.getStructParam(item.label, structName + "_" + item.label), 2);
    });
    // 1' append class end
    tracer.append(tracer.endBracket, 1);
    // 5 handle struct params to class
    members.forEach(function (item) {
        handleBasic(structName + "_" + item.label, item, tracer, true);
    });
}
exports.handleStruct = handleStruct;
function handleMapping(item, tracer, structNameSet, obj) {
    // 1 append class start
    tracer.append(tracer.getClass(item.label), 1);
    // 2 append addr and prefix
    tracer.append(tracer.argsTemplage, 2);
    // 3 append constructor
    tracer.append(tracer.constructorTemplate, 2);
    // 4 handle map second param
    var secondParamIsBasic = true;
    var secondParamType = getStrBetLastCommaAndParen(item.type);
    if (secondParamType.startsWith("t_struct")) {
        secondParamIsBasic = false;
    }
    if (!secondParamIsBasic) {
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
        tracer.append(tracer.getBeforeFuncMap(getTypeTag(item.type), getParamPrefix(item), getValueFunc(item.type), isNumber(item.type)), 2);
        // 4.2 append changes func
        tracer.append(tracer.getChangesFuncMap(getTypeTag(item.type), getParamPrefix(item), getValueFunc(item.type), isNumber(item.type)), 2);
        // 4.3 append lastest func
        tracer.append(tracer.getLatestFuncMap(getTypeTag(item.type), getParamPrefix(item), getValueFunc(item.type), isNumber(item.type)), 2);
        // 4.4 append diff func (only for number type)
        if (isNumber(item.type)) {
            tracer.append(tracer.getDiffFuncMap(getTypeTag(item.type), getParamPrefix(item), getValueFunc(item.type), isNumber(item.type)), 2);
        }
    }
    // 1' append class end
    tracer.append(tracer.endBracket, 1);
}
exports.handleMapping = handleMapping;
