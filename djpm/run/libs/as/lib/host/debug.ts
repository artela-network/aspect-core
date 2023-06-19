import {AString, AUint8Array} from "../message";
import {utils} from "../utils";

declare namespace __debug__ {
    function log(ptr: i32): void;
}

export namespace debug {
    export function log(data: string): void {
        let dataPtr = new AString(data).store();
        __debug__.log(dataPtr);
    }
}