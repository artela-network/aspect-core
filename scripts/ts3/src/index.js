"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var Generator_1 = require("./Generator");
var util = require("./Utils");
var readline = require("readline");
var sourceFilePath = "";
var targetFilePath = "";
var rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});
rl.question('Input your storage layout json file path:', function (sourcePath) {
    rl.question('Input your target generated ts file path:', function (targetPath) {
        sourceFilePath = sourcePath;
        targetFilePath = targetPath;
        rl.close();
        if (util.isStringEmpty(sourceFilePath) || util.isStringEmpty(targetFilePath) ||
            !sourceFilePath.endsWith(".json") || !targetFilePath.endsWith(".ts")) {
            console.log('Illegal input!');
            process.exit(0);
        }
        var tracer = new Generator_1.default(sourceFilePath, targetFilePath);
        var structNameSet = new Set();
        var jsonStr = tracer.getLayoutJson();
        var obj = tracer.getStorage(jsonStr);
        var items = obj.storage;
        // 1. append reference
        tracer.append(tracer.refLib, 0);
        // 2.1 append namespace start
        tracer.append(tracer.getNameSpace(util.getStrAfterLastColon(items[0].contract)), 0);
        // ----- 3.1 Loop to handle multi params start ------
        items.forEach(function (item) {
            var structName = util.getStructName(item.type);
            if (structName.startsWith("t_mapping")) {
                util.handleMapping(item, tracer, structNameSet, obj);
            }
            else if (!util.isStringEmpty(structName)) {
                if (!structNameSet.has(structName)) {
                    var members = obj.types[item.type].members;
                    structNameSet.add(structName);
                    util.handleStruct(item, tracer, structName, members);
                }
            }
            else {
                util.handleBasic(item.label, item, tracer, false);
            }
        });
        // ----- 3.2 Loop to handle multi params end ------
        // 2.2 append namespace end
        tracer.append(tracer.endBracket, 0);
    });
});
