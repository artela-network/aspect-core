"use strict"

const dot = require('dot')
const fs = require('fs')
const path = require('path')

const store = require("./store.json");

let aspectCode = fs.readFileSync('./store.jst',"UTF-8");

var tpl = dot.template(aspectCode);
var resultText = tpl(store);

fs.writeFileSync(path.resolve('./data.ts'),Buffer.from(resultText))
console.log(resultText);